package decoder

import (
	"encoding/xml"
	"io"

	models "github.com/MartinsIrbe/national-rail-go-client/internal/models"
	"github.com/pkg/errors"
)

// NationalRailResponseDecoder used to decode national rail responses
type NationalRailResponseDecoder struct{}

// NewNationalRailResponseDecoder used to create a new instance of national rail response decoder
func NewNationalRailResponseDecoder() *NationalRailResponseDecoder {
	return &NationalRailResponseDecoder{}
}

// DecodeDepartureResponse decodes XML response reader to departures response struct
func (d *NationalRailResponseDecoder) DecodeDepartureResponse(resp io.Reader) (*models.DepartureResponse, error) {
	decoder := xml.NewDecoder(resp)

	for {
		token, err := decoder.Token()
		if err != nil || token == nil {
			break
		}

		switch startElement := token.(type) {
		case xml.StartElement:
			return decodeDepartureResponse(startElement, decoder)
		}
	}

	return nil, errors.New("failed to decode departure xml response")
}

func decodeDepartureResponse(se xml.StartElement, decoder *xml.Decoder) (*models.DepartureResponse, error) {
	if se.Name.Local == "Envelope" && se.Name.Space == "http://schemas.xmlsoap.org/soap/envelope/" {
		var departureResponse models.DepartureResponse
		if err := decoder.DecodeElement(&departureResponse, &se); err != nil {
			return nil, errors.Wrapf(err, "failed to decode departure response")
		}

		return &departureResponse, nil
	}

	return nil, errors.New("failed to decode departure response as the expected envelope xml start element is missing")
}
