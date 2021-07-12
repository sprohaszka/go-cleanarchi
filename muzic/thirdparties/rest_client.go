package thirdparties

import (
	"net/http"
	"io/ioutil"
)

//-- STRUCT
type RestClient struct {
	client http.Client
}

//-- BUILDER
func BuildRestClient() RestClient {
	return RestClient {
		client: http.Client{},
	}
}

//-- METHODS
func (r RestClient) Get(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "LiveDemo(0.0)")
	req.Header.Add("Accept", "application/json")

	resp, _ := r.client.Do(req)
	return ioutil.ReadAll(resp.Body)
}