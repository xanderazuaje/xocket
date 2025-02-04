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
	Oven    []*CookieOven
	Cookies []*http.Cookie
}

func (j *ProgramJar) FillJar() *cookiejar.Jar {
	client := http.Client{}
	u, _ := url.Parse("http://0.0.0.0")
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("failed to create cookie jar:", err.Error())
	}
	colors.Log("@*g(GETTING COOKIES...)")
	var wg sync.WaitGroup
	stdout := make(chan string)
	for _, oven := range j.Oven {
		wg.Add(1)
		go oven.RequestCookie(jar, u, client, &wg, stdout)
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
	jar.SetCookies(u, j.Cookies)
	return jar
}
