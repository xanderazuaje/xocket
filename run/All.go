package run

import (
	"github.com/xanderazuaje/xocket/types"
	"log"
	"net/http/cookiejar"
	"strconv"
)

func All(program types.Program) (ok bool) {
	ok = true
	var jar *cookiejar.Jar
	if program.CookieJar != nil {
		jar = program.CookieJar.FillJar()
	}
	//Do each test
	for i, test := range program.Tests {
		if test.Path == "" {
			requiredProperty(i, "path", &test)
		}
		if test.Method == "" {
			requiredProperty(i, "method", &test)
		}
		if test.Form != nil && (test.Form.Type != "multipart" && test.Form.Type != "urlencoded") {
			if test.Name != "" {
				log.Fatal("invalid form type executing '" + test.Name + "', only 'multipart' or 'urlencoded' is valid")
			}
			log.Fatal("invalid form type executing test number " + strconv.Itoa(i+1) + ", only 'multipart' or 'urlencoded' is valid")
		}
		var err error
		if !test.IgnoreCookieJar && program.CookieJar != nil {
			err, ok = DoTest(test, program.Endpoint, jar)
		} else {
			err, ok = DoTest(test, program.Endpoint, nil)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return ok
}

func requiredProperty(i int, p string, test *types.Test) {
	if test.Name != "" {
		log.Fatal(p + " is required for executing '" + test.Name + "'")
	}
	log.Fatal(p + " is required for executing test number " + strconv.Itoa(i+1))
}
