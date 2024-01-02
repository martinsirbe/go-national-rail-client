package nationalrail_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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
				data, err := os.ReadFile(filepath.Join("testdata", "GetStationBoard_Success.xml"))
				require.NoError(t, err)

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					_, err := w.Write(data)
					require.NoError(t, err)
				}))
				return ts
			},
			expectedStationBoard: &nr.StationBoard{
				GeneratedAt:       timeFromRFC3339(t, "2024-01-01T15:53:39.2335168Z"),
				LocationName:      "Gillingham (Kent)",
				CRS:               "GLM",
				PlatformAvailable: true,
				TrainServices: []*nr.TrainService{
					{
						STA:          pstr("15:57"),
						ETA:          pstr("On time"),
						Platform:     pstr("2"),
						Operator:     "Southeastern",
						OperatorCode: "SE",
						ServiceType:  "train",
						ServiceID:    "3033940GLNGHMK_",
						Origin: &nr.Location{
							Name: "Ramsgate",
							CRS:  "RAM",
							Type: "origin",
						},
						Destination: &nr.Location{
							Name: "London Cannon Street",
							CRS:  "CST",
							Type: "destination",
						},
						PreviousCallingPoints: []*nr.Location{
							{
								Name: "Ramsgate",
								CRS:  "RAM",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("14:52"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Dumpton Park",
								CRS:  "DMP",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("14:55"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Broadstairs",
								CRS:  "BSR",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("14:58"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Margate",
								CRS:  "MAR",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:04"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Westgate-on-Sea",
								CRS:  "WGA",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:08"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Birchington-on-Sea",
								CRS:  "BCH",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:12"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Herne Bay",
								CRS:  "HNB",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:21"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Chestfield & Swalecliffe",
								CRS:  "CSW",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:24"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Whitstable",
								CRS:  "WHI",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:28"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Faversham",
								CRS:  "FAV",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:37"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Sittingbourne",
								CRS:  "SIT",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:45"),
								At:   pstr("On time"),
								Et:   nil,
							},
							{
								Name: "Rainham (Kent)",
								CRS:  "RAI",
								Type: "previous calling point",
								Via:  nil,
								St:   pstr("15:52"),
								At:   pstr("On time"),
								Et:   nil,
							},
						},
					},
					{
						STA:          pstr("16:03"),
						ETA:          pstr("On time"),
						Platform:     pstr("2"),
						Operator:     "Thameslink",
						OperatorCode: "TL",
						ServiceType:  "train",
						ServiceID:    "3034907GLNGHMK_",
						Origin: &nr.Location{
							Name: "Rainham (Kent)",
							CRS:  "RAI",
							Type: "origin",
						},
						Destination: &nr.Location{
							Name: "Kentish Town",
							CRS:  "KTN",
							Type: "destination",
						},
						PreviousCallingPoints: []*nr.Location{
							{
								Name: "Rainham (Kent)",
								CRS:  "RAI",
								Type: "previous calling point",
								St:   pstr("15:58"),
								Et:   pstr("On time"),
							},
						},
					},
				},
				NRCCMessages: []string{
					"The ticket office is currently closed at this station.",
					"The lifts are out of order between Platform 3 and the overbridge at Gillingham Kent station.",
					"No Southeastern services will run to/from London Victoria from Saturday 23 December until Monday 1 January. More details can be found in Latest Travel News.",
				},
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

			assert.Equal(t, tc.expectedStationBoard.LocationName, stationBoard.LocationName)
			assert.Equal(t, tc.expectedStationBoard.CRS, stationBoard.CRS)
			assert.Equal(t, tc.expectedStationBoard.PlatformAvailable, stationBoard.PlatformAvailable)
			assert.ElementsMatch(t, tc.expectedStationBoard.TrainServices, stationBoard.TrainServices)
			assert.ElementsMatch(t, tc.expectedStationBoard.NRCCMessages, stationBoard.NRCCMessages)
		})
	}
}

func timeFromRFC3339(t *testing.T, ts string) time.Time {
	pts, err := time.Parse(time.RFC3339, ts)
	require.NoError(t, err)

	return pts
}

func pstr(s string) *string {
	return &s
}
