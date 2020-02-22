package services

import "news/internal/httprequest"

type Neo struct {
	des string `json:"des"`
	orbitId string `json:"orbitId"`
	closeApproachJd string `json:"closeApproachJd"`
	closeApproachCd string `json:"close_approach_cd"`
	distance string `json:"distance"`
	distanceMin string `json:"distanceMin"`
	distanceMax string `json:"distanceMax"`
	relativeVelocity string `json:"relativeVelocity"`
	inferredVelocity string `json:"inferredVelocity"`
	threeSigma string `json:"threeSigma"`
	body string `json:"body"`
	h string `json:"h"`
	fullName string `json:"fullName"`
}

type NeoGetter interface {
	RequestMostRecent() *Neo
}

type NeoService struct {
	BaseUrl string
	Neos map[string]Neo
	httprequest.Getter
	httprequest.Mapper
}

func NewNeoService() *NeoService {
	return &NeoService {
		BaseUrl: "https://ssd-api.jpl.nasa.gov/cad.api",
		Neos:   make(map[string]Neo),
		Getter: new(httprequest.Requester),
	}
}

func (ns *NeoService) RequestDefault() ([]Neo, error) {
	res, err := ns.Getter.Get(ns.BaseUrl)
	if err != nil {
		return nil, err
	}

	return ns.Mapper.Map(res)
}
