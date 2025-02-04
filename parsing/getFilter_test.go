package parsing_test

import (
	"github.com/xanderazuaje/xocket/parsing"
	"github.com/xanderazuaje/xocket/types"
	"reflect"
	"regexp"
	"testing"
)

func getValue(f float64) *float64 {
	return &f
}

func TestGetFilter(t *testing.T) {
	testCases := []struct {
		in   string
		want types.Field
	}{
		{
			"<integer>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{Type: types.FilterInteger},
				}},
		},
		{
			"<integer:min=2>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{
						Type: types.FilterInteger,
						Min:  getValue(2),
					},
				}},
		},
		{
			"<integer:min=-5>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{
						Type: types.FilterInteger,
						Min:  getValue(-5),
					},
				}},
		},
		{
			"<float>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{Type: types.FilterFloat},
				}},
		},
		{
			"<string>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{Type: types.FilterString},
				}},
		},
		{
			"hola <string> que tal",
			types.Field{
				Data: []string{"hola ", " que tal"},
				Filter: []types.Filter{
					{Type: types.FilterString},
				}},
		},
		{
			"<nil>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{Type: types.FilterNil},
				}},
		},
		{
			"<string> <integer>",
			types.Field{
				Data: []string{"", " ", ""},
				Filter: []types.Filter{
					{Type: types.FilterString},
					{Type: types.FilterInteger},
				}},
		},
		{
			"<string:r='[aA-zZ]\\w+'>",
			types.Field{
				Data: []string{"", ""},
				Filter: []types.Filter{
					{Type: types.FilterString, Regex: regexp.MustCompile("[aA-zZ]\\w+")},
				}},
		},
		{
			"<string",
			types.Field{
				Data: []string{"<string"},
			},
		},
		{
			"<>",
			types.Field{
				Data: []string{"<>"},
			},
		},
	}
	for _, tc := range testCases {
		ret, err := parsing.GetFilterField(tc.in)
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
		_, err := parsing.GetFilterField(s)
		if err == nil {
			t.Error("invalid syntax not returning an error: ", s)
			continue
		}
	}
}
