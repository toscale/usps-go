package usps

import (
	"bytes"
	"encoding/xml"
	"strings"

	"github.com/pkg/errors"
)

type Address struct {
	Address1 string `xml:"Address1"`
	Address2 string `xml:"Address2"`
	City     string `xml:"City"`
	State    string `xml:"State"`
	Zip5     string `xml:"Zip5"`
	Zip4     string `xml:"Zip4"`
}

type ZipCode struct {
	Zip5 string `xml:"Zip5"`
}

type AddressValidateResponse struct {
	Address struct {
		Address1 string `xml:"Address1"`
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		State    string `xml:"State"`
		Zip5     string `xml:"Zip5"`
		Zip4     string `xml:"Zip4"`
	} `xml:"Address"`
}

type ZipCodeLookupResponse struct {
	Address struct {
		Address1 string `xml:"Address1"`
		Address2 string `xml:"Address2"`
		City     string `xml:"City"`
		State    string `xml:"State"`
		Zip5     string `xml:"Zip5"`
		Zip4     string `xml:"Zip4"`
	} `xml:"Address"`
}

type CityStateLookupResponse struct {
	ZipC struct {
		Zip5  string `xml:"Zip5"`
		City  string `xml:"City"`
		State string `xml:"State"`
	} `xml:"ZipCode"`
}

func (U *USPS) AddressVerification(address Address) (*AddressValidateResponse, error) {
	result := &AddressValidateResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(address)
	if err != nil {
		return nil, errors.Wrap(err, "marshal address")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("Verify&XML=")
	urlToEncode := "<AddressValidateRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</AddressValidateRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequest(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get http")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal response")
	}

	return result, nil
}

func (U *USPS) ZipCodeLookup(address Address) (*ZipCodeLookupResponse, error) {
	result := &ZipCodeLookupResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(address)
	if err != nil {
		return nil, errors.Wrap(err, "marshal address")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("ZipCodeLookup&XML=")
	urlToEncode := "<ZipCodeLookupRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</ZipCodeLookupRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequest(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get request")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return result, nil
}

func (U *USPS) CityStateLookup(zipcode ZipCode) (*CityStateLookupResponse, error) {
	result := &CityStateLookupResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(zipcode)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CityStateLookup&XML=")
	urlToEncode := "<CityStateLookupRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CityStateLookupRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequest(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get http")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	return result, nil
}
