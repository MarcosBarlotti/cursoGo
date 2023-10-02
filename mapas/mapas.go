package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println("Argentina")

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  30,
		"Chivas":       37,
		"Boca juniors": 30,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	delete(campeonato, "Real Madrid")

	fmt.Println(campeonato)

}
