package nationalrail

import "time"

type StationBoardRequest struct {
	CRS        CRSCode
	NumRows    int
	FilterCRS  *string
	FilterType *string
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
	LocationName      string         `json:"location_name"`
	Services          []*Service     `json:"services"`
	PlatformAvailable bool           `json:"platform_available"`
	NRCCMessages      []string       `json:"nrcc_messages"`
	Filters           RequestFilters `json:"filters"`
	GeneratedAt       time.Time      `json:"generated_at"`
}

type DeparturesBoard struct {
	CRS          CRSCode      `json:"crs"`
	Departures   []*Departure `json:"departures"`
	LocationName string       `json:"location_name"`
	NRCCMessages []string     `json:"nrcc_messages"`
	GeneratedAt  time.Time    `json:"generated_at"`
}

type Departure struct {
	CRS         CRSCode   `json:"crs"`
	Destination *Location `json:"location"`
	Service     *Service  `json:"services"`
}

type Service struct {
	ID                             string            `json:"id"`
	Type                           string            `json:"type"`
	Destination                    Location          `json:"destination"`
	Origins                        []*Location       `json:"origins"` // e.g. RAM & DVP
	ETA                            string            `json:"eta"`
	Formation                      Formation         `json:"formation"`
	Length                         int               `json:"length"`
	Operator                       Operator          `json:"operator"`
	Platform                       int               `json:"platform"`
	PreviousCallingPointsPerOrigin [][]*CallingPoint `json:"previous_calling_points"`
	RSID                           *string           `json:"rsid"`
	STA                            string            `json:"sta"`
	ETD                            *string           `json:"etd"`
	STD                            *string           `json:"std"`
	SubsequentCallingPoints        []*CallingPoint   `json:"subsequent_calling_points"`
	DelayReason                    *string           `json:"delay_reason"`
}

type ServiceDetails struct {
	Type                           string            `json:"type"`
	Formation                      Formation         `json:"formation"`
	Length                         int               `json:"length"`
	LocationName                   string            `json:"location_name"`
	Operator                       Operator          `json:"operator"`
	Platform                       int               `json:"platform"`
	PreviousCallingPointsPerOrigin [][]*CallingPoint `json:"previous_calling_points"`
	CRS                            CRSCode           `json:"crs"`
	ATA                            string            `json:"ata"`
	ATD                            string            `json:"atd"`
	STA                            string            `json:"sta"`
	ETD                            *string           `json:"etd"`
	STD                            string            `json:"std"`
	SubsequentCallingPoints        []*CallingPoint   `json:"subsequent_calling_points"`
	DelayReason                    *string           `json:"delay_reason"`
	GeneratedAt                    time.Time         `json:"generated_at"`
}

type Location struct {
	CRS  CRSCode `json:"crs"`
	Name string  `json:"name"`
	Via  *string `json:"via"`
}

type CallingPoint struct {
	CRS         CRSCode   `json:"crs"`
	Name        string    `json:"name"`
	At          *string   `json:"at"`
	Et          *string   `json:"et"`
	Formation   Formation `json:"formation"`
	Length      int       `json:"length"`
	St          string    `json:"st"`
	DelayReason *string   `json:"delay_reason"`
}

type RequestFilters struct {
	CRS          CRSCode `json:"crs"`
	LocationName string  `json:"location_name"`
	Type         string  `json:"type"`
}

type Toilet struct {
	Status   *string `json:"status"`
	CharData string  `json:"char_data"`
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

type Operator struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
