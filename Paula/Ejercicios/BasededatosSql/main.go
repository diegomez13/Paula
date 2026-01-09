/*
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := MysqlConnect()
	if err != nil {

		fmt.Println(err)
	}


		supermercados := "Lider"
		categoria := "Lacteos"
		productos := "Leche Condensada"

		sql, err := db.Exec("INSERT INTO supermercados (categoria, productos) VALUES (?,?,?)", categoria, productos)
		if err != nil {
			fmt.Println(err)

		}

		id, err := sql.LastInsertId()
		if err != nil {
			fmt.Println(err)

		}

		_, err = db.Exec("INSERT INTO categoria (productos) VALUES (NOW(),?)", id)
		if err != nil {
			fmt.Println(err)

		}

		fmt.Println("categoria ingresada", id)


		rows, err := db.Query("SELECT id_supermercado, supermercado, categoria, productos FROM supermercado WHERE categorias=?", productos)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		var id_supermercado int
		var Categorias string
		var productos string

		for rows.Next() {
			if err := rows.Scan(&id_supermercado, &Categorias, &productos); err != nil {

				fmt.Println(err)

			}
			fmt.Println(id_supermercado, Categorias, productos)
		}


	id_supermercado := 1
	_, err = db.Exec("DELETE FROM usuarios WHERE id_supermercado = ?", id_supermercado)
	if err != nil {
		fmt.Println(err)

	}

}

func MysqlConnect() (*sql.DB, error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("ERROR 1A")
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Prueba la conexión
	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR 1B")
		return nil, err
	}

	return db, nil
}

*/

/*
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := MysqlConnect()
	if err != nil {
		fmt.Println(err)
	}

	nombre := "Luz"
	correo := "luz@luz"
	password := "90909090"

	sql, err := db.Exec("INSERT INTO usuarios (nombre, correo, password) VALUES (?,?,?)", nombre, correo, password)
	if err != nil {
		fmt.Println(err)
	}

	id, err := sql.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	_, err = db.Exec("INSERT INTO registro_usuarios (fecha, id_usr) VALUES (NOW(),?)", id)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("usuario creado", id)
}

func MysqlConnect() (*sql.DB, error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("ERROR 1A")
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Prueba la conexión
	err = db.Ping()
	if err != nil {
		fmt.Println("ERROR 1B")
		return nil, err
	}

	return db, nil
}
*/

