package main

import (
	"cursoGo/variables"
	"fmt"
)

func main() {
	estado, texto := variables.ConviertoText(100)
	fmt.Println(estado)
	fmt.Println(texto)
}
