package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	"github.com/fasthttp/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

var templates *template.Template

type App struct {
	db *sql.DB
}

func main() {
	db := Connect()
	defer db.Close()

	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("ERROR CRÍTICO: No se pudieron cargar las plantillas HTML. Detalle: ", err)
	}
	log.Println("¡Plantillas cargadas correctamente! Encontrados:", len(templates.Templates()), "archivos")
	app := &App{db: db}
	r := router.New()

	r.GET("/", app.Home)
	r.ServeFiles("/static/{filepath:*}", "./static")
	r.GET("/supermercado", app.AddSuper)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		app.corsMiddleware(r.Handler)(ctx)
	}))
}

func (a *App) Home(ctx *fasthttp.RequestCtx) {
	a.renderHTML(ctx, "admin.html", nil)
}

func (a *App) AddSuper(ctx *fasthttp.RequestCtx) {

	rows, err := a.db.Query(`SELECT id_sup, nombre FROM supermercados ORDER BY fecha`)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprint(ctx, "Error al cargar supermercados")
		return
	}
	defer rows.Close()

	type ListaSuper struct {
		ID     int
		Nombre string
	}
	type Supermercados struct {
		ID     int
		Nombre string
		Lista  []ListaSuper
	}
	var super Supermercados
	for rows.Next() {
		var s ListaSuper
		rows.Scan(&s.ID, &s.Nombre)
		super.Lista = append(super.Lista, s)
	}
	fmt.Println(super.Lista)
	a.renderHTML(ctx, "agregar-supermercado.html", super)

}

func (a *App) corsMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		origin := string(ctx.Request.Header.Peek("Origin"))
		if origin != "" {
			ctx.Response.Header.Set("Access-Control-Allow-Origin", origin)
			ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		}
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusOK)
			return
		}
		next(ctx)
	}
}

func (a *App) renderHTML(ctx *fasthttp.RequestCtx, name string, data any) {
	ctx.SetContentType("text/html; charset=utf-8")
	if templates == nil {
		fmt.Fprint(ctx, "<h1>Error fatal: Plantillas no cargadas</h1>")
		return
	}
	err := templates.ExecuteTemplate(ctx, name, data)
	if err != nil {
		// ESTE ES EL IMPORTANTE: te dice EXACTAMENTE qué línea del HTML falla
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "<h1>Error al renderizar la plantilla: %s</h1><pre>%v</pre>", name, err)
		log.Printf("ERROR RENDER %s: %v", name, err)
		return
	}
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/supermercado?parseTime=true")
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping falló:", err)
	}
	log.Println("✓ Base de datos conectada")
	return db
}
