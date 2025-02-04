package colors

import (
	"fmt"
	"github.com/xanderazuaje/xocket/flags"
	"regexp"
	"strings"
	"time"
)

type color string

const (
	Reset   color = "\033[0m"
	Bold    color = "\033[1m"
	Red     color = "\033[31m"
	Green   color = "\033[32m"
	Yellow  color = "\033[33m"
	Blue    color = "\033[34m"
	Magenta color = "\033[35m"
	Cyan    color = "\033[36m"
	White   color = "\033[37m"
)

var colorMap = map[rune]color{
	'*': Bold,
	'r': Red,
	'g': Green,
	'y': Yellow,
	'b': Blue,
	'm': Magenta,
	'c': Cyan,
	'w': White,
}

func Printf(format string, data ...any) {
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
	if flags.This.RunType.Contains(flags.RunDebug) {
		fmt.Print(time.Now().Format("06-01-02 15:04:05   "))
	}
	fmt.Printf(str, data...)
	fmt.Print("\n")
}
