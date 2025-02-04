package diffPrint

import (
	"fmt"
	"github.com/xanderazuaje/xocket/colors"
	"strings"
)

func printStrArrDiff(sa1, sa2 []string) {
	sa1len := len(sa1)
	sa2len := len(sa2)
	str := "\t[ %s ]\n\t[ %s ]"
	strA1 := make([]string, sa1len)
	strA2 := make([]string, sa2len)
	for i, v := range sa1 {
		if v != sa2[i] {
			strA1[i] = fmt.Sprintf("@r*(%s)", v)
			strA2[i] = fmt.Sprintf("@r*(%s)", sa2[i])
		} else {
			strA1[i] = v
			strA2[i] = sa2[i]
		}
	}
	if sa1len > sa2len {
		s := sa1[sa2len:]
		for i, v := range s {
			s[i] = fmt.Sprintf("@r*(%s)", v)
		}
		strA2 = append(strA2, s...)
	} else if sa2len > sa1len {
		s := sa2[sa1len:]
		for i, v := range s {
			s[i] = fmt.Sprintf("@r*(%s)", v)
		}
		strA2 = append(strA1, s...)
	}
	colors.Printf(fmt.Sprintf(str, strings.Join(strA1, ", "), strings.Join(strA2, ", ")))
}
