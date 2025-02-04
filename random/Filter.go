package random

import (
	"github.com/xanderazuaje/xocket/types"
	"strconv"
)

func Filter(filter *types.Filter) string {
	switch filter.Type {
	case types.FilterString:
		return String(filter)
	case types.FilterInteger:
		return strconv.Itoa(Integer(filter))
	case types.FilterFloat:
		return strconv.FormatFloat(Float(filter), 'f', -1, 64)
	}
	return ""
}
