package decoder

import (
	"encoding/xml"
	"io"

	models "github.com/MartinsIrbe/national-rail-go-client/internal/models"
	"github.com/pkg/errors"
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

// NationalRailResponseDecoder used to decode national rail responses
type NationalRailResponseDecoder struct{}

// NewNationalRailResponseDecoder used to create a new instance of national rail response decoder
func NewNationalRailResponseDecoder() *NationalRailResponseDecoder {
	return &NationalRailResponseDecoder{}
}

func (d *NationalRailResponseDecoder) DecodeArrBoardWithDetailsResponse(resp io.Reader) (*models.GetArrBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, ArrBoardWithDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, ArrBoardWithDetailsResponse)
	}

	return decodedResp.(*models.GetArrBoardWithDetailsResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeArrDepBoardWithDetailsResponse(resp io.Reader) (*models.GetArrDepBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, ArrDepBoardWithDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, ArrDepBoardWithDetailsResponse)
	}

	return decodedResp.(*models.GetArrDepBoardWithDetailsResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeArrivalBoardResponse(resp io.Reader) (*models.GetArrivalBoardResponse, error) {
	decodedResp, err := decode(resp, ArrivalBoardResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, ArrivalBoardResponse)
	}

	return decodedResp.(*models.GetArrivalBoardResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeArrivalDepartureBoardResponse(resp io.Reader) (*models.GetArrivalDepartureBoardResponse, error) {
	decodedResp, err := decode(resp, ArrivalDepartureBoardResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, ArrivalDepartureBoardResponse)
	}

	return decodedResp.(*models.GetArrivalDepartureBoardResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeDepartureBoardResponse(resp io.Reader) (*models.GetDepartureBoardResponse, error) {
	decodedResp, err := decode(resp, DepartureBoardResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, DepartureBoardResponse)
	}

	return decodedResp.(*models.GetDepartureBoardResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeDepBoardWithDetailsResponse(resp io.Reader) (*models.GetDepBoardWithDetailsResponse, error) {
	decodedResp, err := decode(resp, DepBoardWithDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, DepBoardWithDetailsResponse)
	}

	return decodedResp.(*models.GetDepBoardWithDetailsResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeFastestDeparturesResponse(resp io.Reader) (*models.GetFastestDeparturesResponse, error) {
	decodedResp, err := decode(resp, FastestDeparturesResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, FastestDeparturesResponse)
	}

	return decodedResp.(*models.GetFastestDeparturesResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeFastestDeparturesWithDetailsResponse(resp io.Reader) (*models.GetFastestDeparturesWithDetailsResponse, error) {
	decodedResp, err := decode(resp, FastestDeparturesWithDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, FastestDeparturesWithDetailsResponse)
	}

	return decodedResp.(*models.GetFastestDeparturesWithDetailsResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeNextDeparturesResponse(resp io.Reader) (*models.GetNextDeparturesResponse, error) {
	decodedResp, err := decode(resp, NextDeparturesResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, NextDeparturesResponse)
	}

	return decodedResp.(*models.GetNextDeparturesResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeNextDeparturesWithDetailsResponse(resp io.Reader) (*models.GetNextDeparturesWithDetailsResponse, error) {
	decodedResp, err := decode(resp, NextDeparturesWithDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, NextDeparturesWithDetailsResponse)
	}

	return decodedResp.(*models.GetNextDeparturesWithDetailsResponse), nil
}

func (d *NationalRailResponseDecoder) DecodeServiceDetailsResponse(resp io.Reader) (*models.GetServiceDetailsResponse, error) {
	decodedResp, err := decode(resp, ServiceDetailsResponse)
	if err != nil {
		return nil, errors.Wrapf(err, decodeErr, ServiceDetailsResponse)
	}

	return decodedResp.(*models.GetServiceDetailsResponse), nil
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
					var resp models.GetArrBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrDepBoardWithDetailsResponse:
					var resp models.GetArrDepBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrivalBoardResponse:
					var resp models.GetArrivalBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ArrivalDepartureBoardResponse:
					var resp models.GetArrivalDepartureBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case DepartureBoardResponse:
					var resp models.GetDepartureBoardResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case DepBoardWithDetailsResponse:
					var resp models.GetDepBoardWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case FastestDeparturesResponse:
					var resp models.GetFastestDeparturesResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case FastestDeparturesWithDetailsResponse:
					var resp models.GetFastestDeparturesWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case NextDeparturesResponse:
					var resp models.GetNextDeparturesResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case NextDeparturesWithDetailsResponse:
					var resp models.GetNextDeparturesWithDetailsResponse
					if err := decoder.DecodeElement(&resp, &startElement); err != nil {
						return nil, err
					}
					return &resp, nil
				case ServiceDetailsResponse:
					var resp models.GetServiceDetailsResponse
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
