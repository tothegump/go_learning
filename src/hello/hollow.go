package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)


type token struct {
	kind, value string
}

func tokenizer(line string) []token {
	fmt.Println("start..")
	line += "\n"
	current := 0
	tokens := []token{}

	for current < len([]rune(line)) {
		ch := string([]rune(line)[current])
		if ch == "(" {
			tokens = append(tokens, token{
				kind: "paren",
				value: "(",
			})
			current ++
			continue
		}
		if ch == ")" {
			tokens = append(tokens, token{
				kind: "paren",
				value: ")",
			})
			current ++
			continue
		}
		if ch == " " {
			current ++
			continue
		}
		if strings.Index("+-*/><=", ch) > -1 {
			tokens = append(tokens, token{
				kind: "op",
				value: ch,
			})
			current ++
			continue
		}
		if isNumber(ch) {
			num := ""
			c := ""
			for ; isNumber(c); current++ {
				c = string(rune(line[current]))
				num += c
			}
			tokens = append(tokens, token {
				kind: "number",
				value: num,
			})
			continue
		}
		break
	}
	return tokens
}

func isNumber(ch string) bool {
	return strings.Index("0123456789", ch) > -1
}

func main() {
	fmt.Println(n)
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	//line := "2134 + 33"
	//n := [100]token{}
	//for pos, ch := range line {
	//	fmt.Println(pos, unicode.IsDigit(rune(ch)))
	//}
	tokens := tokenizer(line)
	for i :=0; i < len(tokens); i++ {
		fmt.Println(tokens[i].kind, tokens[i].value)
	}

}
