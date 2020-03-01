package sbdb_cad

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type SbCADResponse struct {
	Signature struct {
		Version string `json:"version"`
		Source  string `json:"source"`
	} `json:"signature"`
	Count  string     `json:"count"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
}

type Decoder interface {
	Decode(input interface{}, output interface{}) error
}

type SbCADDecoder struct{}

func (sd *SbCADDecoder) Decode(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return err
	}

	return nil
}

type Mapper interface {
	Map(response *http.Response) ([]SB, error)
}

type SbCADMapper struct {
	Decoder
}

func NewSbCADMapper() *SbCADMapper {
	return &SbCADMapper{
		Decoder: new(SbCADDecoder),
	}
}

func (sb *SbCADMapper) Map(res *http.Response) ([]SB, error) {
	sbRes := new(SbCADResponse)

	err := json.NewDecoder(res.Body).Decode(sbRes)
	if err != nil {
		return nil, err
	}

	if sbRes.Count == "0" {
		return nil, errors.New("no results were found for this search")
	}

	return sb.mapSBCADResToSBCAD(sbRes)
}

func (sb *SbCADMapper) mapSBCADResToSBCAD(sbRes *SbCADResponse) ([]SB, error) {
	var sbs []SB

	for _, res := range sbRes.Data {
		s, err := sb.mapSbCADResArrayToStruct(res, sbRes.Fields)
		if err != nil {
			return nil, err
		}

		sbs = append(sbs, *s)
	}

	return sbs, nil
}

func (sb *SbCADMapper) mapSbCADResArrayToStruct(res []string, fields []string) (*SB, error) {
	mappedSB := make(map[string]string)
	for i, field := range res {
		mappedSB[fields[i]] = field
	}

	result := &SB{}
	err := sb.Decoder.Decode(mappedSB, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
