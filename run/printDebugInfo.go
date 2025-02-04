package run

import (
	"bytes"
	"encoding/json"
	"github.com/xanderazuaje/xocket/colors"
	"io"
	"net/http"
)

func printDebugInfo(res *http.Response) error {
	colors.Printf("@*g(CONTENT-LENGTH:) %d", res.ContentLength)
	//print headers
	headerJson, _ := json.MarshalIndent(res.Header, "", "   ")
	colors.Printf("@*b(HEADERS:)\n%+v", string(headerJson))
	//print cookies
	cookies := res.Cookies()
	if len(cookies) > 0 {
		colors.Printf("@*b(COOKIES:)")
		for _, v := range cookies {
			colors.Printf("   - %s", v.String())
		}
	} else {
		colors.Printf("@*b(COOKIES:) none")
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
		colors.Printf("@*b(BODY (json):)\n%+v", indentedJson.String())
	} else {
		colors.Printf("@*b(BODY:)\n%+v", body)
	}
	return nil
}
