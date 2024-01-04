package nationalrail

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Client struct {
	accessToken string
	httpClient  *http.Client
	url         string
}

var ErrWebServiceResponseMissing = errors.New("web-service response is missing")

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

	if client.accessToken == "" {
		client.accessToken = os.Getenv("NR_ACCESS_TOKEN")
		if client.accessToken == "" {
			return nil, errors.New("access token is missing, provide it via option or env var")
		}
	}

	if client.url == "" {
		client.url = "https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb12.asmx"
	}

	return client, nil
}

// GetArrivalsWithDetails returns all public arrivals for the supplied
// CRS code, including service details.
func (c *Client) GetArrivalsWithDetails(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetArrBoardWithDetailsRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetArrBoardWithDetailsResponse(resp)
}

// GetArrivalsAndDeparturesWithDetails returns all public arrivals and
// departures for the supplied CRS code, including service details.
func (c *Client) GetArrivalsAndDeparturesWithDetails(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetArrDepBoardWithDetailsRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetArrDepBoardWithDetailsResponse(resp)
}

// GetArrivals returns all public arrivals for the supplied CRS code.
func (c *Client) GetArrivals(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetArrivalBoardRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetArrivalBoardResponse(resp)
}

// GetArrivalsAndDepartures returns all public arrivals and departures for
// the supplied CRS code.
func (c *Client) GetArrivalsAndDepartures(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetArrivalDepartureBoardRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetArrivalDepartureBoardResponse(resp)
}

// GetDepartures returns all public departures for the supplied CRS code.
func (c *Client) GetDepartures(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetDepartureBoardRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetDepartureBoardResponse(resp)
}

// GetDeparturesWithDetails returns all public departures for the supplied
// CRS code, including service details.
func (c *Client) GetDeparturesWithDetails(
	crs CRSCode,
	opts ...RequestOption,
) (*StationBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructStationBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(CreateGetDepBoardWithDetailsRequest(c.accessToken, req))
	if err != nil {
		return nil, err
	}

	return MapGetDepBoardWithDetailsResponse(resp)
}

// GetFastestDepartures returns the public departure for the supplied CRS code
// to the locations specified in the filter with the earliest arrival time at
// the filtered location.
func (c *Client) GetFastestDepartures(
	crs CRSCode,
	destinations []CRSCode,
	opts ...RequestOption,
) (*DeparturesBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructDeparturesBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	if len(destinations) < 1 || len(destinations) > 15 {
		return nil, errors.New("destinations must be in range of [1, 15]")
	}

	resp, err := c.makeRequest(CreateGetFastestDeparturesRequest(c.accessToken, req, destinations))
	if err != nil {
		return nil, err
	}

	return MapGetFastestDeparturesResponse(resp)
}

// GetFastestDeparturesWithDetails returns the public departure for the supplied
// CRS code to the locations specified in the filter with the earliest arrival time
// at the filtered location, including service details.
func (c *Client) GetFastestDeparturesWithDetails(
	crs CRSCode,
	destinations []CRSCode,
	opts ...RequestOption,
) (*DeparturesBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructDeparturesBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	if len(destinations) < 1 || len(destinations) > 10 {
		return nil, errors.New("destinations must be in range of [1, 10]")
	}

	resp, err := c.makeRequest(CreateGetFastestDeparturesWithDetailsRequest(c.accessToken, req,
		destinations))
	if err != nil {
		return nil, err
	}

	return MapGetFastestDeparturesWithDetailsResponse(resp)
}

// GetNextDepartures returns the next public departure for the supplied CRS code
// to the locations specified in the filter.
func (c *Client) GetNextDepartures(
	crs CRSCode,
	destinations []CRSCode,
	opts ...RequestOption,
) (*DeparturesBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructDeparturesBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	if len(destinations) < 1 || len(destinations) > 25 {
		return nil, errors.New("destinations must be in range of [1, 25]")
	}

	resp, err := c.makeRequest(CreateGetNextDeparturesRequest(c.accessToken, req, destinations))
	if err != nil {
		return nil, err
	}

	return MapGetNextDeparturesResponse(resp)
}

// GetNextDeparturesWithDetails returns the next public departure for the supplied
// CRS code to the locations specified in the filter, including service details.
func (c *Client) GetNextDeparturesWithDetails(
	crs CRSCode,
	destinations []CRSCode,
	opts ...RequestOption,
) (*DeparturesBoard, error) {
	if StationCodeToNameMap[crs] == "" {
		return nil, errors.New("a valid CRS code is required")
	}

	req, err := constructDeparturesBoardRequest(crs, opts)
	if err != nil {
		return nil, err
	}

	if len(destinations) < 1 || len(destinations) > 10 {
		return nil, errors.New("destinations must be in range of [1, 10]")
	}

	resp, err := c.makeRequest(CreateGetNextDeparturesWithDetailsRequest(c.accessToken, req,
		destinations))
	if err != nil {
		return nil, err
	}

	return MapGetNextDeparturesWithDetailsResponse(resp)
}

// GetServiceDetails returns service details for a specific service identified by
// a station board. These details are supplied relative to the station board from
// which the Service ID field value was generated. Service details are only available
// while the service appears on the station board from which it was obtained.
// This is normally for two minutes after it is expected to have departed, or after
// a terminal arrival. If a request is made for a service that is no longer available
// then ErrGone error is returned.
func (c *Client) GetServiceDetails(serviceID string) (*ServiceDetails, error) {
	resp, err := c.makeRequest(CreateGetServiceDetailsRequest(c.accessToken, serviceID))
	if err != nil {
		return nil, err
	}

	return MapGetServiceDetailsResponse(resp)
}

// GetCRSFromKeyword used to match the given keyword to train station names
// to obtain matching CRS codes.
func (c *Client) GetCRSFromKeyword(keyword string) map[string]CRSCode {
	crsMap := make(map[string]CRSCode)
	for station, crs := range StationNameToCodeMap {
		if strings.Contains(strings.ToLower(station), strings.ToLower(keyword)) {
			crsMap[station] = crs
		}
	}

	return crsMap
}

func (c *Client) makeRequest(body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new HTTP POST request: %w", err)
	}

	req.Header.Add("Content-Type", "text/xml")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected HTTP 200 response, got %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return respBody, nil
}

func constructStationBoardRequest(
	crs CRSCode,
	opts []RequestOption,
) (*StationBoardRequest, error) {
	props := &RequestProperties{}
	for _, opt := range opts {
		if err := opt(props); err != nil {
			return nil, err
		}
	}

	req := StationBoardRequest{
		CRS:        crs,
		FilterCRS:  props.FilterCRS,
		TimeOffset: props.TimeOffset,
		TimeWindow: props.TimeWindow,
	}
	if props.NumRows != nil {
		req.NumRows = *props.NumRows
	}

	return &req, nil
}

func constructDeparturesBoardRequest(
	crs CRSCode,
	opts []RequestOption,
) (*DeparturesBoardRequest, error) {
	props := &RequestProperties{}
	for _, opt := range opts {
		if err := opt(props); err != nil {
			return nil, err
		}
	}

	req := DeparturesBoardRequest{
		CRS:        crs,
		TimeOffset: props.TimeOffset,
		TimeWindow: props.TimeWindow,
	}

	return &req, nil
}
