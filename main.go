// calc is simple caliculator
//
// It's functionality is based on big.Rat.
//
// Evaluator is based on this reference (Japanese only): http://www.oki-osk.jp/esc/golang/calc.html
//
// Copyright (c) 2013 OKI Software Co., Ltd.
// Copyright (c) 2015 Kenta Suzuki
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL
// THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.
package main

import (
	"bufio"
	"fmt"
	"github.com/suzuken/calc/lex"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewScanner(os.Stdin)
	printHeader()
	for buf.Scan() {
		calc(buf.Text())
		printHeader()
	}
}

func printHeader() {
	fmt.Print("(calc) > ")
}

func printError(err error) {
	fmt.Printf("[ERROR] %s\n", err)
}

func calc(line string) {
	r := strings.NewReader(line)
	l := lex.New()
	l.Init(r)

	rat, err := lex.Evaluate(l)
	if err != nil {
		printError(err)
	} else {
		lex.Print(rat)
	}
}
