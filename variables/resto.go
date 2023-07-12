package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nomnbre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nomnbre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()

	fmt.Println(Nomnbre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoText(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
