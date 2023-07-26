package tvtid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestGetChannelsUrl(t *testing.T) {
	mock := newMock(nil, channelsResponse)
	client := NewClient(mock, "localhost")
	expectedUrl := "localhost/schedules/channels"

	_, err := client.GetChannels()

	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}

	if mock.url != expectedUrl {
		t.Errorf("want %s, got %s", expectedUrl, mock.url)
	}
}

func TestGetChannelsDeserialize(t *testing.T) {
	mock := newMock(nil, channelsResponse)
	client := NewClient(mock, "localhost")
	expectedValues := "1, 2, 3, 4, 5, 6, 7"

	channels, err := client.GetChannels()

	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}

	if len(channels) != 1 {
		t.Errorf("want %d, got %d", 1, len(channels))
	}

	actualValues := fmt.Sprintf("%s, %s, %s, %s, %s, %d, %s",
		channels[0].Id,
		channels[0].Title,
		channels[0].IconUrl,
		channels[0].LogoUrl,
		channels[0].SvgLogoUrl,
		channels[0].Sort,
		channels[0].Language)

	if actualValues != expectedValues {
		t.Errorf("want %s, got %s", expectedValues, actualValues)
	}
}

func TestGetProgramsUrl(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")
	expectedUrl := "localhost/epg/dayviews/2023-12-24?ch=1"

	_, err := client.GetPrograms("1", date)

	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}

	if mock.url != expectedUrl {
		t.Errorf("want %s, got %s", expectedUrl, mock.url)
	}
}

func TestGetTodaysProgramsUrl(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date := time.Now()
	expectedUrl := fmt.Sprintf("localhost/epg/dayviews/%v?ch=1", date.Format("2006-01-02"))

	_, err := client.GetPrograms("1", date)

	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}

	if mock.url != expectedUrl {
		t.Errorf("want %s, got %s", expectedUrl, mock.url)
	}
}

func TestGetProgramsDeserialize(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")
	expectedValues := "1, 2, 3, 4, true, 5, true, true, true, 6"

	programs, err := client.GetPrograms("1", date)

	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}

	if len(programs) != 1 {
		t.Errorf("want %d, got %d", 1, len(programs))
	}

	actualValues := fmt.Sprintf("%s, %d, %d, %s, %v, %d, %v, %v, %v, %s",
		programs[0].Id,
		programs[0].StartTimeUnix,
		programs[0].StopTimeUnix,
		programs[0].Title,
		programs[0].AvailableAsVod,
		programs[0].ProgramPartIndex,
		programs[0].Live,
		programs[0].Premiere,
		programs[0].Rerun,
		programs[0].Categories[0])

	if actualValues != expectedValues {
		t.Errorf("want %s, got %s", expectedValues, actualValues)
	}
}

type mockHttpClient struct {
	response string
	err      error
	url      string
}

func newMock(err error, response string) *mockHttpClient {
	return &mockHttpClient{
		response,
		err,
		"localhost",
	}
}

func (m *mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	m.url = req.URL.String()
	return &http.Response{
		Body: ioutil.NopCloser(strings.NewReader(m.response)),
	}, m.err
}

var (
	channelsResponse = "{\"channels\": [{\"id\": \"1\", \"title\": \"2\", \"icon\": \"3\", \"logo\": \"4\", \"svgLogo\": \"5\", \"sort\": 6, \"language\": \"7\"}]}"
	programsResponse = "[{\"id\": \"1\", \"programs\": [{\"id\": \"1\", \"start\": 2, \"stop\": 3, \"title\": \"4\", \"availableAsVod\": true, \"programPartIndex\": 5, \"live\": true, \"premiere\": true, \"rerun\": true, \"categories\": [\"6\"]}]}]"
)
