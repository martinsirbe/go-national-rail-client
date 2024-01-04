# go-national-rail-client

[![Go Report Card](https://goreportcard.com/badge/github.com/martinsirbe/go-national-rail-client)](https://goreportcard.com/report/github.com/martinsirbe/go-national-rail-client)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmartinsirbe%2Fgo-national-rail-client.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmartinsirbe%2Fgo-national-rail-client?ref=badge_shield)
[![codecov](https://codecov.io/gh/martinsirbe/go-national-rail-client/branch/main/graph/badge.svg)](https://codecov.io/gh/martinsirbe/go-national-rail-client)

An open-source Golang library for interfacing with the National Rail's Darwin OpenLDBWS web service. This client 
enables the retrieval of real-time train information, focusing on departures and arrivals between specific stations.

The client have been developed based on [Live Departure Boards Web Service][1] documentation.

## Getting Started
To use this client, register for an OpenLDBWS access token at [National Rail Enquiries][2].

## Configuration Access Token

The access token is mandatory and can be set in one of two ways:
1. By setting the `NR_ACCESS_TOKEN` environment variable.
2. By using the `AccessTokenOpt("YOUR_ACCESS_TOKEN")` client option.

If the client option is not provided, the client will attempt to use the access token from 
the environment variable. Please note that the access token is crucial for OpenLDBWS authentication; without it, 
client initialisation will fail.  
If an invalid access token is provided, requests will fail authentication.

## Examples
### Obtain Departure Board With Details
To retrieve the departure board for a specific train station, provide the station code and the desired number of 
departures. This example demonstrates how to obtain the departure board for Gillingham (Kent) [GLM], displaying 5 
upcoming train departures.

```golang
package main

import (
	"fmt"
	"strconv"

	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

func main() {
  client, err := nr.NewClient(
    nr.AccessTokenOpt("e995265f-df60-4787-bafc-af5a433f9b22"),
  )
  if err != nil {
    panic(err)
  }

  board, err := client.GetDeparturesWithDetails(nr.StationCodeGillinghamKent, nr.NumRowsOpt(5))
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s [%s] Departure Board:\n", board.LocationName, board.CRS)
  fmt.Println("----------------------------------------")
  fmt.Println(" Time\t| Platform\t| Departure Status\t| Destination\t")
  fmt.Println("----------------------------------------")

  for _, s := range board.Services {
    std := "?"
    if s.STD != nil {
      std = *s.STD
    }

    platform := "X"
    if s.Platform != 0 {
      platform = strconv.Itoa(s.Platform)
    }

    etd := "?"
    if s.ETD != nil {
      etd = *s.ETD
    }

    fmt.Printf(" %s\t| %s\t| %s\t| %s [%s]\t\n", std, platform, etd, s.Destination.Name, s.Destination.CRS)
  }
}
```

Output:
```shell
Gillingham (Kent) [GLM] Departure Board:
----------------------------------------
 Time   | Platform      | Departure Status      | Destination   
----------------------------------------
 09:24  | X     | Cancelled             | London St Pancras (Intl) [STP]        
 09:30  | 2     | 09:32         | London Victoria [VIC] 
 09:30  | 3     | 09:32         | Ramsgate [RAM]        
 09:34  | 2     | On time               | Luton [LUT]   
 09:36  | 3     | 09:47         | Faversham [FAV]    
```

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[1]: https://lite.realtime.nationalrail.co.uk/openldbws/
[2]: https://realtime.nationalrail.co.uk/OpenLDBWSRegistration/
