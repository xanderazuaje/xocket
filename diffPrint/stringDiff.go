package diffPrint

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/types"
)

func bodyStringDiff(exp *types.ExpectedResponse, ok *bool, bodyRaw []byte) *bool {
	expStr, ok2 := exp.Body.(string)
	if !ok2 {
		colors.Printf("@r*(ERROR:) Test body type don't match")
		*ok = false
	} else if expStr != string(bodyRaw) {
		colors.Printf("@b*(BODY:) @r*(DIFF)")
		colors.Printf("@b*(EXPECTED:) %s", expStr)
		colors.Printf("@b*(GOT:) %s", bodyRaw)
		*ok = false
	}
	return ok
}
