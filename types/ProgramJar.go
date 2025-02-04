package types

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
)

type ProgramJar struct {
	Oven    map[string][]*CookieOven
	Cookies map[string][]*http.Cookie
}

func (j *ProgramJar) FillJar() *cookiejar.Jar {
	// http request cookies
	client := http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("failed to create cookie jar:", err.Error())
	}
	colors.Printf("@*g(GETTING COOKIES...)")
	var wg sync.WaitGroup
	stdout := make(chan string)
	for path, oven := range j.Oven {
		u, err := url.Parse(path)
		if err != nil {
			colors.Printf("@r(error:) invalid url '%s'", path)
		}
		for _, o := range oven {
			wg.Add(1)
			go o.RequestCookie(jar, u, client, &wg, stdout)
		}
	}
	go func() {
		for {
			select {
			case str, ok := <-stdout:
				if !ok {
					break
				}
				fmt.Println(str)
			}
		}
	}()
	wg.Wait()
	close(stdout)
	client.CloseIdleConnections()
	// static written cookies
	for path, cookies := range j.Cookies {
		u, err := url.Parse(path)
		if err != nil {
			colors.Printf("@r(error:) invalid url '%s'", path)
		}
		jar.SetCookies(u, cookies)
	}
	return jar
}
