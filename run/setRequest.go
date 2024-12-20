package run

import (
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
	infoStr := setLinePrompt(test, endpoint)
	colors.Log(infoStr)
	req, err := http.NewRequest(test.Method, endpoint+test.Path, nil)
	if err != nil {
		return nil, err
	}
	req.Header = test.Header
	req.URL.RawQuery = test.Params.Encode()
	return req, nil
}
