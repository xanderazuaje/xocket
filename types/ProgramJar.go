package types

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"log"
	"net/http"
	"net/http/cookiejar"
	"sync"
)

type ProgramJar struct {
	Oven    []*CookieOven
	Cookies []*http.Cookie
}

func (j *ProgramJar) GetJar() *cookiejar.Jar {
	client := http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal("failed to create cookie jar:", err.Error())
	}
	colors.Log("@*g(GETTING COOKIES...)")
	var wg sync.WaitGroup
	stdout := make(chan string)
	var totalCalls int
	for _, oven := range j.Oven {
		totalCalls++
		wg.Add(totalCalls)
		go oven.RequestCookie(jar, client, &wg, stdout)
	}
	go func() {
		for {
			select {
			case str := <-stdout:
				fmt.Println(str)
				totalCalls--
				if totalCalls == 0 {
					return
				}
			}
		}
	}()
	wg.Wait()
	client.CloseIdleConnections()
	return jar
}
