package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type MyHandler struct {
	Db *sql.DB
}
type silentLogger struct{}

func (sl *silentLogger) Printf(format string, args ...interface{}) {}

func main() {

	ports := flag.Int("port", 8000, "Puerto en el que se iniciará el servidor")
	flag.Parse()
	port := fmt.Sprintf(":%v", *ports)

	db, err := MysqlConnect()
	if err != nil {
		fmt.Println(err)
	}

	h := MyHandler{
		Db: db,
	}

	servHTTP := &fasthttp.Server{Handler: h.HandleFastHTTPS, Logger: &silentLogger{}, Name: "", NoDefaultServerHeader: true}
	servHTTP.ListenAndServe(port)
}

func (h *MyHandler) HandleFastHTTPS(ctx *fasthttp.RequestCtx) {

	if string(ctx.Method()) == "GET" {
		switch string(ctx.Path()) {
		case "/":

		case "/status":

		}
	}
	if string(ctx.Method()) == "POST" {

	}

}

func MysqlConnect() (*sql.DB, error) {

	dsn := "root:12345678@tcp(127.0.0.1:3306)/ControladorAI"
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
