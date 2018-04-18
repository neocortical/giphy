package giphy

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (c *client) get(path string, query url.Values, response interface{}) (err error) {
	requestURL := c.baseURL + path
	if len(query) > 0 {
		requestURL = requestURL + "?" + query.Encode()
	}

	data, err := doAPICall(c.httpClient, "GET", requestURL, nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, response)
	return
}

func buildQuery(apiKey string, opts *Options) (result url.Values) {
	result = url.Values{}
	result.Set("api_key", apiKey)
	if opts != nil {
		if opts.Limit > 0 {
			result.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Offset > 0 {
			result.Set("offset", strconv.Itoa(opts.Offset))
		}
		if len(opts.Rating) > 0 {
			result.Set("rating", opts.Rating)
		}
		if len(opts.Lang) > 0 {
			result.Set("lang", opts.Lang)
		}
	}
	return
}

func doAPICall(client *http.Client, method, urlStr string, body io.Reader) ([]byte, error) {
	httpRequest, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	if method == "POST" {
		httpRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	statusCode, status := resp.StatusCode, resp.Status
	if statusCode != 200 {
		return nil, fmt.Errorf("API returned an error: %s: %s", status, urlStr)
	}

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v: %s", err, urlStr)
	}

	return raw, nil
}