/*
package main

import (

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

func main() {

		db, err := MysqlConnect()
		if err != nil {
			fmt.Println(err)
		}

		nombre := "Renato"
		correo := "Renato@mato"
		password := "181818181"

		sql, err := db.Exec("INSERT INTO usuarios (nombre, correo, password) VALUES (?,?,?)", nombre, correo, password)
		if err != nil {
			fmt.Println(err)
		}

		id, err := sql.LastInsertId()
		if err != nil {
			fmt.Println(err)
		}

		_, err = db.Exec("INSERT INTO registro_usuarios (fecha, id_usr) VALUES (NOW(),?)", id)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("usuario creado", id)
	}

func MysqlConnect() (*sql.DB, error) {

		dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("ERROR 1A")
			return nil, err
		}

		db.SetConnMaxLifetime(0)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(5)

		// Prueba la conexión
		err = db.Ping()
		if err != nil {
			fmt.Println("ERROR 1B")
			return nil, err
		}

		return db, nil
	}
*/
/*
package main

import (
	"db-mysql/database"
	"db-mysql/handlers"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Establecer conexión a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ID del contacto que deseas eliminar
	contactID := 5

	// Eliminar el contacto de la base de datos
	handlers.DeleteContact(db, contactID)

	// Listar contactos (opcional, solo para verificar que el contacto fue eliminado)
	handlers.ListContacts(db)
}
*/
/*
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := MysqlConnect()
	if err != nil {
		fmt.Println("Error de conexión:", err)
		return
	}

	supermercado := "Lider"
	categoria := "Lacteos"
	producto := "Queso"

	// INSERT correcto
	res, err := db.Exec(`
		INSERT INTO supermercados (supermercado, categoria, producto)
		VALUES (?, ?, ?)`,
		supermercado, categoria, producto)
	if err != nil {
		fmt.Println("Error insertando supermercado:", err)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error obteniendo ID:", err)
		return
	}

	// Segundo insert
	_, err = db.Exec(`
		INSERT INTO categorias (producto_id, fecha)
		VALUES (?, NOW())`,
		id)
	if err != nil {
		fmt.Println("Error insertando categoría:", err)
		return
	}

	fmt.Println("Categoría ingresada con ID:", id)

	// SELECT corregido
	rows, err := db.Query(`
		SELECT id, supermercado, categoria, producto
		FROM supermercados
		WHERE producto = ?`, producto)
	if err != nil {
		fmt.Println("Error en SELECT:", err)
		return
	}
	defer rows.Close()

	var (
		idSuper int
		super   string
		cat     string
		prod    string
	)

	for rows.Next() {
		if err := rows.Scan(&idSuper, &super, &cat, &prod); err != nil {
			fmt.Println("Error leyendo filas:", err)
		}
		fmt.Println(idSuper, super, cat, prod)
	}

	// DELETE corregido
	_, err = db.Exec("DELETE FROM supermercados WHERE id = ?", idSuper)
	if err != nil {
		fmt.Println("Error en DELETE:", err)
	}
}

func MysqlConnect() (*sql.DB, error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

*/
/*
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := MysqlConnect()
	if err != nil {
		fmt.Println("Error conectando:", err)
		return
	}

	// DATOS QUE VAMOS A INSERTAR
	supermercadoNombre := "liderIII"
	supermercadoDireccion := "Av. Principal 123"

	categoriaNombre := "Snack"

	productoNombre := "Papas Fritas"

	// =============================
	// 1) INSERTAR SUPERMERCADO
	// =============================
	resSup, err := db.Exec(`
		INSERT INTO supermercados (nombre, direccion)
		VALUES (?, ?)`,
		supermercadoNombre, supermercadoDireccion)

	if err != nil {
		fmt.Println("Error insertando supermercado:", err)
		return
	}

	idSup, _ := resSup.LastInsertId()
	fmt.Println("Supermercado insertado con ID:", idSup)

	// =============================
	// 2) INSERTAR CATEGORIA
	// =============================
	resCat, err := db.Exec(`
		INSERT INTO categorias (nombre, id_sup)
		VALUES (?, ?)`,
		categoriaNombre, idSup)

	if err != nil {
		fmt.Println("Error insertando categoría:", err)
		return
	}

	idCat, _ := resCat.LastInsertId()
	fmt.Println("Categoría insertada con ID:", idCat)

	// =============================
	// 3) INSERTAR PRODUCTO
	// =============================
	resProd, err := db.Exec(`
		INSERT INTO productos (nombre, id_cat)
		VALUES (?, ?)`,
		productoNombre, idCat)

	if err != nil {
		fmt.Println("Error insertando producto:", err)
		return
	}

	idProd, _ := resProd.LastInsertId()
	fmt.Println("Producto insertado con ID:", idProd)

	// =============================
	// 4) CONSULTAR PRODUCTOS DEL SUPERMERCADO
	// =============================
	rows, err := db.Query(`
		SELECT productos.id_pro, productos.nombre, categorias.nombre, supermercados.nombre
		FROM productos
		INNER JOIN categorias ON productos.id_cat = categorias.id_cat
		INNER JOIN supermercados ON categorias.id_sup = supermercados.Id_sup
		WHERE supermercados.Id_sup = ?`, idSup)

	if err != nil {
		fmt.Println("Error en SELECT:", err)
		return
	}

	fmt.Println("\nProductos del supermercado creado:")
	for rows.Next() {
		var idProducto int
		var producto string
		var categoria string
		var super string

		rows.Scan(&idProducto, &producto, &categoria, &super)
		fmt.Println(idProducto, producto, categoria, super)
	}
}

func MysqlConnect() (*sql.DB, error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
*/

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := MysqlConnect()
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	// Captura señales del sistema (Ctrl+C)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	exitCh := make(chan struct{})
	reader := bufio.NewReader(os.Stdin)

	go func() {
		for {
			fmt.Println("\n===== MENÚ =====")
			fmt.Println("1) Crear supermercado")
			fmt.Println("2) Agregar categoría a supermercado existente")
			fmt.Println("3) Agregar producto a categoría existente")
			fmt.Println("4) Listar productos por supermercado")
			fmt.Println("5) Listar categorías por supermercado")
			fmt.Println("6) Listar supermercados")
			fmt.Println("7) Salir")
			fmt.Print("Selecciona opción: ")

			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nLectura terminada. Saliendo...")
				close(exitCh)
				return
			}
			line = strings.TrimSpace(line)
			if line == "" {
				fmt.Println("No ingresaste nada.")
				continue
			}

			lower := strings.ToLower(line)
			if lower == "exit" || lower == "salir" {
				close(exitCh)
				return
			}

			op, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Por favor ingresa un número válido.")
				continue
			}

			switch op {
			case 1:
				fmt.Print("Nombre del supermercado: ")
				nombre := mustReadLineTrim(reader)
				fmt.Print("Dirección: ")
				direccion := mustReadLineTrim(reader)
				if nombre == "" || direccion == "" {
					fmt.Println("Nombre o dirección vacía. Cancelado.")
					continue
				}
				id, err := crearSupermercado(db, nombre, direccion)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("Supermercado creado con ID:", id)
				}

			case 2:
				listarSupermercados(db)
				fmt.Print("ID del supermercado: ")
				idSup, ok := readIntPrompt(reader)
				if !ok {
					fmt.Println("ID inválido. Cancelado.")
					continue
				}
				fmt.Print("Nombre de la categoría: ")
				nombreCat := mustReadLineTrim(reader)
				if nombreCat == "" {
					fmt.Println("Nombre vacío. Cancelado.")
					continue
				}
				if err := agregarCategoria(db, nombreCat, idSup); err != nil {
					fmt.Println("Error:", err)
				}

			case 3:
				listarCategorias(db)
				fmt.Print("ID de la categoría: ")
				idCat, ok := readIntPrompt(reader)
				if !ok {
					fmt.Println("ID inválido. Cancelado.")
					continue
				}
				fmt.Print("Nombre del producto: ")
				nombreProd := mustReadLineTrim(reader)
				if nombreProd == "" {
					fmt.Println("Nombre vacío. Cancelado.")
					continue
				}
				if err := agregarProducto(db, nombreProd, idCat); err != nil {
					fmt.Println("Error:", err)
				}

			case 4:
				listarSupermercados(db)
				fmt.Print("ID del supermercado para listar productos: ")
				idSup, ok := readIntPrompt(reader)
				if !ok {
					fmt.Println("ID inválido. Cancelado.")
					continue
				}
				listarProductos(db, idSup)

			case 5:
				listarSupermercados(db)
				fmt.Print("ID del supermercado para listar categorías: ")
				idSup, ok := readIntPrompt(reader)
				if !ok {
					fmt.Println("ID inválido. Cancelado.")
					continue
				}
				listarCategoriasPorSupermercado(db, idSup)

			case 6:
				listarSupermercados(db)

			case 7:
				fmt.Println("Salida solicitada. Saliendo...")
				close(exitCh)
				return

			default:
				fmt.Println("Opción inválida")
			}
		}
	}()

	select {
	case <-sigCh:
		fmt.Println("\nSeñal recibida. Saliendo...")
	case <-exitCh:
		fmt.Println("Salida por usuario. Saliendo...")
	}
}

