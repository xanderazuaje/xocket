package random

import (
	"github.com/xanderazuaje/xocket/parsing"
	"math/rand"
)

func Float(filter *parsing.Filter) float64 {
	r := rand.Float64()
	if filter.Max != nil {
		r = *filter.Max * r
	}
	if filter.Min != nil {
		r += *filter.Min
	}
	return r
}
