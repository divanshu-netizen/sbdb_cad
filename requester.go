package neos

import (
	"net/http"
)

type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

type Requester struct{}

func (r *Requester) Get(url string) (resp *http.Response, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return res, nil
}
