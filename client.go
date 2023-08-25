package tvtid

import (
	"encoding/json"
	"errors"
	"io"
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

	programs := programsResponse[0].Programs

	for i := range programs {
		program := &programs[i]
		program.StartTime = time.Unix(program.StartTimeUnix, 0).In(apiLocation)
		program.StopTime = time.Unix(program.StopTimeUnix, 0).In(apiLocation)
	}

	return programs, nil
}

func (c *tvTidClient) GetProgramDetails(channelId string, programId string) (*ProgramDetails, error) {
	url := c.baseUrl + "/schedules/channels/" + channelId + "/programs/" + programId
	programDetailsResponse := getProgramDetailsResponse{}
	err := c.getFromJson(url, &programDetailsResponse)

	if err != nil {
		return nil, err
	}

	return &programDetailsResponse.Program, nil
}

func (c *tvTidClient) getFromJson(url string, v interface{}) error {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Accept", "application/json")
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &v)
	return err
}
