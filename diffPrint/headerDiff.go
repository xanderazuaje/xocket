package diffPrint

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/types"
	"net/http"
)

func headerDiff(res *http.Response, exp *types.ExpectedResponse, ok *bool) {
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
