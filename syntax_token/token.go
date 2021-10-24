package syntax_token

import (
	"strconv"
	"strings"
)

type Token struct {
	Type  TokenType
	Value interface{}
}

func Parse(row string) []Token {
	var tkArr []Token = make([]Token, 0)
	var words []string = strings.Split(row, " ")

	for _, word := range words {
		tk := Token{}

		switch word {
		case "+":
			tk = Token{Type: _ADD()}

		case "-":
			tk = Token{Type: _SUB()}

		case "*":
			tk = Token{Type: _MUL()}

		case "/":
			tk = Token{Type: _DIV()}

		case "(":
			tk = Token{Type: _LPARENTHESIS()}

		case ")":
			tk = Token{Type: _RPARENTHESIS()}
		}

		if tk != (Token{}) {
			tkArr = append(tkArr, tk)
			continue
		}

		if parsedInt, err := strconv.ParseInt(word, 10, 64); err == nil {
			switch {
			case parsedInt >= _UINT8().From && parsedInt <= _UINT8().To:
				tk = Token{_UINT8(), uint8(parsedInt)}

			case parsedInt >= _UINT16().From && parsedInt <= _UINT16().To:
				tk = Token{_UINT16(), uint16(parsedInt)}

			case parsedInt >= _SINT8().From && parsedInt <= _SINT8().To:
				tk = Token{_SINT8(), int8(parsedInt)}

			case parsedInt >= _SINT16().From && parsedInt <= _SINT16().To:
				tk = Token{_SINT16(), int16(parsedInt)}
			}
		}

		if tk != (Token{}) {
			tkArr = append(tkArr, tk)
			continue
		}
	}

	return tkArr
}
