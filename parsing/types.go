package parsing

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"regexp"
)

type Program struct {
	Endpoint string
	Tests    []Test
}

type Test struct {
	Name     string
	Path     string
	Method   string
	Params   url.Values
	Header   http.Header
	Expected *ExpectedResponse
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
