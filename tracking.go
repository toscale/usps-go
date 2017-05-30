package usps

import (
	"bytes"
	"encoding/xml"
	"strings"

	"github.com/pkg/errors"
)

type TrackResponse struct {
	TrackInfo struct {
		TrackSummary string `xml:"TrackSummary"`
	} `xml:"TrackInfo"`
}

func (U *USPS) TrackPackage(trackingID string) (*TrackResponse, error) {
	result := &TrackResponse{}
	if U.Username == "" {
		return nil, errors.New("username missing")
	}

	var requestURL bytes.Buffer
	requestURL.WriteString("TrackV2&XML=")
	urlToEncode := "<TrackRequest USERID=\"" + U.Username + "\">"
	urlToEncode += "<TrackID ID=\"" + trackingID + "\"></TrackID>"
	urlToEncode += "</TrackRequest>"
	requestURL.WriteString(URLEncode(urlToEncode))

	body, err := U.GetRequest(requestURL.String())
	if err != nil {
		return nil, errors.Wrap(err, "get http")
	}

	bodyHeaderless := strings.Replace(string(body), xml.Header, "", 1)
	err = xml.Unmarshal([]byte(bodyHeaderless), &result)
	return result, errors.Wrap(err, "unmarshal")
}
