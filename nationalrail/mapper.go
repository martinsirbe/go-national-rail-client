package nationalrail

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/martinsirbe/go-national-rail-client/nationalrail/soap"
)

func MapGetArrBoardWithDetailsResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetArrBoardWithDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetArrDepBoardWithDetailsResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetArrDepBoardWithDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetArrivalBoardResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetArrivalBoardResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetArrivalDepartureBoardResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetArrivalDepartureBoardResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetDepartureBoardResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetDepartureBoardResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetDepBoardWithDetailsResponse(resp []byte) (*StationBoard, error) {
	var env soap.GetDepBoardWithDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.StationBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapStationBoard(env.Body.StationBoard.GetResult)
}

func MapGetFastestDeparturesResponse(resp []byte) (*DeparturesBoard, error) {
	var env soap.GetFastestDeparturesResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.DeparturesBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapDeparturesBoard(env.Body.DeparturesBoard.GetResult)
}

func MapGetFastestDeparturesWithDetailsResponse(resp []byte) (*DeparturesBoard, error) {
	var env soap.GetFastestDeparturesWithDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.DeparturesBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapDeparturesBoard(env.Body.DeparturesBoard.GetResult)
}

func MapGetNextDeparturesResponse(resp []byte) (*DeparturesBoard, error) {
	var env soap.GetNextDeparturesResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.DeparturesBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapDeparturesBoard(env.Body.DeparturesBoard.GetResult)
}

func MapGetNextDeparturesWithDetailsResponse(resp []byte) (*DeparturesBoard, error) {
	var env soap.GetNextDeparturesWithDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.DeparturesBoard == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapDeparturesBoard(env.Body.DeparturesBoard.GetResult)
}

func MapGetServiceDetailsResponse(resp []byte) (*ServiceDetails, error) {
	var env soap.GetServiceDetailsResponse
	if err := xml.Unmarshal(resp, &env); err != nil {
		return nil, fmt.Errorf("failed to unmarshal envelope: %w", err)
	}

	if env.Body.ServiceDetails == nil {
		return nil, ErrWebServiceResponseMissing
	}

	return mapServiceDetails(&env.Body.ServiceDetails.GetResult)
}

func mapStationBoard(sbr soap.GetStationBoardResult) (*StationBoard, error) {
	stationBoard := StationBoard{
		CRS:               CRSCode(sbr.Crs),
		LocationName:      sbr.LocationName,
		PlatformAvailable: sbr.PlatformAvailable,
		Filters: RequestFilters{
			CRS:          CRSCode(sbr.Filtercrs),
			LocationName: sbr.FilterLocationName,
			Type:         sbr.FilterType,
		},
	}

	if sbr.GeneratedAt != "" {
		t, err := time.Parse(time.RFC3339Nano, sbr.GeneratedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse generated at: %w", err)
		}

		stationBoard.GeneratedAt = t
	}

	if len(sbr.NrccMessages.Message) > 0 {
		for _, msg := range sbr.NrccMessages.Message {
			stationBoard.NRCCMessages = append(stationBoard.NRCCMessages, cleanNRCCMessage(msg))
		}
	}

	for _, service := range sbr.TrainServices.Service {
		stationBoard.Services = append(stationBoard.Services, mapService(&service))
	}

	return &stationBoard, nil
}

func mapService(s *soap.Service) *Service {
	service := Service{
		ID:   s.ServiceID,
		Type: s.ServiceType,
		Destination: Location{
			CRS:  CRSCode(s.Destination.Location.Crs),
			Name: s.Destination.Location.LocationName,
			Via:  s.Destination.Location.Via,
		},
		ETA:    s.Eta,
		Length: s.Length,
		Operator: Operator{
			Code: s.OperatorCode,
			Name: s.Operator,
		},
		Platform: s.Platform,
		STA:      s.Sta,
		RSID:     s.Rsid,
		ETD:      s.Etd,
		STD:      s.Std,
		Formation: Formation{
			Coaches: mapCoaches(s.Formation.Coaches.Coach),
		},
		DelayReason: s.DelayReason,
	}

	var origins []*Location
	for _, origin := range s.Origin.Location {
		origins = append(origins, &Location{
			CRS:  CRSCode(origin.Crs),
			Name: origin.LocationName,
			Via:  origin.Via,
		})
	}
	service.Origins = origins

	if s.PreviousCallingPoints != nil {
		previousCallingPoints := make([][]*CallingPoint, len(s.PreviousCallingPoints.CallingPointList))
		for i, cpl := range s.PreviousCallingPoints.CallingPointList {
			var callingPoints []*CallingPoint
			for _, cp := range cpl.CallingPoint {
				callingPoints = append(callingPoints, &CallingPoint{
					CRS:  CRSCode(cp.Crs),
					Name: cp.LocationName,
					At:   cp.At,
					Et:   cp.Et,
					Formation: Formation{
						Coaches: mapCoaches(cp.Formation.Coaches.Coach),
					},
					Length:      cp.Length,
					St:          cp.St,
					DelayReason: cp.DelayReason,
				})
			}

			previousCallingPoints[i] = callingPoints
		}
		service.PreviousCallingPointsPerOrigin = previousCallingPoints
	}

	if s.SubsequentCallingPoints != nil {
		var callingPoints []*CallingPoint
		for _, cp := range s.SubsequentCallingPoints.CallingPointList.CallingPoint {
			callingPoints = append(callingPoints, &CallingPoint{
				CRS:  CRSCode(cp.Crs),
				Name: cp.LocationName,
				At:   cp.At,
				Et:   cp.Et,
				Formation: Formation{
					Coaches: mapCoaches(cp.Formation.Coaches.Coach),
				},
				Length:      cp.Length,
				St:          cp.St,
				DelayReason: cp.DelayReason,
			})
		}
		service.SubsequentCallingPoints = callingPoints
	}

	return &service
}

