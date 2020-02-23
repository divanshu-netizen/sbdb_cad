package sbdb

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNeoDecoder_Decode(t *testing.T) {
	type args struct {
		input  interface{}
		output interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nd := &SBDecoder{}
			if err := nd.Decode(tt.args.input, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNeoMapper_Map(t *testing.T) {
	type fields struct {
		Decoder Decoder
	}
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []SB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &SBMapper{
				Decoder: tt.fields.Decoder,
			}
			got, err := nm.Map(tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("Map() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeoMapper_mapNeoResArrayToStruct(t *testing.T) {
	type fields struct {
		Decoder Decoder
	}
	type args struct {
		res    []string
		fields []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *SB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &SBMapper{
				Decoder: tt.fields.Decoder,
			}
			got, err := nm.mapNeoResArrayToStruct(tt.args.res, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapNeoResArrayToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapNeoResArrayToStruct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNeoMapper_mapNeoResToNeo(t *testing.T) {
	type fields struct {
		Decoder Decoder
	}
	type args struct {
		neoRes *NeoResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []SB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := &SBMapper{
				Decoder: tt.fields.Decoder,
			}
			got, err := nm.mapNeoResToNeo(tt.args.neoRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapNeoResToNeo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapNeoResToNeo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNeoMapper(t *testing.T) {
	tests := []struct {
		name string
		want *SBMapper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSBMapper(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSBMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}