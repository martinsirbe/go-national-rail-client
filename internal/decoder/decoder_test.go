package decoder_test

import (
	"os"
	"testing"

	"strings"

	"github.com/MartinsIrbe/national-rail-go-client/internal/decoder"
	"gopkg.in/check.v1"
)

func TestDecoder(t *testing.T) { check.TestingT(t) }

type DecoderSuite struct {
	decoder *decoder.NationalRailResponseDecoder
}

var _ = check.Suite(&DecoderSuite{})

func (s *DecoderSuite) SetUpSuite(c *check.C) {
	s.decoder = decoder.NewNationalRailResponseDecoder()
}

func (s *DecoderSuite) TestSuccessfullyDecodedDepartureResponse(c *check.C) {
	dr, err := os.Open("../../test_data/departure_response.xml")
	if err != nil {
		c.Errorf("failed to load departure_response.xml, err - %v", err)
	}

	resp, decodeErr := s.decoder.DecodeArrBoardWithDetailsResponse(dr)
	if decodeErr != nil {
		c.Errorf("got an error while decoding departure response")
	}

	c.Assert(resp.AttrXmlnssoap, check.Equals, "http://schemas.xmlsoap.org/soap/envelope/")
	c.Assert(resp.AttrXmlnsxsd, check.Equals, "http://www.w3.org/2001/XMLSchema")
	c.Assert(resp.AttrXmlnsxsi, check.Equals, "http://www.w3.org/2001/XMLSchema-instance")

	c.Assert(resp.Body, check.NotNil)
	c.Assert(resp.Body.XMLName.Space, check.Equals, "http://schemas.xmlsoap.org/soap/envelope/")
	c.Assert(resp.Body.XMLName.Local, check.Equals, "Body")

	c.Assert(resp.Body.GetDepartureBoardResponse, check.NotNil)
	c.Assert(resp.Body.GetDepartureBoardResponse.XMLName, check.NotNil)
	c.Assert(resp.Body.GetDepartureBoardResponse.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/")
	c.Assert(resp.Body.GetDepartureBoardResponse.XMLName.Local, check.Equals, "GetDepartureBoardResponse")

	c.Assert(resp.Body.GetDepartureBoardResponse.Attrxmlns, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/")

	sbr := resp.Body.GetDepartureBoardResponse.GetStationBoardResult
	c.Assert(sbr, check.NotNil)
	c.Assert(sbr.XMLName, check.NotNil)
	c.Assert(sbr.AttrXmlnslt, check.Equals, "http://thalesgroup.com/RTTI/2012-01-13/ldb/types")
	c.Assert(sbr.AttrXmlnslt2, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(sbr.AttrXmlnslt3, check.Equals, "http://thalesgroup.com/RTTI/2015-05-14/ldb/types")
	c.Assert(sbr.AttrXmlnslt4, check.Equals, "http://thalesgroup.com/RTTI/2015-11-27/ldb/types")
	c.Assert(sbr.AttrXmlnslt5, check.Equals, "http://thalesgroup.com/RTTI/2016-02-16/ldb/types")
	c.Assert(sbr.AttrXmlnslt6, check.Equals, "http://thalesgroup.com/RTTI/2017-02-02/ldb/types")
	c.Assert(sbr.AttrXmlnslt7, check.Equals, "http://thalesgroup.com/RTTI/2017-10-01/ldb/types")

	c.Assert(sbr.OriginLocationCode.Text, check.Equals, "CHX")
	c.Assert(sbr.OriginLocationCode.XMLName.Local, check.Equals, "crs")
	c.Assert(sbr.OriginLocationCode.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.FilterLocationName.Text, check.Equals, "Dartford")
	c.Assert(sbr.FilterLocationName.XMLName.Local, check.Equals, "filterLocationName")
	c.Assert(sbr.FilterLocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.DestinationLocationCode.Text, check.Equals, "DFD")
	c.Assert(sbr.DestinationLocationCode.XMLName.Local, check.Equals, "filtercrs")
	c.Assert(sbr.DestinationLocationCode.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(len(sbr.WarningMessages.Messages), check.Equals, 2)
	c.Assert(sbr.WarningMessages.Messages[0].Text, check.Equals, "Test warning 1")
	c.Assert(sbr.WarningMessages.Messages[0].XMLName.Local, check.Equals, "message")
	c.Assert(sbr.WarningMessages.Messages[0].XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2012-01-13/ldb/types")
	c.Assert(sbr.WarningMessages.Messages[1].Text, check.Equals, "Test warning 2")
	c.Assert(sbr.WarningMessages.Messages[1].XMLName.Local, check.Equals, "message")
	c.Assert(sbr.WarningMessages.Messages[1].XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2012-01-13/ldb/types")
	c.Assert(sbr.WarningMessages.XMLName.Local, check.Equals, "nrccMessages")
	c.Assert(sbr.WarningMessages.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.GeneratedAt.Text, check.Equals, "2019-09-04T15:07:03.2618621+00:00")
	c.Assert(sbr.GeneratedAt.XMLName.Local, check.Equals, "generatedAt")
	c.Assert(sbr.GeneratedAt.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.LocationName.Text, check.Equals, "London Charing Cross")
	c.Assert(sbr.LocationName.XMLName.Local, check.Equals, "locationName")
	c.Assert(sbr.LocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.PlatformAvailable.Text, check.Equals, "true")
	c.Assert(sbr.PlatformAvailable.XMLName.Local, check.Equals, "platformAvailable")
	c.Assert(sbr.PlatformAvailable.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(sbr.TrainServiceDetails.XMLName.Local, check.Equals, "trainServices")
	c.Assert(sbr.TrainServiceDetails.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")

	c.Assert(len(sbr.TrainServiceDetails.TrainServices), check.Equals, 2)
	service1 := sbr.TrainServiceDetails.TrainServices[0]
	c.Assert(service1.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.XMLName.Local, check.Equals, "service")
	c.Assert(service1.Platform.XMLName.Local, check.Equals, "platform")
	c.Assert(service1.Platform.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Platform.Text, check.Equals, "1")
	c.Assert(service1.Operator.Text, check.Equals, "Southeastern")
	c.Assert(service1.Operator.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Operator.XMLName.Local, check.Equals, "operator")
	c.Assert(service1.ServiceStatus.Text, check.Equals, "On time")
	c.Assert(service1.ServiceStatus.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.ServiceStatus.XMLName.Local, check.Equals, "etd")
	c.Assert(service1.Origin.Location.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Origin.Location.XMLName.Local, check.Equals, "location")
	c.Assert(service1.Origin.Location.LocationName.Text, check.Equals, "London Charing Cross")
	c.Assert(service1.Origin.Location.LocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Origin.Location.LocationName.XMLName.Local, check.Equals, "locationName")
	c.Assert(service1.Origin.Location.Via, check.IsNil)
	c.Assert(service1.Origin.Location.CRS.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Origin.Location.CRS.XMLName.Local, check.Equals, "crs")
	c.Assert(service1.Origin.Location.CRS.Text, check.Equals, "CHX")
	c.Assert(service1.Destination.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Destination.XMLName.Local, check.Equals, "destination")
	c.Assert(service1.Destination.Location.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Destination.Location.XMLName.Local, check.Equals, "location")
	c.Assert(service1.Destination.Location.LocationName.Text, check.Equals, "Gillingham (Kent)")
	c.Assert(service1.Destination.Location.LocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Destination.Location.LocationName.XMLName.Local, check.Equals, "locationName")
	c.Assert(service1.Destination.Location.Via.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Destination.Location.Via.XMLName.Local, check.Equals, "via")
	c.Assert(service1.Destination.Location.Via.Text, check.Equals, "via Lewisham & Woolwich Arsenal")
	c.Assert(service1.Destination.Location.CRS.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.Destination.Location.CRS.XMLName.Local, check.Equals, "crs")
	c.Assert(service1.Destination.Location.CRS.Text, check.Equals, "GLM")
	c.Assert(service1.DepartureTime.Text, check.Equals, "15:09")
	c.Assert(service1.DepartureTime.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.DepartureTime.XMLName.Local, check.Equals, "std")
	c.Assert(service1.OperatorCode.Text, check.Equals, "SE")
	c.Assert(service1.OperatorCode.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.OperatorCode.XMLName.Local, check.Equals, "operatorCode")
	c.Assert(service1.ServiceID.Text, check.Equals, "Al1SZfqGo0gwRATVpXALWg==")
	c.Assert(service1.ServiceID.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service1.ServiceID.XMLName.Local, check.Equals, "serviceID")

	service2 := sbr.TrainServiceDetails.TrainServices[1]
	c.Assert(service2.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.XMLName.Local, check.Equals, "service")
	c.Assert(service2.Platform.XMLName.Local, check.Equals, "platform")
	c.Assert(service2.Platform.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Platform.Text, check.Equals, "5")
	c.Assert(service2.Operator.Text, check.Equals, "Southeastern")
	c.Assert(service2.Operator.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Operator.XMLName.Local, check.Equals, "operator")
	c.Assert(service2.ServiceStatus.Text, check.Equals, "On time")
	c.Assert(service2.ServiceStatus.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.ServiceStatus.XMLName.Local, check.Equals, "etd")
	c.Assert(service2.Origin.Location.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Origin.Location.XMLName.Local, check.Equals, "location")
	c.Assert(service2.Origin.Location.LocationName.Text, check.Equals, "London Charing Cross")
	c.Assert(service2.Origin.Location.LocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Origin.Location.LocationName.XMLName.Local, check.Equals, "locationName")
	c.Assert(service2.Origin.Location.Via, check.IsNil)
	c.Assert(service2.Origin.Location.CRS.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Origin.Location.CRS.XMLName.Local, check.Equals, "crs")
	c.Assert(service2.Origin.Location.CRS.Text, check.Equals, "CHX")
	c.Assert(service2.Destination.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Destination.XMLName.Local, check.Equals, "destination")
	c.Assert(service2.Destination.Location.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Destination.Location.XMLName.Local, check.Equals, "location")
	c.Assert(service2.Destination.Location.LocationName.Text, check.Equals, "Dartford")
	c.Assert(service2.Destination.Location.LocationName.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Destination.Location.LocationName.XMLName.Local, check.Equals, "locationName")
	c.Assert(service2.Destination.Location.Via.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Destination.Location.Via.XMLName.Local, check.Equals, "via")
	c.Assert(service2.Destination.Location.Via.Text, check.Equals, "via Bexleyheath")
	c.Assert(service2.Destination.Location.CRS.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.Destination.Location.CRS.XMLName.Local, check.Equals, "crs")
	c.Assert(service2.Destination.Location.CRS.Text, check.Equals, "DFD")
	c.Assert(service2.DepartureTime.Text, check.Equals, "15:14")
	c.Assert(service2.DepartureTime.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.DepartureTime.XMLName.Local, check.Equals, "std")
	c.Assert(service2.OperatorCode.Text, check.Equals, "SE")
	c.Assert(service2.OperatorCode.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.OperatorCode.XMLName.Local, check.Equals, "operatorCode")
	c.Assert(service2.ServiceID.Text, check.Equals, "iMkz1KfNnhruLFmjJ6T+Iw==")
	c.Assert(service2.ServiceID.XMLName.Space, check.Equals, "http://thalesgroup.com/RTTI/2014-02-20/ldb/types")
	c.Assert(service2.ServiceID.XMLName.Local, check.Equals, "serviceID")
}

func (s *DecoderSuite) TestDecodeDepartureResponseFailedInvalidResponseString(c *check.C) {
	_, decodeErr := s.decoder.DecodeArrBoardWithDetailsResponse(strings.NewReader("bad string"))
	c.Assert(decodeErr, check.NotNil)
	c.Assert(decodeErr, check.ErrorMatches, "failed to decode departure xml response")
}

func (s *DecoderSuite) TestDecodeDepartureResponseFailedMissingEnvelopeFromResponseString(c *check.C) {
	dr, err := os.Open("../../test_data/departure_response_missing_envelope.xml")
	if err != nil {
		c.Errorf("failed to load departure_response_missing_envelope.xml, err - %v", err)
	}

	_, decodeErr := s.decoder.DecodeArrBoardWithDetailsResponse(dr)
	c.Assert(decodeErr, check.NotNil)
	c.Assert(decodeErr, check.ErrorMatches, "failed to decode departure response as the expected envelope xml start element is missing")
}
