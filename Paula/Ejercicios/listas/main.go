/*package main

import "fmt"

func main() {
	// Creamos un slice de strings
	frutas := []string{"manzana", "pera", "uva"}

	fmt.Println("Lista inicial:", frutas)

	// Agregar un elemento (usando append)
	frutas = append(frutas, "naranja")
	fmt.Println("Después de agregar:", frutas)

	// Acceder a un elemento por índice
	fmt.Println("La primera fruta es:", frutas[0])

	// Recorrer la lista
	for i, fruta := range frutas {
		fmt.Println(i, fruta)
	}
}

*/

package main

import "fmt"

func main() {

	paises := []string{"Chile", "Peru", "Argentina"}

	fmt.Println("Lista de Paises:", paises)

	paises = append(paises, "Bolivia")
	fmt.Println("Pais recien agreagado:", paises)

	fmt.Println("El primer pais es:", paises[0])
	fmt.Println("El segundo pais es:", paises[1])
	fmt.Println("El tercero pais es:", paises[2])
	fmt.Println("El cuarto pais es:", paises[3])

	for i, pais := range paises {
		fmt.Println(i, pais)
	}

}
