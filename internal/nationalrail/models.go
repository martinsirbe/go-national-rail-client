package nationalrail

import "encoding/xml"

type GetArrBoardWithDetailsResponse struct {
	AttrXmlnssoap string                              `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                              `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                              `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetArrBoardWithDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetArrBoardWithDetailsResponseBody struct {
	ArrBoardWithDetails *ArrBoardWithDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrBoardWithDetailsResponse,omitempty" json:"GetArrBoardWithDetailsResponse,omitempty"`
	XMLName             xml.Name             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetArrivalBoardResponse struct {
	AttrXmlnssoap string                       `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                       `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                       `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetArrivalBoardResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetArrivalBoardResponseBody struct {
	ArrivalBoard *ArrivalBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrivalBoardResponse,omitempty" json:"GetArrivalBoardResponse,omitempty"`
	XMLName      xml.Name      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetArrivalDepartureBoardResponse struct {
	AttrXmlnssoap string                                `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                                `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                                `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetArrivalDepartureBoardResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetArrivalDepartureBoardResponseBody struct {
	ArrivalDepartureBoard *ArrivalDepartureBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrivalDepartureBoardResponse,omitempty" json:"GetArrivalDepartureBoardResponse,omitempty"`
	XMLName               xml.Name               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetDepartureBoardResponse struct {
	AttrXmlnssoap string                         `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                         `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                         `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetDepartureBoardResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetDepartureBoardResponseBody struct {
	DepartureBoard *DepartureBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetDepartureBoardResponse,omitempty" json:"GetDepartureBoardResponse,omitempty"`
	XMLName        xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetDepBoardWithDetailsResponse struct {
	AttrXmlnssoap string                              `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                              `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                              `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetDepBoardWithDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                            `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetDepBoardWithDetailsResponseBody struct {
	DepBoardWithDetails *DepBoardWithDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetDepBoardWithDetailsResponse,omitempty" json:"GetDepBoardWithDetailsResponse,omitempty"`
	XMLName             xml.Name             `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetFastestDeparturesResponse struct {
	AttrXmlnssoap string                            `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                            `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                            `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetFastestDeparturesResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                          `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetFastestDeparturesResponseBody struct {
	FastestDepartures *FastestDepartures `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetFastestDeparturesResponse,omitempty" json:"GetFastestDeparturesResponse,omitempty"`
	XMLName           xml.Name           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetFastestDeparturesWithDetailsResponse struct {
	AttrXmlnssoap string                                       `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                                       `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                                       `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetFastestDeparturesWithDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                                     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetFastestDeparturesWithDetailsResponseBody struct {
	FastestDeparturesWithDetails *FastestDeparturesWithDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetFastestDeparturesWithDetailsResponse,omitempty" json:"GetFastestDeparturesWithDetailsResponse,omitempty"`
	XMLName                      xml.Name                      `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetNextDeparturesResponse struct {
	AttrXmlnssoap string                         `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                         `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                         `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetNextDeparturesResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetNextDeparturesResponseBody struct {
	NextDepartures *NextDepartures `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetNextDeparturesResponse,omitempty" json:"GetNextDeparturesResponse,omitempty"`
	XMLName        xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetNextDeparturesWithDetailsResponse struct {
	AttrXmlnssoap string                                    `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                                    `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                                    `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetNextDeparturesWithDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                                  `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetNextDeparturesWithDetailsResponseBody struct {
	NextDeparturesWithDetails *NextDeparturesWithDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetNextDeparturesWithDetailsResponse,omitempty" json:"GetNextDeparturesWithDetailsResponse,omitempty"`
	XMLName                   xml.Name                   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetServiceDetailsResponse struct {
	AttrXmlnssoap string                         `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                         `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                         `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetServiceDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetServiceDetailsResponseBody struct {
	ServiceDetails *ServiceDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetServiceDetailsResponse,omitempty" json:"GetServiceDetailsResponse,omitempty"`
	XMLName        xml.Name        `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type GetArrDepBoardWithDetailsResponse struct {
	AttrXmlnssoap string                                 `xml:"xmlns soap,attr"  json:",omitempty"`
	AttrXmlnsxsd  string                                 `xml:"xmlns xsd,attr"  json:",omitempty"`
	AttrXmlnsxsi  string                                 `xml:"xmlns xsi,attr"  json:",omitempty"`
	Body          *GetArrDepBoardWithDetailsResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName       xml.Name                               `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

type GetArrDepBoardWithDetailsResponseBody struct {
	ArrDepBoardWithDetails *ArrDepBoardWithDetails `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrDepBoardWithDetailsResponse,omitempty" json:"GetArrDepBoardWithDetailsResponse,omitempty"`
	XMLName                xml.Name                `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

type ArrDepBoardWithDetails struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrDepBoardWithDetailsResponse,omitempty" json:"GetArrDepBoardWithDetailsResponse,omitempty"`
}

type ServiceDetails struct {
	Attrxmlns            string                `xml:"xmlns,attr"  json:",omitempty"`
	ServiceDetailsResult *ServiceDetailsResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetServiceDetailsResult,omitempty" json:"GetServiceDetailsResult,omitempty"`
	XMLName              xml.Name              `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetServiceDetailsResponse,omitempty" json:"GetServiceDetailsResponse,omitempty"`
}

type ServiceDetailsResult struct {
	AttrXmlnslt             string                   `xml:"xmlns lt,attr"  json:",omitempty"`
	AttrXmlnslt2            string                   `xml:"xmlns lt2,attr"  json:",omitempty"`
	AttrXmlnslt3            string                   `xml:"xmlns lt3,attr"  json:",omitempty"`
	AttrXmlnslt4            string                   `xml:"xmlns lt4,attr"  json:",omitempty"`
	AttrXmlnslt5            string                   `xml:"xmlns lt5,attr"  json:",omitempty"`
	AttrXmlnslt6            string                   `xml:"xmlns lt6,attr"  json:",omitempty"`
	AttrXmlnslt7            string                   `xml:"xmlns lt7,attr"  json:",omitempty"`
	ETD                     *ETDLT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types etd,omitempty" json:"etd,omitempty"`
	ETA                     *ETALT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types eta,omitempty" json:"eta,omitempty"`
	ATA                     *ATA                     `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types ata,omitempty" json:"ata,omitempty"`
	ATD                     *ATD                     `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types atd,omitempty" json:"atd,omitempty"`
	CRS                     *CRS                     `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types crs,omitempty" json:"crs,omitempty"`
	STA                     *STA                     `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types sta,omitempty" json:"sta,omitempty"`
	STD                     *STD                     `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types std,omitempty" json:"std,omitempty"`
	GeneratedAt             *GeneratedAt             `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
	LocationName            *LocationName            `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	OperatorCode            *OperatorCode            `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
	Operator                *Operator                `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types operator,omitempty" json:"operator,omitempty"`
	Platform                *Platform                `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types platform,omitempty" json:"platform,omitempty"`
	RSID                    *RSID                    `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types rsid,omitempty" json:"rsid,omitempty"`
	ServiceType             *ServiceType             `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types serviceType,omitempty" json:"serviceType,omitempty"`
	PreviousCallingPoints   *PreviousCallingPoints   `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types previousCallingPoints,omitempty" json:"previousCallingPoints,omitempty"`
	SubsequentCallingPoints *SubsequentCallingPoints `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types subsequentCallingPoints,omitempty" json:"subsequentCallingPoints,omitempty"`
	XMLName                 xml.Name                 `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetServiceDetailsResult,omitempty" json:"GetServiceDetailsResult,omitempty"`
}

type At struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types at,omitempty" json:"at,omitempty"`
}

type LocationName struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types locationName,omitempty" json:"locationName,omitempty"`
}

type ATA struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types ata,omitempty" json:"ata,omitempty"`
}

type ATD struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types atd,omitempty" json:"atd,omitempty"`
}

type CallingPoint struct {
	At           *At           `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types at,omitempty" json:"at,omitempty"`
	CRS          *CRS          `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types crs,omitempty" json:"crs,omitempty"`
	ET           *ET           `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types et,omitempty" json:"et,omitempty"`
	LocationName *LocationName `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	St           *St           `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types st,omitempty" json:"st,omitempty"`
	XMLName      xml.Name      `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types callingPoint,omitempty" json:"callingPoint,omitempty"`
}

type CallingPointList struct {
	CallingPoints []*CallingPoint `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types callingPoint,omitempty" json:"callingPoint,omitempty"`
	XMLName       xml.Name        `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types callingPointList,omitempty" json:"callingPointList,omitempty"`
}

type CRS struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types crs,omitempty" json:"crs,omitempty"`
}

type ET struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types et,omitempty" json:"et,omitempty"`
}

type GeneratedAt struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
}

type Operator struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types operator,omitempty" json:"operator,omitempty"`
}

type OperatorCode struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
}

type Platform struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types platform,omitempty" json:"platform,omitempty"`
}

type PreviousCallingPoints struct {
	CallingPointList *CallingPointList `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types callingPointList,omitempty" json:"callingPointList,omitempty"`
	XMLName          xml.Name          `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types previousCallingPoints,omitempty" json:"previousCallingPoints,omitempty"`
}

type RSID struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types rsid,omitempty" json:"rsid,omitempty"`
}

