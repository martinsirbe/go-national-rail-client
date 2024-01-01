package nationalrail

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	nr "github.com/martinsirbe/national-rail-go-client/nationalrail/soap"
)

func MapArrivalBoardWithDetails(r *nr.GetArrBoardWithDetailsResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.ArrBoardWithDetails == nil:
		fallthrough
	case r.Body.ArrBoardWithDetails.StationBoardResult == nil:
		return nil, errors.New("failed to map ArrivalBoardWithDetails, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.ArrBoardWithDetails.StationBoardResult)

	return mappedResponse, nil
}

func MapArrDepBoardWithDetails(r *nr.GetArrDepBoardWithDetailsResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.ArrDepBoardWithDetails == nil:
		fallthrough
	case r.Body.ArrDepBoardWithDetails.StationBoardResult == nil:
		return nil, errors.New("failed to map ArrDepBoardWithDetails, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.ArrDepBoardWithDetails.StationBoardResult)

	return mappedResponse, nil
}

func MapArrivalBoard(r *nr.GetArrivalBoardResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.ArrivalBoard == nil:
		fallthrough
	case r.Body.ArrivalBoard.StationBoardResult == nil:
		return nil, errors.New("failed to map ArrivalBoard, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.ArrivalBoard.StationBoardResult)

	return mappedResponse, nil
}

func MapArrivalDepartureBoard(r *nr.GetArrivalDepartureBoardResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.ArrivalDepartureBoard == nil:
		fallthrough
	case r.Body.ArrivalDepartureBoard.StationBoardResult == nil:
		return nil, errors.New("failed to map ArrivalDepartureBoard, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.ArrivalDepartureBoard.StationBoardResult)

	return mappedResponse, nil
}

func MapDepartureBoard(r *nr.GetDepartureBoardResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.DepartureBoard == nil:
		fallthrough
	case r.Body.DepartureBoard.StationBoardResult == nil:
		return nil, errors.New("failed to map DepartureBoard, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.DepartureBoard.StationBoardResult)

	return mappedResponse, nil
}

func MapDepBoardWithDetails(r *nr.GetDepBoardWithDetailsResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.DepBoardWithDetails == nil:
		fallthrough
	case r.Body.DepBoardWithDetails.StationBoardResult == nil:
		return nil, errors.New("failed to map DepBoardWithDetails, expected objects were missing")
	}

	mappedResponse := mapStationBoardWithDetails(r.Body.DepBoardWithDetails.StationBoardResult)

	return mappedResponse, nil
}

func MapFastestDepartures(r *nr.GetFastestDeparturesResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.FastestDepartures == nil:
		fallthrough
	case r.Body.FastestDepartures.DeparturesBoard == nil:
		return nil, errors.New("failed to map FastestDepartures, expected objects were missing")
	}

	mappedResponse := mapDeparturesBoard(r.Body.FastestDepartures.DeparturesBoard)

	return mappedResponse, nil
}

func MapFastestDeparturesWithDetails(r *nr.GetFastestDeparturesWithDetailsResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.FastestDeparturesWithDetails == nil:
		fallthrough
	case r.Body.FastestDeparturesWithDetails.DeparturesBoard == nil:
		return nil, errors.New("failed to map FastestDeparturesWithDetails, expected objects were missing")
	}

	mappedResponse := mapDeparturesBoard(r.Body.FastestDeparturesWithDetails.DeparturesBoard)

	return mappedResponse, nil
}

func MapNextDepartures(r *nr.GetNextDeparturesResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.NextDepartures == nil:
		fallthrough
	case r.Body.NextDepartures.DeparturesBoard == nil:
		return nil, errors.New("failed to map NextDepartures, expected objects were missing")
	}

	mappedResponse := mapDeparturesBoard(r.Body.NextDepartures.DeparturesBoard)

	return mappedResponse, nil
}

func MapNextDeparturesWithDetails(r *nr.GetNextDeparturesWithDetailsResponse) (*StationBoard, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.NextDeparturesWithDetails == nil:
		fallthrough
	case r.Body.NextDeparturesWithDetails.DeparturesBoard == nil:
		return nil, errors.New("failed to map NextDeparturesWithDetails, expected objects were missing")
	}

	mappedResponse := mapDeparturesBoard(r.Body.NextDeparturesWithDetails.DeparturesBoard)

	return mappedResponse, nil
}

