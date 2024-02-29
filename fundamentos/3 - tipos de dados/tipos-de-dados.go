package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 rune = 1234
	fmt.Println(numero2)

	var numero3 byte = 123
	fmt.Println(numero3)

	var numero4 float64 = 123.45
	fmt.Println(numero4)

	var texto string = "fafsfsafsa"
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Internal server error")
	fmt.Println(erro)
}
