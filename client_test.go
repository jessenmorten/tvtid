package tvtid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetChannelsUrl(t *testing.T) {
	mock := newMock(nil, channelsResponse)
	client := NewClient(mock, "localhost")
	expectedUrl := "localhost/schedules/channels"

	_, err := client.GetChannels()

	assert.Nil(t, err)
	assert.Equal(t, mock.url, expectedUrl)
}

func TestGetChannelsDeserialize(t *testing.T) {
	mock := newMock(nil, channelsResponse)
	client := NewClient(mock, "localhost")

	channels, err := client.GetChannels()

	assert.Nil(t, err)
	assert.Equal(t, len(channels), 1)
	assert.Equal(t, channels[0].Id, "1")
	assert.Equal(t, channels[0].Title, "2")
	assert.Equal(t, channels[0].IconUrl, "3")
	assert.Equal(t, channels[0].LogoUrl, "4")
	assert.Equal(t, channels[0].SvgLogoUrl, "5")
	assert.Equal(t, channels[0].Sort, 6)
	assert.Equal(t, channels[0].Language, "7")
}

func TestGetProgramsUrl(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")
	expectedUrl := "localhost/epg/dayviews/2023-12-24?ch=1"

	_, err := client.GetPrograms("1", date)

	assert.Nil(t, err)
	assert.Equal(t, mock.url, expectedUrl)
}

func TestGetTodaysProgramsUrl(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date := time.Now()
	expectedUrl := fmt.Sprintf("localhost/epg/dayviews/%v?ch=1", date.Format("2006-01-02"))

	_, err := client.GetPrograms("1", date)

	assert.Nil(t, err)
	assert.Equal(t, mock.url, expectedUrl)
}

func TestGetProgramsDeserialize(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")

	programs, err := client.GetPrograms("1", date)

	assert.Nil(t, err)
	assert.Equal(t, len(programs), 1)
	assert.Equal(t, programs[0].Id, "1")
	assert.Equal(t, programs[0].StartTimeUnix, int64(2))
	assert.Equal(t, programs[0].StopTimeUnix, int64(3))
	assert.Equal(t, programs[0].Title, "4")
	assert.Equal(t, programs[0].AvailableAsVod, true)
	assert.Equal(t, programs[0].ProgramPartIndex, 5)
	assert.Equal(t, programs[0].Live, true)
	assert.Equal(t, programs[0].Premiere, true)
	assert.Equal(t, programs[0].Rerun, true)
	assert.Equal(t, programs[0].Categories, []string{"6"})
	assert.Equal(t, programs[0].StartTime.Location(), apiLocation)
	assert.Equal(t, programs[0].StopTime.Location(), apiLocation)
	assert.Equal(t, programs[0].StartTime.Unix(), programs[0].StartTimeUnix)
	assert.Equal(t, programs[0].StopTime.Unix(), programs[0].StopTimeUnix)
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
