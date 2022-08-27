package cepgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
	Do(req *http.Request) (*http.Response, error)
}

var defaultHttpClient HttpClient

func requester(requestUrl string) (body []byte, err error) {
	response, err := http.DefaultClient.Get(requestUrl)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 100 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("GET %s - status code: %d %s", requestUrl, response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func poster(requestUrl string, data map[string]string) (body []byte, err error) {
	if defaultHttpClient == nil {
		defaultHttpClient = http.DefaultClient
	}
	urlData := url.Values{}
	for key, value := range data {
		urlData.Set(key, value)
	}
	encodedData := urlData.Encode()

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(encodedData))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(encodedData)))
	response, err := defaultHttpClient.Do(req)
	if err != nil {
		return nil, err

	}
	if response.StatusCode < 100 || response.StatusCode >= 400 {
		return nil, fmt.Errorf("POST %s - status code: %d %s", requestUrl, response.StatusCode, response.Status)
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
