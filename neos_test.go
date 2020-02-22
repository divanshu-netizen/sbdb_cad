package neos

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

const baseUrl = "baseUrl"

func TestFindNeoBy(t *testing.T) {
	tests := []struct {
		baseUrl         string
		name            string
		args            NeoQueryOptions
		getterData      *http.Response
		getterErr       error
		mapperData      []Neo
		mapperErr       error
		QueryStringData string
		want            []Neo
		wantErr         bool
		err             error
	}{
		{
			baseUrl:         baseUrl,
			name:            "Neos returned, no error occurs",
			args:            NeoQueryOptions{},
			getterData:      new(http.Response),
			getterErr:       nil,
			mapperData:      []Neo{Neo{}, Neo{}},
			mapperErr:       nil,
			QueryStringData: "test query string data",
			want:            []Neo{Neo{}, Neo{}},
			wantErr:         false,
			err:             nil,
		}, {
			baseUrl:         baseUrl,
			name:            "Neos not returned, error occurs in getter",
			args:            NeoQueryOptions{},
			getterData:      nil,
			getterErr:       errors.New("mapper encountered an error"),
			mapperData:      nil,
			mapperErr:       nil,
			QueryStringData: "test query string data",
			want:            nil,
			wantErr:         true,
			err:             errors.New("mapper encountered an error"),
		}, {
			baseUrl:         baseUrl,
			name:            "Neos not returned, error occurs in mapper",
			args:            NeoQueryOptions{},
			getterData:      new(http.Response),
			getterErr:       nil,
			mapperData:      nil,
			mapperErr:       errors.New("mapper encountered an error"),
			QueryStringData: "test query string data",
			want:            nil,
			wantErr:         true,
			err:             errors.New("mapper encountered an error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &neoService{
				BaseUrl:            tt.baseUrl,
				Getter:             &MockGetter{tt.getterData, tt.getterErr},
				Mapper:             &MockMapper{tt.mapperData, tt.mapperErr},
				QueryStringBuilder: &MockQueryStringBuilder{tt.QueryStringData},
			}
			got, err := ns.FindNeoBy(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindNeoBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindNeoBy() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type MockGetter struct {
	data *http.Response
	err  error
}

func (mg *MockGetter) Get(url string) (resp *http.Response, err error) {
	return mg.data, mg.err
}

type MockMapper struct {
	data []Neo
	err  error
}

func (mm *MockMapper) Map(response *http.Response) ([]Neo, error) {
	return mm.data, mm.err
}

type MockQueryStringBuilder struct {
	data string
}

func (mqs *MockQueryStringBuilder) Build(nqo *NeoQueryOptions) string {
	return mqs.data
}
