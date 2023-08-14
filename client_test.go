package tvtid

import (
	"io"
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
	assert.Equal(t, expectedUrl, mock.url)
}

func TestGetChannelsDeserialize(t *testing.T) {
	mock := newMock(nil, channelsResponse)
	client := NewClient(mock, "localhost")

	channels, err := client.GetChannels()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(channels))
	assert.Equal(t, "1", channels[0].Id)
	assert.Equal(t, "2", channels[0].Title)
	assert.Equal(t, "3", channels[0].IconUrl)
	assert.Equal(t, "4", channels[0].LogoUrl)
	assert.Equal(t, "5", channels[0].SvgLogoUrl)
	assert.Equal(t, 6, channels[0].Sort)
	assert.Equal(t, "7", channels[0].Language)
}

func TestGetProgramsUrl(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")
	expectedUrl := "localhost/epg/dayviews/2023-12-24?ch=1"

	_, err := client.GetPrograms("1", date)

	assert.Nil(t, err)
	assert.Equal(t, expectedUrl, mock.url)
}

func TestGetProgramsDeserialize(t *testing.T) {
	mock := newMock(nil, programsResponse)
	client := NewClient(mock, "localhost")
	date, _ := time.Parse("2006-01-02", "2023-12-24")

	programs, err := client.GetPrograms("1", date)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(programs))
	assert.Equal(t, "1", programs[0].Id)
	assert.Equal(t, int64(2), programs[0].StartTimeUnix)
	assert.Equal(t, int64(3), programs[0].StopTimeUnix)
	assert.Equal(t, "4", programs[0].Title)
	assert.Equal(t, true, programs[0].AvailableAsVod)
	assert.Equal(t, 5, programs[0].ProgramPartIndex)
	assert.Equal(t, true, programs[0].Live)
	assert.Equal(t, true, programs[0].Premiere)
	assert.Equal(t, true, programs[0].Rerun)
	assert.Equal(t, []string{"6"}, programs[0].Categories)
	assert.Equal(t, apiLocation, programs[0].StartTime.Location())
	assert.Equal(t, apiLocation, programs[0].StopTime.Location())
	assert.Equal(t, programs[0].StartTimeUnix, programs[0].StartTime.Unix())
	assert.Equal(t, programs[0].StopTimeUnix, programs[0].StopTime.Unix())
}

func TestGetProgramDetailsUrl(t *testing.T) {
	mock := newMock(nil, programDetailsResponse)
	client := NewClient(mock, "localhost")

	_, err := client.GetProgramDetails("1", "2")

	assert.Nil(t, err)
	assert.Equal(t, "localhost/schedules/channels/1/programs/2", mock.url)
}

func TestGetProgramDetailsDeserialize(t *testing.T) {
	mock := newMock(nil, programDetailsResponse)
	client := NewClient(mock, "localhost")

	program, err := client.GetProgramDetails("1", "2")

	assert.Nil(t, err)
	assert.Equal(t, "1", program.Id)
	assert.Equal(t, "2", program.Url)
	assert.Equal(t, "3", program.SeriesId)
	assert.Equal(t, "4", program.Title)
	assert.Equal(t, []string{"6"}, program.Categories)
	assert.Equal(t, "7", program.Description)
	assert.Equal(t, "8", program.OrgiginalTitle)
	assert.Equal(t, 9, program.ProductionYear)
	assert.Equal(t, "10", program.ProductionCountry)
	assert.Equal(t, "11", program.Teaser)
	assert.Equal(t, "12", program.Audio)
	assert.Equal(t, true, program.TtvTexted)
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
		Body: io.NopCloser(strings.NewReader(m.response)),
	}, m.err
}

var (
	channelsResponse       = "{\"channels\": [{\"id\": \"1\", \"title\": \"2\", \"icon\": \"3\", \"logo\": \"4\", \"svgLogo\": \"5\", \"sort\": 6, \"language\": \"7\"}]}"
	programsResponse       = "[{\"id\": \"1\", \"programs\": [{\"id\": \"1\", \"start\": 2, \"stop\": 3, \"title\": \"4\", \"availableAsVod\": true, \"programPartIndex\": 5, \"live\": true, \"premiere\": true, \"rerun\": true, \"categories\": [\"6\"]}]}]"
	programDetailsResponse = "{\"program\": {\"id\": \"1\", \"url\": \"2\", \"seriesId\": \"3\", \"title\": \"4\", \"categories\": [\"6\"], \"desc\": \"7\", \"orgTitle\": \"8\", \"prodYear\": 9, \"prodCountry\": \"10\", \"teaser\": \"11\", \"audio\": \"12\", \"ttvTexted\": true}}"
)
