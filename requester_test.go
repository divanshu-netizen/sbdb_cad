package sbdb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequester_Get(t *testing.T) {
	tests := []struct {
		name        string
		server      httptest.Server
		wantResp    string
		wantErr     bool
		expectedErr error
	}{
		{
			name: "Return http.Response with no error",
			server: *httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "test")
			})),
			wantResp: "test",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &Requester{}

			gotResp, err := r.Get(tt.server.URL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}
			gotRespString, err := ioutil.ReadAll(gotResp.Body)
			gotResp.Body.Close()
			if err != nil {
				t.Errorf("Something went wrong in the test server")

			}
			if string(gotRespString) != tt.wantResp {
				t.Errorf("Get() gotResp = %v, want %v", string(gotRespString), tt.wantResp)
			}
		})
	}
}
