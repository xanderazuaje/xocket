package run

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"net/http"
)

func DoTest(test parsing.Test, endpoint string) (err error, ok bool) {
	req, err := setRequest(test, endpoint)
	if err != nil {
		return err, false
	}
	client := http.Client{}
	colors.Log("@*g(RETURNED:)")
	res, err := client.Do(req)
	if err != nil {
		return err, false
	}
	colors.Log("@*b(STATUS:) %d - %s", res.StatusCode, res.Status)
	if flags.This.RunType.Contains(flags.RunDebug) {
		err := printDebugInfo(res)
		if err != nil {
			return err, false
		}
	}
	ok = compareResponse(res, test.Expected)
	return nil, ok
}