type ServiceType struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types serviceType,omitempty" json:"serviceType,omitempty"`
}

type St struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types st,omitempty" json:"st,omitempty"`
}

type STA struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types sta,omitempty" json:"sta,omitempty"`
}

type STD struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types std,omitempty" json:"std,omitempty"`
}

type SubsequentCallingPoints struct {
	CallingPointList *CallingPointList `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types callingPointList,omitempty" json:"callingPointList,omitempty"`
	XMLName          xml.Name          `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types subsequentCallingPoints,omitempty" json:"subsequentCallingPoints,omitempty"`
}

type Message struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2012-01-13/ldb/types message,omitempty" json:"message,omitempty"`
}

type CRSLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types crs,omitempty" json:"crs,omitempty"`
}

type ETDLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types etd,omitempty" json:"etd,omitempty"`
}

type GeneratedAtLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
}

type LocationLT4 struct {
	CRS          *CRSLT4          `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types crs,omitempty" json:"crs,omitempty"`
	LocationName *LocationNameLT4 `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	Via          *ViaLT4          `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types via,omitempty" json:"via,omitempty"`
	XMLName      xml.Name         `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types location,omitempty" json:"location,omitempty"`
}

type LocationNameLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types locationName,omitempty" json:"locationName,omitempty"`
}

