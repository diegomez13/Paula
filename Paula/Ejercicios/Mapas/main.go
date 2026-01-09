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
/*
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
		Numero:    1154,
		Salas: []Sala{
			Sala{Numero: 1, Curso: "Primero", Capacidad: 30},
			Sala{Numero: 2, Curso: "Segundo", Capacidad: 28},
		},
	}

	colegios := []Colegio{
		Colegio{
			Direccion: "Larrain",
			Numero:    1289,
			Salas: []Sala{
				Sala{Numero: 3, Curso: "Primero", Capacidad: 25},
				Sala{Numero: 4, Curso: "Segundo", Capacidad: 20},
				Sala{Numero: 5, Curso: "Tercero", Capacidad: 22},
				Sala{Numero: 6, Curso: "Cuarto", Capacidad: 33},
				Sala{Numero: 11, Curso: "Quinto", Capacidad: 30},
				Sala{Numero: 12, Curso: "Sexto", Capacidad: 40},
			},
		},
		colegio,
	}

	Supercolegio := [][]Colegio{[]Colegio{
		Colegio{
			Direccion: "Providencia",
			Numero:    1522,
			Salas: []Sala{
				Sala{Numero: 7, Curso: "Primero", Capacidad: 21},
				Sala{Numero: 8, Curso: "Segundo", Capacidad: 25},
				Sala{Numero: 9, Curso: "Tercero", Capacidad: 12},
				Sala{Numero: 10, Curso: "Cuarto", Capacidad: 13},
			},
		},
	},
	}

	Supercolegio = append(Supercolegio, colegios)

	for i := 0; i < len(Supercolegio); i++ {
		//fmt.Println(i, Supercolegio[i])
		for j := 0; j < len(Supercolegio[i]); j++ {
			//fmt.Println(j, Supercolegio[i][j])
			//fmt.Printf("direccion: %v Numero: %v Cantidad de Salas: %v \n", Supercolegio[i][j].Direccion, Supercolegio[i][j].Numero, len(Supercolegio[i][j].Salas))
			for k := 0; k < len(Supercolegio[i][j].Salas); k++ {
				//fmt.Printf("Numero: %v Curso: %v Capacidad: %v \n", Supercolegio[i][j].Salas[k].Numero, Supercolegio[i][j].Salas[k].Curso, Supercolegio[i][j].Salas[k].Capacidad)
				if Supercolegio[i][j].Salas[k].Capacidad < 30 && len(Supercolegio[i][j].Salas) > 1 {
					fmt.Printf("Direccion: %v Numero: %v Curso: %v \n", Supercolegio[i][j].Direccion, Supercolegio[i][j].Numero, Supercolegio[i][j].Salas[k].Curso)

				}
			}
		}

	}
	//fmt.Println("Colegio:", colegio)

	fmt.Println("---------------------------")

	mapSC := make(map[string][][]Colegio, 0)
	mapSC["Colegio1"] = Supercolegio

	if mape, found := mapSC["Colegio2"]; found {
		fmt.Println(mape)
	} else {
		fmt.Println("No se encontro")
	}
}
*/
/*
package main

import "fmt"

type Veterinaria struct {
	Nombre    string
	Direccion string
	Mascotas  []Mascota
}

type Mascota struct {
	Nombre string
	Tipo   string
	Edad   int
}

func main() {

	veterinaria := Veterinaria{
		Nombre:    "Perritos",
		Direccion: "Bilbao 1181",
		Mascotas: []Mascota{Mascota{},
			{Nombre: "Lila", Tipo: "perro", Edad: 8},
			{Nombre: "Margarita", Tipo: "gato", Edad: 5},
		},
	}

	veterinarias := []Veterinaria{Veterinaria{
			Nombre:    "Gatitos",
			Direccion: "Manuel Montt 5566",
			Mascotas: []Mascota{
				{Nombre: "Roky", Tipo: "perro", Edad: 12},
				{Nombre: "Mila", Tipo: "gato", Edad: 9},
				{Nombre: "Pupi", Tipo: "gato", Edad: 1},
			},
		},
	}

	Superveterinaria := [][]Veterinaria{[]Veterinaria{}, veterinarias}
			{
				Nombre:    "Superpet"
				Direccion: "Tobalaba 1567"
				Mascotas: []Mascota{
					Mascota{Nombre: "Rex", Tipo: "gato", Edad: 15},
					Mascota{Nombre: "Coco", Tipo: "gato", Edad: 10},
					Mascota{Nombre: "Luna", Tipo: "gato", Edad: 8},
				}
			}
		}

	Superveterinaria = append(Superveterinaria, veterinarias)
	Superveterinaria = append(Superveterinaria, []Veterinaria{veterinaria})

	for i := 0; i < len(Superveterinaria); i++ {
		fmt.Println(Superveterinaria[i])
		for j := 0; j < len(Superveterinaria[i]); j++ {
			//fmt.Println(Superveterinaria[i][j])
			for k := 0; k < len(Superveterinaria[i][j]); k++ {
			fmt.Println((Superveterinaria[i][j][k]))
			fmt.Printf("Nombre: %v Direccion: %v Mascota: %v \n", Superveterinaria[i][j].Nombre, Superveterinaria[i][j].Direccion, Superveterinaria[i][j].Mascotas)
			fmt.Println("---------------------------")

		}
	}
	}

//fmt.Println(Superveterinaria)

*/
/*
package main

import "fmt"

type Veterinaria struct {
	Nombre    string
	Direccion string
	Comuna    string
	Mascotas  []Mascota
}

type Mascota struct {
	Nombre string
	Tipo   string
	Edad   int
}

func main() {

	veterinaria := Veterinaria{
		Nombre:    "Perritos",
		Direccion: "Bilbao 1181",
		Comuna:    "Providencia",
		Mascotas: []Mascota{
			{Nombre: "Lila", Tipo: "perro", Edad: 8},
			{Nombre: "Margarita", Tipo: "gato", Edad: 5},
			{Nombre: "Balto", Tipo: "perro", Edad: 11},
			{Nombre: "Tito", Tipo: "gato", Edad: 2},
			{Nombre: "Oso", Tipo: "perro", Edad: 10},
		},
	}

	veterinarias := []Veterinaria{
		{
			Nombre:    "Gatitos",
			Direccion: "Manuel Montt 5566",
			Comuna:    "Nunoa",
			Mascotas: []Mascota{
				{Nombre: "Roky", Tipo: "perro", Edad: 12},
				{Nombre: "Mila", Tipo: "gato", Edad: 9},
				{Nombre: "Pupi", Tipo: "gato", Edad: 1},
			},
		},
	}

	superpet := Veterinaria{
		Nombre:    "Superpet",
		Direccion: "Tobalaba 1567",
		Comuna:    "Santiago",
		Mascotas: []Mascota{
			{Nombre: "Rex", Tipo: "gato", Edad: 15},
			{Nombre: "Coco", Tipo: "gato", Edad: 10},
			{Nombre: "Luna", Tipo: "gato", Edad: 8},
		},
	}

	Superveterinaria := [][]Veterinaria{
		veterinarias,
		{veterinaria},
		{superpet},
	}

	totalPerros := 0
	totalGatos := 0

	for i := 0; i < len(Superveterinaria); i++ {
		for j := 0; j < len(Superveterinaria[i]); j++ {

			v := Superveterinaria[i][j]
			perros := 0
			gatos := 0

			EdadPerros := 0
			EdadGatos := 0

			fmt.Println("================================")
			fmt.Printf("Veterinaria: %s\n", v.Nombre)
			fmt.Printf("Comuna: %s\n", v.Comuna)

			for k := 0; k < len(v.Mascotas); k++ {
				m := v.Mascotas[k]

				if m.Tipo == "perro" {
					EdadPerros = EdadPerros + m.Edad
					perros++
					totalPerros++
				} else if m.Tipo == "gato" {
					EdadGatos = EdadGatos + m.Edad
					gatos++
					totalGatos++
				}

				if m.Tipo == "perro" && m.Nombre == "Lila" {
					//fmt.Println("Nombre:", m.Nombre, "Edad:", m.Edad)
				}

			}

			fmt.Println("Edad General Perros: ", EdadPerros)
			fmt.Println("Edad General Gatos: ", EdadGatos)
			fmt.Printf("Perros: %d\n", perros)
			fmt.Printf("Gatos: %d\n", gatos)
			//fmt.Println("Mascotas:")

		}
	}

	fmt.Println("==== TOTAL GENERAL =====")
	fmt.Printf("Total de perros: %d\n", totalPerros)
	fmt.Printf("Total de gatos: %d\n", totalGatos)

	/*mapSV := make(map[string][][]Veterinaria, 0)
	mapSV["Vet"] = Superveterinaria

	if mape, found := mapSV["Vet"]; found {
		fmt.Println(mape)
	} else {
		fmt.Println("No se encontro")
	}
}
*/

