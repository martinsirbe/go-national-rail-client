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
	FilterList []string
	TimeOffset *int
	TimeWindow *int
}

type StationBoard struct {
	GeneratedAt        time.Time       `json:"generated_at"`
	LocationName       string          `json:"location_name"`
	CRS                string          `json:"crs"`
	FilterLocationName *string         `json:"filter_location_name"`
	FilterCRS          *string         `json:"filter_crs"`
	PlatformAvailable  bool            `json:"platform_available"`
	TrainServices      []*TrainService `json:"train_services"`
	NRCCMessages       []string        `json:"nrcc_messages"`
}

type TrainService struct {
	STA                     *string     `json:"sta"`
	ETA                     *string     `json:"eta"`
	STD                     string      `json:"std"`
	ETD                     string      `json:"etd"`
	Platform                *string     `json:"platform"`
	Operator                string      `json:"operator"`
	OperatorCode            string      `json:"operator_code"`
	ServiceType             string      `json:"service_type"`
	ServiceID               string      `json:"service_id"`
	RSID                    string      `json:"rsid"`
	Origin                  *Location   `json:"origin"`
	Destination             *Location   `json:"destination"`
	DelayReason             *string     `json:"delay_reason"`
	PreviousCallingPoints   []*Location `json:"previous_calling_points"`
	SubsequentCallingPoints []*Location `json:"subsequent_calling_points"`
	GeneratedAt             *time.Time  `json:"generated_at"`
}

type TrainServiceDetails struct {
	GeneratedAt             time.Time   `json:"generated_at"`
	LocationName            string      `json:"location_name"`
	ServiceType             string      `json:"service_type"`
	CRS                     string      `json:"crs"`
	Operator                string      `json:"operator"`
	OperatorCode            string      `json:"operator_code"`
	RSID                    string      `json:"rsid"`
	Platform                *string     `json:"platform"`
	STA                     *string     `json:"sta"`
	ETA                     *string     `json:"eta"`
	STD                     *string     `json:"std"`
	ETD                     *string     `json:"etd"`
	PreviousCallingPoints   []*Location `json:"previous_calling_points"`
	SubsequentCallingPoints []*Location `json:"subsequent_calling_points"`
}

type Location struct {
	Name string  `json:"name"`
	CRS  string  `json:"crs"`
	Type string  `json:"type"`
	Via  *string `json:"via"`
	St   *string `json:"st"`
	At   *string `json:"at"`
	Et   *string `json:"et"`
}
