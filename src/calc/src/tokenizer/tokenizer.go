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
	add    = "ADD"
	sub    = "SUB"
	mul    = "MUL"
	div    = "DIV"
	lParen = "L_PAREN"
	rParen = "R_PAREN"
	number = "NUM"
)


type Token struct {
	kind, op string
	value float64
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
				kind: lParen,
				op:   ch,
			})
			current ++
			continue
		case rParenLabel:
			tokens = append(tokens, Token{
				kind: rParen,
				op:   ch,
			})
			current ++
			continue
		case addLabel:
			tokens = append(tokens, Token{
				kind: add,
				op:   ch,
			})
			current ++
			continue
		case subLabel:
			tokens = append(tokens, Token{
				kind: sub,
				op: ch,
			})
			current ++
			continue
		case mulLabel:
			tokens = append(tokens, Token{
				kind: mul,
				op: ch,
			})
			current ++
			continue
		case divLabel:
			tokens = append(tokens, Token{
				kind: div,
				op: ch,
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
				kind:  number,
				value: numValue,
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
