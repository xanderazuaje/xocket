package run

import (
	"fmt"
	"github.com/xanderazuaje/xocket/parsing"
	"strings"
)

func setLinePrompt(test parsing.Test, endpoint string) string {
	var infoStr string
	if test.Name != "" {
		infoStr = strings.Join([]string{
			fmt.Sprintf("@*b(%s:) %s -> %s", test.Name, test.Method, endpoint+test.Path),
			test.Params.Encode(),
		}, "?")
	} else {
		infoStr = strings.Join([]string{
			fmt.Sprintf("@*b(EXECUTING:) %s -> %s", test.Method, endpoint+test.Path),
			test.Params.Encode(),
		}, "?")
	}
	infoStr = strings.TrimRight(infoStr, "?")
	return infoStr
}
