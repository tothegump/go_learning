package parser

import (
	"testing"
	"fmt"
)

func TestExpression(t *testing.T) {
	extractTokensFromLine("8 + (24 + 9 - 6 * 2) / 7 + 100")
	r, e := parseExpression()
	if e != nil {
		t.Error(e)
	}
	fmt.Println(r)
	if r != 111.0 {
		t.Error("compute error")
	}
}

func TestTerm(t *testing.T) {
	extractTokensFromLine("24 * 9 / 6")
	r, e := parseTerm()
	if e != nil {
		t.Error(e)
	}
	if r != 36.0 {
		t.Error("compute error")
	}
}

func TestTermFailed(t *testing.T) {
	extractTokensFromLine("34.3 / 0")
	_, e := parseTerm()
	if e == nil {
		t.Error("expect error here")
	}
}

func TestPrimaryExpression(t *testing.T) {
	extractTokensFromLine("(34.3 - 9 + 10 -1)")
	r, e := parsePrimaryExpression()
	if e != nil {
		t.Error(e)
	}
	if r != 34.3 {
		t.Error("unexpect number value")
	}
}

func TestPrimaryExpressionFailed(t *testing.T) {
	extractTokensFromLine("+")
	r, e := parsePrimaryExpression()
	if r != 0.0 {
		t.Error("expected a 0.0 when parse failed")
	}
	if e == nil {
		t.Error("expected a error here")
	}
}
