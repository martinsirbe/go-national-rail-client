package nationalrail

import (
	"errors"
	"net/http"
)

type ClientOption func(*Client)

func URLOpt(url string) ClientOption {
	return func(c *Client) {
		c.url = url
	}
}

func HTTPClientOpt(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// AccessTokenOpt sets the access token for the National Rail API client.
// To obtain an access token, register at https://realtime.nationalrail.co.uk/OpenLDBWSRegistration/.
// This registration is necessary for accessing the Darwin Web Service, which provides real-time
// train information.
func AccessTokenOpt(accessToken string) ClientOption {
	return func(c *Client) {
		c.accessToken = accessToken
	}
}

type RequestProperties struct {
	NumRows    *int
	FilterCRS  *string
	FilterType *string
	TimeOffset *int
	TimeWindow *int
}

type RequestOption func(*RequestProperties) error

// OriginatingFromOpt filters services to include only those originating
// from the specified location based on the provided CRS code.
func OriginatingFromOpt(crs CRSCode) RequestOption {
	return func(p *RequestProperties) error {
		from := "from"
		fcrs := string(crs)
		p.FilterType = &from
		p.FilterCRS = &fcrs
		return nil
	}
}

// TerminatingAtOpt filters services to include only those terminating
// at the specified location based on the provided CRS code.
func TerminatingAtOpt(crs CRSCode) RequestOption {
	return func(p *RequestProperties) error {
		// When the filter type is not provided, defaults to `to`.
		fcrs := string(crs)
		p.FilterCRS = &fcrs
		return nil
	}
}

// TimeOffsetMinutesOpt an offset in minutes (in range of -120 to 120 inclusive),
// against the current time to provide the station board for. When not
// provided, defaults to 0.
func TimeOffsetMinutesOpt(n int) RequestOption {
	return func(p *RequestProperties) error {
		if n < -120 || n > 120 {
			return errors.New("time offset must be in range of [-120, 120]")
		}
		p.TimeOffset = &n
		return nil
	}
}

// TimeWindowMinutesOpt time window of how far into the future in minutes (in range
// of -120 to 120 inclusive), relative to time offset, to return services for.
// When not provided, defaults to 0.
func TimeWindowMinutesOpt(n int) RequestOption {
	return func(p *RequestProperties) error {
		if n < -120 || n > 120 {
			return errors.New("time window must be in range of [-120, 120]")
		}
		p.TimeWindow = &n
		return nil
	}
}

// NumRowsOpt the number of services to return in the resulting station board.
// Must be in range of 1 to 10 inclusive. When not provided, defaults to 10.
func NumRowsOpt(n int) RequestOption {
	return func(p *RequestProperties) error {
		if n < 1 || n > 10 {
			return errors.New("number of rows must be in range of [1, 10]")
		}
		p.NumRows = &n
		return nil
	}
}
