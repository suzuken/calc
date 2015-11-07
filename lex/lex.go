package lex

import (
	"fmt"
	"io"
	"math/big"
	"text/scanner"
)

// Lex is lexical analyzer of calc
type Lex struct {
	Scanner *scanner.Scanner
	// Token is current token
	Token rune
}

// New returns new lexer
func New() *Lex {
	var s scanner.Scanner
	return &Lex{
		Scanner: &s,
	}
}

// Init initialize lexer
func (l *Lex) Init(r io.Reader) {
	l.Scanner.Init(r)
}

// NextToken gets next token to lexer
func (lex *Lex) NextToken() {
	lex.Token = lex.Scanner.Scan()
}

// Error creates error including current token context.
func (lex *Lex) Error(msg string) error {
	return fmt.Errorf("%s: %q", msg, lex.Token)
}

// Zero returns 0 value of big.Rat
func Zero() *big.Rat {
	return big.NewRat(0, 0)
}

// AddSubExp read summuation and subtraction
func AddSubExp(lex *Lex) *big.Rat {
	unaryMinus := false
	switch lex.Token {
	case '+':
		lex.NextToken()
	case '-':
		unaryMinus = true
		lex.NextToken()
	}
	// multiplication and division is prior than subtract and addition
	x := MulDivExp(lex)
	if unaryMinus {
		x = x.Sub(Zero(), x)
	}
LOOP:
	for {
		switch lex.Token {
		case '+':
			lex.NextToken()
			y := MulDivExp(lex)
			x = y.Add(x, y)
		case '-':
			lex.NextToken()
			y := MulDivExp(lex)
			x = y.Sub(x, y)
		default:
			break LOOP
		}
	}
	return x
}

// MulDivExp evaluate multiplication and division
func MulDivExp(lex *Lex) *big.Rat {
	x := UnaryExp(lex)
LOOP:
	for {
		switch lex.Token {
		case '*':
			lex.NextToken()
			y := UnaryExp(lex)
			x = y.Mul(x, y)
		case '/':
			lex.NextToken()
			y := UnaryExp(lex)
			// zero division check is included
			x = y.Quo(x, y)
		default:
			break LOOP
		}
	}
	return x
}

// Unary evaluate unary expression
// include () expression
func UnaryExp(lex *Lex) *big.Rat {
	if lex.Token == '(' {
		lex.NextToken()
		x := AddSubExp(lex)
		if lex.Token != ')' {
			panic(lex.Error("')' expected but not found. exit."))
		}
		lex.NextToken()
		return x
	}
	if lex.Token != scanner.Int && lex.Token != scanner.Float {
		panic(lex.Error(fmt.Sprintf("number expected. Given token %v is not number.", lex.Token)))
	}

	// found if string can represent as rational.
	var r big.Rat
	rat, ok := r.SetString(string(lex.Token))
	if !ok {
		panic(lex.Error("invalid number"))
	}
	lex.NextToken()
	return rat
}