package main

import "fmt"

type Supermercado struct {
	Nombre     string
	Direccion  string
	Categorias []Categoria
}

type Categoria struct {
	Nombre    string
	Productos []Producto
}

type Producto struct {
	Nombre string
	Tipo   string
}

func main() {

	super1 := Supermercado{
		Nombre:    "Lider",
		Direccion: "Bilbao 123",
		Categorias: []Categoria{
			{
				Nombre: "Frutas",
				Productos: []Producto{
					{Nombre: "Manzana", Tipo: "Fruta"},
					{Nombre: "Plátano", Tipo: "Fruta"},
					{Nombre: "Pera", Tipo: "Fruta"},
				},
			},
			{
				Nombre: "Lácteos",
				Productos: []Producto{
					{Nombre: "Leche", Tipo: "Lácteo"},
					{Nombre: "Yogurt", Tipo: "Lácteo"},
					{Nombre: "Queso", Tipo: "Lácteo"},
				},
			},
		},
	}

	super2 := Supermercado{
		Nombre:    "Jumbo",
		Direccion: "Las Condes 456",
		Categorias: []Categoria{
			{
				Nombre: "Carnes",
				Productos: []Producto{
					{Nombre: "Pollo", Tipo: "Carne"},
					{Nombre: "Vacuno", Tipo: "Carne"},
				},
			},
			{
				Nombre: "Panadería",
				Productos: []Producto{
					{Nombre: "Marraqueta", Tipo: "Pan"},
					{Nombre: "Hallulla", Tipo: "Pan"},
					{Nombre: "Pan rallado", Tipo: "Pan"},
					{Nombre: "Pan de Molde", Tipo: "Pan"},
				},
			},
			{
				Nombre: "Verduras",
				Productos: []Producto{
					{Nombre: "Papas", Tipo: "Verdura"},
					{Nombre: "Zanahorias", Tipo: "Verdura"},
					{Nombre: "Lechuga", Tipo: "Verdura"},
					{Nombre: "Tomates", Tipo: "Verdura"},
				},
			},
		},
	}

	super3 := Supermercado{
		Nombre:    "Unimarc",
		Direccion: "Pocuro 777",
		Categorias: []Categoria{
			{
				Nombre: "Aseo",
				Productos: []Producto{
					{Nombre: "Esponja", Tipo: "Limpieza"},
					{Nombre: "Cloro", Tipo: "Limpieza"},
				},
			},
			{
				Nombre: "Verduras",
				Productos: []Producto{
					{Nombre: "Lechuga", Tipo: "Verdura"},
					{Nombre: "Pepino", Tipo: "Verdura"},
					{Nombre: "Cilantro", Tipo: "Verdura"},
					{Nombre: "Zapallo", Tipo: "Verdura"},
				},
			},
		},
	}

	superMercados := [][]Supermercado{
		{super1},
		{super2},
		{super3},
	}

	totalProductos := 0
	supermercadosConVerduras := 0

	fmt.Println("\n==== INFORMACIÓN DE SUPERMERCADOS =====")

	for i := 0; i < len(superMercados); i++ {
		for j := 0; j < len(superMercados[i]); j++ {

			s := superMercados[i][j]
			productosLocal := 0
			tieneVerduras := false

			fmt.Println("================================")
			fmt.Printf("Supermercado: %s\n", s.Nombre)
			fmt.Printf("Dirección: %s\n", s.Direccion)
			fmt.Println("Productos por categoría:")

			maxCat := 0
			nombreMaxCat := ""

			for k := 0; k < len(s.Categorias); k++ {
				c := s.Categorias[k]
				productosCategoria := len(c.Productos)

				productosLocal = productosLocal + productosCategoria
				//productosLocal += productosCategoria
				totalProductos += productosCategoria

				fmt.Printf(" - %s: %d productos\n", c.Nombre, productosCategoria)

				if productosCategoria > maxCat {
					maxCat = productosCategoria
					nombreMaxCat = c.Nombre
				}

				if c.Nombre == "Verduras" {
					tieneVerduras = true
				}
			}

			fmt.Printf("Categoría más grande: %s (%d productos)\n", nombreMaxCat, maxCat)
			fmt.Printf("Total en este supermercado: %d productos\n", productosLocal)

			if tieneVerduras {
				fmt.Printf(">> Este supermercado TIENE Verduras\n")
				supermercadosConVerduras++
			}
		}
	}

	fmt.Println("\n==== TOTAL GENERAL =====")
	fmt.Printf("Total de productos: %d\n", totalProductos)
	fmt.Printf("Supermercados que tienen Verduras: %d\n", supermercadosConVerduras)
}
