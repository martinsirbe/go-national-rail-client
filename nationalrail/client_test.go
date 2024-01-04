package nationalrail_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nr "github.com/martinsirbe/go-national-rail-client/nationalrail"
)

const testAccessToken = "93f1cb45-b961-4906-bd89-f0158053e696"

func TestGetArrivalsWithDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetArrBoardWithDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Services: []*nr.Service{
					{
						ID:   "3110547GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonStPancrasIntl,
							Name: nr.StationNameLondonStPancrasIntl,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeFaversham,
								Name: nr.StationNameFaversham,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A5",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A6",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
							},
						},
						Length: 6,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 2,
						PreviousCallingPointsPerOrigin: [][]*nr.CallingPoint{
							{
								{
									CRS:  nr.StationCodeFaversham,
									Name: nr.StationNameFaversham,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A5",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A6",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
										},
									},
									Length: 6,
									St:     "09:00",
								},
								{
									CRS:  nr.StationCodeSittingbourne,
									Name: nr.StationNameSittingbourne,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A5",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A6",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
										},
									},
									Length: 6,
									St:     "09:10",
								},
								{
									CRS:  nr.StationCodeRainhamKent,
									Name: nr.StationNameRainhamKent,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A5",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A6",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
										},
									},
									Length: 6,
									St:     "09:19",
								},
							},
						},
						RSID: nil,
						STA:  "09:23",
					},
					{
						ID:   "3110958GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRamsgate,
								Name: nr.StationNameRamsgate,
							},
							{
								CRS:  nr.StationCodeDoverPriory,
								Name: nr.StationNameDoverPriory,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 2,
						PreviousCallingPointsPerOrigin: [][]*nr.CallingPoint{
							{
								{
									CRS:  nr.StationCodeRamsgate,
									Name: nr.StationNameRamsgate,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(9),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(35),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(18),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:12",
								},
								{
									CRS:  nr.StationCodeDumptonPark,
									Name: nr.StationNameDumptonPark,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(22),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(35),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(18),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:15",
								},
								{
									CRS:  nr.StationCodeBroadstairs,
									Name: nr.StationNameBroadstairs,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:18",
								},
								{
									CRS:  nr.StationCodeMargate,
									Name: nr.StationNameMargate,
									At:   pstr("08:25"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(23),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(23),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(9),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(10),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:23",
								},
								{
									CRS:  nr.StationCodeWestgateOnSea,
									Name: nr.StationNameWestgateOnSea,
									At:   pstr("08:29"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(23),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(23),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:27",
								},
								{
									CRS:  nr.StationCodeBirchingtonOnSea,
									Name: nr.StationNameBirchingtonOnSea,
									At:   pstr("08:32"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(16),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(30),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:30",
								},
								{
									CRS:  nr.StationCodeHerneBay,
									Name: nr.StationNameHerneBay,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(15),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(18),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(27),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(27),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:39",
								},
								{
									CRS:  nr.StationCodeChestfieldAndSwalecliffe,
									Name: nr.StationNameChestfieldAndSwalecliffe,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:43",
								},
								{
									CRS:  nr.StationCodeWhitstable,
									Name: nr.StationNameWhitstable,
									At:   pstr("08:48"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(28),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(32),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(14),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(20),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(25),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(14),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:46",
								},
								{
									CRS:  nr.StationCodeFaversham,
									Name: nr.StationNameFaversham,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "09:04",
								},
								{
									CRS:  nr.StationCodeTeynham,
									Name: nr.StationNameTeynham,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "09:10",
								},
								{
									CRS:  nr.StationCodeSittingbourne,
									Name: nr.StationNameSittingbourne,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "09:15",
								},
								{
									CRS:  nr.StationCodeNewington,
									Name: nr.StationNameNewington,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "09:20",
								},
								{
									CRS:  nr.StationCodeRainhamKent,
									Name: nr.StationNameRainhamKent,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "09:24",
								},
							},
							{
								{
									CRS:  nr.StationCodeDoverPriory,
									Name: nr.StationNameDoverPriory,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(35),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(18),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:16",
								},
								{
									CRS:  nr.StationCodeKearsney,
									Name: nr.StationNameKearsney,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(9),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(10),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:20",
								},
								{
									CRS:  nr.StationCodeShepherdsWell,
									Name: nr.StationNameShepherdsWell,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:26",
								},
								{
									CRS:  nr.StationCodeSnowdown,
									Name: nr.StationNameSnowdown,
									At:   pstr("No report"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:29",
								},
								{
									CRS:  nr.StationCodeAylesham,
									Name: nr.StationNameAylesham,
									At:   pstr("No report"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:32",
								},
								{
									CRS:  nr.StationCodeAdisham,
									Name: nr.StationNameAdisham,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(17),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(8),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:35",
								},
								{
									CRS:  nr.StationCodeBekesbourne,
									Name: nr.StationNameBekesbourne,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(27),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:39",
								},
								{
									CRS:  nr.StationCodeCanterburyEast,
									Name: nr.StationNameCanterburyEast,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(25),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(14),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(13),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:45",
								},
								{
									CRS:  nr.StationCodeSelling,
									Name: nr.StationNameSelling,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(27),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(14),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(9),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "08:54",
								},
								{
									CRS:  nr.StationCodeFaversham,
									Name: nr.StationNameFaversham,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 4,
									St:     "09:04",
								},
							},
						},
						RSID: pstr("SE512001"),
						STA:  "09:29",
					},
				},
				PlatformAvailable: true,
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
					"Severe weather to impact Southern, Gatwick Express and Thameslink services today. More details can be found in Latest Travel News",
				},
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeRainhamKent,
					LocationName: nr.StationNameRainhamKent,
					Type:         "from",
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T08:55:05.8471861Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetArrivalsWithDetails(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)

			sort.Slice(tc.expectedStationBoard.Services, func(i, j int) bool {
				return tc.expectedStationBoard.Services[i].ID < tc.expectedStationBoard.Services[j].ID
			})
			sort.Slice(stationBoard.Services, func(i, j int) bool {
				return stationBoard.Services[i].ID < stationBoard.Services[j].ID
			})

			for i, expectedService := range tc.expectedStationBoard.Services {
				actualService := stationBoard.Services[i]

				assert.Equal(t, expectedService.ID, actualService.ID)
				assert.Equal(t, expectedService.Type, actualService.Type)
				assert.Equal(t, expectedService.Destination, actualService.Destination)
				assert.Equal(t, expectedService.ETA, actualService.ETA)
				assert.Equal(t, expectedService.Length, actualService.Length)
				assert.Equal(t, expectedService.Operator, actualService.Operator)
				assert.Equal(t, expectedService.Platform, actualService.Platform)
				assert.Equal(t, expectedService.RSID, actualService.RSID)
				assert.Equal(t, expectedService.STA, actualService.STA)

				assert.ElementsMatch(t, expectedService.Origins, actualService.Origins)
				assert.ElementsMatch(t, expectedService.Formation.Coaches, actualService.Formation.Coaches)

				for expi, expectedCallingPoints := range expectedService.PreviousCallingPointsPerOrigin {
					for expj, expectedCallingPoint := range expectedCallingPoints {
						actualCallingPoint := actualService.PreviousCallingPointsPerOrigin[expi][expj]
						compareCallingPoint(t, expectedCallingPoint, actualCallingPoint)
					}
				}
			}
		})
	}
}

func TestGetArrivalsAndDeparturesWithDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetArrDepBoardWithDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Services: []*nr.Service{
					{
						ID:   "3111018GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRamsgate,
								Name: nr.StationNameRamsgate,
							},
						},
						ETA: "19:59",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},

								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 2,
						PreviousCallingPointsPerOrigin: [][]*nr.CallingPoint{
							{
								{
									CRS:  nr.StationCodeRamsgate,
									Name: nr.StationNameRamsgate,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(8),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(9),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(10),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(48),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(26),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(10),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},

											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(25),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "18:52",
									DelayReason: nil,
								},
								{
									CRS:  nr.StationCodeDumptonPark,
									Name: nr.StationNameDumptonPark,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(18),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},

											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(4),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "18:55",
									DelayReason: nil,
								},
								{
									CRS:  nr.StationCodeBroadstairs,
									Name: nr.StationNameBroadstairs,
									At:   pstr("On time"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},

											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "18:58",
									DelayReason: nil,
								},
								{
									CRS:  nr.StationCodeMargate,
									Name: nr.StationNameMargate,
									At:   pstr("19:09"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(25),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(28),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(2),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(0),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:04",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeWestgateOnSea,
									Name: nr.StationNameWestgateOnSea,
									At:   pstr("19:14"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:08",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeBirchingtonOnSea,
									Name: nr.StationNameBirchingtonOnSea,
									At:   pstr("19:17"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(25),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(11),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(4),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(33),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(7),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(5),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:12",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeHerneBay,
									Name: nr.StationNameHerneBay,
									At:   pstr("19:25"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(28),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(4),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(45),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(7),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:21",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeChestfieldAndSwalecliffe,
									Name: nr.StationNameChestfieldAndSwalecliffe,
									At:   pstr("19:29"),
									Et:   nil,
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: pint(28),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: pint(12),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: pint(4),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: pint(45),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: pint(7),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: pint(6),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: pint(3),
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:24",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeWhitstable,
									Name: nr.StationNameWhitstable,
									At:   nil,
									Et:   pstr("19:31"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:28",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeFaversham,
									Name: nr.StationNameFaversham,
									At:   nil,
									Et:   pstr("19:40"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:37",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeSittingbourne,
									Name: nr.StationNameSittingbourne,
									At:   nil,
									Et:   pstr("19:48"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:45",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
								{
									CRS:  nr.StationCodeRainhamKent,
									Name: nr.StationNameRainhamKent,
									At:   nil,
									Et:   pstr("19:55"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length:      8,
									St:          "19:52",
									DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
								},
							},
						},
						RSID: nil,
						STA:  "19:57",
						ETD:  pstr("19:59"),
						STD:  pstr("19:57"),
						SubsequentCallingPoints: []*nr.CallingPoint{
							{
								CRS:  nr.StationCodeChatham,
								Name: nr.StationNameChatham,
								At:   nil,
								Et:   pstr("20:03"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},

										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:01",
								DelayReason: pstr("This train has been delayed by a safety inspection of the track"),
							},
							{
								CRS:  nr.StationCodeRochester,
								Name: nr.StationNameRochester,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:05",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeMeopham,
								Name: nr.StationNameMeopham,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:16",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeLongfield,
								Name: nr.StationNameLongfield,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:20",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeBromleySouth,
								Name: nr.StationNameBromleySouth,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:32",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeLondonVictoria,
								Name: nr.StationNameLondonVictoria,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: []*nr.Coach{
										{
											Number:  "A1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "A2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "A3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "A4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B1",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
										{
											Number:  "B2",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Standard",
											},
										},
										{
											Number:  "B3",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   nil,
												CharData: "Accessible",
											},
										},
										{
											Number:  "B4",
											Class:   "Standard",
											Loading: nil,
											Toilet: nr.Toilet{
												Status:   pstr("Unknown"),
												CharData: "None",
											},
										},
									},
								},
								Length:      8,
								St:          "20:51",
								DelayReason: nil,
							},
						},
					},
					{
						ID:   "3108044GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeKentishTown,
							Name: nr.StationNameKentishTown,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRainhamKent,
								Name: nr.StationNameRainhamKent,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: nil,
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "TL",
							Name: "Thameslink",
						},
						Platform: 2,
						PreviousCallingPointsPerOrigin: [][]*nr.CallingPoint{
							{
								{
									CRS:  nr.StationCodeRainhamKent,
									Name: nr.StationNameRainhamKent,
									At:   nil,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: nil,
									},
									Length:      8,
									St:          "19:58",
									DelayReason: nil,
								},
							},
						},
						RSID: nil,
						STA:  "20:03",
						ETD:  pstr("On time"),
						STD:  pstr("20:04"),
						SubsequentCallingPoints: []*nr.CallingPoint{
							{
								CRS:  nr.StationCodeChatham,
								Name: nr.StationNameChatham,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:08",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeRochester,
								Name: nr.StationNameRochester,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:11",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeStrood,
								Name: nr.StationNameStrood,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:15",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeHigham,
								Name: nr.StationNameHigham,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:20",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeGravesend,
								Name: nr.StationNameGravesend,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:27",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeNorthfleet,
								Name: nr.StationNameNorthfleet,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:31",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeSwanscombe,
								Name: nr.StationNameSwanscombe,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:33",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeGreenhitheForBluewater,
								Name: nr.StationNameGreenhitheForBluewater,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:37",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeStoneCrossing,
								Name: nr.StationNameStoneCrossing,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:39",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeDartford,
								Name: nr.StationNameDartford,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:45",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeSladeGreen,
								Name: nr.StationNameSladeGreen,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:52",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeAbbeyWood,
								Name: nr.StationNameAbbeyWood,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "20:58",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodePlumstead,
								Name: nr.StationNamePlumstead,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:01",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeWoolwichArsenal,
								Name: nr.StationNameWoolwichArsenal,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:04",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeCharlton,
								Name: nr.StationNameCharlton,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:10",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeWestcombePark,
								Name: nr.StationNameWestcombePark,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:12",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeMazeHill,
								Name: nr.StationNameMazeHill,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:14",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeGreenwich,
								Name: nr.StationNameGreenwich,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:17",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeDeptford,
								Name: nr.StationNameDeptford,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:20",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeLondonBridge,
								Name: nr.StationNameLondonBridge,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:27",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeLondonBlackfriars,
								Name: nr.StationNameLondonBlackfriars,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:34",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeCityThameslink,
								Name: nr.StationNameCityThameslink,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:36",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeFarringdon,
								Name: nr.StationNameFarringdon,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:38",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeLondonStPancrasIntl,
								Name: nr.StationNameLondonStPancrasIntl,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:43",
								DelayReason: nil,
							},
							{
								CRS:  nr.StationCodeKentishTown,
								Name: nr.StationNameKentishTown,
								At:   nil,
								Et:   pstr("On time"),
								Formation: nr.Formation{
									Coaches: nil,
								},
								Length:      8,
								St:          "21:48",
								DelayReason: nil,
							},
						},
					},
				},
				PlatformAvailable: true,
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
				},
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeRainhamKent,
					LocationName: nr.StationNameRainhamKent,
					Type:         "from",
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T19:31:00.0482269Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetArrivalsAndDeparturesWithDetails(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)

			sort.Slice(tc.expectedStationBoard.Services, func(i, j int) bool {
				return tc.expectedStationBoard.Services[i].ID < tc.expectedStationBoard.Services[j].ID
			})
			sort.Slice(stationBoard.Services, func(i, j int) bool {
				return stationBoard.Services[i].ID < stationBoard.Services[j].ID
			})

			for i, expectedService := range tc.expectedStationBoard.Services {
				actualService := stationBoard.Services[i]

				assert.Equal(t, expectedService.ID, actualService.ID)
				assert.Equal(t, expectedService.Type, actualService.Type)
				assert.Equal(t, expectedService.Destination, actualService.Destination)
				assert.Equal(t, expectedService.ETA, actualService.ETA)
				assert.Equal(t, expectedService.Length, actualService.Length)
				assert.Equal(t, expectedService.Operator, actualService.Operator)
				assert.Equal(t, expectedService.Platform, actualService.Platform)
				assert.Equal(t, expectedService.RSID, actualService.RSID)
				assert.Equal(t, expectedService.STA, actualService.STA)

				assert.ElementsMatch(t, expectedService.Origins, actualService.Origins)
				assert.ElementsMatch(t, expectedService.Formation.Coaches, actualService.Formation.Coaches)

				for expi, expectedCallingPoints := range expectedService.PreviousCallingPointsPerOrigin {
					for expj, expectedCallingPoint := range expectedCallingPoints {
						actualCallingPoint := actualService.PreviousCallingPointsPerOrigin[expi][expj]
						compareCallingPoint(t, expectedCallingPoint, actualCallingPoint)
					}
				}

				for expi, expectedCallingPoint := range expectedService.SubsequentCallingPoints {
					actualCallingPoint := actualService.SubsequentCallingPoints[expi]
					compareCallingPoint(t, expectedCallingPoint, actualCallingPoint)
				}

				assert.Equal(t, expectedService.ETD, actualService.ETD)
				assert.Equal(t, expectedService.STD, actualService.STD)
			}
		})
	}
}

