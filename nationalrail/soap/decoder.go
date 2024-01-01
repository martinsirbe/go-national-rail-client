package soap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

const (
	ArrBoardWithDetailsResponse          = "GetArrBoardWithDetailsResponse"
	ArrDepBoardWithDetailsResponse       = "GetArrDepBoardWithDetailsResponse"
	ArrivalBoardResponse                 = "GetArrivalBoardResponse"
	ArrivalDepartureBoardResponse        = "GetArrivalDepartureBoardResponse"
	DepartureBoardResponse               = "GetDepartureBoardResponse"
	DepBoardWithDetailsResponse          = "GetDepBoardWithDetailsResponse"
	FastestDeparturesResponse            = "GetFastestDeparturesResponse"
	FastestDeparturesWithDetailsResponse = "GetFastestDeparturesWithDetailsResponse"
	NextDeparturesResponse               = "GetNextDeparturesResponse"
	NextDeparturesWithDetailsResponse    = "GetNextDeparturesWithDetailsResponse"
	ServiceDetailsResponse               = "GetServiceDetailsResponse"

	decodeErr = "failed to decode %s"
)

func DecodeArrBoardWithDetailsResponse(resp io.Reader) (*GetArrBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, ArrBoardWithDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, ArrBoardWithDetailsResponse, err)
	}

	return decodedResp.(*GetArrBoardWithDetailsResponse), nil
}

func DecodeArrDepBoardWithDetailsResponse(resp io.Reader) (*GetArrDepBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, ArrDepBoardWithDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, ArrDepBoardWithDetailsResponse, err)
	}

	return decodedResp.(*GetArrDepBoardWithDetailsResponse), nil
}

func DecodeArrivalBoardResponse(resp io.Reader) (*GetArrivalBoardResponse, error) {
	decodedResp, err := decode(resp, ArrivalBoardResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, ArrivalBoardResponse, err)
	}

	return decodedResp.(*GetArrivalBoardResponse), nil
}

func DecodeArrivalDepartureBoardResponse(resp io.Reader) (*GetArrivalDepartureBoardResponse, error) {
	decodedResp, err := decode(resp, ArrivalDepartureBoardResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, ArrivalDepartureBoardResponse, err)
	}

	return decodedResp.(*GetArrivalDepartureBoardResponse), nil
}

func DecodeDepartureBoardResponse(resp io.Reader) (*GetDepartureBoardResponse, error) {
	decodedResp, err := decode(resp, DepartureBoardResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, DepartureBoardResponse, err)
	}

	return decodedResp.(*GetDepartureBoardResponse), nil
}

func DecodeDepBoardWithDetailsResponse(resp io.Reader) (*GetDepBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, DepBoardWithDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, DepBoardWithDetailsResponse, err)
	}

	return decodedResp.(*GetDepBoardWithDetailsResponse), nil
}

func DecodeFastestDeparturesResponse(resp io.Reader) (*GetFastestDeparturesResponse, error) {
	decodedResp, err := decode(resp, FastestDeparturesResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, FastestDeparturesResponse, err)
	}

	return decodedResp.(*GetFastestDeparturesResponse), nil
}

func DecodeFastestDeparturesWithDetailsResponse(resp io.Reader) (*GetFastestDeparturesWithDetailsResponse, error) {
	decodedResp, err := decode(resp, FastestDeparturesWithDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, FastestDeparturesWithDetailsResponse, err)
	}

	return decodedResp.(*GetFastestDeparturesWithDetailsResponse), nil
}

func DecodeNextDeparturesResponse(resp io.Reader) (*GetNextDeparturesResponse, error) {
	decodedResp, err := decode(resp, NextDeparturesResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, NextDeparturesResponse, err)
	}

	return decodedResp.(*GetNextDeparturesResponse), nil
}

func DecodeNextDeparturesWithDetailsResponse(resp io.Reader) (*GetNextDeparturesWithDetailsResponse, error) {
	decodedResp, err := decode(resp, NextDeparturesWithDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, NextDeparturesWithDetailsResponse, err)
	}

	return decodedResp.(*GetNextDeparturesWithDetailsResponse), nil
}

func DecodeServiceDetailsResponse(resp io.Reader) (*GetServiceDetailsResponse, error) {
	decodedResp, err := decode(resp, ServiceDetailsResponse)
	if err != nil {
		return nil, fmt.Errorf(decodeErr, ServiceDetailsResponse, err)
	}

	return decodedResp.(*GetServiceDetailsResponse), nil
}

func decode(resp io.Reader, respType string) (interface{}, error) {
	decoder := xml.NewDecoder(resp)

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		if token == nil {
			break
		}

		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == "Envelope" && startElement.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" {
				switch respType {
				case ArrBoardWithDetailsResponse:
					var resp GetArrBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrDepBoardWithDetailsResponse:
					var resp GetArrDepBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrivalBoardResponse:
					var resp GetArrivalBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrivalDepartureBoardResponse:
					var resp GetArrivalDepartureBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case DepartureBoardResponse:
					var resp GetDepartureBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case DepBoardWithDetailsResponse:
					var resp GetDepBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case FastestDeparturesResponse:
					var resp GetFastestDeparturesResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case FastestDeparturesWithDetailsResponse:
					var resp GetFastestDeparturesWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case NextDeparturesResponse:
					var resp GetNextDeparturesResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case NextDeparturesWithDetailsResponse:
					var resp GetNextDeparturesWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ServiceDetailsResponse:
					var resp GetServiceDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				}
			}
			return nil, errors.New("expected envelope xml start element is missing")
		}
	}
	return nil, errors.New("xml token is missing")
}
