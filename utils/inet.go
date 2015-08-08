package utils

import (
	"errors"
	"net/http"
	"time"
)

func RetryRequest(tries int, client *http.Client, request *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error
	for try := 0; try < tries; try++ {
		resp, err = client.Do(request)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			time.Sleep(2 * time.Second)
			continue
		}
		return resp, err
	}
	return resp, errors.New("Service unavailable")
}
