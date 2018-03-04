package mapper_test

import (
	"encoding/xml"
	"testing"

	"github.com/MartinsIrbe/national-rail-go-client/internal/mapper"
	im "github.com/MartinsIrbe/national-rail-go-client/internal/models"
	"gopkg.in/check.v1"
)

func TestMapper(t *testing.T) { check.TestingT(t) }

type MapperSuite struct {
	mapper *mapper.NationalRailResponseMapper
}

var _ = check.Suite(&MapperSuite{})

func (s *MapperSuite) SetUpSuite(c *check.C) {
	s.mapper = mapper.NewNationalRailResponseMapper()
}

func (s *MapperSuite) TestErrorReturnedOnMissingDepartureResponse(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(nil)
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have departure response but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have response body but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingGetDepartureBoardResponseFromResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{},
	})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have departure board response but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingGetStationBoardResultFromResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{
			GetDepartureBoardResponse: &im.GetDepartureBoardResponse{},
		},
	})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have station board result but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingGetStationBoardResultLocationNameFromResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{
			GetDepartureBoardResponse: &im.GetDepartureBoardResponse{
				GetStationBoardResult: &im.GetStationBoardResult{},
			},
		},
	})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have location name but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingGetStationBoardResultCurrentLocationCodeFromResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{
			GetDepartureBoardResponse: &im.GetDepartureBoardResponse{
				GetStationBoardResult: &im.GetStationBoardResult{
					LocationName: &im.LocationName{
						Text: "test",
					},
				},
			},
		},
	})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have current location code but was missing")
}

func (s *MapperSuite) TestErrorReturnedOnMissingGetStationBoardResultDestinationLocationCodeFromResponseBody(c *check.C) {
	_, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{
			GetDepartureBoardResponse: &im.GetDepartureBoardResponse{
				GetStationBoardResult: &im.GetStationBoardResult{
					LocationName: &im.LocationName{
						Text: "test",
					},
					OriginLocationCode: &im.OriginLocationCode{
						Text: "test",
					},
				},
			},
		},
	})
	c.Assert(err, check.ErrorMatches, "failed to map: expected to have destination location code but was missing")
}

