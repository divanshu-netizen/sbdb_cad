package sbdb_cad

import (
	"encoding/json"
	"errors"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type SbCadResponse struct {
	Signature `json:"signature"`
	Count     string     `json:"count"`
	Fields    []string   `json:"fields"`
	Data      [][]string `json:"data"`
}

type Signature struct {
	Version string `json:"version"`
	Source  string `json:"source"`
}

type Decoder interface {
	Decode(input interface{}, output interface{}) error
}

type SbCadDecoder struct{}

func (sd *SbCadDecoder) Decode(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return err
	}

	return nil
}

type Mapper interface {
	Map(response *http.Response) ([]SbCad, error)
}

type SbCadMapper struct {
	Decoder
}

func NewSbCadMapper() *SbCadMapper {
	return &SbCadMapper{
		Decoder: new(SbCadDecoder),
	}
}

func (sb *SbCadMapper) Map(res *http.Response) ([]SbCad, error) {
	sbRes := new(SbCadResponse)

	err := json.NewDecoder(res.Body).Decode(sbRes)
	if err != nil {
		return nil, err
	}

	if sbRes.Count == "0" {
		return nil, errors.New("no results were found for this search")
	}

	if sbRes.Signature.Version != "1.1" {
		return nil, errors.New("api version has been updated. please contact maintainer of this library")
	}

	return sb.mapSbCadResToSbCad(sbRes)
}

func (sb *SbCadMapper) mapSbCadResToSbCad(sbRes *SbCadResponse) ([]SbCad, error) {
	var sbs []SbCad

	for _, res := range sbRes.Data {
		s, err := sb.mapSbCadResArrayToStruct(res, sbRes.Fields)
		if err != nil {
			return nil, err
		}

		sbs = append(sbs, *s)
	}

	return sbs, nil
}

func (sb *SbCadMapper) mapSbCadResArrayToStruct(res []string, fields []string) (*SbCad, error) {
	mappedSb := make(map[string]string)
	for i, field := range res {
		mappedSb[fields[i]] = field
	}

	result := &SbCad{}
	err := sb.Decoder.Decode(mappedSb, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
