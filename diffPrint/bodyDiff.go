package diffPrint

import (
	"github.com/xanderazuaje/xocket/types"
	"io"
	"log"
	"net/http"
)

func BodyDiff(res *http.Response, exp *types.ExpectedResponse, ok *bool) {
	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	switch exp.BodyType {
	case types.BodyJson:
		jsonDiff(bodyRaw, exp, ok)
	case types.BodyString:
		ok = bodyStringDiff(exp, ok, bodyRaw)
	case types.BodyHTML:
		tagDiff(bodyRaw, ok, exp, "html")
	case types.BodyXML:
		tagDiff(bodyRaw, ok, exp, "xml")
	}
}