func TestGetArrivals(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetArrivalBoardResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Services: []*nr.Service{
					{
						ID:   "3111023GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRamsgate,
								Name: nr.StationNameRamsgate,
							},
						},
						ETA: "21:03",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: pint(0),
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: pint(6),
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: pint(5),
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: pint(20),
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: pint(5),
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: pint(9),
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: pint(0),
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: pint(0),
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform:    2,
						STA:         "20:57",
						DelayReason: pstr("This train has been delayed by a problem currently under investigation"),
					},
					{
						ID:   "3111024GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeDoverPriory,
								Name: nr.StationNameDoverPriory,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 4,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 2,
						STA:      "21:29",
					},
					{
						ID:   "3111026GLNGHMK_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRamsgate,
								Name: nr.StationNameRamsgate,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 0,
						STA:      "21:57",
					},
				},
				PlatformAvailable: true,
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
				},
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeLondonVictoria,
					LocationName: nr.StationNameLondonVictoria,
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T21:03:35.4326061Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetArrivals(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)

			sort.Slice(tc.expectedStationBoard.Services, func(i, j int) bool {
				return tc.expectedStationBoard.Services[i].ID < tc.expectedStationBoard.Services[j].ID
			})
			sort.Slice(stationBoard.Services, func(i, j int) bool {
				return stationBoard.Services[i].ID < stationBoard.Services[j].ID
			})

			for i, expectedService := range tc.expectedStationBoard.Services {
				actualService := stationBoard.Services[i]

				assert.Equal(t, expectedService.ID, actualService.ID)
				assert.Equal(t, expectedService.Type, actualService.Type)
				assert.Equal(t, expectedService.Destination, actualService.Destination)
				assert.Equal(t, expectedService.ETA, actualService.ETA)
				assert.Equal(t, expectedService.Length, actualService.Length)
				assert.Equal(t, expectedService.Operator, actualService.Operator)
				assert.Equal(t, expectedService.Platform, actualService.Platform)
				assert.Equal(t, expectedService.STA, actualService.STA)

				assert.ElementsMatch(t, expectedService.Origins, actualService.Origins)
				assert.ElementsMatch(t, expectedService.Formation.Coaches, actualService.Formation.Coaches)
			}
		})
	}
}

