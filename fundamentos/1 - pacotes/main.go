package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("lucas@gmail.com")
	fmt.Println(erro)
}
