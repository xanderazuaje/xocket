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
	ok = true
	req, err := setters.SetRequest(test, endpoint)
	if err != nil {
		return err, false
	}
	client := http.Client{}
	if jar != nil {
		client.Jar = jar
	}
	colors.Printf("@*g(RETURNED:)")
	res, err := client.Do(req)
	if err != nil {
		return err, false
	}
	colors.Printf("@*b(STATUS:) %d - %s", res.StatusCode, res.Status)
	if flags.This.RunType.Contains(flags.RunDebug) {
		err := printDebugInfo(res)
		if err != nil {
			return err, false
		}
	}
	compareResponse(res, test.Expected, &ok)
	return nil, ok
}

func compareResponse(res *http.Response, exp *types.ExpectedResponse, ok *bool) {
	diffPrint.PrintHttpDiff(res, exp, ok)
	if *ok {
		colors.Printf("test @g(OK)")
	} else {
		colors.Printf("test @r(ERROR)")
	}
	fmt.Println("----------------------------")
}
