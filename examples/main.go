package main

import (
	"encoding/json"

	c "github.com/MartinsIrbe/national-rail-go-client/pkg"
	"github.com/sirupsen/logrus"
)

func main() {
	client := c.NewNationalRailClient("YOUR_ACCESS_TOKEN")

	getDeparturesExample(client)
}

func getDeparturesExample(client *c.NationalRailClient) {
	originCode := "CHX"
	destinationCode := "DFD"
	numberOfDepartures := 5

	departures, err := client.GetDepartures(originCode, destinationCode, numberOfDepartures)
	if err != nil {
		logrus.WithError(err).
			Panicf("failed to get departure for %s-%s with %d departures",
				originCode, destinationCode, numberOfDepartures)
	}

	b, err := json.Marshal(departures)
	if err != nil {
		logrus.WithError(err).
			Panicf("failed to marshal departure struct to a json string, departures - %+v",
				departures)
	}

	logrus.Infof("departures: %s", string(b))
}
