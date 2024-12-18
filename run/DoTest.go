package run

import (
	"bytes"
	"encoding/json"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"io"
	"net/http"
	"os"
	"strings"
)

func DoTest(test parsing.Test, endpoint string) (err error) {
	req, err := setRequest(test, endpoint)
	if err != nil {
		return err
	}
	client := http.Client{}
	colors.Log("@*g(RETURNED:)")
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	colors.Log("@*b(STATUS:) %d - %s", res.StatusCode, res.Status)
	if flags.This.RunType.Contains(flags.RunDebug) {
		err := printDebugInfo(res)
		if err != nil {
			return err
		}
	}
	compareResponse(res, test.Expected)
	return nil
}

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

func printDebugInfo(res *http.Response) error {
	colors.Log("@*g(CONTENT-LENGTH:) %d", res.ContentLength)
	//print headers
	headerJson, _ := json.MarshalIndent(res.Header, "", "   ")
	colors.Log("@*b(HEADERS:)\n%+v", string(headerJson))
	//print cookies
	cookies := res.Cookies()
	if len(cookies) > 0 {
		colors.Log("@*b(COOKIES:)")
		for _, v := range cookies {
			colors.Log("   - %s", v.String())
		}
	} else {
		colors.Log("@*b(COOKIES:) none")
	}
	//print body
	var body []byte
	body, err := io.ReadAll(res.Body)
	if json.Valid(body) {
		var indentedJson bytes.Buffer
		err = json.Indent(&indentedJson, body, "", "   ")
		if err != nil {
			return err
		}
		colors.Log("@*b(BODY (json):)\n%+v", indentedJson.String())
	} else {
		colors.Log("@*b(BODY:)\n%+v", body)
	}
	return nil
}
