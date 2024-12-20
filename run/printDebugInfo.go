package run

import (
	"bytes"
	"encoding/json"
	"github.com/xanderazuaje/xocket/colors"
	"io"
	"net/http"
)

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