func MapServiceDetails(r *nr.GetServiceDetailsResponse) (*TrainServiceDetails, error) {
	switch {
	case r == nil:
		fallthrough
	case r.Body == nil:
		fallthrough
	case r.Body.ServiceDetails == nil:
		fallthrough
	case r.Body.ServiceDetails.ServiceDetailsResult == nil:
		return nil, errors.New("failed to map ServiceDetails, expected objects were missing")
	}

	mappedResponse := mapServiceDetails(r.Body.ServiceDetails.ServiceDetailsResult)

	return mappedResponse, nil
}

func mapServiceDetails(sdr *nr.ServiceDetailsResult) *TrainServiceDetails {
	mappedResponse := &TrainServiceDetails{}

	if sdr.GeneratedAt != nil && sdr.GeneratedAt.Text != "" {
		if t, err := time.Parse(time.RFC3339Nano, sdr.GeneratedAt.Text); err != nil {
			logrus.WithError(err).Errorf("failed to map generated at time date string - %s", sdr.GeneratedAt.Text)
		} else {
			mappedResponse.GeneratedAt = t
		}
	}

	if sdr.ServiceType != nil && sdr.ServiceType.Text != "" {
		mappedResponse.ServiceType = sdr.ServiceType.Text
	}

	if sdr.LocationName != nil && sdr.LocationName.Text != "" {
		mappedResponse.LocationName = sdr.LocationName.Text
	}

	if sdr.CRS != nil && sdr.CRS.Text != "" {
		mappedResponse.CRS = sdr.LocationName.Text
	}

	if sdr.STA != nil && sdr.STA.Text != "" {
		mappedResponse.STA = &sdr.STA.Text
	}

	if sdr.ETA != nil && sdr.ETA.Text != "" {
		mappedResponse.ETA = &sdr.ETA.Text
	}

	if sdr.STD != nil && sdr.STD.Text != "" {
		mappedResponse.STD = &sdr.STD.Text
	}

	if sdr.ETD != nil && sdr.ETD.Text != "" {
		mappedResponse.ETD = &sdr.ETD.Text
	}

	if sdr.Platform != nil && sdr.Platform.Text != "" {
		mappedResponse.Platform = &sdr.Platform.Text
	}

	if sdr.Operator != nil && sdr.Operator.Text != "" {
		mappedResponse.Operator = sdr.Operator.Text
	}

	if sdr.OperatorCode != nil && sdr.OperatorCode.Text != "" {
		mappedResponse.OperatorCode = sdr.OperatorCode.Text
	}

	if sdr.RSID != nil && sdr.RSID.Text != "" {
		mappedResponse.RSID = sdr.RSID.Text
	}

	if sdr.PreviousCallingPoints != nil && sdr.PreviousCallingPoints.CallingPointList != nil {
		previousCallingPoints := mapCallingPoints(
			sdr.PreviousCallingPoints.CallingPointList.CallingPoints,
			"previous calling point")
		mappedResponse.PreviousCallingPoints = previousCallingPoints
	}

	if sdr.SubsequentCallingPoints != nil && sdr.SubsequentCallingPoints.CallingPointList != nil {
		subsequentCallingPoints := mapCallingPoints(
			sdr.SubsequentCallingPoints.CallingPointList.CallingPoints,
			"subsequent calling point")
		mappedResponse.SubsequentCallingPoints = subsequentCallingPoints
	}

	return mappedResponse
}

