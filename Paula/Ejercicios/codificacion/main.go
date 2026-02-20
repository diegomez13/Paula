package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	AnimalEncode("Animales.json")
	AnimalDecode("Animales.json")
}

type Animales struct {
	Especie string
	Raza    string
	Edad    int
}

func AnimalEncode(file string) {
	animales := []Animales{}
	animales = append(animales, Animales{Especie: "perro", Raza: "Chiguagua", Edad: 1})
	animales = append(animales, Animales{Especie: "perro", Raza: "Chiguagua", Edad: 1})
	animales = append(animales, Animales{Especie: "perro", Raza: "Chiguagua", Edad: 1})

	//fmt.Println(perro)
	data, err := json.Marshal(animales)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func AnimalDecode(file string) {

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var u []Animales
	err = json.Unmarshal(data, &u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
