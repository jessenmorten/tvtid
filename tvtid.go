package tvtid

import (
	"net/http"
	"time"
)

var (
	baseUrl        = "https://tvtid-api.api.tv2.dk/api/tvtid/v1"
	apiLocation, _ = time.LoadLocation("Europe/Copenhagen")
)

func NewDefaultClient() TvTidClient {
	return NewClient(http.DefaultClient, baseUrl)
}

func NewClient(httpClient httpClient, baseUrl string) TvTidClient {
	return &tvTidClient{httpClient, baseUrl}
}

type TvTidClient interface {
	GetChannels() ([]Channel, error)
	GetPrograms(channelId string, date time.Time) ([]Program, error)
	GetTodaysPrograms(channelId string) ([]Program, error)
}
