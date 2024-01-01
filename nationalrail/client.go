package nationalrail

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

	"github.com/martinsirbe/national-rail-go-client/nationalrail/soap"
)

const nationalRailWebService = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx"

type Client struct {
	Token                string
	StationToCRSCodeMaps map[string]string
	httpClient           *http.Client
}

func NewClient(token string) *Client {
	return &Client{
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

func (c *Client) GetArrBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeArrBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrBoardWithDetailsResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapArrivalBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetArrDepBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrDepBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeArrDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrDepBoardWithDetailsResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapArrDepBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetArrivalBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrivalBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeArrivalBoardResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrivalBoardResponse xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapArrivalBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetArrivalDepartureBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrivalDepartureBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeArrivalDepartureBoardResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ArrivalDepartureBoard xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapArrivalDepartureBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetDepartureBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetDepartureBoardRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	r := string(resp)
	fmt.Println(r)

	decodedResp, decodeErr := soap.DecodeDepartureBoardResponse(strings.NewReader(r))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode DepartureBoard xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapDepartureBoard(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetDepBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetDepBoardWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode DepBoardWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapDepBoardWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	logrus.Infof("GetDepBoardWithDetails - %+v", mappedResp)

	return mappedResp, nil
}

func (c *Client) GetFastestDepartures(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetFastestDeparturesRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeFastestDeparturesResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode FastestDepartures xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapFastestDepartures(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetFastestDeparturesWithDetails(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetFastestDeparturesWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeFastestDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode FastestDeparturesWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapFastestDeparturesWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetNextDepartures(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetNextDeparturesRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeNextDeparturesResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode NextDepartures xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapNextDepartures(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetNextDeparturesWithDetails(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetNextDeparturesWithDetailsRequest(c.Token, req))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeNextDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode NextDeparturesWithDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapNextDeparturesWithDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

func (c *Client) GetServiceDetails(serviceID string) (*TrainServiceDetails, error) {
	resp, err := c.makeRequest(CreateGetServiceDetailsRequest(c.Token, serviceID))
	if err != nil {
		return nil, err
	}

	decodedResp, decodeErr := soap.DecodeServiceDetailsResponse(strings.NewReader(string(resp)))
	if decodeErr != nil {
		return nil, errors.Wrapf(decodeErr, "failed to decode ServiceDetails xml response, response body - %s", string(resp))
	}

	mappedResp, mapErr := MapServiceDetails(decodedResp)
	if mapErr != nil {
		return nil, errors.Wrap(mapErr, "failed to map national rail response to the simplified model")
	}

	return mappedResp, nil
}

// GetCRSFromKeyword used to match the given keyword to train station names to obtain matching CRS codes
func (c *Client) GetCRSFromKeyword(keyword string) map[string]string {
	crsMap := make(map[string]string)
	for station, crs := range c.StationToCRSCodeMaps {
		if strings.Contains(strings.ToLower(station), strings.ToLower(keyword)) {
			crsMap[station] = crs
		}
	}

	return crsMap
}

func (c *Client) makeRequest(body io.Reader) ([]byte, error) {
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
