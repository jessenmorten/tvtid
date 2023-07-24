package tvtid

import (
	"net/http"
)

var (
	baseUrl = "https://tvtid-api.api.tv2.dk/api/tvtid/v1"
)

func NewDefaultClient() TvTidClient {
	return NewClient(http.DefaultClient, baseUrl)
}

func NewClient(httpClient httpClient, baseUrl string) TvTidClient {
	return &tvTidClient{httpClient, baseUrl}
}

type TvTidClient interface {
	GetChannels() ([]Channel, error)
}
