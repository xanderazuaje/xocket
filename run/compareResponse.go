package run

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"io"
	"log"
	"net/http"
	"reflect"
)

func compareResponse(res *http.Response, exp *parsing.ExpectedResponse) (ok bool) {
	ok = true
	if exp.Status != "" {
		printStringDiff(&ok, "status", exp.Status, res.Status)
	}
	if exp.StatusCode != 0 {
		printIntDiff(&ok, "status code", exp.StatusCode, res.StatusCode)
	}
	if exp.Proto != "" {
		printStringDiff(&ok, "proto", exp.Proto, res.Proto)
	}
	if exp.ProtoMajor != nil {
		printIntDiff(&ok, "proto major", *exp.ProtoMajor, res.ProtoMajor)
	}
	if exp.ProtoMinor != nil {
		printIntDiff(&ok, "proto minor", *exp.ProtoMinor, res.ProtoMinor)
	}
	if len(exp.Cookies) > 0 {
		cookiesDiff(res, exp, &ok)
	}
	if len(exp.Header) > 0 {
		headerDiff(res, exp, &ok)
	}
	if exp.Body != nil {
		bodyDiff(res, exp, &ok)
	}
	if ok {
		colors.Log("test @g(OK)")
	} else {
		colors.Log("test @r(ERROR)")
	}
	fmt.Println("----------------------------")
	return ok
}

func bodyDiff(res *http.Response, exp *parsing.ExpectedResponse, ok *bool) {
	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	switch exp.BodyType {
	case parsing.BodyJson:
		jsonDiff(bodyRaw, exp, ok)
	case parsing.BodyString:
		ok = stringDiff(exp, ok, bodyRaw)
	case parsing.BodyHTML:
		tagDiff(bodyRaw, ok, exp, "html")
	case parsing.BodyXML:
		tagDiff(bodyRaw, ok, exp, "xml")
	}
}

func stringDiff(exp *parsing.ExpectedResponse, ok *bool, bodyRaw []byte) *bool {
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

func jsonDiff(bodyRaw []byte, exp *parsing.ExpectedResponse, ok *bool) {
	if json.Valid(bodyRaw) {
		var resBody interface{}
		err := json.Unmarshal(bodyRaw, &resBody)
		if err != nil {
			log.Fatal(err.Error())
		}
		if !reflect.DeepEqual(exp.Body, resBody) {
			*ok = false
			colors.Log("@b*(BODY:) @r*(DIFF)")
			if flags.This.RunType.Contains(flags.RunDebug) {
				colors.Log("@b*(EXPECTED:)")
			} else {
				colors.Log("expected json and received json @*r(doesn't) match")
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
				colors.Log("@b*(GOT:)")
				buff.Reset()
				err = json.Indent(&buff, bodyRaw, "", "  ")
				if err != nil {
					log.Fatal(err.Error())
				}
				fmt.Println(buff.String())
			}
		}
	} else {
		colors.Log("@b*(BODY:) @r*(DIFF)")
		colors.Log("Response's body is not a valid json")
		if flags.This.RunType.Contains(flags.RunDebug) {
			colors.Log("@b*(GOT:)")
			fmt.Println(bodyRaw)
		}
		*ok = false
	}
}
