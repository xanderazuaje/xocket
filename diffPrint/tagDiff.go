package diffPrint

import (
	"encoding/xml"
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/types"
	"log"
	"reflect"
)

func tagDiff(bodyRaw []byte, ok *bool, exp *types.ExpectedResponse, format string) {
	var resBody any
	if err := xml.Unmarshal(bodyRaw, &resBody); err != nil {
		colors.Printf("@b*(BODY:) @r*(ERROR) response's body is not a valid html")
		*ok = false
		return
	}
	if !reflect.DeepEqual(exp.Body, resBody) {
		*ok = false
		colors.Printf("@b*(BODY:) @r*(DIFF)")
		if flags.This.RunType.Contains(flags.RunDebug) {
			colors.Printf("@b*(EXPECTED:)")
		} else {
			colors.Printf("Expected %s and received %s @*r(doesn't) match", format, format)
		}
		// Printing expected body
		if flags.This.RunType.Contains(flags.RunDebug) {
			str, err := xml.Marshal(exp.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			colors.Printf("@b*(GOT:)")
			fmt.Println(str)
		}
	}
}