func TestGetArrivalsAndDepartures(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetArrivalDepartureBoardResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:          nr.StationCodeRochester,
				LocationName: nr.StationNameRochester,
				Services: []*nr.Service{
					{
						ID:   "3111026RCHT____",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeRamsgate,
								Name: nr.StationNameRamsgate,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 1,
						STA:      "22:05",
						ETD:      pstr("On time"),
						STD:      pstr("22:05"),
					},
					{
						ID:   "3090708RCHT____",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeGillinghamKent,
								Name: nr.StationNameGillinghamKent,
							},
						},
						ETA: "On time",
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 6,
						Operator: nr.Operator{
							Code: "SE",
							Name: "Southeastern",
						},
						Platform: 1,
						STA:      "22:22",
						ETD:      pstr("On time"),
						STD:      pstr("22:23"),
					},
				},
				PlatformAvailable: true,
				NRCCMessages: []string{
					"The disabled toilets are out of order at Rochester station.",
					"Severe weather to affect services in the South West of England today. More details can be found in Latest Travel News.",
				},
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeLondonVictoria,
					LocationName: nr.StationNameLondonVictoria,
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T21:54:19.4101724Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetArrivalsAndDepartures(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)

			sort.Slice(tc.expectedStationBoard.Services, func(i, j int) bool {
				return tc.expectedStationBoard.Services[i].ID < tc.expectedStationBoard.Services[j].ID
			})
			sort.Slice(stationBoard.Services, func(i, j int) bool {
				return stationBoard.Services[i].ID < stationBoard.Services[j].ID
			})

			for i, expectedService := range tc.expectedStationBoard.Services {
				actualService := stationBoard.Services[i]

				assert.Equal(t, expectedService.ID, actualService.ID)
				assert.Equal(t, expectedService.Type, actualService.Type)
				assert.Equal(t, expectedService.Destination, actualService.Destination)
				assert.Equal(t, expectedService.ETA, actualService.ETA)
				assert.Equal(t, expectedService.Length, actualService.Length)
				assert.Equal(t, expectedService.Operator, actualService.Operator)
				assert.Equal(t, expectedService.Platform, actualService.Platform)
				assert.Equal(t, expectedService.STA, actualService.STA)

				assert.ElementsMatch(t, expectedService.Origins, actualService.Origins)
				assert.ElementsMatch(t, expectedService.Formation.Coaches, actualService.Formation.Coaches)

				assert.Equal(t, expectedService.ETD, actualService.ETD)
				assert.Equal(t, expectedService.STD, actualService.STD)
			}
		})
	}
}

