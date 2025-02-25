package types

import (
	"github.com/xanderazuaje/xocket/colors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
)

type CookieOven struct {
	Name    string
	Cookies []*http.Cookie
	Url     string
	Method  string
	Params  url.Values
	Header  http.Header
	Form    *Form
}

func (oven *CookieOven) RequestCookie(jar *cookiejar.Jar, u *url.URL, client http.Client, wg *sync.WaitGroup, stdout chan<- string) {
	defer wg.Done()
	var requestName string
	if oven.Name != "" {
		requestName = oven.Name
	} else {
		requestName = oven.Url
	}
	var req *http.Request
	if oven.Form != nil {
		body, err := oven.Form.GetBodyBuff(&oven.Header)
		if err != nil {
			log.Fatal("failed to create request from oven at '"+oven.Name+"':", err.Error())
		}
		req, err = http.NewRequest(strings.ToUpper(oven.Method), oven.Url+"?"+oven.Params.Encode(), body)
		if err != nil {
			log.Fatal("failed to create request from oven at '"+oven.Name+"':", err.Error())
		}
	} else {
		var err error
		req, err = http.NewRequest(oven.Method, oven.Url, nil)
		if err != nil {
			log.Fatal("failed to create request from oven at '"+oven.Name+"':", err.Error())
		}
	}
	req.Header = oven.Header
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("failed to get cookies from oven at '"+req.URL.RawQuery+"' :", err.Error())
	}
	cookies := res.Cookies()
	if len(cookies) == 0 {
		var body []byte
		_, _ = res.Body.Read(body)
		stdout <- colors.Sprintf("@*y(WARNING:) no cookies returned from '%s', request returned status: '%s'\n@y(with body:) '%s'",
			requestName, res.Status, body)
	} else {
		stdout <- colors.Sprintf("@*g(SUCCESS:) %s", requestName)
	}
	jar.SetCookies(u, res.Cookies())
}