type NRCCMessagesLT4 struct {
	Messages []*Message `xml:"http://thalesgroup.com/RTTI/2012-01-13/ldb/types message,omitempty" json:"message,omitempty"`
	XMLName  xml.Name   `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types nrccMessages,omitempty" json:"nrccMessages,omitempty"`
}

type OperatorLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types operator,omitempty" json:"operator,omitempty"`
}

type OperatorCodeLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
}

type PlatformLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types platform,omitempty" json:"platform,omitempty"`
}

type PlatformAvailableLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types platformAvailable,omitempty" json:"platformAvailable,omitempty"`
}

type ServiceIDLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types serviceID,omitempty" json:"serviceID,omitempty"`
}

type ServiceTypeLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types serviceType,omitempty" json:"serviceType,omitempty"`
}

type STDLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types std,omitempty" json:"std,omitempty"`
}

type ViaLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types via,omitempty" json:"via,omitempty"`
}

type DestinationLT5 struct {
	Location *LocationLT4 `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types location,omitempty" json:"location,omitempty"`
	XMLName  xml.Name     `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types destination,omitempty" json:"destination,omitempty"`
}

type RSIDLT5 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types rsid,omitempty" json:"rsid,omitempty"`
}

type DeparturesBoard struct {
	AttrXmlnslt       string                `xml:"xmlns lt,attr"  json:",omitempty"`
	AttrXmlnslt2      string                `xml:"xmlns lt2,attr"  json:",omitempty"`
	AttrXmlnslt3      string                `xml:"xmlns lt3,attr"  json:",omitempty"`
	AttrXmlnslt4      string                `xml:"xmlns lt4,attr"  json:",omitempty"`
	AttrXmlnslt5      string                `xml:"xmlns lt5,attr"  json:",omitempty"`
	AttrXmlnslt6      string                `xml:"xmlns lt6,attr"  json:",omitempty"`
	AttrXmlnslt7      string                `xml:"xmlns lt7,attr"  json:",omitempty"`
	CRS               *CRSLT4               `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types crs,omitempty" json:"crs,omitempty"`
	Departures        *DeparturesLT7        `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types departures,omitempty" json:"departures,omitempty"`
	GeneratedAt       *GeneratedAtLT4       `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
	LocationName      *LocationNameLT4      `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	NRCCMessages      *NRCCMessagesLT4      `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types nrccMessages,omitempty" json:"nrccMessages,omitempty"`
	PlatformAvailable *PlatformAvailableLT4 `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types platformAvailable,omitempty" json:"platformAvailable,omitempty"`
	XMLName           xml.Name              `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ DeparturesBoard,omitempty" json:"DeparturesBoard,omitempty"`
}

