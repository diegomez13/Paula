package main

import "fmt"

type Edificio struct {
	Direccion     string
	Numero        int
	Departamentos []Departamento
}

type Departamento struct {
	Numero       int
	MtsCuadrados int
	Piezas       int
}

func main() {

	edificio := Edificio{Direccion: "Jose Tomas Rider", Numero: 5566, Departamentos: []Departamento{Departamento{Numero: 333, MtsCuadrados: 120, Piezas: 4}, Departamento{Numero: 111, MtsCuadrados: 4, Piezas: 1}}}

	edificios := []Edificio{Edificio{Direccion: "Pedro de Valdicia", Numero: 999, Departamentos: []Departamento{Departamento{Numero: 777, MtsCuadrados: 67, Piezas: 9}}}, Edificio{}, Edificio{}, edificio}

	Superedificio := [][]Edificio{[]Edificio{}, edificios}

	Hiperedificio := [][][]Edificio{[][]Edificio{[]Edificio{edificio, Edificio{Direccion: "Bilabao", Numero: 55, Departamentos: []Departamento{Departamento{Numero: 323, MtsCuadrados: 88, Piezas: 1}}}}, []Edificio{}}, [][]Edificio{}, Superedificio}

	/*	for i := 0; i < len(Hiperedificio); i++ {
			//fmt.Println(i, Hiperedificio[i])
			for j := 0; j < len(Hiperedificio[i]); j++ {
				//fmt.Println(i, j, Hiperedificio[i][j])
				for k := 0; k < len(Hiperedificio[i][j]); k++ {
					//fmt.Println(i, j, k, Hiperedificio[i][j][k])
					fmt.Println(i, j, k, Hiperedificio[i][j][k].Departamentos)
					for l := 0; l < len(Hiperedificio[i][j][k].Departamentos); l++ {
						fmt.Println("Mts Cuadrados", Hiperedificio[i][j][k].Departamentos[l].MtsCuadrados)
						fmt.Println("Piezas", Hiperedificio[i][j][k].Departamentos[l].Piezas)
						fmt.Println("Numero", Hiperedificio[i][j][k].Departamentos[l].Numero)
						fmt.Println("-----------------------------------")

					}
					fmt.Println("===================================")
				}
			}
		}
	*/
	mapHE := make(map[string][][][]Edificio, 0)
	mapHE["Edif1"] = Hiperedificio

	if mape, found := mapHE["Edif2"]; found {
		fmt.Println(mape)
	} else {
		fmt.Println("No se encontro")
	}
}