func (s *MapperSuite) TestSuccessfullyMappedDepartureResponse(c *check.C) {
	mappedResp, err := s.mapper.MapDepartureResponse(&im.DepartureResponse{
		Body: &im.Body{
			GetDepartureBoardResponse: &im.GetDepartureBoardResponse{
				GetStationBoardResult: &im.GetStationBoardResult{
					WarningMessages: &im.WarningMessages{
						Messages: []*im.WarningMessage{
							{
								Text: "test-WarningMessage-1",
								XMLName: xml.Name{
									Space: "space-WarningMessage-1",
									Local: "local-WarningMessage-1",
								},
							},
							{
								Text: "test-WarningMessage-2",
								XMLName: xml.Name{
									Space: "space-WarningMessage-2",
									Local: "local-WarningMessage-2",
								},
							},
							{
								Text: "test-WarningMessage-3",
								XMLName: xml.Name{
									Space: "space-WarningMessage-3",
									Local: "local-WarningMessage-3",
								},
							},
						},
						XMLName: xml.Name{
							Space: "space-LocationName",
							Local: "local-LocationName",
						},
					},
					LocationName: &im.LocationName{
						Text: "test-LocationName",
						XMLName: xml.Name{
							Space: "space-LocationName",
							Local: "local-LocationName",
						},
					},
					OriginLocationCode: &im.OriginLocationCode{
						Text: "test-OriginLocationCode",
						XMLName: xml.Name{
							Space: "space-OriginLocationCode",
							Local: "local-OriginLocationCode",
						},
					},
					DestinationLocationCode: &im.DestinationLocationCode{
						Text: "test-DestinationLocationCode",
						XMLName: xml.Name{
							Space: "space-DestinationLocationCode",
							Local: "local-DestinationLocationCode",
						},
					},
					TrainServiceDetails: &im.TrainServiceDetails{
						TrainServices: []*im.TrainService{
							{
								Destination: &im.Destination{
									Location: &im.Location{
										CRS: &im.OriginLocationCode{
											Text: "test-OriginLocationCode-1",
											XMLName: xml.Name{
												Space: "space-OriginLocationCode",
												Local: "local-OriginLocationCode",
											},
										},
										LocationName: &im.LocationName{
											Text: "test-Destination-LocationName-1",
											XMLName: xml.Name{
												Space: "space-LocationName",
												Local: "local-LocationName",
											},
										},
										Via: &im.Via{
											Text: "test-Via-1",
											XMLName: xml.Name{
												Space: "space-Via",
												Local: "local-Via",
											},
										},
										XMLName: xml.Name{
											Space: "space-Destination-Location-1",
											Local: "local-Destination-Location",
										},
									},
									XMLName: xml.Name{
										Space: "space-Destination-1",
										Local: "local-Destination",
									},
								},
								ServiceStatus: &im.ServiceStatus{
									Text: "test-ServiceStatus-1",
									XMLName: xml.Name{
										Space: "space-ServiceStatus",
										Local: "local-ServiceStatus",
									},
								},
								OperatorCode: &im.OperatorCode{
									Text: "test-OperatorCode-1",
									XMLName: xml.Name{
										Space: "space-OperatorCode",
										Local: "local-OperatorCode",
									},
								},
								Operator: &im.Operator{
									Text: "test-Operator-1",
									XMLName: xml.Name{
										Space: "space-Operator",
										Local: "local-Operator",
									},
								},
								Origin: &im.Origin{
									Location: &im.Location{
										CRS: &im.OriginLocationCode{
											Text: "test-OriginLocationCode-1",
											XMLName: xml.Name{
												Space: "space-OriginLocationCode",
												Local: "local-OriginLocationCode",
											},
										},
										LocationName: &im.LocationName{
											Text: "test-Origin-LocationName-1",
											XMLName: xml.Name{
												Space: "space-LocationName",
												Local: "local-LocationName",
											},
										},
										Via: &im.Via{
											Text: "test-Via-1",
											XMLName: xml.Name{
												Space: "space-Via",
												Local: "local-Via",
											},
										},
										XMLName: xml.Name{
											Space: "space-Origin-Location",
											Local: "local-Origin-Location",
										},
									},
									XMLName: xml.Name{
										Space: "space-Origin",
										Local: "local-Origin",
									},
								},
								Platform: &im.Platform{
									Text: "test-Platform-1",
									XMLName: xml.Name{
										Space: "space-Platform",
										Local: "local-Platform",
									},
								},
								ServiceID: &im.ServiceID{
									Text: "test-ServiceID-1",
									XMLName: xml.Name{
										Space: "space-ServiceID",
										Local: "local-ServiceID",
									},
								},
								DepartureTime: &im.DepartureTime{
									Text: "21:20",
									XMLName: xml.Name{
										Space: "space-DepartureTime",
										Local: "local-DepartureTime",
									},
								},
								XMLName: xml.Name{
									Space: "space-TrainService",
									Local: "local-TrainService",
								},
							},
							{
								Destination: &im.Destination{
									Location: &im.Location{
										CRS: &im.OriginLocationCode{
											Text: "test-OriginLocationCode-2",
											XMLName: xml.Name{
												Space: "space-OriginLocationCode",
												Local: "local-OriginLocationCode",
											},
										},
										LocationName: &im.LocationName{
											Text: "test-Destination-LocationName-2",
											XMLName: xml.Name{
												Space: "space-LocationName",
												Local: "local-LocationName",
											},
										},
										Via: &im.Via{
											Text: "test-Via-2",
											XMLName: xml.Name{
												Space: "space-Via",
												Local: "local-Via",
											},
										},
										XMLName: xml.Name{
											Space: "space-Destination-Location-2",
											Local: "local-Destination-Location",
										},
									},
									XMLName: xml.Name{
										Space: "space-Destination-2",
										Local: "local-Destination",
									},
								},
								ServiceStatus: &im.ServiceStatus{
									Text: "test-ServiceStatus-2",
									XMLName: xml.Name{
										Space: "space-ServiceStatus",
										Local: "local-ServiceStatus",
									},
								},
								OperatorCode: &im.OperatorCode{
									Text: "test-OperatorCode-2",
									XMLName: xml.Name{
										Space: "space-OperatorCode",
										Local: "local-OperatorCode",
									},
								},
								Operator: &im.Operator{
									Text: "test-Operator-2",
									XMLName: xml.Name{
										Space: "space-Operator",
										Local: "local-Operator",
									},
								},
								Origin: &im.Origin{
									Location: &im.Location{
										CRS: &im.OriginLocationCode{
											Text: "test-OriginLocationCode-2",
											XMLName: xml.Name{
												Space: "space-OriginLocationCode",
												Local: "local-OriginLocationCode",
											},
										},
										LocationName: &im.LocationName{
											Text: "test-Origin-LocationName-2",
											XMLName: xml.Name{
												Space: "space-LocationName",
												Local: "local-LocationName",
											},
										},
										Via: &im.Via{
											Text: "test-Via-2",
											XMLName: xml.Name{
												Space: "space-Via",
												Local: "local-Via",
											},
										},
										XMLName: xml.Name{
											Space: "space-Origin-Location",
											Local: "local-Origin-Location",
										},
									},
									XMLName: xml.Name{
										Space: "space-Origin",
										Local: "local-Origin",
									},
								},
								ServiceID: &im.ServiceID{
									Text: "test-ServiceID-2",
									XMLName: xml.Name{
										Space: "space-ServiceID",
										Local: "local-ServiceID",
									},
								},
								XMLName: xml.Name{
									Space: "space-TrainService",
									Local: "local-TrainService",
								},
							},
						},
					},
				},
			},
		},
	})

	c.Assert(err, check.IsNil)

	c.Assert(mappedResp.DestinationLocationCode, check.Equals, "test-DestinationLocationCode")
	c.Assert(mappedResp.OriginLocationCode, check.Equals, "test-OriginLocationCode")
	c.Assert(len(mappedResp.WarningMessages), check.Equals, 3)
	c.Assert(mappedResp.WarningMessages[0], check.Equals, "test-WarningMessage-1")
	c.Assert(mappedResp.WarningMessages[1], check.Equals, "test-WarningMessage-2")
	c.Assert(mappedResp.WarningMessages[2], check.Equals, "test-WarningMessage-3")

	c.Assert(len(mappedResp.DepartureDetails), check.Equals, 2)

	dd1 := mappedResp.DepartureDetails[0]
	c.Assert(*dd1.Via, check.Equals, "test-Via-1")
	c.Assert(dd1.Destination, check.Equals, "test-Destination-LocationName-1")
	c.Assert(dd1.Operator, check.Equals, "test-Operator-1")
	c.Assert(dd1.Platform, check.Equals, "test-Platform-1")
	c.Assert(dd1.Status, check.Equals, "test-ServiceStatus-1")
	c.Assert(dd1.Time.Hour(), check.Equals, 21)
	c.Assert(dd1.Time.Minute(), check.Equals, 20)
	c.Assert(dd1.OriginLocation, check.Equals, "test-Origin-LocationName-1")

	dd2 := mappedResp.DepartureDetails[1]
	c.Assert(*dd2.Via, check.Equals, "test-Via-2")
	c.Assert(dd2.Destination, check.Equals, "test-Destination-LocationName-2")
	c.Assert(dd2.Operator, check.Equals, "test-Operator-2")
	c.Assert(dd2.Platform, check.Equals, "not available")
	c.Assert(dd2.Status, check.Equals, "test-ServiceStatus-2")
	c.Assert(dd2.Time, check.IsNil)
	c.Assert(dd2.OriginLocation, check.Equals, "test-Origin-LocationName-2")
}