type NextDeparturesWithDetails struct {
	Attrxmlns       string           `xml:"xmlns,attr"  json:",omitempty"`
	DeparturesBoard *DeparturesBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ DeparturesBoard,omitempty" json:"DeparturesBoard,omitempty"`
	XMLName         xml.Name         `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetNextDeparturesWithDetailsResponse,omitempty" json:"GetNextDeparturesWithDetailsResponse,omitempty"`
}

type DeparturesLT7 struct {
	Destination []*DestinationLT7 `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types destination,omitempty" json:"destination,omitempty"`
	XMLName     xml.Name          `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types departures,omitempty" json:"departures,omitempty"`
}

type DestinationLT7 struct {
	Attrcrs string      `xml:"crs,attr"  json:",omitempty"`
	Service *ServiceLT7 `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types service,omitempty" json:"service,omitempty"`
	XMLName xml.Name    `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types destination,omitempty" json:"destination,omitempty"`
}

type ServiceLT7 struct {
	AttrXsiSpacenil         string                   `xml:"http://www.w3.org/2001/XMLSchema-instance nil,attr"  json:",omitempty"`
	Destination             *DestinationLT5          `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types destination,omitempty" json:"destination,omitempty"`
	DelayReason             *DelayReasonLT4          `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types delayReason,omitempty" json:"delayReason,omitempty"`
	ETD                     *ETDLT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types etd,omitempty" json:"etd,omitempty"`
	ETA                     *ETALT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types eta,omitempty" json:"eta,omitempty"`
	OperatorCode            *OperatorCodeLT4         `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types operatorCode,omitempty" json:"operatorCode,omitempty"`
	Operator                *OperatorLT4             `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types operator,omitempty" json:"operator,omitempty"`
	Origin                  *OriginLT5               `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types origin,omitempty" json:"origin,omitempty"`
	Platform                *PlatformLT4             `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types platform,omitempty" json:"platform,omitempty"`
	RSID                    *RSIDLT5                 `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types rsid,omitempty" json:"rsid,omitempty"`
	ServiceID               *ServiceIDLT4            `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types serviceID,omitempty" json:"serviceID,omitempty"`
	ServiceType             *ServiceTypeLT4          `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types serviceType,omitempty" json:"serviceType,omitempty"`
	STD                     *STDLT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types std,omitempty" json:"std,omitempty"`
	SubsequentCallingPoints *SubsequentCallingPoints `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types subsequentCallingPoints,omitempty" json:"subsequentCallingPoints,omitempty"`
	STA                     *STALT4                  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types sta,omitempty" json:"sta,omitempty"`
	PreviousCallingPoints   *PreviousCallingPoints   `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types previousCallingPoints,omitempty" json:"previousCallingPoints,omitempty"`
	XMLName                 xml.Name                 `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types service,omitempty" json:"service,omitempty"`
}

type DelayReasonLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types delayReason,omitempty" json:"delayReason,omitempty"`
}

type STLT7 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types st,omitempty" json:"st,omitempty"`
}

type NextDepartures struct {
	Attrxmlns       string           `xml:"xmlns,attr"  json:",omitempty"`
	DeparturesBoard *DeparturesBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ DeparturesBoard,omitempty" json:"DeparturesBoard,omitempty"`
	XMLName         xml.Name         `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetNextDeparturesResponse,omitempty" json:"GetNextDeparturesResponse,omitempty"`
}

type FastestDeparturesWithDetails struct {
	Attrxmlns       string           `xml:"xmlns,attr"  json:",omitempty"`
	DeparturesBoard *DeparturesBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ DeparturesBoard,omitempty" json:"DeparturesBoard,omitempty"`
	XMLName         xml.Name         `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetFastestDeparturesWithDetailsResponse,omitempty" json:"GetFastestDeparturesWithDetailsResponse,omitempty"`
}

