package syntax_token

type TokenType struct {
	Name, Description string
	Numeric           bool
	From              int64
	To                int64
}

// These functions returns a constant value
// that is used to "simulate a constant struct"

func _UINT8() TokenType {
	return TokenType{
		Name:        "uint8",
		Description: "Unsigned integer of 8 bits",
		Numeric:     true,
		From:        0,
		To:          255,
	}
}

func _UINT16() TokenType {
	return TokenType{
		Name:        "uint16",
		Description: "Unsigned integer of 16 bits",
		Numeric:     true,
		From:        0,
		To:          65535,
	}
}

func _SINT8() TokenType {
	return TokenType{
		Name:        "sint8",
		Description: "Signed integer of 8 bits",
		Numeric:     true,
		From:        -128,
		To:          127,
	}
}

func _SINT16() TokenType {
	return TokenType{
		Name:        "sint16",
		Description: "Signed integer of 16 bits",
		Numeric:     true,
		From:        -32768,
		To:          32767,
	}
}

func _ADD() TokenType {
	return TokenType{
		Name:        "addition",
		Description: "Addition operator",
	}
}

func _SUB() TokenType {
	return TokenType{
		Name:        "subtraction",
		Description: "Subtraction operator",
	}
}

func _MUL() TokenType {
	return TokenType{
		Name:        "multiplication",
		Description: "Multiplication operator",
	}
}

func _DIV() TokenType {
	return TokenType{
		Name:        "division",
		Description: "Division operator",
	}
}

func _LPARENTHESIS() TokenType {
	return TokenType{
		Name:        "left parenthesis",
		Description: "Left parentesis",
	}
}

func _RPARENTHESIS() TokenType {
	return TokenType{
		Name:        "right parenthesis",
		Description: "Right parentesis",
	}
}
