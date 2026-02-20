package main

import (
	"bufio"
	"fmt"
	"modules/bases"
	"os"
	"strings"
)

func main() {

	var accion string

	bases.Test()

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

	reader := bufio.NewReader(os.Stdin)
	userSistema, _ := reader.ReadString('\n')
	userSistema = strings.TrimSpace(userSistema)

}
