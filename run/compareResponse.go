package run

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
	"github.com/xanderazuaje/xocket/parsing"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
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
	var bodyRaw []byte
	bodyRaw, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	switch exp.BodyType {
	case parsing.BodyJson:
		jsonDiff(bodyRaw, exp, ok)
	case parsing.BodyString:
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
	case parsing.BodyHTML:
		//TODO: check if are the same
		hasHTML := false
		t := html.NewTokenizer(strings.NewReader(string(bodyRaw)))
		for {
			tt := t.Next()
			if tt == html.ErrorToken {
				err := t.Err()
				if err != nil && err.Error() == "EOF" {
					if !hasHTML {
						*ok = false
						colors.Log("@b*(BODY:) @r*(ERROR) not a valid html")
					}
					break
				}
			} else if tt == html.StartTagToken ||
				tt == html.EndTagToken ||
				tt == html.SelfClosingTagToken {
				hasHTML = true
			}
		}
	case parsing.BodyXML:
		//TODO: check if are the same
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
			colors.Log("@b*(EXPECTED:)")
			var buff bytes.Buffer
			encoder := json.NewEncoder(&buff)
			encoder.SetEscapeHTML(false)
			encoder.SetIndent("", "  ")
			err := encoder.Encode(exp.Body)
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(buff.String())
			colors.Log("@b*(GOT:)")
			buff.Reset()
			err = json.Indent(&buff, bodyRaw, "", "  ")
			if err != nil {
				log.Fatal(err.Error())
			}
			fmt.Println(buff.String())
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
