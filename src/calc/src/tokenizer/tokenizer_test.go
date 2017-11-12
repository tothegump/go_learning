package tokenizer

import (
	"testing"
	"fmt"
)

func TestTokenizer(t *testing.T) {
	tokens, err := Tokenizer("3 + ( 15.3 - 9.25 ) * 10")
	if err != nil {
		t.Error("unexpect err")
	}
	fmt.Println(tokens)
}

func TestTokenizerForNumber(t *testing.T) {
	tokens, err := Tokenizer("12 134 33.3")
	values := []float64{12, 134, 33.3}
	if err != nil {
		fmt.Println(err)
		t.Error("unexpect err")
	}
	if len(tokens) != 3 {
		t.Error("unexpect count")
	}
	for i, token := range tokens {
		if token.kind != number {
			t.Error("unexpect kind")
		}
		if token.value != values[i] {
			t.Error("unexpect value")
		}
	}
}

func TestTokenizerForOP(t *testing.T) {
	tokens, err := Tokenizer("(+-")
	ops := []string{lParen, add, sub}
	values := []string{"(", "+", "-"}
	if err != nil {
		fmt.Println(err)
		t.Error("unexpect err")
	}
	if len(tokens) != 3 {
		t.Error("unexpect count")
	}
	for i, token := range tokens {
		if token.kind != ops[i] {
			t.Error("unexpect kind")
		}
		if token.op != values[i] {
			t.Error("unexpect value")
		}
	}
}

func TestIsNumber(t *testing.T) {
	r := isNumber("0")
	if !r {
		t.Error("0 is a number")
	}

	r = isNumber("12")
	if !r {
		t.Error("0 is a number")
	}

	r = isNumber("a")
	if r {
		t.Error("0 is a number")
	}

}
