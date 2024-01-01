package client

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/martinsirbe/national-rail-go-client/internal/nationalrail"
	m "github.com/martinsirbe/national-rail-go-client/pkg/models"
	"github.com/martinsirbe/national-rail-go-client/pkg/requests"
)

const nationalRailWebService = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx"

type NationalRailClient struct {
	Token                string
	StationToCRSCodeMaps map[string]string
	httpClient           *http.Client
}

func NewNationalRailClient(token string) *NationalRailClient {
	return &NationalRailClient{
		Token:                token,
		StationToCRSCodeMaps: loadStations(),
		httpClient:           &http.Client{},
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

func (c *NationalRailClient) GetArrBoardWithDetails(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetArrBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeArrBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrBoardWithDetailsResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapArrivalBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetArrDepBoardWithDetails(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetArrDepBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeArrDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrDepBoardWithDetailsResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapArrDepBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetArrivalBoard(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetArrivalBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeArrivalBoardResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrivalBoardResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapArrivalBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetArrivalDepartureBoard(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetArrivalDepartureBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeArrivalDepartureBoardResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrivalDepartureBoard xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapArrivalDepartureBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetDepartureBoard(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetDepartureBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	r := string(resp)
	fmt.Println(r)

	decodedResp, decodeErr := nationalrail.DecodeDepartureBoardResponse(strings.NewReader(r))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode DepartureBoard xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapDepartureBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetDepBoardWithDetails(req m.StationBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetDepBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode DepBoardWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapDepBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	logrus.Infof("GetDepBoardWithDetails - %+v", mappedResp)

	return mappedResp, nil
}

func (c *NationalRailClient) GetFastestDepartures(req m.DeparturesBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetFastestDeparturesRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeFastestDeparturesResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode FastestDepartures xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapFastestDepartures(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetFastestDeparturesWithDetails(req m.DeparturesBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetFastestDeparturesWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeFastestDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode FastestDeparturesWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapFastestDeparturesWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetNextDepartures(req m.DeparturesBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetNextDeparturesRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeNextDeparturesResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode NextDepartures xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapNextDepartures(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetNextDeparturesWithDetails(req m.DeparturesBoardRequest) (*m.StationBoard, error) {
	resp, err := c.makeRequest(requests.CreateGetNextDeparturesWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeNextDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode NextDeparturesWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapNextDeparturesWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *NationalRailClient) GetServiceDetails(serviceID string) (*m.TrainServiceDetails, error) {
	resp, err := c.makeRequest(requests.CreateGetServiceDetailsRequest(c.Token, serviceID))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := nationalrail.DecodeServiceDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ServiceDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := nationalrail.MapServiceDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
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

func (c *NationalRailClient) makeRequest(body io.Reader) ([]byte, error) {
	req, newReqErr := http.NewRequest(http.MethodPost, nationalRailWebService, body)
	if newReqErr != nil {
		return nil, errors.Wrapf(
			newReqErr,
			"failed to create a new HTTP POST request for %s",
			nationalRailWebService)
	}

	req.Header.Add("Content-Type", "text/xml")

	resp, reqErr := c.httpClient.Do(req)
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
		return nil, errors.Wrap(err, "fail")
	}

	return respBody, nil
}
