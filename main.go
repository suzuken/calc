// calc is simple caliculator
//
// reference (Japanese only): http://www.oki-osk.jp/esc/golang/calc.html
package main

import (
	"github.com/suzuken/calc/lex"
	"strings"
)

func main() {
	line := "hoge"
	r := strings.NewReader(line)
	l := lex.New()
	l.Init(r)
}
