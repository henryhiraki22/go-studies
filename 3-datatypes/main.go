package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 1
	fmt.Println(numero)

	var str string = "Texto"
	fmt.Println(str)

	var booleano1 bool = true
	var booleano2 bool = false

	fmt.Println(booleano1)
	fmt.Println(booleano2)

	var erro error = errors.New("Erro interno")

	fmt.Println(erro)
}
