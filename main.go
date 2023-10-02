package main

import (
	e "cursoGo/ejer_interfaces"
	"cursoGo/modelos"
)

func main() {
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)

}
