package run

import (
	"bytes"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/parsing"
	"net/http"
	"os"
	"strings"
)

func setRequest(test parsing.Test, endpoint string) (*http.Request, error) {
	var err error
	test.Method = strings.ToUpper(os.ExpandEnv(test.Method))
	test.Path = os.ExpandEnv(test.Path)
	test.Params, err = expandMap(test.Params)
	if err != nil {
		return nil, err
	}
	test.Header, err = expandMap(test.Header)
	if err != nil {
		return nil, err
	}
	for _, c := range test.Cookies {
		c.Value = os.ExpandEnv(c.Value)
	}
	infoStr := setLinePrompt(test, endpoint)
	colors.Log(infoStr)
	var b bytes.Buffer
	if test.Form != nil {
		err = AddFormData(&test, &b)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(test.Method, endpoint+test.Path, &b)
	if err != nil {
		return nil, err
	}
	if test.Header != nil {
		req.Header = test.Header
	}
	req.URL.RawQuery = test.Params.Encode()
	for _, c := range test.Cookies {
		req.AddCookie(c)
	}
	return req, nil
}
