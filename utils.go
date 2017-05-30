package usps

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

func URLEncode(urlToEncode string) string {
	return url.QueryEscape(urlToEncode)
}

func (U *USPS) GetRequest(requestURL string) ([]byte, error) {
	currentURL := ""
	if U.Production {
		currentURL += prodbase
	} else {
		currentURL += devbase
	}
	currentURL += requestURL

	resp, err := http.Get(currentURL)
	if err != nil {
		return nil, errors.Wrap(err, "http get")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, errors.Wrap(err, "read all")
}

func (U *USPS) GetRequestHTTPS(requestURL string) ([]byte, error) {
	currentURL := ""
	if U.Production {
		currentURL += prodhttpsbase
	} else {
		currentURL += devhttpsbase
	}
	currentURL += requestURL

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(currentURL)
	if err != nil {
		return nil, errors.Wrap(err, "get https")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, errors.Wrap(err, "read all")
}