func mapCoaches(c []soap.Coach) []*Coach {
	var coaches []*Coach
	for _, coach := range c {
		coaches = append(coaches, &Coach{
			Number:  coach.Number,
			Class:   coach.CoachClass,
			Loading: coach.Loading,
			Toilet: Toilet{
				Status:   coach.Toilet.Status,
				CharData: coach.Toilet.CharData,
			},
		})
	}

	return coaches
}

func mapDeparturesBoard(sbr soap.GetDeparturesBoardResult) (*DeparturesBoard, error) {
	departuresBoard := DeparturesBoard{
		CRS:          CRSCode(sbr.Crs),
		LocationName: sbr.LocationName,
	}

	var departures []*Departure
	for _, d := range sbr.Departures.Destination {
		var departure Departure
		if d.Crs != nil {
			departure.CRS = CRSCode(*d.Crs)
		}

		if d.Location != nil {
			departure.Destination = &Location{
				CRS:  CRSCode(d.Location.Crs),
				Name: d.Location.LocationName,
				Via:  d.Location.Via,
			}
		}

		if d.Service != nil {
			departure.Service = mapService(d.Service)
		}

		departures = append(departures, &departure)
	}
	departuresBoard.Departures = departures

	if sbr.GeneratedAt != "" {
		t, err := time.Parse(time.RFC3339Nano, sbr.GeneratedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse generated at: %w", err)
		}

		departuresBoard.GeneratedAt = t
	}

	if len(sbr.NrccMessages.Message) > 0 {
		for _, msg := range sbr.NrccMessages.Message {
			departuresBoard.NRCCMessages = append(departuresBoard.NRCCMessages, cleanNRCCMessage(msg))
		}
	}

	return &departuresBoard, nil
}

func mapServiceDetails(s *soap.GetServiceDetailsResult) (*ServiceDetails, error) {
	serviceDetails := ServiceDetails{
		ATA: s.Ata,
		ATD: s.Atd,
		CRS: CRSCode(s.Crs),
		Formation: Formation{
			Coaches: mapCoaches(s.Formation.Coaches.Coach),
		},
		Length:       s.Length,
		LocationName: s.LocationName,
		Operator: Operator{
			Code: s.OperatorCode,
			Name: s.Operator,
		},
		Platform: s.Platform,
		Type:     s.ServiceType,
		STA:      s.Sta,
		STD:      s.Std,
	}

	if s.GeneratedAt != "" {
		t, err := time.Parse(time.RFC3339Nano, s.GeneratedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to parse generated at: %w", err)
		}

		serviceDetails.GeneratedAt = t
	}

	previousCallingPoints := make([][]*CallingPoint, len(s.PreviousCallingPoints.CallingPointList))
	for i, cpl := range s.PreviousCallingPoints.CallingPointList {
		var callingPoints []*CallingPoint
		for _, cp := range cpl.CallingPoint {
			callingPoints = append(callingPoints, &CallingPoint{
				CRS:  CRSCode(cp.Crs),
				Name: cp.LocationName,
				At:   cp.At,
				Et:   cp.Et,
				Formation: Formation{
					Coaches: mapCoaches(cp.Formation.Coaches.Coach),
				},
				Length:      cp.Length,
				St:          cp.St,
				DelayReason: cp.DelayReason,
			})
		}

		previousCallingPoints[i] = callingPoints
	}
	serviceDetails.PreviousCallingPointsPerOrigin = previousCallingPoints

	var callingPoints []*CallingPoint
	for _, cp := range s.SubsequentCallingPoints.CallingPointList.CallingPoint {
		callingPoints = append(callingPoints, &CallingPoint{
			CRS:  CRSCode(cp.Crs),
			Name: cp.LocationName,
			At:   cp.At,
			Et:   cp.Et,
			Formation: Formation{
				Coaches: mapCoaches(cp.Formation.Coaches.Coach),
			},
			Length:      cp.Length,
			St:          cp.St,
			DelayReason: cp.DelayReason,
		})
	}
	serviceDetails.SubsequentCallingPoints = callingPoints

	return &serviceDetails, nil
}
