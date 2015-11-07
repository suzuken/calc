package lex

import (
	"strings"
	"testing"
)

// calcIter helps testing calculate the given expression and its result.
func calcIter(t *testing.T, given, expected string) {
	l := New()
	l.Init(strings.NewReader(given))
	rat, err := Evaluate(l)
	if err != nil {
		t.Fatalf("evaluation error %s", err)
	}
	actual := String(rat)
	if expected != actual {
		t.Fatalf("caliculation failed. given: %s, expected: %s, actual %s", given, expected, actual)
	}
}

func TestCalc(t *testing.T) {
	testSet := [][]string{
		[]string{"1+1", "2"},
		[]string{"1-5", "-4"},
		[]string{"2.1+1", "31/10"},
		[]string{"2*10", "20"},
		[]string{"2/10", "1/5"},
		[]string{"2/10 * 3 + 2", "13/5"},
		[]string{"(1+2) * (3-4)", "-3"},
	}
	for _, ts := range testSet {
		calcIter(t, ts[0], ts[1])
	}
}
