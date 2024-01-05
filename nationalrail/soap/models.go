package soap

type GetArrBoardWithDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetArrBoardWithDetailsResponse"`
	} `xml:"Body"`
}

type GetArrDepBoardWithDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetArrDepBoardWithDetailsResponse"`
	} `xml:"Body"`
}

type GetArrivalBoardResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetArrivalBoardResponse"`
	} `xml:"Body"`
}

type GetArrivalDepartureBoardResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetArrivalDepartureBoardResponse"`
	} `xml:"Body"`
}

type GetDepartureBoardResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetDepartureBoardResponse"`
	} `xml:"Body"`
}

type GetDepBoardWithDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		StationBoard *StationBoard `xml:"GetDepBoardWithDetailsResponse"`
	} `xml:"Body"`
}

type GetFastestDeparturesResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		DeparturesBoard *DeparturesBoard `xml:"GetFastestDeparturesResponse"`
	} `xml:"Body"`
}

type GetFastestDeparturesWithDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		DeparturesBoard *DeparturesBoard `xml:"GetFastestDeparturesWithDetailsResponse"`
	} `xml:"Body"`
}

type GetNextDeparturesResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		DeparturesBoard *DeparturesBoard `xml:"GetNextDeparturesResponse"`
	} `xml:"Body"`
}

type GetNextDeparturesWithDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		DeparturesBoard *DeparturesBoard `xml:"GetNextDeparturesWithDetailsResponse"`
	} `xml:"Body"`
}

type GetServiceDetailsResponse struct {
	Soap string `xml:"soap,attr"`
	Xsd  string `xml:"xsd,attr"`
	Xsi  string `xml:"xsi,attr"`
	Body struct {
		ServiceDetails *ServiceDetails `xml:"GetServiceDetailsResponse"`
	} `xml:"Body"`
}

type CallingPoint struct {
	At           *string   `xml:"at"`
	Crs          string    `xml:"crs"`
	Et           *string   `xml:"et"`
	Formation    Formation `xml:"formation"`
	Length       int       `xml:"length"`
	LocationName string    `xml:"locationName"`
	St           string    `xml:"st"`
	DelayReason  *string   `xml:"delayReason"`
}

type CallingPointList struct {
	CallingPoint []CallingPoint `xml:"callingPoint"`
}

type Coach struct {
	Number     string `xml:"number,attr"`
	CoachClass string `xml:"coachClass"`
	Loading    *int   `xml:"loading"`
	Toilet     Toilet `xml:"toilet"`
}

type Coaches struct {
	Coach []Coach `xml:"coach"`
}

type Destination struct {
	Location Location `xml:"location"`

	Crs     *string  `xml:"crs,attr"`
	Service *Service `xml:"service"`
}

type Formation struct {
	Coaches Coaches `xml:"coaches"`
}

type StationBoard struct {
	Xmlns     string                `xml:"xmlns,attr"`
	GetResult GetStationBoardResult `xml:"GetStationBoardResult"`
}

type DeparturesBoard struct {
	Xmlns     string                   `xml:"xmlns,attr"`
	GetResult GetDeparturesBoardResult `xml:"DeparturesBoard"`
}

type GetDeparturesBoardResult struct {
	Departures        Departures   `xml:"departures"`
	NrccMessages      NrccMessages `xml:"nrccMessages"`
	Lt                string       `xml:"lt,attr"`
	Lt2               string       `xml:"lt2,attr"`
	Lt3               string       `xml:"lt3,attr"`
	Lt4               string       `xml:"lt4,attr"`
	Lt5               string       `xml:"lt5,attr"`
	Lt6               string       `xml:"lt6,attr"`
	Lt7               string       `xml:"lt7,attr"`
	Lt8               string       `xml:"lt8,attr"`
	Crs               string       `xml:"crs"`
	GeneratedAt       string       `xml:"generatedAt"`
	LocationName      string       `xml:"locationName"`
	PlatformAvailable bool         `xml:"platformAvailable"`
}

type ServiceDetails struct {
	Xmlns     string                  `xml:"xmlns,attr"`
	GetResult GetServiceDetailsResult `xml:"GetServiceDetailsResult"`
}

