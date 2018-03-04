# National Rail Go client
A simple National Rail web-service client written in Golang that allows to communicate with National Rail SOAP web-service.
The current implementation includes only support for checking for departures by providing the origin and destination train station codes.
To obtain National Rail web-service access token you can register for OpenLDBWS and obtain an access token over [here][1].

## Download
```bash
go get github.com/MartinsIrbe/national-rail-go-client
```

## Initialise the client
```golang
client := c.NewNationalRailClient("YOUR_ACCESS_TOKEN")
``` 

## Get departures
```golang
originCode := "CHX"
destinationCode := "DFD"

departures, err := client.GetDepartures(originCode, destinationCode)
if err != nil {
    logrus.WithError(err).
        Panicf("failed to get departure for %s-%s", originCode, destinationCode)
}
```

### Example JSON response
```json
{
    "origin_code": "CHX",
    "destination_code": "DFD",
    "departure_details": [
        {
            "origin_location": "London Charing Cross",
            "destination_location": "Gillingham (Kent)",
            "via": "via Lewisham & Woolwich Arsenal",
            "departure_time": "0000-01-01T17:09:00Z",
            "status": "On time",
            "operator": "Southeastern",
            "platform": "1"
        },
        {
            "origin_location": "London Charing Cross",
            "destination_location": "Dartford",
            "via": "via Bexleyheath",
            "departure_time": "0000-01-01T17:14:00Z",
            "status": "On time",
            "operator": "Southeastern",
            "platform": "5"
        },
        {
            "origin_location": "London Charing Cross",
            "destination_location": "Gravesend",
            "via": "via Lewisham & Sidcup",
            "departure_time": "0000-01-01T17:18:00Z",
            "status": "On time",
            "operator": "Southeastern",
            "platform": "2"
        },
        {
            "origin_location": "London Charing Cross",
            "destination_location": "Gillingham (Kent)",
            "via": "via Lewisham & Woolwich Arsenal",
            "departure_time": "0000-01-01T17:39:00Z",
            "status": "On time",
            "operator": "Southeastern",
            "platform": "not available"
        }
    ]
}
```

## Run tests
```golang
go test -v --cover ./...
```

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

[1]: http://realtime.nationalrail.co.uk/OpenLDBWSRegistration/