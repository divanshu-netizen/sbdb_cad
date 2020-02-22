package neos

type Neo struct {
	Des       string `json:"destination"`
	Orbit_id  string `json:"orbitId"`
	Jd        string `json:"closeApproachJd"`
	Cd        string `json:"closeApproachCd"`
	Dist      string `json:"distance"`
	Dist_max  string `json:"distanceMax"`
	Dist_min  string `json:"distanceMin"`
	V_rel     string `json:"relativeVelocity"`
	V_inf     string `json:"inferredVelocity"`
	T_sigma_f string `json:"threeSigma"`
	Body      string `json:"body"`
	H         string `json:"h"`
	FullName  string `json:"fullName"`
}

type NeoQueryOptions struct {
	DateMin             string `json:"dateMin"`
	DateMax             string `json:"dateMax"`
	DistanceMin         string `json:"distanceMin"`
	DistanceMax         string `json:"distanceMax"`
	HMin                string `json:"hMin"`
	HMax                string `json:"hMax"`
	VelocityInfMax      string `json:"velocityInfMax"`
	VelocityInfMin      string `json:"velocityInfMin"`
	VelocityRelativeMax string `json:"velocityRelativeMax"`
	VelocityRelativeMin string `json:"velocityRelativeMin"`
	Class               string `json:"class"`
	Pha                 bool   `json:"pha"`
	Nea                 bool   `json:"nea"`
	Comet               bool   `json:"comet"`
	NeaComet            bool   `json:"neaComet"`
	Neo                 bool   `json:"neos"`
	Kind                string `json:"kind"`
	Spk                 string `json:"spk"`
	Designation         string `json:"designation"`
	Body                string `json:"body"`
	Sort                string `json:"sort"`
	Limit               string `json:"limit"`
	FullName            bool   `json:"fullName"`
}

type NeoFinder interface {
	FindNeoBy(nqo NeoQueryOptions) ([]Neo, error)
}

type neoService struct {
	BaseUrl string
	Getter
	Mapper
	QueryStringBuilder
}

func NewNeoService() *neoService {
	return &neoService{
		BaseUrl:            "https://ssd-api.jpl.nasa.gov/cad.api?",
		Getter:             new(Requester),
		Mapper:             NewNeoMapper(),
		QueryStringBuilder: NewQueryBuilder(),
	}
}

func (ns *neoService) FindNeoBy(nqo NeoQueryOptions) ([]Neo, error) {
	res, err := ns.Getter.Get(ns.BaseUrl + ns.QueryStringBuilder.Build(&nqo))
	if err != nil {
		return nil, err
	}

	return ns.Mapper.Map(res)
}
