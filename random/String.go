package random

import (
	"github.com/xanderazuaje/xocket/types"
	"math/rand/v2"
)

const letters = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\n"

func String(filter *types.Filter) string {
	lenght := 42
	if filter.Len != nil {
		lenght = *filter.Len
	}
	s := make([]byte, lenght)
	for i := range lenght {
		s[i] = letters[rand.IntN(len(letters))]
	}
	return string(s)
}
