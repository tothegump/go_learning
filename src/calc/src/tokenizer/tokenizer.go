package tokenizer

import (
	"strings"
	"strconv"
)

const (
	addLabel = "+"
	subLabel    = "-"
	mulLabel    = "*"
	divLabel    = "/"
	lParenLabel = "("
	rParenLabel = ")"
)

const (
	Add    = "ADD"
	Sub    = "SUB"
	Mul    = "MUL"
	Div    = "DIV"
	LParen = "L_PAREN"
	RParen = "R_PAREN"
	Number = "NUM"
)


type Token struct {
	Kind, Op string
	Value    float64
}


func Tokenizer(line string) ([]Token, error) {
	line += "\n"
	current := 0
	tokens := []Token{}

	for current < len([]rune(line)) {
		ch := string([]rune(line)[current])
		if ch == " " || ch == "\t"{
			current ++
			continue
		}
		switch ch {
		case lParenLabel:
			tokens = append(tokens, Token{
				Kind: LParen,
				Op:   ch,
			})
			current ++
			continue
		case rParenLabel:
			tokens = append(tokens, Token{
				Kind: RParen,
				Op:   ch,
			})
			current ++
			continue
		case addLabel:
			tokens = append(tokens, Token{
				Kind: Add,
				Op:   ch,
			})
			current ++
			continue
		case subLabel:
			tokens = append(tokens, Token{
				Kind: Sub,
				Op:   ch,
			})
			current ++
			continue
		case mulLabel:
			tokens = append(tokens, Token{
				Kind: Mul,
				Op:   ch,
			})
			current ++
			continue
		case divLabel:
			tokens = append(tokens, Token{
				Kind: Div,
				Op:   ch,
			})
			current ++
			continue
		}
		if isNumber(ch) {
			numStr := ""
			hasInt := false
			inFraction := false
			for ; current < len([]rune(line)) && isNumber(ch) || (hasInt && !inFraction && ch == "."); {
				ch := string([]rune(line)[current])
				if hasInt && !inFraction && ch == "." {
					inFraction = true
				} else if !inFraction && isNumber(ch) {
					hasInt = true
				} else if hasInt && inFraction && isNumber(ch) {
				} else {
					break
				}
				c := string(rune(line[current]))
				numStr += c
				current ++
			}
			numValue, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, Token{
				Kind:  Number,
				Value: numValue,
			})
			continue
		}
		break
	}
	return tokens, nil
}

func isNumber(ch string) bool {
	return strings.Index("0123456789", ch) > -1
}
