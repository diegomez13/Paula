package main

import (
	"fmt"
)

func main() {

	var accion string

	fmt.Println("Que quieres hacer? (mover/atacar/salir)")
	fmt.Scanln(&accion)

	if accion == "mover" {
		fmt.Println("Te mueves hacia adelante")
	} else if accion == "Atacar" {
		fmt.Println("Lanzar un ataque poderoso")
	} else if accion == "Salir" {
		fmt.Println("Haz salido del juego")
	} else {
		fmt.Println("No entiendo la acción")

	}

}
