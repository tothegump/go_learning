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
		if token.Kind != Number {
			t.Error("unexpect Kind")
		}
		if token.Value != values[i] {
			t.Error("unexpect Value")
		}
	}
}

func TestTokenizerForOP(t *testing.T) {
	tokens, err := Tokenizer("(+-")
	ops := []string{LParen, Add, Sub}
	values := []string{"(", "+", "-"}
	if err != nil {
		fmt.Println(err)
		t.Error("unexpect err")
	}
	if len(tokens) != 3 {
		t.Error("unexpect count")
	}
	for i, token := range tokens {
		if token.Kind != ops[i] {
			t.Error("unexpect Kind")
		}
		if token.Op != values[i] {
			t.Error("unexpect Value")
		}
	}
}

func TestIsNumber(t *testing.T) {
	r := isNumber("0")
	if !r {
		t.Error("0 is a Number")
	}

	r = isNumber("12")
	if !r {
		t.Error("0 is a Number")
	}

	r = isNumber("a")
	if r {
		t.Error("0 is a Number")
	}

}
