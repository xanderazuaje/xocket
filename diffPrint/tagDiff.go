package diffPrint

import (
	"encoding/xml"
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"log"
	"reflect"
)

func tagDiff(bodyRaw []byte, ok *bool, exp *parsing.ExpectedResponse, format string) {
	var resBody any
	if err := xml.Unmarshal(bodyRaw, &resBody); err != nil {
		colors.Log("@b*(BODY:) @r*(ERROR) response's body is not a valid html")
		*ok = false
		return
	}
	if !reflect.DeepEqual(exp.Body, resBody) {
		*ok = false
		colors.Log("@b*(BODY:) @r*(DIFF)")
		if flags.This.RunType.Contains(flags.RunDebug) {
			colors.Log("@b*(EXPECTED:)")
		} else {
			colors.Log("Expected %s and received %s @*r(doesn't) match", format, format)
		}
		// Printing expected body
		if flags.This.RunType.Contains(flags.RunDebug) {
			str, err := xml.Marshal(exp.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			colors.Log("@b*(GOT:)")
			fmt.Println(str)
		}
	}
}
