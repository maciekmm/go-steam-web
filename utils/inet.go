package utils

import (
	"errors"
	"net/http"
)

func RetryRequest(tries int, client *http.Client, request *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for try := 0; try < tries; try++ {
		resp, err = client.Do(request)
		if resp.StatusCode != http.StatusOK {
			continue
		}
		return resp, err
	}
	return resp, errors.New("Service unavailable")
}
