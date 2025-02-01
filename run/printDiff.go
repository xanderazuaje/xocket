package run

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"net/http"
	"strings"
)

func printStringDiff(isOk *bool, name, s1, s2 string) {
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
}

func printIntDiff(isOk *bool, name string, d1, d2 int) {
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
}

func headerDiff(res *http.Response, exp *parsing.ExpectedResponse, ok *bool) {
	for k, v := range exp.Header {
		for i := range v {
			if len(res.Header[k]) != len(v) {
				colors.Log("@r*(HEADER MISMATCH:) %s", k)
				printStrArrDiff(res.Header[k], v)
				*ok = false
				break
			}
			if res.Header[k][i] != v[i] {
				colors.Log("@r*(HEADER MISMATCH:) %s", k)
				printStrArrDiff(res.Header[k], v)
				*ok = false
				break
			}
		}
	}
}

func cookiesDiff(res *http.Response, exp *parsing.ExpectedResponse, ok *bool) {
	resCookies := res.Cookies()
	if len(resCookies) == 0 {
		colors.Log("@b*(COOKIES:) @r(ERROR) - response has no cookies")
		*ok = false
	} else {
		cookiesMap := map[string]*http.Cookie{}
		for _, c := range resCookies {
			cookiesMap[c.Name] = c
		}
		for i, c := range exp.Cookies {
			rc := cookiesMap[c.Name]
			*ok = c.PrintDifference(i, rc)
		}
	}
}

func printStrArrDiff(sa1, sa2 []string) {
	sa1len := len(sa1)
	sa2len := len(sa2)
	str := "\t[ %s ]\n\t[ %s ]"
	strA1 := make([]string, sa1len)
	strA2 := make([]string, sa2len)
	for i, v := range sa1 {
		if v != sa2[i] {
			strA1[i] = fmt.Sprintf("@r*(%s)", v)
			strA2[i] = fmt.Sprintf("@r*(%s)", sa2[i])
		} else {
			strA1[i] = v
			strA2[i] = sa2[i]
		}
	}
	if sa1len > sa2len {
		s := sa1[sa2len:]
		for i, v := range s {
			s[i] = fmt.Sprintf("@r*(%s)", v)
		}
		strA2 = append(strA2, s...)
	} else if sa2len > sa1len {
		s := sa2[sa1len:]
		for i, v := range s {
			s[i] = fmt.Sprintf("@r*(%s)", v)
		}
		strA2 = append(strA1, s...)
	}
	colors.Log(fmt.Sprintf(str, strings.Join(strA1, ", "), strings.Join(strA2, ", ")))
}
