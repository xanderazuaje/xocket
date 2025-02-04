package diffPrint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/types"
	"log"
	"reflect"
)

func jsonDiff(bodyRaw []byte, exp *types.ExpectedResponse, ok *bool) {
	if json.Valid(bodyRaw) {
		var resBody interface{}
		err := json.Unmarshal(bodyRaw, &resBody)
		if err != nil {
			log.Fatal(err.Error())
		}
		if !reflect.DeepEqual(exp.Body, resBody) {
			*ok = false
			colors.Printf("@b*(BODY:) @r*(DIFF)")
			if flags.This.RunType.Contains(flags.RunDebug) {
				colors.Printf("@b*(EXPECTED:)")
			} else {
				colors.Printf("expected json and received json @*r(doesn't) match")
			}
			// Printing expected body
			if flags.This.RunType.Contains(flags.RunDebug) {
				var buff bytes.Buffer
				encoder := json.NewEncoder(&buff)
				encoder.SetEscapeHTML(false)
				encoder.SetIndent("", "  ")
				err := encoder.Encode(exp.Body)
				if err != nil {
					log.Fatal(err.Error())
				}
				colors.Printf("@b*(GOT:)")
				buff.Reset()
				err = json.Indent(&buff, bodyRaw, "", "  ")
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(buff.String())
			}
		}
	} else {
		colors.Printf("@b*(BODY:) @r*(DIFF)")
		colors.Printf("Response's body is not a valid json")
		if flags.This.RunType.Contains(flags.RunDebug) {
			colors.Printf("@b*(GOT:)")
			fmt.Println(bodyRaw)
		}
		*ok = false
	}
}
