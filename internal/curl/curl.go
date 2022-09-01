package curl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func GetRequest(url string) ([]byte, error) {

	token := os.Getenv("GCP_TOKEN")
	if token == "" {
		return nil, errors.New("internalError: invalid token, export GCP_TOKEN")
	}

	bearer := "Bearer " + token

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	req.Header.Set("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		errorMsg := fmt.Sprintf("too many redirects: %v", err)
		return nil, errors.New(errorMsg)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	r := regexp.MustCompile(`20([0-9])`)
	if !r.Match([]byte(string(resp.StatusCode))) {
		errorMsg := fmt.Sprintf(
			"expected status code 2XX, but received: %d.\n %v",
			resp.StatusCode, err)
		return data, errors.New(errorMsg)
	}
	return data, nil
}
