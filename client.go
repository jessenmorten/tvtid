package tvtid

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type tvTidClient struct {
	httpClient httpClient
	baseUrl    string
}

func (c *tvTidClient) GetChannels() ([]Channel, error) {
	url := c.baseUrl + "/schedules/channels"
	channelsResponse := getChannelsResponse{}
	err := c.getFromJson(url, &channelsResponse)
	return channelsResponse.Channels, err
}

func (c *tvTidClient) GetTodaysPrograms(channelId string) ([]Program, error) {
	return c.GetPrograms(channelId, time.Now())
}

func (c *tvTidClient) GetPrograms(channelId string, date time.Time) ([]Program, error) {
	url := c.baseUrl + "/epg/dayviews/" + date.Format("2006-01-02") + "?ch=" + channelId
	programsResponse := []getProgramsResponse{}
	err := c.getFromJson(url, &programsResponse)

	if err != nil {
		return nil, err
	}

	if len(programsResponse) == 0 {
		return []Program{}, nil
	}

	if len(programsResponse) > 1 {
		return nil, errors.New("Unexpected response from server")
	}

	return programsResponse[0].Programs, nil
}

func (c *tvTidClient) getFromJson(url string, v interface{}) error {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return errors.New("Error creating request: " + err.Error())
	}

	request.Header.Add("Accept", "application/json")

	response, err := c.httpClient.Do(request)

	if err != nil {
		return errors.New("Error sending request: " + err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return errors.New("Error reading response: " + err.Error())
	}

	err = json.Unmarshal(body, &v)

	if err != nil {
		return errors.New("Error parsing response: " + err.Error())
	}

	return nil
}