func mapDeparturesBoard(depBoard *nr.DeparturesBoard) *StationBoard {
	mappedResponse := &StationBoard{}

	if depBoard.GeneratedAt != nil && depBoard.GeneratedAt.Text != "" {
		if t, err := time.Parse(time.RFC3339Nano, depBoard.GeneratedAt.Text); err != nil {
			logrus.WithError(err).Errorf("failed to map generated at time date string - %s", depBoard.GeneratedAt.Text)
		} else {
			mappedResponse.GeneratedAt = t
		}
	}

	if depBoard.LocationName != nil && depBoard.LocationName.Text != "" {
		mappedResponse.LocationName = depBoard.LocationName.Text
	}

	if depBoard.CRS != nil && depBoard.CRS.Text != "" {
		mappedResponse.CRS = depBoard.CRS.Text
	}

	if depBoard.PlatformAvailable != nil && depBoard.PlatformAvailable.Text != "" {
		if b, err := strconv.ParseBool(depBoard.PlatformAvailable.Text); err != nil {
			logrus.WithError(err).Errorf("failed to map platform available boolean string - %s", depBoard.PlatformAvailable.Text)
		} else {
			mappedResponse.PlatformAvailable = b
		}
	}

	if depBoard.Departures != nil && depBoard.Departures.Destination != nil {
		var trainServices []*TrainService
		for _, d := range depBoard.Departures.Destination {
			if d == nil || d.Service == nil {
				continue
			}

			trainServices = append(trainServices, mapTrainService(d.Service))
		}

		mappedResponse.TrainServices = trainServices
	}

	if depBoard.NRCCMessages != nil && depBoard.NRCCMessages.Messages != nil {
		var msgs []string
		for _, msg := range depBoard.NRCCMessages.Messages {
			msgs = append(msgs, msg.Text)
		}
		mappedResponse.NRCCMessages = &msgs
	}

	return mappedResponse
}

func mapStationBoardWithDetails(sbr *nr.GetStationBoardResult) *StationBoard {
	mappedResponse := &StationBoard{}
	if sbr.GeneratedAt != nil && sbr.GeneratedAt.Text != "" {
		if t, err := time.Parse(time.RFC3339Nano, sbr.GeneratedAt.Text); err != nil {
			logrus.WithError(err).Errorf("failed to map generated at time date string - %s", sbr.GeneratedAt.Text)
		} else {
			mappedResponse.GeneratedAt = t
		}
	}

	if sbr.LocationName != nil && sbr.LocationName.Text != "" {
		mappedResponse.LocationName = sbr.LocationName.Text
	}

	if sbr.CRS != nil && sbr.CRS.Text != "" {
		mappedResponse.CRS = sbr.CRS.Text
	}

	if sbr.FilterLocationName != nil && sbr.FilterLocationName.Text != "" {
		mappedResponse.FilterLocationName = &sbr.FilterLocationName.Text
	}

	if sbr.FilterCRS != nil && sbr.FilterCRS.Text != "" {
		mappedResponse.FilterCRS = &sbr.FilterCRS.Text
	}

	if sbr.PlatformAvailable != nil && sbr.PlatformAvailable.Text != "" {
		if b, err := strconv.ParseBool(sbr.PlatformAvailable.Text); err != nil {
			logrus.WithError(err).Errorf("failed to map platform available boolean string - %s", sbr.PlatformAvailable.Text)
		} else {
			mappedResponse.PlatformAvailable = b
		}
	}

	if sbr.TrainServices != nil && sbr.TrainServices.Services != nil {
		mappedResponse.TrainServices = mapTrainServices(sbr)
	}

	if sbr.NRCCMessages != nil && sbr.NRCCMessages.Messages != nil {
		var msgs []string
		for _, msg := range sbr.NRCCMessages.Messages {
			msgs = append(msgs, msg.Text)
		}
		mappedResponse.NRCCMessages = &msgs
	}

	return mappedResponse
}

func mapTrainServices(sbr *nr.GetStationBoardResult) []*TrainService {
	var trainServices []*TrainService
	for _, ts := range sbr.TrainServices.Services {
		if ts == nil {
			continue
		}

		trainServices = append(trainServices, mapTrainService(ts))
	}

	return trainServices
}

