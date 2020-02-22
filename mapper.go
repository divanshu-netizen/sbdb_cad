package neos

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

type NeoDecoder struct{}

func (nd *NeoDecoder) Decode(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return err
	}

	return nil
}

type Mapper interface {
	Map(response *http.Response) ([]Neo, error)
}

type NeoMapper struct {
	Decoder
}

func NewNeoMapper() *NeoMapper {
	return &NeoMapper{
		Decoder: new(NeoDecoder),
	}
}

func (nm *NeoMapper) Map(res *http.Response) ([]Neo, error) {
	neoRes := new(NeoResponse)

	err := json.NewDecoder(res.Body).Decode(neoRes)
	if err != nil {
		return nil, err
	}

	if neoRes.Count == "0" {
		return nil, errors.New("no results were found for this search")
	}

	return nm.mapNeoResToNeo(neoRes)
}

func (nm *NeoMapper) mapNeoResToNeo(neoRes *NeoResponse) ([]Neo, error) {
	var neos []Neo

	for _, res := range neoRes.Data {
		n, err := nm.mapNeoResArrayToStruct(res, neoRes.Fields)
		if err != nil {
			return nil, err
		}

		neos = append(neos, *n)
	}

	return neos, nil
}

func (nm *NeoMapper) mapNeoResArrayToStruct(res []string, fields []string) (*Neo, error) {
	mappedNeo := make(map[string]string)
	for i, field := range res {
		mappedNeo[fields[i]] = field
	}

	result := &Neo{}
	err := mapstructure.Decode(mappedNeo, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
