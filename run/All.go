package run

import (
	"github.com/xanderazuaje/xocket/parsing"
	"log"
	"strconv"
)

func All(program parsing.Program) (ok bool) {
	ok = true
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
		err, tOk := DoTest(test, program.Endpoint)
		if err != nil {
			log.Fatal(err)
		}
		if !tOk {
			ok = false
		}
	}
	return ok
}

func requiredProperty(i int, p string, test *parsing.Test) {
	if test.Name != "" {
		log.Fatal(p + " is required for executing '" + test.Name + "'")
	}
	log.Fatal(p + " is required for executing test number " + strconv.Itoa(i+1))
}
