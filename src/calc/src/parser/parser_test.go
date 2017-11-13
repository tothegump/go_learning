package parser

import (
	"testing"
)


func TestTerm(t *testing.T) {
	extractTokensFromLine("34.3 + 5")
	r, e := parseTerm()
	if e != nil {
		t.Error(e)
	}
	if r != 39.3 {
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
	extractTokensFromLine("34.3")
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
