package nationalrail

import (
	"fmt"
	"strings"
)

const (
	soapEnvelope = `<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
xmlns:typ="http://thalesgroup.com/RTTI/2013-11-28/Token/types" 
xmlns:ldb="http://thalesgroup.com/RTTI/2021-11-01/ldb/">
   <soap:Header>
      <typ:AccessToken>
         <typ:TokenValue>%s</typ:TokenValue>
      </typ:AccessToken>
   </soap:Header>
   <soap:Body>
      %s
   </soap:Body>
</soap:Envelope>`

	getArrBoardWithDetailsRequest = `<ldb:GetArrBoardWithDetailsRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
		 %s
      </ldb:GetArrBoardWithDetailsRequest>`

	getArrDepBoardWithDetailsRequest = `<ldb:GetArrDepBoardWithDetailsRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
         %s
      </ldb:GetArrDepBoardWithDetailsRequest>`

	getArrivalBoardRequest = `<ldb:GetArrivalBoardRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
         %s
      </ldb:GetArrivalBoardRequest>`

	getArrivalDepartureBoardRequest = `<ldb:GetArrivalDepartureBoardRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
         %s
      </ldb:GetArrivalDepartureBoardRequest>`

	getDepartureBoardRequest = `<ldb:GetDepartureBoardRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
         %s
      </ldb:GetDepartureBoardRequest>`

	getDepBoardWithDetailsRequest = `<ldb:GetDepBoardWithDetailsRequest>
         <ldb:numRows>%d</ldb:numRows>
         <ldb:crs>%s</ldb:crs>
         %s
      </ldb:GetDepBoardWithDetailsRequest>`

	getFastestDeparturesRequest = `<ldb:GetFastestDeparturesRequest>
         <ldb:crs>%s</ldb:crs>
		 <ldb:filterList>
		 	%s
		 </ldb:filterList>
         %s
      </ldb:GetFastestDeparturesRequest>`

	getFastestDeparturesWithDetailsRequest = `<ldb:GetFastestDeparturesWithDetailsRequest>
         <ldb:crs>%s</ldb:crs>
         <ldb:filterList>
            %s
         </ldb:filterList>
         %s
      </ldb:GetFastestDeparturesWithDetailsRequest>`

	getNextDeparturesRequest = `<ldb:GetNextDeparturesRequest>
         <ldb:crs>%s</ldb:crs>
         <ldb:filterList>
            %s
         </ldb:filterList>
         %s
      </ldb:GetNextDeparturesRequest>`

	getNextDeparturesWithDetailsRequest = `<ldb:GetNextDeparturesWithDetailsRequest>
         <ldb:crs>%s</ldb:crs>
         <ldb:filterList>
            %s
         </ldb:filterList>
         %s
      </ldb:GetNextDeparturesWithDetailsRequest>`

	getServiceDetailsRequest = `<ldb:GetServiceDetailsRequest>
		 <ldb:serviceID>%s</ldb:serviceID>
	  </ldb:GetServiceDetailsRequest>`
)

func CreateGetArrBoardWithDetailsRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getArrBoardWithDetailsRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetArrDepBoardWithDetailsRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getArrDepBoardWithDetailsRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetArrivalBoardRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getArrivalBoardRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetArrivalDepartureBoardRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getArrivalDepartureBoardRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetDepartureBoardRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getDepartureBoardRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetDepBoardWithDetailsRequest(token string, req *StationBoardRequest) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(req.FilterCRS, req.FilterType, req.TimeOffset, req.TimeWindow)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getDepBoardWithDetailsRequest,
			req.NumRows,
			req.CRS,
			optionalProperties,
		),
	))
}

func CreateGetFastestDeparturesRequest(
	token string,
	req *DeparturesBoardRequest,
	destinations []CRSCode,
) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(nil, nil, req.TimeOffset, req.TimeWindow)
	filterListProperties := getFilterListProperties(destinations)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getFastestDeparturesRequest,
			req.CRS,
			filterListProperties,
			optionalProperties,
		),
	))
}

func CreateGetFastestDeparturesWithDetailsRequest(
	token string,
	req *DeparturesBoardRequest,
	destinations []CRSCode,
) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(nil, nil, req.TimeOffset, req.TimeWindow)
	filterListProperties := getFilterListProperties(destinations)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getFastestDeparturesWithDetailsRequest,
			req.CRS,
			filterListProperties,
			optionalProperties,
		),
	))
}

func CreateGetNextDeparturesRequest(
	token string,
	req *DeparturesBoardRequest,
	destinations []CRSCode,
) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(nil, nil, req.TimeOffset, req.TimeWindow)
	filterListProperties := getFilterListProperties(destinations)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getNextDeparturesRequest,
			req.CRS,
			filterListProperties,
			optionalProperties,
		),
	))
}

func CreateGetNextDeparturesWithDetailsRequest(
	token string,
	req *DeparturesBoardRequest,
	destinations []CRSCode,
) *strings.Reader {
	optionalProperties := getOptionalRequestProperties(nil, nil, req.TimeOffset, req.TimeWindow)
	filterListProperties := getFilterListProperties(destinations)

	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getNextDeparturesWithDetailsRequest,
			req.CRS,
			filterListProperties,
			optionalProperties,
		),
	))
}

func CreateGetServiceDetailsRequest(token, serviceID string) *strings.Reader {
	return strings.NewReader(fmt.Sprintf(
		soapEnvelope,
		token,
		fmt.Sprintf(
			getServiceDetailsRequest,
			serviceID,
		),
	))
}

func getOptionalRequestProperties(filterCRS *string, filterType *string, timeOffset *int, timeWindow *int) string {
	optionalProperties := strings.Join([]string{
		getFilterCRSProp(filterCRS),
		getFilterTypeProp(filterType),
		getTimeOffsetProp(timeOffset),
		getTimeWindowProp(timeWindow),
	}, "\n		 ")
	return optionalProperties
}

func getFilterCRSProp(filterCRS *string) string {
	if filterCRS == nil {
		return ""
	}
	return fmt.Sprintf("<ldb:filterCrs>%s</ldb:filterCrs>", *filterCRS)
}

func getFilterTypeProp(filterType *string) string {
	if filterType == nil {
		return ""
	}
	return fmt.Sprintf("<ldb:filterType>%s</ldb:filterType>", *filterType)
}

func getTimeOffsetProp(timeOffset *int) string {
	if timeOffset == nil {
		return ""
	}
	return fmt.Sprintf("<ldb:timeOffset>%d</ldb:timeOffset>", *timeOffset)
}

func getTimeWindowProp(timeWindow *int) string {
	if timeWindow == nil {
		return ""
	}
	return fmt.Sprintf("<ldb:timeWindow>%d</ldb:timeWindow>", *timeWindow)
}

func getFilterListProperties(filterList []CRSCode) string {
	var filterListProperties []string
	for _, f := range filterList {
		filterListProperties = append(filterListProperties, fmt.Sprintf("<ldb:crs>%s</ldb:crs>", f))
	}

	return strings.Join(filterListProperties, "\n			")
}
