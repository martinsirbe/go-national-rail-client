# National Rail Go client
A simple National Rail web-service client written in Golang that allows to communicate with National Rail SOAP web-service.
The current implementation includes only support for checking for departures by providing the origin and destination train station codes.
To obtain National Rail web-service access token you can register for OpenLDBWS and obtain an access token over [here][1].

## Download
```bash
go get github.com/martinsirbe/go-national-rail-client
```

## Initialise the client
```golang
client := nationalrail.NewClient("NR_ACCESS_TOKEN")
``` 

## Examples
### Obtain Departure Board
To retrieve the departure board for a specific train station, provide the station code and the desired number of 
departures. This example demonstrates how to obtain the departure board for Gillingham (Kent) [GLM], displaying 5 
upcoming train departures.

```golang
package main

import (
	"fmt"

	"github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func main() {
	client := nationalrail.NewClient("NR_ACCESS_TOKEN")

	board, err := client.GetDepartureBoard(nationalrail.StationBoardRequest{
		CRS:     nationalrail.StationCodeGillinghamKent,
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