func TestGetDepartures(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetDepartureBoardResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:          nr.StationCodeLondonBridge,
				LocationName: nr.StationNameLondonBridge,
				Services: []*nr.Service{
					{
						ID:   "3108056LNDNBDE_",
						Type: "train",
						Destination: nr.Location{
							CRS:  nr.StationCodeGillinghamKent,
							Name: nr.StationNameGillinghamKent,
							Via:  pstr("via Greenwich & Woolwich Arsenal"),
						},
						Origins: []*nr.Location{
							{
								CRS:  nr.StationCodeKentishTown,
								Name: nr.StationNameKentishTown,
							},
						},
						Length: 8,
						Operator: nr.Operator{
							Code: "TL",
							Name: "Thameslink",
						},
						Platform: 4,
						RSID:     pstr("TL353900"),
						ETD:      pstr("22:22"),
						STD:      pstr("22:18"),
					},
				},
				PlatformAvailable: true,
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeGillinghamKent,
					LocationName: nr.StationNameGillinghamKent,
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T22:14:23.3784418Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetDepartures(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)

			sort.Slice(tc.expectedStationBoard.Services, func(i, j int) bool {
				return tc.expectedStationBoard.Services[i].ID < tc.expectedStationBoard.Services[j].ID
			})
			sort.Slice(stationBoard.Services, func(i, j int) bool {
				return stationBoard.Services[i].ID < stationBoard.Services[j].ID
			})

			for i, expectedService := range tc.expectedStationBoard.Services {
				actualService := stationBoard.Services[i]

				assert.Equal(t, expectedService.ID, actualService.ID)
				assert.Equal(t, expectedService.Type, actualService.Type)
				assert.Equal(t, expectedService.Destination, actualService.Destination)
				assert.Equal(t, expectedService.ETA, actualService.ETA)
				assert.Equal(t, expectedService.Length, actualService.Length)
				assert.Equal(t, expectedService.Operator, actualService.Operator)
				assert.Equal(t, expectedService.Platform, actualService.Platform)
				assert.Equal(t, expectedService.RSID, actualService.RSID)
				assert.Equal(t, expectedService.STA, actualService.STA)

				assert.ElementsMatch(t, expectedService.Origins, actualService.Origins)
				assert.ElementsMatch(t, expectedService.Formation.Coaches, actualService.Formation.Coaches)

				assert.Equal(t, expectedService.ETD, actualService.ETD)
				assert.Equal(t, expectedService.STD, actualService.STD)
			}
		})
	}
}

func TestGetDeparturesWithDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                  nr.CRSCode
		setupMock            func() *httptest.Server
		expectedStationBoard *nr.StationBoard
		errAssert            assert.ErrorAssertionFunc
	}{
		"Success": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetDepBoardWithDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				CRS:               nr.StationCodeLondonBridge,
				LocationName:      nr.StationNameLondonBridge,
				PlatformAvailable: true,
				Filters: nr.RequestFilters{
					CRS:          nr.StationCodeGillinghamKent,
					LocationName: nr.StationNameGillinghamKent,
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T22:27:22.5995898Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs: nr.CRSCode("test"),
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs: nr.StationCodeGillinghamKent,
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			stationBoard, err := client.GetDeparturesWithDetails(tc.crs)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.Equal(t, tc.expectedStationBoard.GeneratedAt, stationBoard.GeneratedAt.UTC())
			assert.Equal(t, tc.expectedStationBoard.Filters, stationBoard.Filters)
		})
	}
}

func TestGetFastestDepartures(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                     nr.CRSCode
		destinations            []nr.CRSCode
		setupMock               func() *httptest.Server
		expectedDeparturesBoard *nr.DeparturesBoard
		errAssert               assert.ErrorAssertionFunc
	}{
		"Success": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetFastestDeparturesResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedDeparturesBoard: &nr.DeparturesBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Departures: []*nr.Departure{
					{
						CRS: nr.StationCodeRochester,
						Service: &nr.Service{
							ID:   "3090727GLNGHMK_",
							Type: "train",
							Destination: nr.Location{
								CRS:  nr.StationCodeLondonVictoria,
								Name: nr.StationNameLondonVictoria,
							},
							Origins: []*nr.Location{
								{
									CRS:  nr.StationCodeRamsgate,
									Name: nr.StationNameRamsgate,
								},
							},
							ETA: "On time",
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							Operator: nr.Operator{
								Code: "SE",
								Name: "Southeastern",
							},
							Platform: 2,
							STA:      "23:13",
							ETD:      pstr("On time"),
							STD:      pstr("23:18"),
						},
					},
				},
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T23:08:53.6080948Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs:          nr.CRSCode("test"),
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_MissingFilterDestinations": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			departuresBoard, err := client.GetFastestDepartures(tc.crs, tc.destinations)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedDeparturesBoard.CRS, departuresBoard.CRS)
			assert.Equal(t, tc.expectedDeparturesBoard.LocationName, departuresBoard.LocationName)
			assert.Equal(t, tc.expectedDeparturesBoard.GeneratedAt, departuresBoard.GeneratedAt.UTC())
			assert.ElementsMatch(t, tc.expectedDeparturesBoard.NRCCMessages, departuresBoard.NRCCMessages)

			for i, expectedDeparture := range tc.expectedDeparturesBoard.Departures {
				actualDeparture := departuresBoard.Departures[i]

				assert.Equal(t, expectedDeparture.CRS, actualDeparture.CRS)
				assert.Equal(t, expectedDeparture.Destination, actualDeparture.Destination)
				assert.Equal(t, expectedDeparture.Service, actualDeparture.Service)
			}
		})
	}
}

func TestGetFastestDeparturesWithDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                     nr.CRSCode
		destinations            []nr.CRSCode
		setupMock               func() *httptest.Server
		expectedDeparturesBoard *nr.DeparturesBoard
		errAssert               assert.ErrorAssertionFunc
	}{
		"Success": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetFastestDeparturesWithDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedDeparturesBoard: &nr.DeparturesBoard{
				CRS:          nr.StationCodeStrood,
				LocationName: nr.StationNameStrood,
				Departures: []*nr.Departure{
					{
						CRS: nr.StationCodeGillinghamKent,
						Service: &nr.Service{
							ID:   "3108056STROOD__",
							Type: "train",
							Destination: nr.Location{
								CRS:  nr.StationCodeGillinghamKent,
								Name: nr.StationNameGillinghamKent,
							},
							Origins: []*nr.Location{
								{
									CRS:  nr.StationCodeKentishTown,
									Name: nr.StationNameKentishTown,
								},
							},
							ETA:    "On time",
							Length: 8,
							Operator: nr.Operator{
								Code: "TL",
								Name: "Thameslink",
							},
							Platform: 1,
							RSID:     pstr("TL353900"),
							STA:      "23:29",
							ETD:      pstr("On time"),
							STD:      pstr("23:29"),
							SubsequentCallingPoints: []*nr.CallingPoint{
								{
									CRS:    nr.StationCodeRochester,
									Name:   nr.StationNameRochester,
									Et:     pstr("On time"),
									Length: 8,
									St:     "23:32",
								},
								{
									CRS:    nr.StationCodeChatham,
									Name:   nr.StationNameChatham,
									Et:     pstr("On time"),
									Length: 8,
									St:     "23:36",
								},
								{
									CRS:    nr.StationCodeGillinghamKent,
									Name:   nr.StationNameGillinghamKent,
									Et:     pstr("On time"),
									Length: 8,
									St:     "23:40",
								},
							},
						},
					},
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T23:25:18.9059932Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs:          nr.CRSCode("test"),
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_MissingFilterDestinations": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			departuresBoard, err := client.GetFastestDeparturesWithDetails(tc.crs, tc.destinations)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedDeparturesBoard.CRS, departuresBoard.CRS)
			assert.Equal(t, tc.expectedDeparturesBoard.LocationName, departuresBoard.LocationName)
			assert.Equal(t, tc.expectedDeparturesBoard.GeneratedAt, departuresBoard.GeneratedAt.UTC())
			assert.ElementsMatch(t, tc.expectedDeparturesBoard.NRCCMessages, departuresBoard.NRCCMessages)

			for i, expectedDeparture := range tc.expectedDeparturesBoard.Departures {
				actualDeparture := departuresBoard.Departures[i]

				assert.Equal(t, expectedDeparture.CRS, actualDeparture.CRS)
				assert.Equal(t, expectedDeparture.Destination, actualDeparture.Destination)
				assert.Equal(t, expectedDeparture.Service, actualDeparture.Service)
			}
		})
	}
}

func TestGetNextDepartures(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                     nr.CRSCode
		destinations            []nr.CRSCode
		setupMock               func() *httptest.Server
		expectedDeparturesBoard *nr.DeparturesBoard
		errAssert               assert.ErrorAssertionFunc
	}{
		"Success": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetNextDeparturesResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedDeparturesBoard: &nr.DeparturesBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Departures: []*nr.Departure{
					{
						CRS: nr.StationCodeRainhamKent,
						Service: &nr.Service{
							ID:   "3092244GLNGHMK_",
							Type: "train",
							Destination: nr.Location{
								CRS:  nr.StationCodeFaversham,
								Name: nr.StationNameFaversham,
							},
							Origins: []*nr.Location{
								{
									CRS:  nr.StationCodeLondonVictoria,
									Name: nr.StationNameLondonVictoria,
								},
							},
							ETA: "On time",
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							Operator: nr.Operator{
								Code: "SE",
								Name: "Southeastern",
							},
							Platform: 3,
							STA:      "00:01",
							ETD:      pstr("On time"),
							STD:      pstr("00:01"),
						},
					},
				},
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T23:47:31.4401482Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs:          nr.CRSCode("test"),
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_MissingFilterDestinationsOpt": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			departuresBoard, err := client.GetNextDepartures(tc.crs, tc.destinations)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedDeparturesBoard.CRS, departuresBoard.CRS)
			assert.Equal(t, tc.expectedDeparturesBoard.LocationName, departuresBoard.LocationName)
			assert.Equal(t, tc.expectedDeparturesBoard.GeneratedAt, departuresBoard.GeneratedAt.UTC())
			assert.ElementsMatch(t, tc.expectedDeparturesBoard.NRCCMessages, departuresBoard.NRCCMessages)

			for i, expectedDeparture := range tc.expectedDeparturesBoard.Departures {
				actualDeparture := departuresBoard.Departures[i]

				assert.Equal(t, expectedDeparture.CRS, actualDeparture.CRS)
				assert.Equal(t, expectedDeparture.Service, actualDeparture.Service)
				assert.Equal(t, expectedDeparture.Destination, actualDeparture.Destination)
			}
		})
	}
}

func TestGetNextDeparturesWithDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		crs                     nr.CRSCode
		destinations            []nr.CRSCode
		setupMock               func() *httptest.Server
		expectedDeparturesBoard *nr.DeparturesBoard
		errAssert               assert.ErrorAssertionFunc
	}{
		"Success": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetNextDeparturesWithDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedDeparturesBoard: &nr.DeparturesBoard{
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Departures: []*nr.Departure{
					{
						CRS: nr.StationCodeRainhamKent,
						Service: &nr.Service{
							ID:   "3092244GLNGHMK_",
							Type: "train",
							Destination: nr.Location{
								CRS:  nr.StationCodeFaversham,
								Name: nr.StationNameFaversham,
							},
							Origins: []*nr.Location{
								{
									CRS:  nr.StationCodeLondonVictoria,
									Name: nr.StationNameLondonVictoria,
								},
							},
							ETA: "On time",
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							Operator: nr.Operator{
								Code: "SE",
								Name: "Southeastern",
							},
							Platform: 3,
							STA:      "00:01",
							ETD:      pstr("On time"),
							STD:      pstr("00:01"),
							SubsequentCallingPoints: []*nr.CallingPoint{
								{
									CRS:  nr.StationCodeRainhamKent,
									Name: nr.StationNameRainhamKent,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "C3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "C4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "D3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "D4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "00:06",
								},
								{
									CRS:  nr.StationCodeNewington,
									Name: nr.StationNameNewington,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "C3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "C4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "D3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "D4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "00:10",
								},
								{
									CRS:  nr.StationCodeSittingbourne,
									Name: nr.StationNameSittingbourne,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "C3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "C4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "D3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "D4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "00:15",
								},
								{
									CRS:  nr.StationCodeTeynham,
									Name: nr.StationNameTeynham,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "C3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "C4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "D3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "D4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "00:20",
								},
								{
									CRS:  nr.StationCodeFaversham,
									Name: nr.StationNameFaversham,
									Et:   pstr("On time"),
									Formation: nr.Formation{
										Coaches: []*nr.Coach{
											{
												Number:  "A1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "A2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "A3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "A4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "B2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "B3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "B4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "C2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "C3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "C4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D1",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
											{
												Number:  "D2",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Standard",
												},
											},
											{
												Number:  "D3",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   nil,
													CharData: "Accessible",
												},
											},
											{
												Number:  "D4",
												Class:   "Standard",
												Loading: nil,
												Toilet: nr.Toilet{
													Status:   pstr("Unknown"),
													CharData: "None",
												},
											},
										},
									},
									Length: 8,
									St:     "00:26",
								},
							},
						},
					},
				},
				NRCCMessages: []string{
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-03T23:47:15.9929307Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_BadCRS": {
			crs:          nr.CRSCode("test"),
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_MissingFilterDestinations": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{},
			setupMock: func() *httptest.Server {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
		"Fail_UnexpectedServerError": {
			crs:          nr.StationCodeGillinghamKent,
			destinations: []nr.CRSCode{nr.StationCodeLondonVictoria},
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			departuresBoard, err := client.GetNextDeparturesWithDetails(tc.crs, tc.destinations)
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedDeparturesBoard.CRS, departuresBoard.CRS)
			assert.Equal(t, tc.expectedDeparturesBoard.LocationName, departuresBoard.LocationName)
			assert.Equal(t, tc.expectedDeparturesBoard.GeneratedAt, departuresBoard.GeneratedAt.UTC())
			assert.ElementsMatch(t, tc.expectedDeparturesBoard.NRCCMessages, departuresBoard.NRCCMessages)

			for i, expectedDeparture := range tc.expectedDeparturesBoard.Departures {
				actualDeparture := departuresBoard.Departures[i]

				assert.Equal(t, expectedDeparture.CRS, actualDeparture.CRS)
				assert.Equal(t, expectedDeparture.CRS, actualDeparture.CRS)
				assert.Equal(t, expectedDeparture.Service, actualDeparture.Service)
			}
		})
	}
}

