package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	emptyArgs string = "Don't panic!"
	gopher    string = `  \
   \
    \    ,_---~~~~~----._
  _,,_,*^____      _____\˴ᐠ*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \ᐠ/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |`
)

func maxLen(text ...string) int {
	ml := 0
	for _, str := range text {
		sl := len(str)
		if ml < sl {
			ml = sl
		}
	}

	return ml
}

func gopherSay(text ...string) string {
	if len(text) == 0 {
		text = append(text, emptyArgs)
	}

	var sb strings.Builder
	ln := maxLen(text...)
	sb.WriteString(fmt.Sprintf("+-%s-+\n", strings.Repeat("-", ln)))
	for _, str := range text {
		sb.WriteString(fmt.Sprintf("| %s%s |\n", str, strings.Repeat(" ", ln-len(str))))
	}
	sb.WriteString(fmt.Sprintf("+-%s-+\n", strings.Repeat("-", ln)))
	sb.WriteString(gopher)
	return sb.String()
}

func main() {
	fmt.Print(gopherSay(os.Args[1:]...))
}
