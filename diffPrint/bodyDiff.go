package diffPrint

import (
	"github.com/xanderazuaje/xocket/parsing"
	"io"
	"log"
	"net/http"
)

func BodyDiff(res *http.Response, exp *parsing.ExpectedResponse, ok *bool) {
	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	switch exp.BodyType {
	case parsing.BodyJson:
		jsonDiff(bodyRaw, exp, ok)
	case parsing.BodyString:
		ok = bodyStringDiff(exp, ok, bodyRaw)
	case parsing.BodyHTML:
		tagDiff(bodyRaw, ok, exp, "html")
	case parsing.BodyXML:
		tagDiff(bodyRaw, ok, exp, "xml")
	}
}