func mapTrainService(ts *nr.ServiceLT7) *TrainService {
	trainService := &TrainService{}

	if ts.STA != nil && ts.STA.Text != "" {
		trainService.STA = &ts.STA.Text
	}

	if ts.ETA != nil && ts.ETA.Text != "" {
		trainService.ETA = &ts.ETA.Text
	}

	if ts.STD != nil && ts.STD.Text != "" {
		trainService.STD = ts.STD.Text
	}

	if ts.ETD != nil && ts.ETD.Text != "" {
		trainService.ETD = ts.ETD.Text
	}

	if ts.Platform != nil && ts.Platform.Text != "" {
		trainService.Platform = &ts.Platform.Text
	}

	if ts.Operator != nil && ts.Operator.Text != "" {
		trainService.Operator = ts.Operator.Text
	}

	if ts.OperatorCode != nil && ts.OperatorCode.Text != "" {
		trainService.OperatorCode = ts.OperatorCode.Text
	}

	if ts.ServiceType != nil && ts.ServiceType.Text != "" {
		trainService.ServiceType = ts.ServiceType.Text
	}

	if ts.ServiceID != nil && ts.ServiceID.Text != "" {
		trainService.ServiceID = ts.ServiceID.Text
	}

	if ts.RSID != nil && ts.RSID.Text != "" {
		trainService.RSID = ts.RSID.Text
	}

	if ts.Origin != nil && ts.Origin.Location != nil {
		origin := mapLocation(ts.Origin.Location, "origin")
		trainService.Origin = origin
	}

	if ts.Destination != nil && ts.Destination.Location != nil {
		destination := mapLocation(ts.Destination.Location, "destination")
		trainService.Destination = destination
	}

	if ts.PreviousCallingPoints != nil && ts.PreviousCallingPoints.CallingPointList != nil {
		previousCallingPoints := mapCallingPoints(
			ts.PreviousCallingPoints.CallingPointList.CallingPoints,
			"previous calling point")
		trainService.PreviousCallingPoints = previousCallingPoints
	}

	if ts.SubsequentCallingPoints != nil && ts.SubsequentCallingPoints.CallingPointList != nil {
		subsequentCallingPoints := mapCallingPoints(
			ts.SubsequentCallingPoints.CallingPointList.CallingPoints,
			"subsequent calling point")
		trainService.SubsequentCallingPoints = subsequentCallingPoints
	}

	if ts.DelayReason != nil && ts.DelayReason.Text != "" {
		trainService.DelayReason = &ts.DelayReason.Text
	}

	return trainService
}

func mapLocation(location *nr.LocationLT4, locationType string) *Location {
	mappedLocation := &Location{Type: locationType}

	if location.LocationName != nil && location.LocationName.Text != "" {
		mappedLocation.Name = location.LocationName.Text
	}
	if location.CRS != nil && location.CRS.Text != "" {
		mappedLocation.CRS = location.CRS.Text
	}
	if location.Via != nil && location.Via.Text != "" {
		mappedLocation.Via = &location.Via.Text
	}

	return mappedLocation
}

func mapCallingPoints(cps []*nr.CallingPoint, locationType string) []*Location {
	var callingPoints []*Location
	for _, cp := range cps {
		if cp == nil {
			continue
		}

		callingPoint := &Location{Type: locationType}

		if cp.CRS != nil {
			callingPoint.CRS = cp.CRS.Text
		}

		if cp.LocationName != nil {
			callingPoint.Name = cp.LocationName.Text
		}

		if cp.St != nil {
			callingPoint.St = &cp.St.Text
		}

		if cp.At != nil {
			callingPoint.At = &cp.At.Text
		}

		if cp.ET != nil {
			callingPoint.Et = &cp.ET.Text
		}

		callingPoints = append(callingPoints, callingPoint)
	}

	return callingPoints
}
