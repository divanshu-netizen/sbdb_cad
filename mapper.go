package sbdb_cad

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type SBResponse struct {
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

type SBDecoder struct{}

func (sd *SBDecoder) Decode(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return err
	}

	return nil
}

type Mapper interface {
	Map(response *http.Response) ([]SB, error)
}

type SBMapper struct {
	Decoder
}

func NewSBMapper() *SBMapper {
	return &SBMapper{
		Decoder: new(SBDecoder),
	}
}

func (sb *SBMapper) Map(res *http.Response) ([]SB, error) {
	sbRes := new(SBResponse)

	err := json.NewDecoder(res.Body).Decode(sbRes)
	if err != nil {
		return nil, err
	}

	if sbRes.Count == "0" {
		return nil, errors.New("no results were found for this search")
	}

	return sb.mapSBResToSB(sbRes)
}

func (sb *SBMapper) mapSBResToSB(sbRes *SBResponse) ([]SB, error) {
	var sbs []SB

	for _, res := range sbRes.Data {
		s, err := sb.mapSBResArrayToStruct(res, sbRes.Fields)
		if err != nil {
			return nil, err
		}

		sbs = append(sbs, *s)
	}

	return sbs, nil
}

func (sb *SBMapper) mapSBResArrayToStruct(res []string, fields []string) (*SB, error) {
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
