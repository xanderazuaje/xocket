package diffPrint

import (
	"github.com/xanderazuaje/xocket/colors"
	"github.com/xanderazuaje/xocket/flags"
)

func printStringDiff(isOk *bool, name, s1, s2 string) {
	if s1 != s2 {
		colors.Printf(
			"@b(%s) - @r*(DIFF:)\n\t@b*(expected:) -> '%s'\n\t@r*(got:) -> '%s'",
			name,
			s1,
			s2,
		)
		*isOk = false
	} else if flags.This.RunType.Contains(flags.RunDebug) {
		colors.Printf("@b(%s) - @g*(OK!)", name)
	}
}
