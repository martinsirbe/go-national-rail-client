package mapper

import (
	"time"

	"github.com/MartinsIrbe/national-rail-go-client/pkg/models"
	im "github.com/MartinsIrbe/national-rail-go-client/internal/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// NationalRailResponseMapper a mapper used to map national rail response to a simpler model
type NationalRailResponseMapper struct{}

// NewNationalRailResponseMapper used to create a new instance of national rail response mapper
func NewNationalRailResponseMapper() *NationalRailResponseMapper {
	return &NationalRailResponseMapper{}
}

// MapDepartureResponse maps national rail departure response to a simple model
func (m *NationalRailResponseMapper) MapDepartureResponse(r *im.DepartureResponse) (*models.DepartureInfo, error) {
	switch {
	case r == nil:
		return nil, errors.New("failed to map: expected to have departure response but was missing")
	case r.Body == nil:
		return nil, errors.New("failed to map: expected to have response body but was missing")
	case r.Body.GetDepartureBoardResponse == nil:
		return nil, errors.New("failed to map: expected to have departure board response but was missing")
	case r.Body.GetDepartureBoardResponse.GetStationBoardResult == nil:
		return nil, errors.New("failed to map: expected to have station board result but was missing")
	case r.Body.GetDepartureBoardResponse.GetStationBoardResult.LocationName == nil:
		return nil, errors.New("failed to map: expected to have location name but was missing")
	case r.Body.GetDepartureBoardResponse.GetStationBoardResult.OriginLocationCode == nil:
		return nil, errors.New("failed to map: expected to have current location code but was missing")
	case r.Body.GetDepartureBoardResponse.GetStationBoardResult.DestinationLocationCode == nil:
		return nil, errors.New("failed to map: expected to have destination location code but was missing")
	case r.Body.GetDepartureBoardResponse.GetStationBoardResult.TrainServiceDetails == nil:
		return nil, errors.New("failed to map: expected to have train services details but was missing")
	}

	sbr := r.Body.GetDepartureBoardResponse.GetStationBoardResult

	return &models.DepartureInfo{
		OriginLocationCode:      sbr.OriginLocationCode.Text,
		DestinationLocationCode: sbr.DestinationLocationCode.Text,
		WarningMessages:         getWarningMessages(sbr),
		DepartureDetails:        getDepartureDetails(sbr),
	}, nil
}

// getWarningMessages used to extract warning messages from the station board result
func getWarningMessages(sbr *im.GetStationBoardResult) []string {
	var warningMessages []string
	if sbr.WarningMessages != nil {
		for _, wm := range sbr.WarningMessages.Messages {
			if wm != nil {
				warningMessages = append(warningMessages, wm.Text)
			}
		}
	}
	return warningMessages
}

// getDepartureDetails used to extract departure details from the station board result
func getDepartureDetails(sbr *im.GetStationBoardResult) []models.DepartureDetails {
	var departureDetails []models.DepartureDetails
	for _, ts := range sbr.TrainServiceDetails.TrainServices {
		dpd := models.DepartureDetails{}

		if ts.Platform != nil {
			dpd.Platform = ts.Platform.Text
		}

		if ts.Origin != nil && ts.Origin.Location != nil {
			loc := ts.Origin.Location

			if loc.LocationName != nil {
				dpd.OriginLocation = loc.LocationName.Text
			}
		}

		if ts.Destination != nil && ts.Destination.Location != nil {
			loc := ts.Destination.Location

			if loc.LocationName != nil {
				dpd.Destination = loc.LocationName.Text
			}

			if loc.Via != nil {
				dpd.Via = &loc.Via.Text
			}
		}

		if ts.DepartureTime != nil {
			if dt, err := time.Parse("15:04", ts.DepartureTime.Text); err != nil {
				logrus.WithError(err).Warn("failed to parse departure time string of %s", ts.DepartureTime.Text)
			} else {
				dpd.Time = &dt
			}
		}

		if ts.ServiceStatus != nil {
			dpd.Status = ts.ServiceStatus.Text
		}

		if ts.Operator != nil {
			dpd.Operator = ts.Operator.Text
		}

		if ts.Platform != nil {
			dpd.Platform = ts.Platform.Text
		} else {
			dpd.Platform = "not available"
		}

		departureDetails = append(departureDetails, dpd)
	}
	return departureDetails
}
