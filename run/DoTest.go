package run

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/diffPrint"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/setters"
	"github.com/xanderazuaje/xocket/types"
	"net/http"
	"net/http/cookiejar"
)

func DoTest(test types.Test, endpoint string, jar *cookiejar.Jar) (err error, ok bool) {
	req, err := setters.SetRequest(test, endpoint)
	if err != nil {
		return err, false
	}
	client := http.Client{Jar: jar}
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

func compareResponse(res *http.Response, exp *types.ExpectedResponse) bool {
	ok := diffPrint.PrintHttpDiff(res, exp)
	if ok {
		colors.Log("test @g(OK)")
	} else {
		colors.Log("test @r(ERROR)")
	}
	fmt.Println("----------------------------")
	return ok
}
