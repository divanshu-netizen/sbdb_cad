package sbdb_cad

type SB struct {
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

type SmallBodyOptions struct {
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
	Neo                 bool   `json:"neo"`
	Kind                string `json:"kind"`
	Spk                 string `json:"spk"`
	Designation         string `json:"designation"`
	Body                string `json:"body"`
	Sort                string `json:"sort"`
	Limit               string `json:"limit"`
	FullName            bool   `json:"fullName"`
}

type SbFinder interface {
	FindSBBy(sbo SmallBodyOptions) ([]SB, error)
}

type sbService struct {
	BaseUrl string
	Getter
	Mapper
	QueryStringBuilder
}

func NewSBService() *sbService {
	return &sbService{
		BaseUrl:            "https://ssd-api.jpl.nasa.gov/cad.api?",
		Getter:             new(Requester),
		Mapper:             NewSBMapper(),
		QueryStringBuilder: NewQueryBuilder(),
	}
}

func (ss *sbService) FindSBBy(sbo SmallBodyOptions) ([]SB, error) {
	res, err := ss.Getter.Get(ss.BaseUrl + ss.QueryStringBuilder.Build(&sbo))
	if err != nil {
		return nil, err
	}

	return ss.Mapper.Map(res)
}
