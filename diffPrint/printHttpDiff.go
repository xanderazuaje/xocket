package diffPrint

import (
	"github.com/xanderazuaje/xocket/parsing"
	"net/http"
)

func PrintHttpDiff(res *http.Response, exp *parsing.ExpectedResponse) bool {
	var ok bool
	if exp.Status != "" {
		printStringDiff(&ok, "status", exp.Status, res.Status)
	}
	if exp.StatusCode != 0 {
		printIntDiff(&ok, "status code", exp.StatusCode, res.StatusCode)
	}
	if exp.Proto != "" {
		printStringDiff(&ok, "proto", exp.Proto, res.Proto)
	}
	if exp.ProtoMajor != nil {
		printIntDiff(&ok, "proto major", *exp.ProtoMajor, res.ProtoMajor)
	}
	if exp.ProtoMinor != nil {
		printIntDiff(&ok, "proto minor", *exp.ProtoMinor, res.ProtoMinor)
	}
	if len(exp.Cookies) > 0 {
		cookiesDiff(res, exp, &ok)
	}
	if len(exp.Header) > 0 {
		headerDiff(res, exp, &ok)
	}
	if exp.Body != nil {
		BodyDiff(res, exp, &ok)
	}
	return ok
}
