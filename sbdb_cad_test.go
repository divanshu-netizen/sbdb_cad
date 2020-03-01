package sbdb_cad

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

const baseUrl = "baseUrl"

func TestFindSBBy(t *testing.T) {
	tests := []struct {
		baseUrl         string
		name            string
		args            SmallBodyOptions
		getterData      *http.Response
		getterErr       error
		mapperData      []SB
		mapperErr       error
		QueryStringData string
		want            []SB
		wantErr         bool
		err             error
	}{
		{
			baseUrl:         baseUrl,
			name:            "SBs returned, no error occurs",
			args:            SmallBodyOptions{},
			getterData:      new(http.Response),
			getterErr:       nil,
			mapperData:      []SB{SB{}, SB{}},
			mapperErr:       nil,
			QueryStringData: "test query string data",
			want:            []SB{SB{}, SB{}},
			wantErr:         false,
			err:             nil,
		}, {
			baseUrl:         baseUrl,
			name:            "SBs not returned, error occurs in getter",
			args:            SmallBodyOptions{},
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
			args:            SmallBodyOptions{},
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
			ns := &sbService{
				BaseUrl:            tt.baseUrl,
				Getter:             &MockGetter{tt.getterData, tt.getterErr},
				Mapper:             &MockMapper{tt.mapperData, tt.mapperErr},
				QueryStringBuilder: &MockQueryStringBuilder{tt.QueryStringData},
			}
			got, err := ns.FindSBBy(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSBBy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSBBy() got = %v, want %v", got, tt.want)
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
	data []SB
	err  error
}

func (mm *MockMapper) Map(response *http.Response) ([]SB, error) {
	return mm.data, mm.err
}

type MockQueryStringBuilder struct {
	data string
}

func (mqs *MockQueryStringBuilder) Build(sbo *SmallBodyOptions) string {
	return mqs.data
}
