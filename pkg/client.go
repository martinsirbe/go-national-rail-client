package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"bufio"
	"encoding/csv"
	"github.com/MartinsIrbe/national-rail-go-client/internal/decoder"
	"github.com/MartinsIrbe/national-rail-go-client/internal/mapper"
	"github.com/MartinsIrbe/national-rail-go-client/pkg/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

const nationalRailWebService = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb6.asmx"

type NationalRailClient struct {
	Token                string
	StationToCRSCodeMaps map[string]string
	ResponseDecoder      *decoder.NationalRailResponseDecoder
	ResponseMapper       *mapper.NationalRailResponseMapper
}

func NewNationalRailClient(token string) *NationalRailClient {
	return &NationalRailClient{
		Token:                token,
		StationToCRSCodeMaps: loadStations(),
		ResponseDecoder:      decoder.NewNationalRailResponseDecoder(),
		ResponseMapper:       mapper.NewNationalRailResponseMapper(),
	}
}

func loadStations() map[string]string {
	stations, err := os.Open("station_codes.csv")
	if err != nil {
		logrus.WithError(err).Panic("failed to load train station codes")
	}
	reader := csv.NewReader(bufio.NewReader(stations))

	stationCodeMap := make(map[string]string)

	for {
		l, readErr := reader.Read()
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			logrus.WithError(readErr).Error("failed to read line from the station codes csv file")
		}

		stationCodeMap[l[0]] = l[1]
	}

	return stationCodeMap
}

// GetDepartures used to obtain departure information
func (c *NationalRailClient) GetDepartures(originCode, destinationCode string, numberOfDepartures int) (*models.DepartureInfo, error) {
	body := c.createRequestBody(originCode, destinationCode, numberOfDepartures)
	client := &http.Client{}

	req, newReqErr := http.NewRequest(http.MethodPost, nationalRailWebService, body)
	if newReqErr != nil {
		return nil, errors.Wrapf(
			newReqErr,
			"failed to create a new HTTP POST request for %s",
			nationalRailWebService)
	}

	req.Header.Add("Content-Type", "text/xml")

	resp, reqErr := client.Do(req)
	if reqErr != nil || resp.StatusCode != 200 {
		return nil, errors.Wrapf(
			newReqErr,
			"failed to perform HTTP POST request to %s, resp - %+v",
			nationalRailWebService,
			resp)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to obtain departure information, current location code - %s, destination code - %s",
			originCode,
			destinationCode)
	}

	decodedResp, decodeErr := c.ResponseDecoder.DecodeDepartureResponse(strings.NewReader(string(respBody)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode national rail xml response, response body - %s", string(respBody))
	}

	mappedResp, mapErr := c.ResponseMapper.MapDepartureResponse(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to a simplified model")
	}

	return mappedResp, nil
}

// GetCRSFromKeyword used to match the given keyword to train station names to obtain matching CRS codes
func (c *NationalRailClient) GetCRSFromKeyword(keyword string) map[string]string {
	crsMap := make(map[string]string)
	for station, crs := range c.StationToCRSCodeMaps {
		if strings.Contains(strings.ToLower(station), strings.ToLower(keyword)) {
			crsMap[station] = crs
		}
	}

	return crsMap
}

func (c *NationalRailClient) createRequestBody(from, to string, dep int) *strings.Reader {
	body := strings.NewReader(
		fmt.Sprintf(`<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="http://thalesgroup.com/RTTI/2014-02-20/ldb/" xmlns:ns2="http://thalesgroup.com/RTTI/2010-11-01/ldb/commontypes">
		  <SOAP-ENV:Header>
			<ns2:AccessToken>
				<ns2:TokenValue>%s</ns2:TokenValue>
			</ns2:AccessToken>
		  </SOAP-ENV:Header>
		  <SOAP-ENV:Body>
			<ns1:GetDepartureBoardRequest>
			  <ns1:numRows>%d</ns1:numRows>
			  <ns1:crs>%s</ns1:crs>
			  <ns1:filterCrs>%s</ns1:filterCrs>
			</ns1:GetDepartureBoardRequest>
		  </SOAP-ENV:Body>
		</SOAP-ENV:Envelope>`, c.Token, dep, from, to))
	return body
}
