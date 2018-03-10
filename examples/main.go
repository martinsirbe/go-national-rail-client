package main

import (
	c "github.com/MartinsIrbe/national-rail-go-client/pkg"
	"github.com/MartinsIrbe/national-rail-go-client/pkg/models"
	"github.com/sirupsen/logrus"
)

func main() {
	client := c.NewNationalRailClient("17809a9c-2578-46b7-827e-08af0cec7e9f")

	//get(client)

	getServiceDetails(client)
}

//
//func getDeparturesExample(client *c.NationalRailClient) {
//	originCode := "CHX"
//	destinationCode := "DFD"
//	numberOfDepartures := 5
//
//	departures, err := client.GetDepartureBoard(originCode, destinationCode, numberOfDepartures)
//	if err != nil {
//		logrus.WithError(err).
//			Panicf("failed to get departure for %s-%s with %d departures",
//				originCode, destinationCode, numberOfDepartures)
//	}
//
//	b, err := json.Marshal(departures)
//	if err != nil {
//		logrus.WithError(err).
//			Panicf("failed to marshal departure struct to a json string, departures - %+v",
//				departures)
//	}
//
//	logrus.Infof("departures: %s", string(b))
//}

func get(client *c.NationalRailClient) {
	timeOffset := 0
	timeWindow := 120
	resp, err := client.GetNextDeparturesWithDetails(models.DeparturesBoardRequest{
		CRS:        "GLM",
		NumRows:    4,
		FilterList: []string{"VIC"},
		TimeOffset: &timeOffset,
		TimeWindow: &timeWindow,
	})
	if err != nil {
		logrus.WithError(err).Error("GetArrBoardWithDetails request failed")
	}

	logrus.Infof("calling points %+v", resp)
}

func getServiceDetails(client *c.NationalRailClient) {
	resp, err := client.GetServiceDetails("+ATOEG3lw8q/kEoAMopEBw==")
	if err != nil {
		logrus.WithError(err).Error("getServiceDetails request failed")
	}

	logrus.Infof("calling points %+v", resp)
}
