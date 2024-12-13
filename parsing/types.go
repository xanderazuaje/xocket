package parsing

import (
	"net/http"
	"net/url"
	"regexp"
)

type Program struct {
	Endpoint string
	Tests    []Test
}

type Test struct {
	Path     string
	Method   string
	Params   url.Values
	Headers  http.Header
	Expected *ExpectedResponse
}

type ExpectedResponse struct {
	Status  int
	Type    ExpectedType
	Body    *ExpectedBody
	Content string
}

type ExpectedBody struct {
	RawShape string        `yaml:"shape"`
	Shape    ExpectedShape `yaml:"-"`
	Filter   string        `yaml:"-"`
	Model    map[string]interface{}
}

type Field struct {
	Data   []string
	Filter []Filter
}

type Filter struct {
	Type  FilterType
	Min   float64
	Max   float64
	Len   int
	Regex *regexp.Regexp
}