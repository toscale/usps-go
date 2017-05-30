package usps

import (
	"bytes"
	"encoding/xml"
	"strings"

	"github.com/pkg/errors"
)

type RateRequest struct {
	XMLName        xml.Name `xml:Package"`
	Revision       string   `xml:"Revision"`
	Service        string   `xml:"Service"`
	ZipOrigination string   `xml:"ZipOrigination"`
	ZipDestination string   `xml:"ZipDestination"`
	Pounds         string   `xml:"Pounds"`
	Ounces         string   `xml:"Ounces"`
	Container      string   `xml:"Container"`
	Size           string   `xml:"Size"`
	Width          string   `xml:"Width"`
	Length         string   `xml:"Length"`
	Height         string   `xml:"Height"`
	Girth          string   `xml:"Girth"`
}

type RateV4Response struct {
}

func (U *USPS) RateDomestic(rate RateRequest) (*RateV4Response, error) {
	result := &RateV4Response{}
	if U.Username == "" {
		return nil, errors.New("username is missing")
	}

	xmlOut, err := xml.Marshal(rate)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("RateV4&XML=")
	urlToEncode := "<RateV4Request USERID=\"" + U.Username + "\">"
	urlToEncode += string(xmlOut)
	urlToEncode += "</RateV4Request>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequest(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get http")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), result)
	return result, errors.Wrap(err, "unmarshal")
}
