package random

import (
	"github.com/xanderazuaje/xocket/parsing"
	"strconv"
)

func Filter(filter *parsing.Filter) string {
	switch filter.Type {
	case parsing.FilterString:
		return String(filter)
	case parsing.FilterInteger:
		return strconv.Itoa(Integer(filter))
	case parsing.FilterFloat:
		return strconv.FormatFloat(Float(filter), 'f', -1, 64)
	}
	return ""
}
