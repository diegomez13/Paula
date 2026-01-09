package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"github.com/valyala/fasthttp"
)

type App struct{}

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("html/*.html"))

	app := &App{}
	log.Fatal(fasthttp.ListenAndServe(":8000", app.handler))

}

type LoginResponse struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

func (a *App) handler(ctx *fasthttp.RequestCtx) {

	if string(ctx.Method()) == "GET" {
		switch string(ctx.Path()) {
		case "/":
			ctx.SetContentType("text/html; charset=utf-8")
			if len(ctx.Request.Header.Cookie("yggy")) > 0 {
				templates.ExecuteTemplate(ctx.Response.BodyWriter(), "home.html", nil)
			} else {
				templates.ExecuteTemplate(ctx.Response.BodyWriter(), "login.html", nil)
			}

		}
	}

	if string(ctx.Method()) == "POST" {
		switch string(ctx.Path()) {
		case "/login":
			ctx.SetContentType("application/json")
			username := string(ctx.FormValue("username"))
			password := string(ctx.FormValue("password"))

			fmt.Println(username, password)

			resp := LoginResponse{
				Username: username,
				Status:   "ok",
			}
			jsonData, _ := json.Marshal(resp)

			ctx.Response.Header.SetCookie(CreateCookie([]byte("yy"), []byte("1"), 100000000))

			ctx.Write(jsonData)

		}

	}
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

/*










 */
