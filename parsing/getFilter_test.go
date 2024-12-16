package parsing_test

import (
	"github.com/xanderazuaje/xocket/parsing"
	"reflect"
	"regexp"
	"testing"
)

func TestGetFilter(t *testing.T) {
	testCases := []struct {
		in   string
		want parsing.Field
	}{
		{
			"<integer>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterInteger},
				}},
		},
		{
			"<integer:min=2>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{
						Type: parsing.FilterInteger,
						Min:  2,
					},
				}},
		},
		{
			"<integer:min=-5>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{
						Type: parsing.FilterInteger,
						Min:  -5,
					},
				}},
		},
		{
			"<float>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterFloat},
				}},
		},
		{
			"<string>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterString},
				}},
		},
		{
			"hola <string> que tal",
			parsing.Field{
				Data: []string{"hola ", " que tal"},
				Filter: []parsing.Filter{
					{Type: parsing.FilterString},
				}},
		},
		{
			"<nil>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterNil},
				}},
		},
		{
			"<string> <integer>",
			parsing.Field{
				Data: []string{"", " ", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterString},
					{Type: parsing.FilterInteger},
				}},
		},
		{
			"<string:r='[aA-zZ]\\w+'>",
			parsing.Field{
				Data: []string{"", ""},
				Filter: []parsing.Filter{
					{Type: parsing.FilterString, Regex: regexp.MustCompile("[aA-zZ]\\w+")},
				}},
		},
		{
			"<string",
			parsing.Field{
				Data: []string{"<string"},
			},
		},
		{
			"<>",
			parsing.Field{
				Data: []string{"<>"},
			},
		},
	}
	for _, tc := range testCases {
		ret, err := parsing.GetFilter(tc.in)
		if err != nil {
			t.Error(err.Error())
			continue
		}
		if !reflect.DeepEqual(*ret, tc.want) {
			t.Log("expected: ", tc.want)
			t.Log("got: ", *ret)
			t.Error("unexpected output")
		}
	}
	errorCases := []string{
		"<hola>",
		"<integer:float>",
	}

	for _, s := range errorCases {
		_, err := parsing.GetFilter(s)
		if err == nil {
			t.Error("invalid syntax not returning an error: ", s)
			continue
		}
	}
}
