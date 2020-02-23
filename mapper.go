package sbdb

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type NeoResponse struct {
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
	neoRes := new(NeoResponse)

	err := json.NewDecoder(res.Body).Decode(neoRes)
	if err != nil {
		return nil, err
	}

	if neoRes.Count == "0" {
		return nil, errors.New("no results were found for this search")
	}

	return sb.mapNeoResToNeo(neoRes)
}

func (sb *SBMapper) mapNeoResToNeo(neoRes *NeoResponse) ([]SB, error) {
	var neos []SB

	for _, res := range neoRes.Data {
		n, err := sb.mapNeoResArrayToStruct(res, neoRes.Fields)
		if err != nil {
			return nil, err
		}

		neos = append(neos, *n)
	}

	return neos, nil
}

func (sb *SBMapper) mapNeoResArrayToStruct(res []string, fields []string) (*SB, error) {
	mappedNeo := make(map[string]string)
	for i, field := range res {
		mappedNeo[fields[i]] = field
	}

	result := &SB{}
	err := mapstructure.Decode(mappedNeo, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
