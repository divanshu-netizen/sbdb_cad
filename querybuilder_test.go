package sbdb

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewQueryBuilder(t *testing.T) {
	tests := []struct {
		name string
		want *queryBuilder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewQueryBuilder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewQueryBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOVC_Values(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    url.Values
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &OVC{}
			got, err := b.Values(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Values() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryBuilder_Build(t *testing.T) {
	type fields struct {
		OptionValueConverter OptionValueConverter
	}
	type args struct {
		nqo *SmallBodyOptions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				OptionValueConverter: tt.fields.OptionValueConverter,
			}
			if got := qb.Build(tt.args.nqo); got != tt.want {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}