func TestGetServiceDetails(t *testing.T) {
	for name, tc := range map[string]struct {
		setupMock              func() *httptest.Server
		expectedServiceDetails *nr.ServiceDetails
		errAssert              assert.ErrorAssertionFunc
	}{
		"Success": {
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "GetServiceDetailsResponse.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedServiceDetails: &nr.ServiceDetails{
				Type: "train",
				Formation: nr.Formation{
					Coaches: []*nr.Coach{
						{
							Number:  "A1",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "A2",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Standard",
							},
						},
						{
							Number:  "A3",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Accessible",
							},
						},
						{
							Number:  "A4",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "B1",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "B2",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Standard",
							},
						},
						{
							Number:  "B3",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Accessible",
							},
						},
						{
							Number:  "B4",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "C1",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "C2",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Standard",
							},
						},
						{
							Number:  "C3",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Accessible",
							},
						},
						{
							Number:  "C4",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "D1",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
						{
							Number:  "D2",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Standard",
							},
						},
						{
							Number:  "D3",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   nil,
								CharData: "Accessible",
							},
						},
						{
							Number:  "D4",
							Class:   "Standard",
							Loading: nil,
							Toilet: nr.Toilet{
								Status:   pstr("Unknown"),
								CharData: "None",
							},
						},
					},
				},
				Length:       8,
				CRS:          nr.StationCodeGillinghamKent,
				LocationName: nr.StationNameGillinghamKent,
				Operator: nr.Operator{
					Code: "SE",
					Name: "Southeastern",
				},
				Platform: 3,
				PreviousCallingPointsPerOrigin: [][]*nr.CallingPoint{
					{
						{
							CRS:  nr.StationCodeLondonVictoria,
							Name: nr.StationNameLondonVictoria,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:10",
						},
						{
							CRS:  nr.StationCodeBromleySouth,
							Name: nr.StationNameBromleySouth,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:27",
						},
						{
							CRS:  nr.StationCodeLongfield,
							Name: nr.StationNameLongfield,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:39",
						},
						{
							CRS:  nr.StationCodeMeopham,
							Name: nr.StationNameMeopham,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:43",
						},
						{
							CRS:  nr.StationCodeRochester,
							Name: nr.StationNameRochester,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:53",
						},
						{
							CRS:  nr.StationCodeChatham,
							Name: nr.StationNameChatham,
							At:   pstr("On time"),
							Formation: nr.Formation{
								Coaches: []*nr.Coach{
									{
										Number:  "A1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "A2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "A3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "A4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "B2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "B3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "B4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "C2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "C3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "C4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D1",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
									{
										Number:  "D2",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Standard",
										},
									},
									{
										Number:  "D3",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   nil,
											CharData: "Accessible",
										},
									},
									{
										Number:  "D4",
										Class:   "Standard",
										Loading: nil,
										Toilet: nr.Toilet{
											Status:   pstr("Unknown"),
											CharData: "None",
										},
									},
								},
							},
							Length: 8,
							St:     "23:57",
						},
					},
				},
				ATA: "On time",
				ATD: "On time",
				STD: "00:01",
				STA: "00:01",
				SubsequentCallingPoints: []*nr.CallingPoint{
					{
						CRS:  nr.StationCodeRainhamKent,
						Name: nr.StationNameRainhamKent,
						At:   pstr("On time"),
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "C3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "C4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "D3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "D4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						St:     "00:06",
					},
					{
						CRS:  nr.StationCodeNewington,
						Name: nr.StationNameNewington,
						Et:   pstr("On time"),
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "C3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "C4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "D3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "D4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						St:     "00:10",
					},
					{
						CRS:  nr.StationCodeSittingbourne,
						Name: nr.StationNameSittingbourne,
						Et:   pstr("On time"),
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "C3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "C4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "D3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "D4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						St:     "00:15",
					},
					{
						CRS:  nr.StationCodeTeynham,
						Name: nr.StationNameTeynham,
						Et:   pstr("On time"),
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "C3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "C4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "D3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "D4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						St:     "00:20",
					},
					{
						CRS:  nr.StationCodeFaversham,
						Name: nr.StationNameFaversham,
						Et:   pstr("On time"),
						Formation: nr.Formation{
							Coaches: []*nr.Coach{
								{
									Number:  "A1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "A2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "A3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "A4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "B2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "B3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "B4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "C2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "C3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "C4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D1",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
								{
									Number:  "D2",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Standard",
									},
								},
								{
									Number:  "D3",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   nil,
										CharData: "Accessible",
									},
								},
								{
									Number:  "D4",
									Class:   "Standard",
									Loading: nil,
									Toilet: nr.Toilet{
										Status:   pstr("Unknown"),
										CharData: "None",
									},
								},
							},
						},
						Length: 8,
						St:     "00:26",
					},
				},
				GeneratedAt: timeFromRFC3339(t, "2024-01-04T00:07:52.5881196Z"),
			},
			errAssert: assert.NoError,
		},
		"Fail_UnexpectedServerError": {
			setupMock: func() *httptest.Server {
				data, err := os.ReadFile(filepath.Join("testdata", "UnexpectedServerError.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			errAssert: assert.Error,
		},
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			ts := tc.setupMock()
			defer ts.Close()

			client, err := nr.NewClient(
				nr.URLOpt(ts.URL),
				nr.AccessTokenOpt(testAccessToken),
			)
			require.NoError(t, err)

			serviceDetails, err := client.GetServiceDetails("test")
			tc.errAssert(t, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.expectedServiceDetails.CRS, serviceDetails.CRS)
			assert.Equal(t, tc.expectedServiceDetails.LocationName, serviceDetails.LocationName)
			assert.Equal(t, tc.expectedServiceDetails.Type, serviceDetails.Type)
			assert.Equal(t, tc.expectedServiceDetails.Length, serviceDetails.Length)
			assert.Equal(t, tc.expectedServiceDetails.Operator, serviceDetails.Operator)
			assert.Equal(t, tc.expectedServiceDetails.PreviousCallingPointsPerOrigin, serviceDetails.PreviousCallingPointsPerOrigin)
			assert.Equal(t, tc.expectedServiceDetails.ATA, serviceDetails.ATA)
			assert.Equal(t, tc.expectedServiceDetails.ATD, serviceDetails.ATD)
			assert.Equal(t, tc.expectedServiceDetails.STA, serviceDetails.STA)
			assert.Equal(t, tc.expectedServiceDetails.ETD, serviceDetails.ETD)
			assert.Equal(t, tc.expectedServiceDetails.STD, serviceDetails.STD)
			assert.Equal(t, tc.expectedServiceDetails.GeneratedAt, serviceDetails.GeneratedAt.UTC())
			assert.ElementsMatch(t, tc.expectedServiceDetails.Formation.Coaches, serviceDetails.Formation.Coaches)

			for expi, expectedCallingPoint := range tc.expectedServiceDetails.SubsequentCallingPoints {
				actualCallingPoint := serviceDetails.SubsequentCallingPoints[expi]
				compareCallingPoint(t, expectedCallingPoint, actualCallingPoint)
			}
		})
	}
}

func compareCallingPoint(
	t *testing.T,
	expectedCallingPoint *nr.CallingPoint,
	actualCallingPoint *nr.CallingPoint,
) {
	assert.Equal(t, expectedCallingPoint.CRS, actualCallingPoint.CRS)
	assert.Equal(t, expectedCallingPoint.Name, actualCallingPoint.Name)
	assert.Equal(t, expectedCallingPoint.Length, actualCallingPoint.Length)
	assert.Equal(t, expectedCallingPoint.St, actualCallingPoint.St)
	assert.Equal(t, expectedCallingPoint.At, actualCallingPoint.At)
	assert.Equal(t, expectedCallingPoint.Et, actualCallingPoint.Et)
	assert.Equal(t, expectedCallingPoint.DelayReason, actualCallingPoint.DelayReason)
	require.ElementsMatch(t, expectedCallingPoint.Formation.Coaches, actualCallingPoint.Formation.Coaches)
}

func timeFromRFC3339(t *testing.T, ts string) time.Time {
	pts, err := time.Parse(time.RFC3339, ts)
	require.NoError(t, err)

	return pts
}

func pstr(s string) *string {
	return &s
}

func pint(i int) *int {
	return &i
}
