package run

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"net/http"
)

func printStringComparison(isOk *bool, name, s1, s2 string) {
	if s1 != s2 {
		colors.Log(
			"@b(%s) - @r*(DIFF:)\n\t@b*(expected:) -> '%s'\n\t@r*(got:) -> '%s'",
			name,
			s1,
			s2,
		)
		*isOk = false
	} else if flags.This.RunType.Contains(flags.RunDebug) {
		colors.Log("@b(%s) - @g*(OK!)", name)
	}
	*isOk = true
}
func printIntComparison(isOk *bool, name string, d1, d2 int) {
	if d1 != d2 {
		colors.Log(
			"@b(%s) - @r*(DIFF:)\n\t@b*(expected:) -> '%d'\n\t@r*(got:) -> '%d'",
			name,
			d1,
			d2,
		)
		*isOk = false
	} else if flags.This.RunType.Contains(flags.RunDebug) {
		colors.Log("@b(%s) - @g*(OK!)", name)
	}
	*isOk = true
}

func compareResponse(res *http.Response, exp *parsing.ExpectedResponse) (ok bool) {
	ok = true
	if exp.Status != "" {
		printStringComparison(&ok, "status", exp.Status, res.Status)
	}
	if exp.StatusCode != 0 {
		printIntComparison(&ok, "status code", exp.StatusCode, res.StatusCode)
	}
	if exp.Proto != "" {
		printStringComparison(&ok, "proto", exp.Proto, res.Proto)
	}
	if exp.ProtoMajor != nil {
		printIntComparison(&ok, "proto major", *exp.ProtoMajor, res.ProtoMajor)
	}
	if exp.ProtoMinor != nil {
		printIntComparison(&ok, "proto minor", *exp.ProtoMinor, res.ProtoMinor)
	}
	if len(exp.Cookies) > 0 {
		resCookies := res.Cookies()
		if len(resCookies) == 0 {
			colors.Log("@b*(COOKIES:) @r(ERROR) - response has no cookies")
			ok = false
		} else {
			cookiesMap := map[string]*http.Cookie{}
			for _, c := range resCookies {
				cookiesMap[c.Name] = c
			}
			for i, c := range exp.Cookies {
				rc := cookiesMap[c.Name]
				ok = c.PrintDifference(i, rc)
			}
		}
	}
	if ok {
		colors.Log("test @g(OK)")
	} else {
		colors.Log("test @r(ERROR)")
	}
	fmt.Println("----------------------------")
	return true
}
