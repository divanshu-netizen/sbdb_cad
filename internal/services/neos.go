package services

import (
	"news/internal/httprequest"
)

type Neo struct {
	Destination      string `json:"destination"`
	OrbitId          string `json:"orbitId"`
	CloseApproachJd  string `json:"closeApproachJd"`
	CloseApproachCd  string `json:"closeApproachCd"`
	Distance         string `json:"distance"`
	DistanceMin      string `json:"distanceMin"`
	DistanceMax      string `json:"distanceMax"`
	RelativeVelocity string `json:"relativeVelocity"`
	InferredVelocity string `json:"inferredVelocity"`
	ThreeSigma       string `json:"threeSigma"`
	Body             string `json:"body"`
	H                string `json:"h"`
	FullName         string `json:"fullName"`
}

type NeoQueryOptions struct {
	DateMin             string `json:"dateMin"`
	DateMax             string `json:"dateMax"`
	DistanceMin         string `json:"distanceMin"`
	DistanceMax         string `json:"distanceMax"`
	HMin                int    `json:"hMin"`
	HMax                int    `json:"hMax"`
	VelocityInfMax      int    `json:"velocityInfMax"`
	VelocityInfMin      int    `json:"velocityInfMin"`
	VelocityRelativeMax int    `json:"velocityRelativeMax"`
	VelocityRelativeMin int    `json:"velocityRelativeMin"`
	Class               string `json:"class"`
	Pha                 bool   `json:"pha"`
	Nea                 bool   `json:"nea"`
	Comet               bool   `json:"comet"`
	NeaComet            bool   `json:"neaComet"`
	Neo                 bool   `json:"neo"`
	Kind                string `json:"kind"`
	Spk                 int    `json:"spk"`
	Designation         string `json:"designation"`
	Body                string `json:"body"`
	Sort                string `json:"sort"`
	Limit               string `json:"limit"`
	FullName            bool   `json:"fullName"`
}

type NeoGetter interface {
	FindNeoBy(nqo NeoQueryOptions) ([]Neo, error)
}

type NeoService struct {
	BaseUrl string
	Neos    map[string]Neo
	httprequest.Getter
	Mapper
	QueryStringBuilder
}

func NewNeoService() *NeoService {
	return &NeoService{
		BaseUrl:            "https://ssd-api.jpl.nasa.gov/cad.api?",
		Neos:               make(map[string]Neo),
		Getter:             new(httprequest.Requester),
		Mapper:             NewNeoMapper(),
		QueryStringBuilder: NewQueryBuilder(),
	}
}

func (ns *NeoService) FindNeoBy(nqo NeoQueryOptions) ([]Neo, error) {
	println(ns.BaseUrl + ns.QueryStringBuilder.Build(&nqo))
	res, err := ns.Getter.Get(ns.BaseUrl + ns.QueryStringBuilder.Build(&nqo))
	if err != nil {
		return nil, err
	}

	return ns.Mapper.Map(res)
}
