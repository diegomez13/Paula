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


		nombre := "Coco"
		correo := "Coco@coco"
		password := "3333"

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

		var password2 string = "00000"

		rows, err := db.Query("SELECT id_usr, nombre, correo, password FROM usuarios WHERE password=?", password2)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()

		var id_usr int
		var nombre string
		var correo string
		var password string

		for rows.Next() {
			if err := rows.Scan(&id_usr, &nombre, &correo, &password); err != nil {

				fmt.Println(err)

			}
			fmt.Println(id_usr, nombre, correo, password)
		}


	id_usr := 4
	_, err = db.Exec("DELETE FROM usuarios WHERE id_usr = ?", id_usr)
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
