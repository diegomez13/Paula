package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
)

type App struct {
	db *sql.DB
}

var templates *template.Template

func main() {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/ejercicio?parseTime=true")
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Ping falló:", err)
	}

	templates = template.Must(template.ParseGlob("*.html"))

	app := &App{db: db}

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", app.handler))
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")

	path := string(ctx.Path())
	method := string(ctx.Method())

	switch {
	case method == "GET" && (path == "/" || path == ""):
		templates.ExecuteTemplate(ctx, "menu.html", nil)

	case method == "GET" && path == "/supermercado":
		templates.ExecuteTemplate(ctx, "supermercado.html", nil)

	case method == "GET" && path == "/categoria":
		a.mostrarFormularioCategoria(ctx)

	case method == "GET" && path == "/producto":
		a.mostrarFormularioProducto(ctx)

	case method == "GET" && path == "/listar-super":
		a.listarSupermercados(ctx)

	case method == "GET" && path == "/listar-categorias":
		a.formListarCategorias(ctx)

	case method == "GET" && path == "/listar-productos":
		a.formListarProductos(ctx)

	case method == "POST" && path == "/crear-supermercado":
		a.crearSupermercado(ctx)

	case method == "POST" && path == "/crear-categoria":
		a.crearCategoria(ctx)

	case method == "POST" && path == "/crear-producto":
		a.crearProducto(ctx)

	case method == "POST" && path == "/ver-categorias":
		a.verCategorias(ctx)

	case method == "POST" && path == "/ver-productos":
		a.verProductos(ctx)

	default:
		ctx.SetStatusCode(404)
		ctx.WriteString("<h2>404 - Página no encontrada</h2><a href='/'>Volver al menú</a>")
	}
}

// ======================= CREACIONES =======================
func (a *App) crearSupermercado(ctx *fasthttp.RequestCtx) {
	nombre := string(ctx.FormValue("nombre"))
	direccion := string(ctx.FormValue("direccion"))
	_, err := a.db.Exec("INSERT INTO supermercados (nombre, direccion) VALUES (?, ?)", nombre, direccion)
	msg := fmt.Sprintf("Supermercado <strong>%s</strong> creado correctamente", nombre)
	if err != nil {
		msg = "<span style='color:red;'>Error: " + err.Error() + "</span>"
	}
	templates.ExecuteTemplate(ctx, "menu.html", map[string]any{"Mensaje": template.HTML(msg)})
}

func (a *App) crearCategoria(ctx *fasthttp.RequestCtx) {
	nombre := string(ctx.FormValue("nombre"))
	idSup := string(ctx.FormValue("id_sup"))

	var superNombre string
	a.db.QueryRow("SELECT nombre FROM supermercados WHERE Id_sup = ?", idSup).Scan(&superNombre)
	if superNombre == "" {
		superNombre = "desconocido"
	}

	_, err := a.db.Exec("INSERT INTO categorias (nombre, id_sup) VALUES (?, ?)", nombre, idSup)
	msg := fmt.Sprintf("Categoría <strong>%s</strong> creada correctamente en el supermercado <strong>%s</strong>", nombre, superNombre)
	if err != nil {
		msg = "<span style='color:red;'>Error: " + err.Error() + "</span>"
	}
	templates.ExecuteTemplate(ctx, "menu.html", map[string]any{"Mensaje": template.HTML(msg)})
}

func (a *App) crearProducto(ctx *fasthttp.RequestCtx) {
	nombre := string(ctx.FormValue("nombre"))
	idCatStr := string(ctx.FormValue("id_cat"))

	var catNombre, superNombre string
	err := a.db.QueryRow(`
        SELECT c.nombre, s.nombre 
        FROM categorias c 
        JOIN supermercados s ON c.id_sup = s.Id_sup 
        WHERE c.id_cat = ?`, idCatStr).
		Scan(&catNombre, &superNombre)

	if err != nil || catNombre == "" {
		catNombre = "Categoría desconocida"
		superNombre = "Supermercado desconocido"
	}

	_, err = a.db.Exec("INSERT INTO productos (nombre, id_cat) VALUES (?, ?)", nombre, idCatStr)

	msg := fmt.Sprintf("Producto <strong>%s</strong> creado correctamente en la categoría <strong>%s</strong> del supermercado <strong>%s</strong>",
		nombre, catNombre, superNombre)

	if err != nil {
		msg = "<span style='color:red;'>Error al crear el producto: " + err.Error() + "</span>"
	}

	templates.ExecuteTemplate(ctx, "menu.html", map[string]any{
		"Mensaje": template.HTML(msg),
	})
}

// ======================= FORMULARIOS DINÁMICOS =======================
func (a *App) mostrarFormularioCategoria(ctx *fasthttp.RequestCtx) {
	rows, err := a.db.Query("SELECT Id_sup, nombre FROM supermercados ORDER BY nombre")
	if err != nil {
		ctx.WriteString("Error al cargar supermercados")
		return
	}
	defer rows.Close()

	type Super struct {
		ID     int
		Nombre string
	}
	var lista []Super
	for rows.Next() {
		var s Super
		rows.Scan(&s.ID, &s.Nombre)
		lista = append(lista, s)
	}
	templates.ExecuteTemplate(ctx, "categoria.html", map[string]any{"Supermercados": lista})
}