type FastestDepartures struct {
	Attrxmlns       string           `xml:"xmlns,attr"  json:",omitempty"`
	DeparturesBoard *DeparturesBoard `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ DeparturesBoard,omitempty" json:"DeparturesBoard,omitempty"`
	XMLName         xml.Name         `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetFastestDeparturesResponse,omitempty" json:"GetFastestDeparturesResponse,omitempty"`
}

type FilterLocationNameLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types filterLocationName,omitempty" json:"filterLocationName,omitempty"`
}

type FilterCRSLT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types filtercrs,omitempty" json:"filtercrs,omitempty"`
}

type DepBoardWithDetails struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetDepBoardWithDetailsResponse,omitempty" json:"GetDepBoardWithDetailsResponse,omitempty"`
}

type GetStationBoardResult struct {
	AttrXmlnslt        string                 `xml:"xmlns lt,attr"  json:",omitempty"`
	AttrXmlnslt2       string                 `xml:"xmlns lt2,attr"  json:",omitempty"`
	AttrXmlnslt3       string                 `xml:"xmlns lt3,attr"  json:",omitempty"`
	AttrXmlnslt4       string                 `xml:"xmlns lt4,attr"  json:",omitempty"`
	AttrXmlnslt5       string                 `xml:"xmlns lt5,attr"  json:",omitempty"`
	AttrXmlnslt6       string                 `xml:"xmlns lt6,attr"  json:",omitempty"`
	AttrXmlnslt7       string                 `xml:"xmlns lt7,attr"  json:",omitempty"`
	CRS                *CRSLT4                `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types crs,omitempty" json:"crs,omitempty"`
	FilterLocationName *FilterLocationNameLT4 `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types filterLocationName,omitempty" json:"filterLocationName,omitempty"`
	FilterCRS          *FilterCRSLT4          `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types filtercrs,omitempty" json:"filtercrs,omitempty"`
	GeneratedAt        *GeneratedAtLT4        `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types generatedAt,omitempty" json:"generatedAt,omitempty"`
	LocationName       *LocationNameLT4       `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types locationName,omitempty" json:"locationName,omitempty"`
	NRCCMessages       *NRCCMessagesLT4       `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types nrccMessages,omitempty" json:"nrccMessages,omitempty"`
	PlatformAvailable  *PlatformAvailableLT4  `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types platformAvailable,omitempty" json:"platformAvailable,omitempty"`
	TrainServices      *TrainServicesLT7      `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types trainServices,omitempty" json:"trainServices,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	Text               string                 `xml:",chardata" json:",omitempty"`
}

type TrainServicesLT7 struct {
	Services []*ServiceLT7 `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types service,omitempty" json:"service,omitempty"`
	XMLName  xml.Name      `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/types trainServices,omitempty" json:"trainServices,omitempty"`
}

type DepartureBoard struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetDepartureBoardResponse,omitempty" json:"GetDepartureBoardResponse,omitempty"`
}

type ETALT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types eta,omitempty" json:"eta,omitempty"`
}

type STALT4 struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types sta,omitempty" json:"sta,omitempty"`
}

type ArrivalDepartureBoard struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrivalDepartureBoardResponse,omitempty" json:"GetArrivalDepartureBoardResponse,omitempty"`
}

type ArrBoardWithDetails struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrBoardWithDetailsResponse,omitempty" json:"GetArrBoardWithDetailsResponse,omitempty"`
}

type OriginLT5 struct {
	Location *LocationLT4 `xml:"http://thalesgroup.com/RTTI/2015-11-27/ldb/types location,omitempty" json:"location,omitempty"`
	XMLName  xml.Name     `xml:"http://thalesgroup.com/RTTI/2016-02-16/ldb/types origin,omitempty" json:"origin,omitempty"`
}

type ArrivalBoard struct {
	Attrxmlns          string                 `xml:"xmlns,attr"  json:",omitempty"`
	StationBoardResult *GetStationBoardResult `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetStationBoardResult,omitempty" json:"GetStationBoardResult,omitempty"`
	XMLName            xml.Name               `xml:"http://thalesgroup.com/RTTI/2017-10-01/ldb/ GetArrivalBoardResponse,omitempty" json:"GetArrivalBoardResponse,omitempty"`
}
