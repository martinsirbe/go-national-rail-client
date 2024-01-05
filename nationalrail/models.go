package nationalrail

import "time"

type StationBoardRequest struct {
	CRS        CRSCode
	FilterCRS  *string
	FilterType *string
	NumRows    int
	TimeOffset *int
	TimeWindow *int
	FilterList []string
}

type DeparturesBoardRequest struct {
	CRS        CRSCode
	NumRows    int
	TimeOffset *int
	TimeWindow *int
}

type StationBoard struct {
	CRS               CRSCode        `json:"crs"`
	GeneratedAt       time.Time      `json:"generated_at"`
	Filters           RequestFilters `json:"filters"`
	Services          []*Service     `json:"services"`
	LocationName      string         `json:"location_name"`
	PlatformAvailable bool           `json:"platform_available"`
	NRCCMessages      []string       `json:"nrcc_messages"`
}

type DeparturesBoard struct {
	CRS          CRSCode      `json:"crs"`
	GeneratedAt  time.Time    `json:"generated_at"`
	Departures   []*Departure `json:"departures"`
	LocationName string       `json:"location_name"`
	NRCCMessages []string     `json:"nrcc_messages"`
}

type Departure struct {
	CRS         CRSCode   `json:"crs"`
	Destination *Location `json:"location"`
	Service     *Service  `json:"services"`
}

type Service struct {
	Destination             Location        `json:"destination"`
	Formation               Formation       `json:"formation"`
	Operator                Operator        `json:"operator"`
	Origins                 []*Location     `json:"origins"` // e.g. RAM & DVP
	SubsequentCallingPoints []*CallingPoint `json:"subsequent_calling_points"`
	// PreviousCallingPointsPerOrigin representing the previous calling points in the journey.
	// A separate calling point list will be present for each origin of the service, relative
	// to the current location.
	PreviousCallingPointsPerOrigin [][]*CallingPoint `json:"previous_calling_points"`
	ID                             string            `json:"id"`
	Type                           string            `json:"type"`
	RetailServiceID                *string           `json:"rsid"`
	// ScheduledTimeOfArrival an optional Scheduled Time of Arrival of the service
	// at the station board location.
	ScheduledTimeOfArrival *string `json:"scheduled_time_of_arrival"`
	// EstimatedTimeOfArrival an optional Estimated Time of Arrival of the service
	// at the station board location.
	EstimatedTimeOfArrival *string `json:"estimated_time_of_arrival"`
	// ScheduledTimeOfDeparture an optional Scheduled Time of Departure of the service
	// at the station board location.
	ScheduledTimeOfDeparture *string `json:"scheduled_time_of_departure"`
	// EstimatedTimeOfDeparture an optional Estimated Time of Departure of the service
	// at the station board location
	EstimatedTimeOfDeparture *string `json:"estimated_time_of_departure"`
	DelayReason              *string `json:"delay_reason"`
	Length                   int     `json:"length"`
	Platform                 int     `json:"platform"`
}

