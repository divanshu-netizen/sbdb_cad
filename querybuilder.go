package main

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/url"
	"reflect"
	"strings"
)

type Options struct {
	Query   string `url:"q"`
	ShowAll bool   `url:"all"`
	Page    int    `url:"page"`
}

type OptionValueConverter interface {
	Values(v interface{}) (url.Values, error)
}

type OVC struct {
}

func (b *OVC) Values(v interface{}) (url.Values, error) {
	return query.Values(v)
}

type QueryStringBuilder interface {
	Build(nqo *NeoQueryOptions) string
}

type queryBuilder struct {
	OptionValueConverter
}

func NewQueryBuilder() *queryBuilder {
	return &queryBuilder{
		OptionValueConverter: new(OVC),
	}
}

func (qb *queryBuilder) Build(nqo *NeoQueryOptions) string {
	fields := reflect.Indirect(reflect.ValueOf(nqo))
	var qp string
	for field := 0; field < fields.Type().NumField(); field++ {
		fieldName := fields.Type().Field(field).Name
		fieldVal := fields.Field(field)
		if len(fieldVal.String()) == 0 {
			fieldName = ""
		}

		switch fieldName {
		case "DateMin":
			qp += "date-min=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "DateMax":
			qp += "date-max=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "DistanceMin":
			qp += "dist-min=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "DistanceMax":
			qp += "dist-max=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "HMin":
			qp += "h-min=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "HMax":
			qp += "h-max=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "VelocityInfMin":
			qp += "v-inf-min=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "VelocityInfMax":
			qp += "v-inf-max=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "VelocityRelativeMin":
			qp += "vel-rel-min=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "VelocityRelativeMax":
			qp += "vel-rel-max=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Pha":
			qp += "pha=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Nea":
			qp += "nea=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Comet":
			qp += "comet=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "NeaComet":
			qp += "nea-comet=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Neo":
			qp += "neo=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Kind":
			qp += "kind=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Spk":
			qp += "spk=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Designation":
			qp += "designation=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Body":
			qp += "body=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "Sort":
			qp += "sort=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "limit":
			qp += "limit=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		case "FullName":
			qp += "fullname=" + fmt.Sprintf("%v", fields.Field(field).Interface()) + "&"
		}
	}


	return strings.TrimSuffix(qp, "&")
}
