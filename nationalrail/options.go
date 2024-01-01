package nationalrail

import "net/http"

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