type ServiceDetails struct {
	CRS                            CRSCode           `json:"crs"`
	Formation                      Formation         `json:"formation"`
	Operator                       Operator          `json:"operator"`
	GeneratedAt                    time.Time         `json:"generated_at"`
	DivertedVia                    *CRSCode          `json:"diverted_via"`
	SubsequentCallingPoints        []*CallingPoint   `json:"subsequent_calling_points"`
	PreviousCallingPointsPerOrigin [][]*CallingPoint `json:"previous_calling_points"`
	Type                           string            `json:"type"`
	LocationName                   string            `json:"location_name"`
	CancelReason                   *string           `json:"cancel_reason"`
	DelayReason                    *string           `json:"delay_reason"`
	DiversionReason                *string           `json:"diversion_reason"`
	OverdueMessage                 *string           `json:"overdue_message"`
	// ScheduledTimeOfArrival of service at location. If not present then the location is
	// the origin of service, or it does not set down passengers at this location.
	ScheduledTimeOfArrival *string `json:"scheduled_time_of_arrival"`
	// EstimatedTimeOfArrival will only be present if ScheduledTimeOfArrival is also present
	// and ActualTimeOfArrival is not present.
	EstimatedTimeOfArrival *string `json:"estimated_time_of_arrival"`
	// ActualTimeOfArrival will only be present if ScheduledTimeOfArrival is also present and
	//EstimatedTimeOfArrival is not present.
	ActualTimeOfArrival *string `json:"actual_time_of_arrival"`
	// ScheduledTimeOfDeparture of service at location. If not present then this is
	// the destination of service, or it does not pick up passengers at this location.
	ScheduledTimeOfDeparture *string `json:"scheduled_time_of_departure"`
	// EstimatedTimeOfDeparture will only be present if ScheduledTimeOfDeparture is also present
	// and atd is not present.
	EstimatedTimeOfDeparture *string `json:"estimated_time_of_departure"`
	// ActualTimeOfDeparture will only be present if ScheduledTimeOfDeparture is also present
	// and EstimatedTimeOfDeparture is not present.
	ActualTimeOfDeparture *string `json:"actual_time_of_departure"`
	Length                int     `json:"length"`
	Platform              int     `json:"platform"`
	IsCancelled           bool    `json:"is_cancelled"`
	DetachFront           bool    `json:"detach_front"`
	IsReverseFormation    bool    `json:"is_reverse_formation"`
}

type Location struct {
	CRS  CRSCode `json:"crs"`
	Name string  `json:"name"`
	Via  *string `json:"via"`
}

type CallingPoint struct {
	CRS       CRSCode   `json:"crs"`
	Formation Formation `json:"formation"`
	Name      string    `json:"name"`
	// ScheduledTime of the service at the location.
	// The time will be either an arrival or departure time, depending on
	// whether it is in the subsequent or previous calling point list.
	ScheduledTime string `json:"scheduled_time"`
	// DelayReason the delay reason for this service at location.
	DelayReason *string `json:"delay_reason"`
	// EstimatedTime of the service at the location.
	// The time will be either an arrival or departure time, depending on
	// whether it is in the subsequent or previous calling point list.
	// Will only be present if an ActualTime is not present.
	EstimatedTime *string `json:"estimated_time"`
	// ActualTime of the service at the location.
	// The time will be either an arrival or departure time, depending on
	// whether it is in the subsequent or previous calling point list.
	// Will only be present if an EstimatedTime is not present.
	ActualTime *string `json:"actual_time"`
	// Length of train (number of units) at the location. If not supplied,
	// or zero, the length is unknown.
	Length int `json:"length"`
}

type RequestFilters struct {
	CRS          CRSCode `json:"crs"`
	LocationName string  `json:"location_name"`
	Type         string  `json:"type"`
}

type ToiletType string

const (
	ToiletTypeUnknown    ToiletType = "unknown"
	ToiletTypeNone       ToiletType = "none"
	ToiletTypeStandard   ToiletType = "standard"
	ToiletTypeAccessible ToiletType = "accessible"
)

type ToiletServiceStatus string

const (
	ToiletServiceStatusUnknown      ToiletServiceStatus = "unknown"
	ToiletServiceStatusInService    ToiletServiceStatus = "in service"
	ToiletServiceStatusNotInService ToiletServiceStatus = "not in service"
)

// Toilet indicates the toilet availability on a coach.
type Toilet struct {
	Status ToiletServiceStatus `json:"status"`
	Type   ToiletType          `json:"type"`
}

type Coach struct {
	Number  string `json:"number"`
	Class   string `json:"class"`
	Loading *int   `json:"loading"`
	Toilet  Toilet `json:"toilet"`
}

type Formation struct {
	Coaches []*Coach `json:"coaches"`
}

// Operator details of the Train Operating Company that operates the service.
type Operator struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
