# National Rail Go client
A simple National Rail web-service client written in Golang that allows to communicate with National Rail SOAP web-service.
The current implementation includes only support for checking for departures by providing the origin and destination train station codes.
To obtain National Rail web-service access token you can register for OpenLDBWS and obtain an access token over [here][1].

## Download
```bash
go get github.com/martinsirbe/go-national-rail-client
```

## Configuration

The API client can be configured using the following options:
- **AccessTokenOpt**: Set using `AccessTokenOpt("YOUR_ACCESS_TOKEN")`. If not provided, the library attempts to 
retrieve the access token from the `NR_ACCESS_TOKEN` environment variable. This token is crucial for API authentication.  
Client initialisation will fail if neither `AccessTokenOpt` nor `NR_ACCESS_TOKEN` environment variable is provided.

- **HTTPClientOpt**: Customised using `HTTPClientOpt(customHttpClient)`. If not provided, a default HTTP client with 
the following configuration is used:
    - `Timeout`: 10 seconds
    - `MaxIdleConns`: 10 connections
    - `MaxConnsPerHost`: 10 connections
    - `MaxIdleConnsPerHost`: 10 connections

- **URLOpt**: Set using `URLOpt("https://example.com")`. If not provided, will use `https://lite.realtime.nationalrail.co.uk/OpenLDBWS/ldb11.asmx`.

## Examples
### Obtain Departure Board
To retrieve the departure board for a specific train station, provide the station code and the desired number of 
departures. This example demonstrates how to obtain the departure board for Gillingham (Kent) [GLM], displaying 5 
upcoming train departures.

```golang
package main

import (
	"fmt"

	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func main() {
	client, err := nr.NewClient(
		nr.AccessTokenOpt("e995265f-df60-4787-bafc-af5a433f9b22"),
	)

	board, err := client.GetDepartures(nr.StationBoardRequest{
		CRS:     nr.StationCodeGillinghamKent,
		NumRows: 5,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s [%s] Departure Board:\n", board.LocationName, board.CRS)
	fmt.Println("----------------------------------------")
	fmt.Println("Time  | Platform | Status  | Destination")
	fmt.Println("----------------------------------------")
	for _, s := range board.TrainServices {
		platform := "?"
		if s.Platform != nil {
			platform = *s.Platform
		}

		fmt.Printf("%s |    %s     | %s | %s [%s]\n", s.STD, platform, s.ETD, s.Destination.Name, s.Destination.CRS)
	}
}
```

Output:
```shell
Gillingham (Kent) [GLM] Departure Board:
----------------------------------------
Time  | Platform | Status  | Destination
----------------------------------------
12:57 |    2     | On time | London Cannon Street [CST]
13:04 |    2     | On time | Kentish Town [KTN]
13:07 |    3     | On time | Ramsgate [RAM]
13:10 |    2     | On time | London Cannon Street [CST]
13:11 |    3     | On time | Rainham (Kent) [RAI]
```

## Station codes
Station codes can be obtained from [National Rail stations destinations][2]

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[1]: http://realtime.nationalrail.co.uk/OpenLDBWSRegistration/
[2]: http://www.nationalrail.co.uk/stations_destinations/48541.aspx