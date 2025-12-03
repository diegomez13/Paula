package main

import (
	"database/sql"
	"fmt"
	"net/mail"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	Db *sql.DB
}

func main() {

	db, err := MysqlConnect()
	if err != nil {
		fmt.Println(err)
	}

	h := MyHandler{
		Db: db,
	}

	servHTTP := &fasthttp.Server{Handler: h.HandleFastHTTPS, Name: "", NoDefaultServerHeader: true}
	servHTTP.ListenAndServe(":8000")
}

func (h *MyHandler) HandleFastHTTPS(ctx *fasthttp.RequestCtx) {

	if string(ctx.Method()) == "GET" {
		ctx.Response.Header.Set("Content-Type", "text/html; charset=utf-8")
		switch string(ctx.Path()) {
		case "/registrarse":
			data, err := os.ReadFile("signup.html")
			if err != nil {
				fmt.Println(err)
			}
			ctx.SetBody(data)
		case "/logearse":
			data, err := os.ReadFile("login.html")
			if err != nil {
				fmt.Println(err)
			}
			ctx.SetBody(data)

		}
	}
	if string(ctx.Method()) == "POST" {
		switch string(ctx.Path()) {
		case "/submit":
			nombre := string(ctx.PostArgs().Peek("nombre"))
			correo := string(ctx.PostArgs().Peek("correo"))
			password := string(ctx.PostArgs().Peek("password"))
			password2 := string(ctx.PostArgs().Peek("password2"))

			err := h.AgregarUsuario(nombre, correo, password, password2)
			if err != nil {
				fmt.Println(err)
				return
			}

			ctx.SetBody([]byte("Usuario Ingresado"))

		case "/login":
			correo := string(ctx.PostArgs().Peek("usuario"))
			password := string(ctx.PostArgs().Peek("password"))

			err := h.Logearse(correo, password)
			if err != nil {
				fmt.Println(err)
				return
			}

			ctx.SetBody([]byte("Usuario logeado"))

		}

	}

}

func (h *MyHandler) Logearse(correo, password string) error {

	rows, err := h.Db.Query("SELECT password FROM usuarios WHERE correo=?", correo)
	if err != nil {
		return err
	}
	defer rows.Close()

	var passworddb string

	for rows.Next() {
		if err := rows.Scan(&passworddb); err != nil {
			return err
		}
		if password != passworddb {
			return fmt.Errorf("Password Invalido")

		} else {
			return nil
		}

	}
	return fmt.Errorf("No encontro a ningun usuario con ese correo")
}

func (h *MyHandler) AgregarUsuario(nombre, correo, password, password2 string) error {

	_, err := mail.ParseAddress(correo)
	if err != nil {
		return err
	}

	count := 0

	err = h.Db.QueryRow("SELECT COUNT(*) FROM usuarios WHERE correo=?", correo).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("Usuario ya existe")
	}
	if password != password2 {
		return fmt.Errorf("Password diferente")
	}
	_, err = h.Db.Exec("INSERT INTO usuarios (nombre, correo, password) VALUES (?,?,?)", nombre, correo, password)
	if err != nil {
		return err
	}
	return nil

}

func MysqlConnect() (*sql.DB, error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/ejercicio"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Prueba la conexión
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateCookie(key []byte, value []byte, expire int) *fasthttp.Cookie {

	authCookie := fasthttp.Cookie{}
	authCookie.SetKeyBytes(key)
	authCookie.SetValueBytes(value)
	authCookie.SetMaxAge(expire)
	authCookie.SetPath("/")
	authCookie.SetHTTPOnly(true)
	authCookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	//fmt.Println("CreateCookie:", authCookie.String())
	return &authCookie
}

func DeleteCookie(ctx *fasthttp.RequestCtx, cookieName string) {
	// Crear la cookie con el mismo nombre
	cookie := fasthttp.AcquireCookie()
	defer fasthttp.ReleaseCookie(cookie)

	// Configurar los valores necesarios para eliminar la cookie
	cookie.SetKey(cookieName)
	cookie.SetMaxAge(-1)     // Fecha en el pasado
	cookie.SetPath("/")      // Asegúrate de que el path coincida con el original
	cookie.SetHTTPOnly(true) // Si la cookie original era HTTPOnly, mantén esta configuración

	// Añadir la cookie de eliminación a la respuesta
	ctx.Response.Header.SetCookie(cookie)
}
