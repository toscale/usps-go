package usps

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type PickUpRequest struct {
	FirmName     string `xml:"FirmName"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	Urbanization string `xml:"Urbanization"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
}

type Package struct {
	ServiceType string `xml:"ServiceType"`
	Count       string `xml:"Count"`
}

type PickupChangeRequest struct {
	FirstName           string `xml:"FirstName"`
	LastName            string `xml:"LastName"`
	FirmName            string `xml:"FirmName"`
	SuiteOrApt          string `xml:"SuiteOrApt"`
	Address2            string `xml:"Address2"`
	Urbanization        string `xml:"Urbanization"`
	City                string `xml:"City"`
	State               string `xml:"State"`
	ZIP5                string `xml:"ZIP5"`
	ZIP4                string `xml:"ZIP4"`
	Phone               string `xml:"Phone"`
	Extension           string `xml:"Extension"`
	Package             `xml:"Package"`
	EstimatedWeight     string `xml:"EstimatedWeight"`
	PackageLocation     string `xml:"PackageLocation"`
	SpecialInstructions string `xml:"SpecialInstructions"`
	ConfirmationNumber  string `xml:"ConfirmationNumber"`
}

type PickUpInquiryRequest struct {
	FirmName           string `xml:"FirmName"`
	SuiteOrApt         string `xml:"SuiteOrApt"`
	Address2           string `xml:"Address2"`
	Urbanization       string `xml:"Urbanization"`
	City               string `xml:"City"`
	State              string `xml:"State"`
	ZIP5               string `xml:"ZIP5"`
	ZIP4               string `xml:"ZIP4"`
	ConfirmationNumber string `xml:"ConfirmationNumber"`
}

type CarrierPickupAvailabilityResponse struct {
	FirmName     string `xml:"FirmName"`
	SuiteOrApt   string `xml:"SuiteOrApt"`
	Address2     string `xml:"Address2"`
	City         string `xml:"City"`
	State        string `xml:"State"`
	ZIP5         string `xml:"ZIP5"`
	ZIP4         string `xml:"ZIP4"`
	DayOfWeek    string `xml:"DayOfWeek"`
	Date         string `xml:"Date"`
	CarrierRoute string `xml:"CarrierRoute"`
	Error        string `xml:"Error"`
}

type CarrierPickupChangeResponse struct {
	FirstName           string `xml:"FirstName"`
	LastName            string `xml:"LastName"`
	FirmName            string `xml:"FirmName"`
	SuiteOrApt          string `xml:"SuiteOrApt"`
	Address2            string `xml:"Address2"`
	Urbanization        string `xml:"Urbanization"`
	City                string `xml:"City"`
	State               string `xml:"State"`
	ZIP5                string `xml:"ZIP5"`
	ZIP4                string `xml:"ZIP4"`
	Phone               string `xml:"Phone"`
	Extension           string `xml:"Extension"`
	Package             `xml:"Package"`
	EstimatedWeight     string `xml:"EstimatedWeight"`
	PackageLocation     string `xml:"PackageLocation"`
	SpecialInstructions string `xml:"SpecialInstructions"`
	ConfirmationNumber  string `xml:"ConfirmationNumber"`
	DayOfWeek           string `xml:"DayOfWeek"`
	Date                string `xml:"Date"`
	Status              string `xml:"Status"`
	Error               string `xml:"Error"`
}

type CarrierPickupInquiryResponse struct {
	FirstName           string    `xml:"FirstName"`
	LastName            string    `xml:"LastName"`
	FirmName            string    `xml:"FirmName"`
	SuiteOrApt          string    `xml:"SuiteOrApt"`
	Address2            string    `xml:"Address2"`
	Urbanization        string    `xml:"Urbanization"`
	City                string    `xml:"City"`
	State               string    `xml:"State"`
	ZIP5                string    `xml:"ZIP5"`
	ZIP4                string    `xml:"ZIP4"`
	Phone               string    `xml:"Phone"`
	Extension           string    `xml:"Extension"`
	Package             []Package `xml:"Package"`
	EstimatedWeight     string    `xml:"EstimatedWeight"`
	PackageLocation     string    `xml:"PackageLocation"`
	SpecialInstructions string    `xml:"SpecialInstructions"`
	ConfirmationNumber  string    `xml:"ConfirmationNumber"`
	DayOfWeek           string    `xml:"DayOfWeek"`
	Date                string    `xml:"Date"`
	Error               string    `xml:"Error"`
}

type Error struct {
	Number      string `xml:"Number"`
	Description string `xml:"Description"`
	Source      string `xml:"Source"`
}

func (U *USPS) PickupAvailability(pickup PickUpRequest) (*CarrierPickupAvailabilityResponse, error) {
	result := &CarrierPickupAvailabilityResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupAvailability&XML=")
	urlToEncode := "<CarrierPickupAvailabilityRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupAvailabilityRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequestHTTPS(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get https")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	fmt.Println(bodyHeaderless)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err != nil {
		errorResult := Error{}
		err = xml.Unmarshal([]byte(bodyHeaderless), &errorResult)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal error")
		}
		result.Error = errorResult.Description
	}

	return result, nil
}

func (U *USPS) PickupChange(pickup PickupChangeRequest) (*CarrierPickupChangeResponse, error) {
	result := &CarrierPickupChangeResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupChange&XML=")
	urlToEncode := "<CarrierPickupChangeRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupChangeRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequestHTTPS(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get https")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err != nil {
		errorResult := Error{}
		err = xml.Unmarshal([]byte(bodyHeaderless), &errorResult)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal error")
		}
		result.Error = errorResult.Description
	}

	return result, nil
}

func (U *USPS) PickupInquiry(pickup PickUpInquiryRequest) (*CarrierPickupInquiryResponse, error) {
	result := &CarrierPickupInquiryResponse{}
	if U.Username == "" {
		return nil, errors.New("Username is missing")
	}

	xmlOut, err := xml.Marshal(pickup)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("CarrierPickupInquiry&XML=")
	urlToEncode := "<CarrierPickupInquiryRequest USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</CarrierPickupInquiryRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequestHTTPS(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get https")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	if err == nil {
		return result, nil
		errorResult := Error{}
		err = xml.Unmarshal([]byte(bodyHeaderless), &errorResult)
		if err != nil {
			return nil, errors.Wrap(err, "unmarshal error")
		}
		result.Error = errorResult.Description
	}

	return result, nil
}
