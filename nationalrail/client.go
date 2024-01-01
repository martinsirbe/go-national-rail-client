package nationalrail

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/martinsirbe/go-national-rail-client/nationalrail/soap"
)

const nationalRailWebService = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx"

type Client struct {
	AccessToken string
	httpClient  *http.Client
}

func NewClient(opts ...ClientOption) (*Client, error) {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}

	if client.httpClient == nil {
		t := http.DefaultTransport.(*http.Transport).Clone()
		t.MaxIdleConns = 10
		t.MaxConnsPerHost = 10
		t.MaxIdleConnsPerHost = 10

		client.httpClient = &http.Client{
			Transport: t,
			Timeout:   10 * time.Second,
		}
	}

	if client.AccessToken == "" {
		client.AccessToken = os.Getenv("NR_ACCESS_TOKEN")
		if client.AccessToken == "" {
			return nil, errors.New("access token is missing, provide it via option or env var")
		}
	}

	return client, nil
}

func (c *Client) GetArrBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrBoardWithDetailsRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeArrBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapArrivalBoardWithDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetArrDepBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrDepBoardWithDetailsRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeArrDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapArrDepBoardWithDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetArrivalBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrivalBoardRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeArrivalBoardResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapArrivalBoard(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetArrivalDepartureBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetArrivalDepartureBoardRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeArrivalDepartureBoardResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapArrivalDepartureBoard(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetDepartureBoard(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetDepartureBoardRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	r := string(resp)
	fmt.Println(r)

	decodedResp, err := soap.DecodeDepartureBoardResponse(strings.NewReader(r))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapDepartureBoard(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetDepBoardWithDetails(req StationBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetDepBoardWithDetailsRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeDepBoardWithDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapDepBoardWithDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetFastestDepartures(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetFastestDeparturesRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeFastestDeparturesResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapFastestDepartures(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetFastestDeparturesWithDetails(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetFastestDeparturesWithDetailsRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeFastestDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapFastestDeparturesWithDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetNextDepartures(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetNextDeparturesRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeNextDeparturesResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapNextDepartures(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetNextDeparturesWithDetails(req DeparturesBoardRequest) (*StationBoard, error) {
	resp, err := c.makeRequest(CreateGetNextDeparturesWithDetailsRequest(c.AccessToken, req))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeNextDeparturesWithDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapNextDeparturesWithDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

func (c *Client) GetServiceDetails(serviceID string) (*TrainServiceDetails, error) {
	resp, err := c.makeRequest(CreateGetServiceDetailsRequest(c.AccessToken, serviceID))
	if err != nil {
		return nil, err
	}

	decodedResp, err := soap.DecodeServiceDetailsResponse(strings.NewReader(string(resp)))
	if err != nil {
		return nil, fmt.Errorf("failed to decode xml response: %w", err)
	}

	mappedResp, err := MapServiceDetails(decodedResp)
	if err != nil {
		return nil, fmt.Errorf("failed to map response to the simplified model: %w", err)
	}

	return mappedResp, nil
}

// GetCRSFromKeyword used to match the given keyword to train station names to obtain matching CRS codes
func (c *Client) GetCRSFromKeyword(keyword string) map[string]string {
	crsMap := make(map[string]string)
	for station, crs := range StationNameToCodeMap {
		if strings.Contains(strings.ToLower(station), strings.ToLower(keyword)) {
			crsMap[station] = crs
		}
	}

	return crsMap
}

func (c *Client) makeRequest(body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, nationalRailWebService, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new HTTP POST request: %w", err)
	}

	req.Header.Add("Content-Type", "text/xml")

	resp, err := c.httpClient.Do(req)
	defer resp.Body.Close()

	switch {
	case err != nil:
		return nil, fmt.Errorf("request failed: %w", err)
	case resp.StatusCode != http.StatusOK:
		return nil, fmt.Errorf("expected HTTP 200 response, got %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}
