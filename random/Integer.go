package random

import (
	"fmt"
	"github.com/xanderazuaje/xocket/parsing"
	"math/rand"
)

func Integer(filter *parsing.Filter) int {
	r := rand.Float64()
	if filter.Max != nil {
		r = *filter.Max * r
	} else {
		r = 100 * r
	}
	if filter.Min != nil {
		r += *filter.Min
		fmt.Printf("%f\n", r)
	}
	return int(r)
}
