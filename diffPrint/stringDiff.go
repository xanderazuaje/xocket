package diffPrint

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/types"
)

func bodyStringDiff(exp *types.ExpectedResponse, ok *bool, bodyRaw []byte) *bool {
	expStr, ok2 := exp.Body.(string)
	if !ok2 {
		colors.Log("@r*(ERROR:) Test body type don't match")
		*ok = false
	} else if expStr != string(bodyRaw) {
		colors.Log("@b*(BODY:) @r*(DIFF)")
		colors.Log("@b*(EXPECTED:) %s", expStr)
		colors.Log("@b*(GOT:) %s", bodyRaw)
		*ok = false
	}
	return ok
}
