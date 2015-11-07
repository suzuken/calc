// calc is simple caliculator
//
// reference (Japanese only): http://www.oki-osk.jp/esc/golang/calc.html
package main

import (
	"bufio"
	"github.com/suzuken/calc/lex"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		calc(buf.Text())
	}
}

func calc(line string) {
	r := strings.NewReader(line)
	l := lex.New()
	l.Init(r)

	rat, err := lex.Evaluate(l)
	if err != nil {
		panic(err)
	}
	lex.Print(rat)
}
