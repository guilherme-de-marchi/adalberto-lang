package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	row := "+ - * / ( ) 1 256 -1 -129"
	tokens := parse_token(row)
	fmt.Println(tokens)

	var input string
	fmt.Scanln(&input)
}

/*
---------------------
	Syntax Token
---------------------
*/

func parse_token(row string) []token {
	var tkArr []token = make([]token, 0)
	var words []string = strings.Split(row, " ")

	for _, word := range words {
		tk := token{}

		switch word {
		case "+":
			tk = token{Type: _ADD()}

		case "-":
			tk = token{Type: _SUB()}

		case "*":
			tk = token{Type: _MUL()}

		case "/":
			tk = token{Type: _DIV()}

		case "(":
			tk = token{Type: _LPARENTHESIS()}

		case ")":
			tk = token{Type: _RPARENTHESIS()}
		}

		if tk != (token{}) {
			tkArr = append(tkArr, tk)
			continue
		}

		if parsedInt, err := strconv.ParseInt(word, 10, 64); err == nil {
			switch {
			case parsedInt >= _UINT8().From && parsedInt <= _UINT8().To:
				tk = token{_UINT8(), uint8(parsedInt)}

			case parsedInt >= _UINT16().From && parsedInt <= _UINT16().To:
				tk = token{_UINT16(), uint16(parsedInt)}

			case parsedInt >= _SINT8().From && parsedInt <= _SINT8().To:
				tk = token{_SINT8(), int8(parsedInt)}

			case parsedInt >= _SINT16().From && parsedInt <= _SINT16().To:
				tk = token{_SINT16(), int16(parsedInt)}
			}
		}

		if tk != (token{}) {
			tkArr = append(tkArr, tk)
			continue
		}
	}

	return tkArr
}

type token struct {
	Type  tokenType
	Value interface{}
}

type tokenType struct {
	Name, Description string
	Numeric           bool
	From              int64
	To                int64
}

// These functions returns a constant value
// that is used to "simulate a constant struct"

func _UINT8() tokenType {
	return tokenType{
		Name:        "uint8",
		Description: "Unsigned integer of 8 bits",
		Numeric:     true,
		From:        0,
		To:          255,
	}
}

func _UINT16() tokenType {
	return tokenType{
		Name:        "uint16",
		Description: "Unsigned integer of 16 bits",
		Numeric:     true,
		From:        0,
		To:          65535,
	}
}

func _SINT8() tokenType {
	return tokenType{
		Name:        "sint8",
		Description: "Signed integer of 8 bits",
		Numeric:     true,
		From:        -128,
		To:          127,
	}
}

func _SINT16() tokenType {
	return tokenType{
		Name:        "sint16",
		Description: "Signed integer of 16 bits",
		Numeric:     true,
		From:        -32768,
		To:          32767,
	}
}

func _ADD() tokenType {
	return tokenType{
		Name:        "addition",
		Description: "Addition operator",
	}
}

func _SUB() tokenType {
	return tokenType{
		Name:        "subtraction",
		Description: "Subtraction operator",
	}
}

func _MUL() tokenType {
	return tokenType{
		Name:        "multiplication",
		Description: "Multiplication operator",
	}
}

func _DIV() tokenType {
	return tokenType{
		Name:        "division",
		Description: "Division operator",
	}
}

func _LPARENTHESIS() tokenType {
	return tokenType{
		Name:        "left parenthesis",
		Description: "Left parentesis",
	}
}

func _RPARENTHESIS() tokenType {
	return tokenType{
		Name:        "right parenthesis",
		Description: "Right parentesis",
	}
}
