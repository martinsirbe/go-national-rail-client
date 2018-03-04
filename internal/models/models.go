package internal_models

import "encoding/xml"

type DepartureResponse struct {
	AttrXmlnssoap string   `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string   `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string   `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *Body    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type Body struct {
	GetDepartureBoardResponse *GetDepartureBoardResponse `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/ GetDepartureBoardResponse,omitempty" json:"GetDepartureBoardResponse,omitempty"`
	XMLName                   xml.Name                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetDepartureBoardResponse struct {
	Attrxmlns             string                 `xml:"xmlns,attr"  json:",omitempty"`
	GetStationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName               xml.Name               `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/ GetDepartureBoardResponse,omitempty" json:"GetDepartureBoardResponse,omitempty"`
}

type GetStationBoardResult struct {
	AttrXmlnslt             string                   `xml:"xmlns lt,attr"  json:",omitempty"`
	AttrXmlnslt2            string                   `xml:"xmlns lt2,attr"  json:",omitempty"`
	AttrXmlnslt3            string                   `xml:"xmlns lt3,attr"  json:",omitempty"`
	AttrXmlnslt4            string                   `xml:"xmlns lt4,attr"  json:",omitempty"`
	AttrXmlnslt5            string                   `xml:"xmlns lt5,attr"  json:",omitempty"`
	AttrXmlnslt6            string                   `xml:"xmlns lt6,attr"  json:",omitempty"`
	AttrXmlnslt7            string                   `xml:"xmlns lt7,attr"  json:",omitempty"`
	OriginLocationCode      *OriginLocationCode      `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types crs,omitempty" json:"crs,omitempty"`
	FilterLocationName      *FilterLocationName      `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types filterLocationName,omitempty" json:"filterLocationName,omitempty"`
	DestinationLocationCode *DestinationLocationCode `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types filtercrs,omitempty" json:"filtercrs,omitempty"`
	WarningMessages         *WarningMessages         `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types nrccMessages,omitempty" json:"nrccMessages,omitempty"`
	GeneratedAt             *GeneratedAt             `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
	LocationName            *LocationName            `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	PlatformAvailable       *PlatformAvailable       `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types platformAvailable,omitempty" json:"platformAvailable,omitempty"`
	TrainServiceDetails     *TrainServiceDetails     `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types trainServices,omitempty" json:"trainServices,omitempty"`
	XMLName                 xml.Name                 `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
}

type OriginLocationCode struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types crs,omitempty" json:"crs,omitempty"`
}

type Destination struct {
	Location *Location `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types location,omitempty" json:"location,omitempty"`
	XMLName  xml.Name  `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types destination,omitempty" json:"destination,omitempty"`
}

type ServiceStatus struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types etd,omitempty" json:"etd,omitempty"`
}

type FilterLocationName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types filterLocationName,omitempty" json:"filterLocationName,omitempty"`
}

type DestinationLocationCode struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types filtercrs,omitempty" json:"filtercrs,omitempty"`
}

type WarningMessages struct {
	Messages []*WarningMessage `xml:"http://thalesgroup.com/RTTI/2012-01-13/ldb/types message,omitempty" json:"message,omitempty"`
	XMLName  xml.Name          `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types nrccMessages,omitempty" json:"nrccMessages,omitempty"`
}

type WarningMessage struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2012-01-13/ldb/types message,omitempty" json:"message,omitempty"`
}

type GeneratedAt struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
}

type Location struct {
	CRS          *OriginLocationCode `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types crs,omitempty" json:"crs,omitempty"`
	LocationName *LocationName       `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	Via          *Via                `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types via,omitempty" json:"via,omitempty"`
	XMLName      xml.Name            `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types location,omitempty" json:"location,omitempty"`
}

type LocationName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types locationName,omitempty" json:"locationName,omitempty"`
}

type Operator struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types operator,omitempty" json:"operator,omitempty"`
}

type OperatorCode struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
}

type Origin struct {
	Location *Location `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types location,omitempty" json:"location,omitempty"`
	XMLName  xml.Name  `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types origin,omitempty" json:"origin,omitempty"`
}

type Platform struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types platform,omitempty" json:"platform,omitempty"`
}

type PlatformAvailable struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types platformAvailable,omitempty" json:"platformAvailable,omitempty"`
}

type TrainService struct {
	Destination   *Destination   `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types destination,omitempty" json:"destination,omitempty"`
	ServiceStatus *ServiceStatus `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types etd,omitempty" json:"etd,omitempty"`
	OperatorCode  *OperatorCode  `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
	Operator      *Operator      `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types operator,omitempty" json:"operator,omitempty"`
	Origin        *Origin        `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types origin,omitempty" json:"origin,omitempty"`
	Platform      *Platform      `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types platform,omitempty" json:"platform,omitempty"`
	ServiceID     *ServiceID     `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types serviceID,omitempty" json:"serviceID,omitempty"`
	DepartureTime *DepartureTime `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types std,omitempty" json:"std,omitempty"`
	XMLName       xml.Name       `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types service,omitempty" json:"service,omitempty"`
}

type ServiceID struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types serviceID,omitempty" json:"serviceID,omitempty"`
}

type DepartureTime struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types std,omitempty" json:"std,omitempty"`
}

type TrainServiceDetails struct {
	TrainServices []*TrainService `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types service,omitempty" json:"service,omitempty"`
	XMLName       xml.Name        `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types trainServices,omitempty" json:"trainServices,omitempty"`
}

type Via struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2014-02-20/ldb/types via,omitempty" json:"via,omitempty"`
}