// ==================== helpers de lectura ====================
func mustReadLineTrim(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func readIntPrompt(r *bufio.Reader) (int, bool) {
	s, err := r.ReadString('\n')
	if err != nil {
		return 0, false
	}
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, false
	}
	return n, true
}

// ==================== Conexión a DB ====================
func MysqlConnect() (*sql.DB, error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// ==================== CRUD ====================
func crearSupermercado(db *sql.DB, nombre, direccion string) (int64, error) {
	res, err := db.Exec("INSERT INTO supermercados (nombre, direccion) VALUES (?, ?)", nombre, direccion)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func agregarCategoria(db *sql.DB, nombre string, idSup int) error {
	_, err := db.Exec("INSERT INTO categorias (nombre, id_sup) VALUES (?, ?)", nombre, idSup)
	if err != nil {
		return err
	}
	fmt.Println("Categoría agregada correctamente")
	return nil
}

func agregarProducto(db *sql.DB, nombre string, idCat int) error {
	_, err := db.Exec("INSERT INTO productos (nombre, id_cat) VALUES (?, ?)", nombre, idCat)
	if err != nil {
		return err
	}
	fmt.Println("Producto agregado correctamente")
	return nil
}

// ==================== Listados ====================
func listarSupermercados(db *sql.DB) {
	rows, err := db.Query("SELECT Id_sup, nombre FROM supermercados")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	tiene := false
	fmt.Println("\nSupermercados:")
	for rows.Next() {
		var id int
		var nombre string
		if err := rows.Scan(&id, &nombre); err != nil {
			fmt.Println("Scan error:", err)
			return
		}
		fmt.Printf("ID: %d | Nombre: %s\n", id, nombre)
		tiene = true
	}
	if !tiene {
		fmt.Println("No hay supermercados registrados.")
	}
}

func listarCategorias(db *sql.DB) {
	rows, err := db.Query(`
		SELECT categorias.id_cat, categorias.nombre AS nombre_cat, supermercados.nombre AS nombre_sup
		FROM categorias
		INNER JOIN supermercados ON categorias.id_sup = supermercados.Id_sup`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	tiene := false
	fmt.Println("\nCategorías:")
	for rows.Next() {
		var id int
		var nombreCat string
		var nombreSup string
		if err := rows.Scan(&id, &nombreCat, &nombreSup); err != nil {
			fmt.Println("Scan error:", err)
			return
		}
		fmt.Printf("ID: %d | Categoría: %s | Supermercado: %s\n", id, nombreCat, nombreSup)
		tiene = true
	}
	if !tiene {
		fmt.Println("No hay categorías registradas.")
	}
}

// Listar categorías de un supermercado específico
func listarCategoriasPorSupermercado(db *sql.DB, idSup int) {
	rows, err := db.Query(`
		SELECT id_cat, nombre AS nombre_cat
		FROM categorias
		WHERE id_sup = ?`, idSup)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	tiene := false
	fmt.Println("\nCategorías del supermercado:")
	for rows.Next() {
		var id int
		var nombreCat string
		if err := rows.Scan(&id, &nombreCat); err != nil {
			fmt.Println("Scan error:", err)
			return
		}
		fmt.Printf("ID: %d | Categoría: %s\n", id, nombreCat)
		tiene = true
	}
	if !tiene {
		fmt.Println("No hay categorías para este supermercado.")
	}
}

func listarProductos(db *sql.DB, idSup int) {
	rows, err := db.Query(`
		SELECT productos.id_pro, productos.nombre AS nombre_prod, categorias.nombre AS nombre_cat
		FROM productos
		INNER JOIN categorias ON productos.id_cat = categorias.id_cat
		WHERE categorias.id_sup = ?`, idSup)
	if err != nil {
		fmt.Println("Error al consultar productos:", err)
		return
	}
	defer rows.Close()

	tieneProductos := false
	fmt.Println("\nProductos del supermercado:")
	for rows.Next() {
		var id int
		var nombreProd string
		var nombreCat string
		if err := rows.Scan(&id, &nombreProd, &nombreCat); err != nil {
			fmt.Println("Scan error:", err)
			continue
		}
		fmt.Printf("ID: %d | Producto: %s | Categoría: %s\n", id, nombreProd, nombreCat)
		tieneProductos = true
	}
	if !tieneProductos {
		fmt.Println("No hay productos para este supermercado.")
	}
}
