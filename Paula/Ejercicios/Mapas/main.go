/*package main

import "fmt"

func main() {

	edades := map[string]int{

		"Paula":  25,
		"Renato": 13,
		"Luz":    28,
	}

	fmt.Println("Mapa completo:", edades)

	fmt.Println("Edad de Paula:", edades["Paula"])

	edades["Juan"] = 22

	valor, existe := edades["Luz"]
	if existe {
		fmt.Println("Luz tiene", valor, "anos")
	} else {

		fmt.Println("Luz no está en el mapa")

	}

	fmt.Println("\nListado de edades:")
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d anos\n", nombre, edad)
	}

}
*/
/*
package main

import "fmt"

type Zoologico struct {
	Nombre string
	Areas  []Area
}

type Area struct {
	Id       int
	Nombre   string
	Animales []Animal
}

type Animal struct {
	Id      int
	Nombre  string
	Especie string
	Edad    int
}

func main() {

	z := Zoologico{Nombre: "Zoologico del centro", Areas: []Area{}}

	z.AgregarArea(1, "Sabana Africana")
	z.AgregarArea(2, "Zona Acuatica")
	z.AgregarArea(3, "Reptiliano")

	z.AgregarAnimal(1, 101, "Simba", "Leon", 5)
	z.AgregarAnimal(1, 102, "Dumbo", "Elefante", 10)
	z.AgregarAnimal(2, 103, "Nemo", "Pez Payaso", 2)
	z.AgregarAnimal(3, 104, "Drilo", "Cocodrilo", 8)
	z.AgregarAnimal(3, 105, "Tuga", "Tortuga", 100)
	z.AgregarAnimal(2, 106, "Doris", "Pez Cirujano", 4)
	z.AgregarAnimal(1, 107, "Sebastian", "Pelicano", 9)

	z.ModificarAnimal(1, 101, "Mufasa", "Leon", 6)

	z.MostrarZoologico()

}

func (z *Zoologico) AgregarArea(id int, nombre string) {
	z.Areas = append(z.Areas, Area{Id: id, Nombre: nombre})
}

func (z *Zoologico) AgregarAnimal(idArea int, idAnimal int, nombre, especie string, edad int) {
	for i := 0; i < len(z.Areas); i++ {
		if z.Areas[i].Id == idArea {
			z.Areas[i].Animales = append(z.Areas[i].Animales, Animal{Id: idAnimal, Nombre: nombre, Especie: especie, Edad: edad})
			return
		}
	}
}

func (z *Zoologico) ModificarAnimal(idArea int, idAnimal int, nuevoNombre, nuevaEspecie string, nuevaEdad int) {
	for i := 0; i < len(z.Areas); i++ {
		if z.Areas[i].Id == idArea {
			for j := 0; j < len(z.Areas[i].Animales); j++ {
				if z.Areas[i].Animales[j].Id == idAnimal {
					z.Areas[i].Animales[j].Nombre = nuevoNombre
					z.Areas[i].Animales[j].Especie = nuevaEspecie
					z.Areas[i].Animales[j].Edad = nuevaEdad
					return

				}
			}

		}

	}
}

func (z *Zoologico) MostrarZoologico() {
	fmt.Println("Zoologico:", z.Nombre)
	for _, area := range z.Areas {
		fmt.Println("- Area:", area.Nombre)
		for _, animal := range area.Animales {
			fmt.Printf(" * Animal: %s (%s), Edad: %d anos\n", animal.Nombre, animal.Especie, animal.Edad)
		}
	}

}
*/

package main

import "fmt"

type Colegio struct {
	Direccion string
	Numero    int
	Salas     []Sala
}

type Sala struct {
	Numero    int
	Curso     string
	Capacidad int
}

func main() {

	colegio := Colegio{
		Direccion: "Vitacura",
		Numero:    33,
		Salas: []Sala{
			Sala{Numero: 1, Curso: "Primero", Capacidad: 30},
			Sala{Numero: 2, Curso: "Segundo", Capacidad: 28},
		},
	}

	colegios := []Colegio{
		Colegio{
			Direccion: "Larrain",
			Numero:    101,
			Salas: []Sala{
				Sala{Numero: 3, Curso: "Primero", Capacidad: 25},
				Sala{Numero: 4, Curso: "Segundo", Capacidad: 20},
				Sala{Numero: 5, Curso: "Tercero", Capacidad: 22},
				Sala{Numero: 6, Curso: "Cuarto", Capacidad: 33},
				Sala{},
				Sala{},
			},
		},
		colegio,
	}

	Supercolegio := [][]Colegio{[]Colegio{
		Colegio{
			Direccion: "Providencia",
			Numero:    11,
			Salas: []Sala{
				Sala{Numero: 7, Curso: "Primero", Capacidad: 21},
				Sala{Numero: 8, Curso: "Segundo", Capacidad: 25},
				Sala{Numero: 9, Curso: "Tercero", Capacidad: 12},
				Sala{Numero: 10, Curso: "Cuarto", Capacidad: 13},
			},
		},
	}, colegios}

	Supercolegio = append(Supercolegio, colegios)

	for i := 0; i < len(Supercolegio); i++ {
		//fmt.Println(i, Supercolegio[i])
		for j := 0; j < len(Supercolegio[i]); j++ {
			fmt.Println(j, Supercolegio[i][j])

		}

	}
	//fmt.Println("Colegio:", colegio)

	fmt.Println("---------------------------")

	mapHE := make(map[string][][]Colegio, 0)
	mapHE["Colegio1"] = Supercolegio

	if mape, found := mapHE["Colegio1"]; found {
		fmt.Println(mape)
	} else {
		fmt.Println("No se encontro")
	}
}
