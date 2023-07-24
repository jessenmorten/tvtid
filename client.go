package tvtid

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
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