func (a *App) mostrarFormularioProducto(ctx *fasthttp.RequestCtx) {
	rows, err := a.db.Query("SELECT id_cat, nombre FROM categorias ORDER BY nombre")
	if err != nil {
		ctx.WriteString("Error al cargar categorías")
		return
	}
	defer rows.Close()

	type Cat struct {
		ID     int
		Nombre string
	}
	var lista []Cat
	for rows.Next() {
		var c Cat
		rows.Scan(&c.ID, &c.Nombre)
		lista = append(lista, c)
	}
	templates.ExecuteTemplate(ctx, "producto.html", map[string]any{"Categorias": lista})
}

// ======================= LISTADOS =======================
func (a *App) listarSupermercados(ctx *fasthttp.RequestCtx) {
	rows, err := a.db.Query("SELECT Id_sup, nombre, direccion FROM supermercados ORDER BY nombre")
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	defer rows.Close()

	html := "<h2>Listado de Supermercados</h2><table border='1' cellpadding='10' style='border-collapse:collapse'>"
	html += "<tr style='background:#f0f0f0'><th>ID</th><th>Nombre</th><th>Dirección</th></tr>"
	for rows.Next() {
		var id int
		var n, d string
		rows.Scan(&id, &n, &d)
		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td></tr>", id, n, d)
	}
	html += "</table><br><a href='/'>Volver al menú</a>"
	ctx.WriteString(html)
}

func (a *App) formListarCategorias(ctx *fasthttp.RequestCtx) {
	rows, err := a.db.Query("SELECT Id_sup, nombre FROM supermercados ORDER BY nombre")
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	defer rows.Close()

	html := "<h2>Listar categorías por supermercado</h2><form method='POST' action='/ver-categorias'>"
	html += "Supermercado: <select name='id_sup'>"
	for rows.Next() {
		var id int
		var nombre string
		rows.Scan(&id, &nombre)
		html += fmt.Sprintf("<option value='%d'>%s</option>", id, nombre)
	}
	html += "</select> <button>Ver categorías</button></form><br><a href='/'>Volver</a>"
	ctx.WriteString(html)
}

func (a *App) verCategorias(ctx *fasthttp.RequestCtx) {
	idSup := string(ctx.FormValue("id_sup"))

	var superNombre string
	a.db.QueryRow("SELECT nombre FROM supermercados WHERE Id_sup = ?", idSup).Scan(&superNombre)

	rows, err := a.db.Query("SELECT id_cat, nombre FROM categorias WHERE id_sup = ? ORDER BY nombre", idSup)
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	defer rows.Close()

	html := fmt.Sprintf("<h2>Categorías del supermercado: <strong>%s</strong></h2>", superNombre)
	html += "<table border='1' cellpadding='10' style='border-collapse:collapse'>"
	html += "<tr style='background:#f0f0f0'><th>ID</th><th>Categoría</th></tr>"
	for rows.Next() {
		var id int
		var nombre string
		rows.Scan(&id, &nombre)
		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td></tr>", id, nombre)
	}
	html += "</table><br><a href='/'>Volver al menú</a>"
	ctx.WriteString(html)
}

func (a *App) formListarProductos(ctx *fasthttp.RequestCtx) {
	rows, err := a.db.Query("SELECT id_cat, nombre FROM categorias ORDER BY nombre")
	if err != nil {
		ctx.WriteString("Error")
		return
	}
	defer rows.Close()

	html := "<h2>Listar productos por categoría</h2><form method='POST' action='/ver-productos'>"
	html += "Categoría: <select name='id_cat'>"
	for rows.Next() {
		var id int
		var nombre string
		rows.Scan(&id, &nombre)
		html += fmt.Sprintf("<option value='%d'>%s</option>", id, nombre)
	}
	html += "</select> <button>Ver productos</button></form><br><a href='/'>Volver al menú</a>"
	ctx.WriteString(html)
}

func (a *App) verProductos(ctx *fasthttp.RequestCtx) {
	idCat := string(ctx.FormValue("id_cat"))

	var catNombre, superNombre string
	err := a.db.QueryRow(`
        SELECT c.nombre, s.nombre 
        FROM categorias c 
        JOIN supermercados s ON c.id_sup = s.Id_sup 
        WHERE c.id_cat = ?`, idCat).
		Scan(&catNombre, &superNombre)

	if err != nil {
		catNombre, superNombre = "desconocida", "desconocido"
	}

	rows, err := a.db.Query(`
        SELECT p.id_pro, p.nombre, s.nombre 
        FROM productos p 
        JOIN categorias c ON p.id_cat = c.id_cat 
        JOIN supermercados s ON c.id_sup = s.Id_sup 
        WHERE p.id_cat = ? 
        ORDER BY p.nombre`, idCat)
	if err != nil {
		ctx.WriteString("Error al cargar productos")
		return
	}
	defer rows.Close()

	html := fmt.Sprintf("<h2>Productos de la categoría: <strong>%s</strong> (Supermercado: <strong>%s</strong>)</h2>", catNombre, superNombre)
	html += "<table border='1' cellpadding='10' style='border-collapse:collapse'>"
	html += "<tr style='background:#f0f0f0'><th>ID</th><th>Producto</th><th>Supermercado</th></tr>"
	for rows.Next() {
		var id int
		var prod, super string
		rows.Scan(&id, &prod, &super)
		html += fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td></tr>", id, prod, super)
	}
	html += "</table><br><a href='/'>Volver al menú</a>"
	ctx.WriteString(html)
}
