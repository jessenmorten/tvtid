package tvtid

import "time"

type Channel struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	IconUrl    string `json:"icon"`
	LogoUrl    string `json:"logo"`
	SvgLogoUrl string `json:"svgLogo"`
	Sort       int    `json:"sort"`
	Language   string `json:"language"`
}

type Program struct {
	Id               string   `json:"id"`
	StartTimeUnix    int64    `json:"start"`
	StopTimeUnix     int64    `json:"stop"`
	Title            string   `json:"title"`
	AvailableAsVod   bool     `json:"availableAsVod"`
	ProgramPartIndex int      `json:"programPartIndex"`
	Live             bool     `json:"live"`
	Premiere         bool     `json:"premiere"`
	Rerun            bool     `json:"rerun"`
	Categories       []string `json:"categories"`
	StartTime        time.Time
	StopTime         time.Time
}

type ProgramDetails struct {
	Id                string           `json:"id"`
	Url               string           `json:"url"`
	SeriesId          string           `json:"seriesId"`
	Title             string           `json:"title"`
	Categories        []string         `json:"categories"`
	Description       string           `json:"desc"`
	OrgiginalTitle    string           `json:"orgTitle"`
	ProductionYear    int              `json:"prodYear"`
	ProductionCountry string           `json:"prodCountry"`
	Teaser            string           `json:"teaser"`
	Audio             string           `json:"audio"`
	TtvTexted         bool             `json:"ttvTexted"`
	ParentalGuidance  ParentalGuidance `json:"parentalGuidance"`
}

type ParentalGuidance struct {
	MinimumAge int `json:"minimumAge"`
}

type getChannelsResponse struct {
	Channels []Channel `json:"channels"`
}

type getProgramsResponse struct {
	ChannelId string    `json:"id"`
	Programs  []Program `json:"programs"`
}

type getProgramDetailsResponse struct {
	Program ProgramDetails `json:"program"`
}
