package types

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

type Program struct {
	Endpoint  string
	CookieJar *ProgramJar `yaml:"cookie-jar"`
	Tests     []Test
}

type Test struct {
	Name            string
	Cookies         []*http.Cookie
	Path            string
	Method          string
	Params          url.Values
	Header          http.Header
	Form            *Form
	Expected        *ExpectedResponse
	IgnoreCookieJar bool
}

type BodyType string

const (
	BodyJson   BodyType = "json"
	BodyXML    BodyType = "xml"
	BodyHTML   BodyType = "html"
	BodyString BodyType = "string"
)

var bodyMap = map[BodyType]struct{}{
	BodyJson:   {},
	BodyXML:    {},
	BodyHTML:   {},
	BodyString: {},
}

func (b *BodyType) UnmarshalYAML(s []byte) error {
	_, ok := bodyMap[BodyType(s)]
	if ok == false {
		return errors.New("invalid body-type property")
	}
	*b = BodyType(s)
	return nil
}

type ExpectedResponse struct {
	Status           string
	StatusCode       int `yaml:"status-code"`
	Proto            string
	ProtoMajor       *int `yaml:"proto-major"`
	ProtoMinor       *int `yaml:"proto-minor"`
	Cookies          []*ExpectedCookie
	Header           http.Header
	Body             any
	BodyType         BodyType `yaml:"body-type"`
	ContentLength    *int64   `yaml:"content-length"`
	TransferEncoding []string `yaml:"transfer-encoding"`
	Close            *bool
	Uncompressed     *bool
	Trailer          http.Header
	Request          *http.Request
	TLS              *tls.ConnectionState
}

type Field struct {
	Data   []string
	Filter []Filter
}

type Filter struct {
	Type  FilterType
	Min   *float64
	Max   *float64
	Len   *int
	Regex *regexp.Regexp
}
