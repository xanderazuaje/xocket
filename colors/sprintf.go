package colors

import (
	"fmt"
	"regexp"
	"strings"
)

func Sprintf(format string, data ...any) string {
	r := regexp.MustCompile("@[rgybmcw*]{1,2}\\(.[^)]*\\)")
	normal := r.Split(format, -1)
	values := r.FindAllString(format, -1)
	var str string
	for i, v := range normal {
		str += v
		if i < len(values) {
			v = values[i][1:]
			for _, c := range v {
				if c == '(' {
					break
				}
				str += string(colorMap[c])
			}
			v = strings.TrimLeft(v, "*rgybmcw")
			v = strings.Trim(v, "()")
			str += v
			str += string(Reset)
		}
	}
	return fmt.Sprintf(str, data...)
}
