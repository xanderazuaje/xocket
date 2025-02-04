package diffPrint

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/types"
	"net/http"
)

func cookiesDiff(res *http.Response, exp *types.ExpectedResponse, ok *bool) {
	resCookies := res.Cookies()
	if len(resCookies) == 0 {
		colors.Printf("@b*(COOKIES:) @r(ERROR) - response has no cookies")
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
