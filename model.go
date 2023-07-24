package tvtid

type Channel struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	IconUrl    string `json:"icon"`
	LogoUrl    string `json:"logo"`
	SvgLogoUrl string `json:"svgLogo"`
	Sort       int    `json:"sort"`
	Language   string `json:"language"`
}

type getChannelsResponse struct {
	Channels []Channel `json:"channels"`
}
