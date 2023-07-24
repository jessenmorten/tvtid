package tvtid

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestGetChannelsUrl(t *testing.T) {
	mock := newMock(nil, "{\"channels\": [{\"id\": \"tv2\", \"title\": \"TV 2\"}]}")
	client := NewClient(mock, "localhost")
	expectedUrl := "localhost/schedules/channels"

	_, _ = client.GetChannels()

	if mock.url != expectedUrl {
		t.Errorf("want %s, got %s", expectedUrl, mock.url)
	}
}

func TestGetChannelsDeserialize(t *testing.T) {
	response := "{\"channels\": [{\"id\": \"1\", \"title\": \"2\", \"icon\": \"3\", \"logo\": \"4\", \"svgLogo\": \"5\", \"sort\": 6, \"language\": \"7\"}]}"
	mock := newMock(nil, response)
	client := NewClient(mock, "localhost")
	expectedValues := "1, 2, 3, 4, 5, 6, 7"

	channels, _ := client.GetChannels()

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
