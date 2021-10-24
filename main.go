package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/adalberto-lang/syntax_token"
)

func main() {
	row := "+ - * / ( ) 1 256 -1 -129"
	tokens := syntax_token.Parse(row)
	fmt.Println(tokens)

	var input string
	fmt.Scanln(&input)
}
