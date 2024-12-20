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