type GetServiceDetailsResult struct {
	Formation               Formation               `xml:"formation"`
	PreviousCallingPoints   PreviousCallingPoints   `xml:"previousCallingPoints"`
	SubsequentCallingPoints SubsequentCallingPoints `xml:"subsequentCallingPoints"`
	Lt                      string                  `xml:"lt,attr"`
	Lt2                     string                  `xml:"lt2,attr"`
	Lt3                     string                  `xml:"lt3,attr"`
	Lt4                     string                  `xml:"lt4,attr"`
	Lt5                     string                  `xml:"lt5,attr"`
	Lt6                     string                  `xml:"lt6,attr"`
	Lt7                     string                  `xml:"lt7,attr"`
	Lt8                     string                  `xml:"lt8,attr"`
	Crs                     string                  `xml:"crs"`
	LocationName            string                  `xml:"locationName"`
	Operator                string                  `xml:"operator"`
	OperatorCode            string                  `xml:"operatorCode"`
	GeneratedAt             string                  `xml:"generatedAt"`
	ServiceType             string                  `xml:"serviceType"`
	CancelReason            *string                 `xml:"cancelReason"`
	DelayReason             *string                 `xml:"delayReason"`
	DiversionReason         *string                 `xml:"diversionReason"`
	DivertedVia             *string                 `xml:"divertedVia"`
	OverdueMessage          *string                 `xml:"overdueMessage"`
	Sta                     *string                 `xml:"sta"`
	Eta                     *string                 `xml:"eta"`
	Ata                     *string                 `xml:"ata"`
	Std                     *string                 `xml:"std"`
	Etd                     *string                 `xml:"etd"`
	Atd                     *string                 `xml:"atd"`
	Length                  int                     `xml:"length"`
	Platform                int                     `xml:"platform"`
	IsCancelled             bool                    `xml:"isCancelled"`
	DetachFront             bool                    `xml:"detachFront"`
	IsReverseFormation      bool                    `xml:"isReverseFormation"`
}

type Departures struct {
	Destination []DepartureDestination `xml:"destination"`
}

type DepartureDestination struct {
	Crs      *string   `xml:"crs,attr"`
	Location *Location `xml:"location"`
	Service  *Service  `xml:"service"`
}

type DestinationLocationWithService struct {
	Crs      *string   `xml:"crs,attr"`
	Location *Location `xml:"location"`
	Service  *Service  `xml:"service"`
}

type GetStationBoardResult struct {
	Lt                 string        `xml:"lt,attr"`
	Lt2                string        `xml:"lt2,attr"`
	Lt3                string        `xml:"lt3,attr"`
	Lt4                string        `xml:"lt4,attr"`
	Lt5                string        `xml:"lt5,attr"`
	Lt6                string        `xml:"lt6,attr"`
	Lt7                string        `xml:"lt7,attr"`
	Lt8                string        `xml:"lt8,attr"`
	Crs                string        `xml:"crs"`
	FilterLocationName string        `xml:"filterLocationName"`
	FilterType         string        `xml:"filterType"`
	Filtercrs          string        `xml:"filtercrs"`
	GeneratedAt        string        `xml:"generatedAt"`
	LocationName       string        `xml:"locationName"`
	NrccMessages       NrccMessages  `xml:"nrccMessages"`
	PlatformAvailable  bool          `xml:"platformAvailable"`
	TrainServices      TrainServices `xml:"trainServices"`
}

type Location struct {
	Crs          string  `xml:"crs"`
	LocationName string  `xml:"locationName"`
	Via          *string `xml:"via"`
}

type NrccMessages struct {
	Message []string `xml:"message"`
}

type Origin struct {
	Location []Location `xml:"location"`
}

type PreviousCallingPoints struct {
	CallingPointList []CallingPointList `xml:"callingPointList"`
}

type Service struct {
	Destination             Destination              `xml:"destination"`
	Formation               Formation                `xml:"formation"`
	Origin                  Origin                   `xml:"origin"`
	PreviousCallingPoints   *PreviousCallingPoints   `xml:"previousCallingPoints"`
	SubsequentCallingPoints *SubsequentCallingPoints `xml:"subsequentCallingPoints"`
	Operator                string                   `xml:"operator"`
	OperatorCode            string                   `xml:"operatorCode"`
	ServiceID               string                   `xml:"serviceID"`
	ServiceType             string                   `xml:"serviceType"`
	Rsid                    *string                  `xml:"rsid"`
	DelayReason             *string                  `xml:"delayReason"`
	Sta                     *string                  `xml:"sta"`
	Eta                     *string                  `xml:"eta"`
	Std                     *string                  `xml:"std"`
	Etd                     *string                  `xml:"etd"`
	Length                  int                      `xml:"length"`
	Platform                int                      `xml:"platform"`
}

type Toilet struct {
	Status   *string `xml:"status,attr"`
	CharData string  `xml:",chardata"`
}

type TrainServices struct {
	Service []Service `xml:"service"`
}

type SubsequentCallingPoints struct {
	CallingPointList CallingPointList `xml:"callingPointList"`
}
