package parser

import (
	"tokenizer"
	"errors"
)

var tokens = []tokenizer.Token{}
var current = 0

//func extractTokensFromLine(line string) (tokens []tokenizer.Token, err error) {
func extractTokensFromLine(line string) {
	tokens, _ = tokenizer.Tokenizer(line)
}

func parseExpression() (result float64, err error) {
	return
}

func parseTerm() (result float64, err error) {
	var LNumber, RNumber float64
	LNumber, err = parsePrimaryExpression()
	current ++
	if current >= len(tokens) {
		return LNumber, nil
	}
	opToken := tokens[current]
	current ++
	RNumber, err = parsePrimaryExpression()
	switch opToken.Kind {
	case tokenizer.Add:
		return LNumber + RNumber, nil
	case tokenizer.Sub:
		return LNumber - RNumber, nil
	case tokenizer.Mul:
		return LNumber * RNumber, nil
	case tokenizer.Div:
		if RNumber == 0.0 {
			return 0.0, errors.New("(╯‵□′)╯︵┻━┻ divided by zero")
		}
		return LNumber / RNumber, nil
	}
	return
}

func parsePrimaryExpression() (result float64, err error) {
	token := tokens[current]
	if token.Kind == tokenizer.Number {
		return token.Value, nil
	}
	return 0.0, errors.New("primary expression is not a number")
}

func ParseLine(line string) (result float64, err error) {
	extractTokensFromLine(line)
	if len(tokens) == 0 {
		return 0.0, nil
	}
	return parseExpression()
}
