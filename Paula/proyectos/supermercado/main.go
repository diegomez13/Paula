package main

import (
	"bufio"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fasthttp/router"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"gopkg.in/gomail.v2"

	"github.com/valyala/fasthttp"
)

var templates *template.Template

type App struct {
	db      *sql.DB
	Sistema Sistema
	Modo    string
}

type Sistema struct {
	Configuracion Configuracion
	Supermercados []Supermercado
}

type Configuracion struct {
	IdSuper      uint64
	IdCat        uint64
	IdPro        uint64
	Userdb       string
	Passdb       string
	UserSistema  string `json:"UserSistema"`
	PassSistema  string `json:"PassSistema"`
	EmailSMTP    string `json:"EmailSMTP"`    // tu email de Gmail
	PasswordSMTP string `json:"PasswordSMTP"` // contraseña de aplicación de Gmail

}
type Supermercado struct {
	IdSuper    uint64
	Nombre     string
	Fecha      time.Time
	Categorias []Categoria
}
type Categoria struct {
	IdCat     uint64
	Nombre    string
	Fecha     time.Time
	Productos []Producto
}
type HTML struct {
	SUPER  Supermercado
	NOMBRE string
	ID_sup int
	ID_cat int
	ID_pro int
}
type Producto struct {
	IdPro  uint64
	Nombre string
	Fecha  time.Time
}

func (s *Sistema) GuardarConfiguracion() {

	data, err := json.Marshal(s.Configuracion)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("./Configuracion.Json", data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func validarCredencialesSistema(user, pass, hashGuardado string) bool {
	if user == "" || hashGuardado == "" {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashGuardado), []byte(pass))
	return err == nil
}
func (s *Sistema) AddSuper(nombre string) {
	id_sup := atomic.AddUint64(&s.Configuracion.IdSuper, 1)
	super := Supermercado{IdSuper: id_sup, Nombre: nombre, Fecha: time.Now(), Categorias: []Categoria{}}
	s.Supermercados = append(s.Supermercados, super)

	data, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}
	path := fmt.Sprintf("./Super/Super_%v", id_sup)
	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
	s.GuardarConfiguracion()
}
func (s *Sistema) ModificarSuper(IdSuper uint64, Nombre string) {

	path := fmt.Sprintf("./Super/Super_%v", IdSuper)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var super Supermercado
	err = json.Unmarshal(data, &super)
	if err != nil {
		fmt.Println(err)
	}

	super.Nombre = Nombre
	data2, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data2, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
func (s *Sistema) AddCategoria(IdSuper uint64, Nombre string) {

	id_cat := atomic.AddUint64(&s.Configuracion.IdCat, 1)
	categoria := Categoria{IdCat: id_cat, Nombre: Nombre, Fecha: time.Now(), Productos: []Producto{}}

	path := fmt.Sprintf("./Super/Super_%v", IdSuper)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var super Supermercado
	err = json.Unmarshal(data, &super)
	if err != nil {
		fmt.Println(err)
	}
	super.Categorias = append(super.Categorias, categoria)

	data2, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data2, 0644)
	if err != nil {
		fmt.Println(err)
	}
	s.GuardarConfiguracion()
}
func (s *Sistema) ModificarCategoria(IdSuper uint64, IdCat uint64, Nombre string) {

	path := fmt.Sprintf("./Super/Super_%v", IdSuper)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var super Supermercado
	err = json.Unmarshal(data, &super)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(super.Categorias); i++ {
		if super.Categorias[i].IdCat == IdCat {
			super.Categorias[i].Nombre = Nombre
		}

	}
	data2, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data2, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func (s *Sistema) AddProducto(IdSuper uint64, IdCat uint64, Nombre string) {

	path := fmt.Sprintf("./Super/Super_%v", IdSuper)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var super Supermercado
	err = json.Unmarshal(data, &super)
	if err != nil {
		fmt.Println(err)
	}

	producto := Producto{IdPro: atomic.AddUint64(&s.Configuracion.IdPro, 1), Nombre: Nombre, Fecha: time.Now()}

	for i := 0; i < len(super.Categorias); i++ {
		if super.Categorias[i].IdCat == IdCat {
			super.Categorias[i].Productos = append(super.Categorias[i].Productos, producto)
		}
	}

	data2, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data2, 0644)
	if err != nil {
		fmt.Println(err)
	}
	s.GuardarConfiguracion()
}
func (s *Sistema) ModificarProducto(IdSuper uint64, IdCat uint64, IdPro uint64, Nombre string) {

	path := fmt.Sprintf("./Super/Super_%v", IdSuper)
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var super Supermercado
	err = json.Unmarshal(data, &super)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(super.Categorias); i++ {
		if super.Categorias[i].IdCat == IdCat {
			for j := 0; j < len(super.Categorias[i].Productos); j++ {
				if super.Categorias[i].Productos[j].IdPro == IdPro {
					super.Categorias[i].Productos[j].Nombre = Nombre
				}
			}

		}

	}
	data2, err := json.Marshal(super)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(path, data2, 0644)
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	reader := bufio.NewReader(os.Stdin)

	// Cargar configuración existente
	var Conf Configuracion
	configExists := false
	data, err := os.ReadFile("./Configuracion.Json")
	if err == nil {
		err = json.Unmarshal(data, &Conf)
		if err == nil {
			configExists = true
		}
	}

	// Verificar si configuración está completa
	configCompleta := configExists &&
		Conf.Userdb != "" &&
		Conf.Passdb != "" &&
		Conf.UserSistema != "" &&
		Conf.PassSistema != ""

	// SI YA ESTÁ TODO CONFIGURADO → Iniciar directo
	if configCompleta && validarDB(Conf.Userdb, Conf.Passdb) {
		fmt.Println("✔ Configuración válida encontrada")
		fmt.Println("✔ Iniciando servidor...")
		iniciarServidor(Conf)
		return
	}

	// SI NO ESTÁ CONFIGURADO O FALLÓ VALIDACIÓN → Mostrar menú
	if configCompleta && !validarDB(Conf.Userdb, Conf.Passdb) {
		fmt.Println("⚠️  Error: No se puede conectar a la base de datos")
		fmt.Println("    Reconfigure las credenciales")
	} else {
		fmt.Println("⚠️  Primera vez o configuración incompleta")
	}

	for {
		fmt.Println("\n=== Configurar sistema ===")
		fmt.Println("1) Base de datos")
		fmt.Println("2) Archivos (Usuario del sistema)")
		fmt.Println("3) Iniciar servidor")
		fmt.Print("> ")

		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		switch op {
		case "1":
			// Configurar Base de Datos
			var dbUser, dbPass string
			for {
				fmt.Print("Usuario DB: ")
				dbUser, _ = reader.ReadString('\n')
				dbUser = strings.TrimSpace(dbUser)

				fmt.Print("Password DB: ")
				p, _ := term.ReadPassword(int(os.Stdin.Fd()))
				fmt.Println()
				dbPass = string(p)

				if validarDB(dbUser, dbPass) {
					fmt.Println("✔ Conexión DB OK")
					Conf.Userdb = dbUser
					Conf.Passdb = dbPass

					// Guardar configuración
					guardarConfiguracion(Conf)
					configExists = true
					configCompleta = Conf.UserSistema != "" && Conf.PassSistema != ""
					data, err := os.ReadFile("./schema.sql")
					if err != nil {
						panic(err)
					}
					db := Connect2(dbUser, dbPass)
					defer db.Close()
					_, err = db.Exec(string(data))
					if err != nil {
						panic(err)
					}

					break

				}
				fmt.Println("❌ Credenciales DB incorrectas. Intente nuevamente.")
			}

		case "2":
			// Configurar Usuario del Sistema
			fmt.Print("Ingrese usuario del sistema: ")
			userSistema, _ := reader.ReadString('\n')
			userSistema = strings.TrimSpace(userSistema)

			fmt.Print("Ingrese contraseña del sistema: ")
			passSistema, _ := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()

			// Hashear la contraseña
			hash, err := bcrypt.GenerateFromPassword([]byte(passSistema), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("❌ Error al procesar la contraseña")
				continue
			}

			Conf.UserSistema = userSistema
			Conf.PassSistema = string(hash)

			// Guardar configuración
			guardarConfiguracion(Conf)
			fmt.Println("✔ Usuario del sistema configurado correctamente")
			configExists = true
			configCompleta = Conf.Userdb != "" && Conf.Passdb != ""

		case "3":
			// Iniciar servidor
			if !configCompleta {
				fmt.Println("❌ ERROR: Debe configurar primero:")
				if Conf.Userdb == "" || Conf.Passdb == "" {
					fmt.Println("   - Base de datos (opción 1)")
				}
				if Conf.UserSistema == "" || Conf.PassSistema == "" {
					fmt.Println("   - Usuario del sistema (opción 2)")
				}
				continue
			}

			// Validar credenciales del sistema antes de iniciar
			fmt.Println("\n=== Autenticación del Sistema ===")
			fmt.Print("Usuario del sistema: ")
			userInput, _ := reader.ReadString('\n')
			userInput = strings.TrimSpace(userInput)

			fmt.Print("Contraseña del sistema: ")
			passInput, _ := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()

			// Validar credenciales
			if userInput != Conf.UserSistema {
				fmt.Println("❌ Usuario incorrecto")
				continue
			}

			err := bcrypt.CompareHashAndPassword([]byte(Conf.PassSistema), passInput)
			if err != nil {
				fmt.Println("❌ Contraseña incorrecta")
				continue
			}

			fmt.Println("✔ Autenticación exitosa")

			// Validar conexión DB antes de iniciar
			if !validarDB(Conf.Userdb, Conf.Passdb) {
				fmt.Println("❌ Error: No se puede conectar a la base de datos")
				fmt.Println("   Reconfigure las credenciales (opción 1)")
				continue
			}

			// Iniciar el servidor
			iniciarServidor(Conf)
			return

		default:
			fmt.Println("Opción no válida")
		}
	}

}

/*
	data, errs := os.ReadFile("./Configuracion.Json")
	if errs == nil {

		var Conf Configuracion
		errs = json.Unmarshal(data, &Conf)
		if errs != nil {
			fmt.Println(errs)
		}
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("1) Configurar sistema")
	fmt.Println("2) Salir")
	fmt.Print("> ")

	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	if op != "1" {
		fmt.Println("Saliendo...")
		return
	}

	var dbUser, dbPass string

	for {
		fmt.Print("Usuario DB: ")
		dbUser, _ = reader.ReadString('\n')
		dbUser = strings.TrimSpace(dbUser)

		fmt.Print("Password DB: ")
		p, _ := term.ReadPassword(int(os.Stdin.Fd()))
		fmt.Println()
		dbPass = string(p)

		if validarDB(dbUser, dbPass) {
			fmt.Println("✔ Conexión DB OK")
			break
		}

		fmt.Println("❌ Credenciales DB incorrectas")
	}
*/
//hash, _ := bcrypt.GenerateFromPassword([]byte(dbPass), bcrypt.DefaultCost)

/*
db := Connect()

		defer db.Close()

		var err error
		templates, err = template.ParseGlob("templates/*.html")
		if err != nil {
			log.Fatal("ERROR CRÍTICO: No se pudieron cargar las plantillas HTML. Detalle: ", err)
		}
		log.Println("¡Plantillas cargadas correctamente! Encontrados:", len(templates.Templates()), "archivos")
		app := &App{db: db, Modo: "1", Sistema: Sistema{Configuracion: Configuracion{IdSuper: 0, IdCat: 0, IdPro: 0, Userdb: dbUser, Passdb: dbPass}, Supermercados: []Supermercado{}}}

		r := router.New()

		r.GET("/", app.mostrarHomeOLogin)
		r.POST("/registro", app.procesarRegistro)
		r.POST("/login", app.procesarLogin)
		r.GET("/logout", app.logout)

		r.GET("/supermercado/{id}", app.authMiddleware(app.AddSuper))
		r.GET("/categoria/{id_sup}/{id_cat}", app.authMiddleware(app.AddCategoria))
		r.GET("/producto/{id_sup}/{id_cat}/{id_pro}", app.authMiddleware(app.AddProducto))
		r.POST("/guardar", app.authMiddleware(app.Guardar))
		r.POST("/eliminar", app.authMiddleware(app.Eliminar))

		r.ServeFiles("/static/{filepath:*}", "./static")

		fmt.Println("Servidor corriendo en http://localhost:8080")
		log.Fatal(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
			app.corsMiddleware(r.Handler)(ctx)
		}))
	}
*/
type GuardarData struct {
	Accion    string `json:"accion"`
	Id        string `json:"id"`
	Nombre    string `json:"n"`
	Direccion string `json:"d"`
	Precio    string `json:"p"`
	Id_sup    string `json:"id_sup"`
	Id_cat    string `json:"id_cat"`
	Id_pro    string `json:"id_pro"`
}
type Respuesta struct {
	Accion  string
	Id      string
	Mensaje string
	Id_cat  string
}

func (a *App) Eliminar(ctx *fasthttp.RequestCtx) {

	var Data GuardarData
	body := ctx.PostBody()
	respuesta := Respuesta{}
	respuesta.Mensaje = "Se borró correctamente"

	if err := json.Unmarshal(body, &Data); err != nil {
		ctx.Error("JSON inválido: "+err.Error(), fasthttp.StatusBadRequest)
		fmt.Println(err)

		return
	}
	if Data.Accion == "super" {
		_, err := a.db.Exec(`
        DELETE FROM supermercados
        WHERE id_sup = ? `, Data.Id_sup)
		if err != nil {
			respuesta.Mensaje = "Hubo un problema al eliminar el supermercado"
		}
	}
	if Data.Accion == "categoria" {
		_, err := a.db.Exec(`
        DELETE FROM categorias
        WHERE id_cat = ? `, Data.Id_cat)
		if err != nil {
			respuesta.Mensaje = "Hubo un problema al eliminar la categoria"
		}
	}
	if Data.Accion == "producto" {
		_, err := a.db.Exec(`
        DELETE FROM productos
        WHERE id_pro = ? `, Data.Id_pro)
		if err != nil {
			respuesta.Mensaje = "Hubo un problema al eliminar el producto"
		}
	}
	ctx.Response.Header.SetContentType("application/json")

	jsonBytes, err := json.Marshal(respuesta)
	if err != nil {
		ctx.Error("Error al generar respuesta", fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(jsonBytes)
}
func (a *App) Guardar(ctx *fasthttp.RequestCtx) {

	var Data GuardarData
	body := ctx.PostBody()

	if err := json.Unmarshal(body, &Data); err != nil {
		ctx.Error("JSON inválido: "+err.Error(), fasthttp.StatusBadRequest)
		fmt.Println(err)
		return
	}
	respuesta := Respuesta{}

	if Data.Accion == "super" {
		respuesta.Accion = "super"
		if Data.Id == "0" {
			respuesta.Mensaje = "Supermercado ingresado correctamente"
			a.Sistema.AddSuper(Data.Nombre)
			result, err := a.db.Exec("INSERT INTO supermercados (nombre, direccion, fecha) VALUES (?, ?, NOW())", Data.Nombre, Data.Direccion)
			if err != nil {
				respuesta.Mensaje = "Error al ingresar supermercados"
				fmt.Println(err)
			} else {
				// Obtener el ID del supermercado recién creado
				newSuperId, _ := result.LastInsertId()
				userID, _ := a.getUserID(ctx)

				// Dar permisos automáticamente al creador
				_, err = a.db.Exec(`
                INSERT INTO permisos (id_usr, id_sup, p_addcat, p_addpro) 
                VALUES (?, ?, 1, 1)`, userID, newSuperId)

				if err != nil {
					log.Printf("Error creando permisos automáticos: %v", err)
				} else {
					log.Printf("✅ Permisos creados para user_id=%d en super_id=%d", userID, newSuperId)
				}
			}
		}
	}

	/*if Data.Accion == "super" {
			respuesta.Accion = "super"
			if Data.Id == "0" {
				respuesta.Mensaje = "Supermercado ingresado correctamente"
				a.Sistema.AddSuper(Data.Nombre)
				_, err := a.db.Exec("INSERT INTO supermercados (nombre, direccion, fecha) VALUES (?, ?, NOW())", Data.Nombre, Data.Direccion)
				if err != nil {
					respuesta.Mensaje = "Error al ingresar supermercados"
					fmt.Println(err)
				}
			} else {
				respuesta.Mensaje = "Supermercado actualizado correctamente"
				a.Sistema.ModificarSuper(getUint(Data.Id), Data.Nombre)
				_, err := a.db.Exec(`
				UPDATE supermercados
	        	SET nombre = ?, direccion = ?
	        	WHERE id_sup = ? `, Data.Nombre, Data.Direccion, Data.Id)
				if err != nil {
					respuesta.Mensaje = "Error al actualizar supermercados"
				}
			}
		}
	*/
	if Data.Accion == "categoria" {
		permisos := a.Permisos(Data.Id_sup, ctx)
		if permisos.Add_cat == 1 {
			respuesta.Accion = "categoria"
			respuesta.Id = Data.Id_sup
			if Data.Id == "0" {
				respuesta.Mensaje = "Categoria ingresada correctamente"
				a.Sistema.AddCategoria(getUint(Data.Id_sup), Data.Nombre)
				_, err := a.db.Exec("INSERT INTO categorias (nombre, fecha, id_sup) VALUES (?, NOW(), ?)", Data.Nombre, Data.Id_sup)
				if err != nil {
					respuesta.Mensaje = "Error al ingresar categoria"
					fmt.Println(err)
				}
			} else {
				respuesta.Mensaje = "Categoria actualizada correctamente"
				a.Sistema.ModificarCategoria(getUint(Data.Id_sup), getUint(Data.Id), Data.Nombre)
				_, err := a.db.Exec(`
    		UPDATE categorias 
        	SET nombre = ?
        	WHERE id_cat = ? `, Data.Nombre, Data.Id)
				if err != nil {
					respuesta.Mensaje = "Error al actualizar categorias"
				}
			}
		} else {
			respuesta.Mensaje = "Error no tiene permiso"
			respuesta.Accion = "error"
		}
	}

	if Data.Accion == "producto" {
		permisos := a.Permisos(Data.Id_sup, ctx)
		if permisos.Add_pro == 1 {
			respuesta.Accion = "producto"
			respuesta.Id = Data.Id_sup
			respuesta.Id_cat = Data.Id_cat
			if Data.Id == "0" {
				respuesta.Mensaje = "Producto ingresado correctamente"
				a.Sistema.AddProducto(getUint(Data.Id_sup), getUint(Data.Id_cat), Data.Nombre)
				_, err := a.db.Exec("INSERT INTO productos (nombre, precio, fecha, id_cat) VALUES (?, ?, NOW(), ?)", Data.Nombre, Data.Precio, Data.Id_cat)
				if err != nil {
					respuesta.Mensaje = "Error al ingresar producto"
					fmt.Println(err)
				}
			} else {
				respuesta.Mensaje = "producto actualizado correctamente"
				a.Sistema.ModificarProducto(getUint(Data.Id_sup), getUint(Data.Id_cat), getUint(Data.Id), Data.Nombre)
				_, err := a.db.Exec(`
    		UPDATE productos
        	SET nombre = ?, precio = ?
        	WHERE id_pro = ? `, Data.Nombre, Data.Precio, Data.Id)
				if err != nil {
					respuesta.Mensaje = "Error al actualizar productos"
				}
			}
		} else {
			respuesta.Mensaje = "Error no tiene permiso"
			respuesta.Accion = "error"
		}
	}

	ctx.Response.Header.SetContentType("application/json")

	jsonBytes, err := json.Marshal(respuesta)
	if err != nil {
		ctx.Error("Error al generar respuesta", fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetBody(jsonBytes)
}
func guardarConfiguracion(conf Configuracion) error {
	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("./Configuracion.Json", data, 0600)
}
func iniciarServidor(Conf Configuracion) {
	db := Connect(Conf.Userdb, Conf.Passdb)
	verificarYCrearTablas(db)
	defer db.Close()

	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("ERROR CRÍTICO: No se pudieron cargar las plantillas HTML. Detalle: ", err)
	}
	log.Println("¡Plantillas cargadas correctamente! Encontrados:", len(templates.Templates()), "archivos")

	app := &App{
		db:   db,
		Modo: "1",
		Sistema: Sistema{
			Configuracion: Conf,
			Supermercados: []Supermercado{},
		},
	}
	// En iniciarServidor, después de db := Connect()
	app.limpiarTokensExpirados()

	r := router.New()
	r.GET("/", app.mostrarHomeOLogin)
	r.POST("/registro", app.procesarRegistro)
	r.POST("/login", app.procesarLogin)
	r.GET("/logout", app.logout)
	r.GET("/supermercado/{id}", app.authMiddleware(app.AddSuper))
	r.GET("/categoria/{id_sup}/{id_cat}", app.authMiddleware(app.AddCategoria))
	r.GET("/producto/{id_sup}/{id_cat}/{id_pro}", app.authMiddleware(app.AddProducto))
	r.POST("/guardar", app.authMiddleware(app.Guardar))
	r.POST("/eliminar", app.authMiddleware(app.Eliminar))
	r.ServeFiles("/static/{filepath:*}", "./static")
	r.GET("/recuperar-password", app.mostrarRecuperacion)
	r.POST("/recuperar-password", app.procesarRecuperacion)
	r.GET("/resetear-password", app.mostrarReseteo)
	r.POST("/resetear-password", app.procesarReseteo)
	//r.POST("/recuperar-password", app.procesarRecuperacion)

	fmt.Println("✔ Servidor corriendo en http://localhost:8080")
	log.Fatal(fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		app.corsMiddleware(r.Handler)(ctx)
	}))
}
func (a *App) AddSuper(ctx *fasthttp.RequestCtx) {
	//fmt.Printf("tipo %T", ctx.UserValue("id"))
	idsuper := getID(ctx.UserValue("id"))

	type ListaSuper struct {
		ID     int
		Nombre string
	}
	type Supermercados struct {
		ID        int
		Nombre    string
		Direccion string
		Lista     []ListaSuper
	}
	var super Supermercados

	tipo := 0
	if tipo == 0 {
		files, err := os.ReadDir("./Super")
		if err != nil {
			fmt.Println(err)
		}
		for _, f := range files {
			if !f.IsDir() {
				fmt.Println(f.Name())

				path := fmt.Sprintf("./Super/%v", f.Name())
				data, err := os.ReadFile(path)
				if err != nil {
					fmt.Println(err)
				}

				var superx Supermercado
				err = json.Unmarshal(data, &superx)
				if err != nil {
					fmt.Println(err)
				}
				super.Lista = append(super.Lista, ListaSuper{ID: int(superx.IdSuper), Nombre: superx.Nombre})
			}
		}
	} else {

		super.ID = idsuper
		if idsuper > 0 {
			err := a.db.QueryRow(`
        SELECT nombre, direccion
        FROM supermercados 
        WHERE id_sup = ? `, idsuper).Scan(&super.Nombre, &super.Direccion)
			if err != nil {
				//a.redirect(ctx, "/listar-super?mensaje=Supermercado+no+encontrado+o+sin+permiso")
				return
			}
		}
		rows, err := a.db.Query(`SELECT id_sup, nombre FROM supermercados ORDER BY fecha`)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			fmt.Fprint(ctx, "Error al cargar supermercados")
			return
		}
		defer rows.Close()

		for rows.Next() {
			var s ListaSuper
			rows.Scan(&s.ID, &s.Nombre)
			super.Lista = append(super.Lista, s)
		}
		fmt.Println(super.Lista)
	}
	a.renderHTML(ctx, "agregar-supermercado.html", super)
}
func (a *App) AddCategoria(ctx *fasthttp.RequestCtx) {
	id_sup := getID(ctx.UserValue("id_sup"))
	id_cat := getID(ctx.UserValue("id_cat"))

	tipo := 0
	html := HTML{ID_cat: id_cat, ID_sup: id_sup}
	if tipo == 0 {
		path := fmt.Sprintf("./Super/Super_%v", id_sup)
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}

		var super Supermercado
		err = json.Unmarshal(data, &super)
		if err != nil {
			fmt.Println(err)
		}
		if id_cat > 0 {
			for i := 0; i < len(super.Categorias); i++ {
				if super.Categorias[i].IdCat == uint64(id_cat) {
					html.NOMBRE = super.Categorias[i].Nombre
				}
			}
		}
		html.SUPER = super
	} else {

		if id_cat > 0 {
			err := a.db.QueryRow(`
	        SELECT nombre
	        FROM categorias
	        WHERE id_cat = ? `, id_cat).Scan(&html.NOMBRE)
			if err != nil {
				//a.redirect(ctx, "/listar-super?mensaje=Supermercado+no+encontrado+o+sin+permiso")
				return
			}
		}

		rows, err := a.db.Query(`SELECT id_cat, nombre FROM categorias WHERE id_sup = ? ORDER BY fecha`, id_sup)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			fmt.Fprint(ctx, "Error al cargar supermercados")
			return
		}
		defer rows.Close()

		for rows.Next() {
			var s Categoria
			rows.Scan(&s.IdCat, &s.Nombre)
			html.SUPER.Categorias = append(html.SUPER.Categorias, s)
		}
	}
	a.renderHTML(ctx, "agregar-categorias.html", html)

}
func (a *App) AddProducto(ctx *fasthttp.RequestCtx) {

	id_sup := getID(ctx.UserValue("id_sup"))
	id_cat := getID(ctx.UserValue("id_cat"))
	id_pro := getID(ctx.UserValue("id_pro"))

	tipo := 0
	html := HTML{ID_cat: id_cat, ID_sup: id_sup, ID_pro: id_pro}
	if tipo == 0 {
		path := fmt.Sprintf("./Super/Super_%v", id_sup)
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		var super Supermercado
		err = json.Unmarshal(data, &super)
		if err != nil {
			fmt.Println(err)
		}
		if id_pro > 0 {
			for i := 0; i < len(super.Categorias); i++ {
				if super.Categorias[i].IdCat == uint64(id_cat) {
					for j := 0; j < len(super.Categorias[i].Productos); j++ {
						if super.Categorias[i].Productos[j].IdPro == uint64(id_pro) {
							html.NOMBRE = super.Categorias[i].Productos[j].Nombre
						}
					}
				}
			}
		}
		html.SUPER = super
	} else {

		if id_pro > 0 {
			err := a.db.QueryRow(`
	        SELECT nombre
	        FROM productos
	        WHERE id_pro = ? `, id_pro).Scan(&html.NOMBRE)
			if err != nil {
				//a.redirect(ctx, "/listar-super?mensaje=Supermercado+no+encontrado+o+sin+permiso")
				return
			}
		}
		html.SUPER = a.GetDbSuper(id_sup)
	}
	a.renderHTML(ctx, "agregar-producto.html", html)

}
func (a *App) GetDbSuper(id_sup int) Supermercado {

	var Supermercado Supermercado

	rows, err := a.db.Query(`SELECT id_cat, nombre FROM categorias WHERE id_sup = ? ORDER BY fecha`, id_sup)
	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var s Categoria
		rows.Scan(&s.IdCat, &s.Nombre)

		rows2, err := a.db.Query(`SELECT id_pro, nombre FROM productos WHERE id_cat = ? ORDER BY fecha`, s.IdCat)
		if err != nil {
			fmt.Println(err)
		}
		defer rows2.Close()

		for rows2.Next() {
			var p Producto
			rows.Scan(&p.IdPro, &s.Nombre)
			s.Productos = append(s.Productos, p)
		}
		Supermercado.Categorias = append(Supermercado.Categorias, s)
	}
	return Supermercado
}

type Permisos struct {
	Add_sup int
	Add_cat int
	Add_pro int
}

/*func (a *App) Permisos(id_sup string, ctx *fasthttp.RequestCtx) Permisos {
	userID, ok := a.getUserID(ctx)
	if !ok {
		log.Printf("Intento de acceso no autenticado a permisos")
		return Permisos{}
	}

	log.Printf("Cargando permisos - userID=%d, id_sup=%s", userID, id_sup)

	if id_sup == "" {
		log.Printf("id_sup vacío")
		return Permisos{}
	}

	idSupInt, err := strconv.Atoi(id_sup)
	if err != nil {
		log.Printf("id_sup no es número: %s", id_sup)
		return Permisos{}
	}

	var permisos Permisos
	err = a.db.QueryRow(`
        SELECT p_addcat, p_addpro
        FROM permisos
        WHERE id_usr = ? AND id_sup = ?`,
		userID, idSupInt).Scan(&permisos.Add_cat, &permisos.Add_pro)

	if err == sql.ErrNoRows {
		log.Printf("No permiso encontrado para userID=%d y id_sup=%d", userID, idSupInt)
		return Permisos{}
	}
	if err != nil {
		log.Printf("Error en consulta permisos: %v", err)
		return Permisos{}
	}

	log.Printf("Permisos encontrados - add_cat=%v, add_pro=%v", permisos.Add_cat, permisos.Add_pro)
	return permisos
}*/

func (a *App) Permisos(id_sup string, ctx *fasthttp.RequestCtx) Permisos {
	userID, ok := a.getUserID(ctx)
	if !ok {
		return Permisos{} // o maneja error
	}

	var permisos Permisos
	err := a.db.QueryRow(`
        SELECT t2.p_addsup, t1.p_addcat, t1.p_addpro
        FROM permisos t1, usuarios t2
        WHERE t1.id_usr = ? AND t1.id_sup = ? AND t1.id_usr = t2.id_usr`,
		userID, id_sup).Scan(&permisos.Add_sup, &permisos.Add_cat, &permisos.Add_pro)
	if err != nil {
		log.Printf("Error cargando permisos: %v", err)
		return Permisos{}
	}
	return permisos
}
func (a *App) procesarRegistro(ctx *fasthttp.RequestCtx) {
	nombre := strings.TrimSpace(string(ctx.PostArgs().Peek("nombre")))
	correo := strings.TrimSpace(string(ctx.PostArgs().Peek("correo")))
	password := strings.TrimSpace(string(ctx.PostArgs().Peek("password")))

	if nombre == "" || correo == "" || password == "" {
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Todos los campos son obligatorios"})
		return
	}

	var exists int
	err := a.db.QueryRow("SELECT 1 FROM usuarios WHERE correo = ?", correo).Scan(&exists)
	if err == nil {
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Este correo ya está registrado"})
		return
	} else if err != sql.ErrNoRows {
		log.Printf("Error verificando correo: %v", err)
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Error interno al verificar correo"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		log.Printf("Error generando hash: %v", err)
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Error interno al crear la cuenta"})
		return
	}

	_, err = a.db.Exec(`
    INSERT INTO usuarios (nombre, correo, password, cookie, p_addsup)
    VALUES (?, ?, ?, ?, ?)
`, nombre, correo, hashedPassword, "", 1)

	if err != nil {
		log.Printf("Error al registrar usuario: %v", err)
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Error al crear la cuenta. Intente con otro correo."})
		return
	}

	a.renderHTML(ctx, "login.html", map[string]string{"Success": "¡Cuenta creada con éxito! Ahora inicia sesión."})
}
func (a *App) procesarLogin(ctx *fasthttp.RequestCtx) {
	correo := strings.TrimSpace(string(ctx.PostArgs().Peek("username")))
	password := strings.TrimSpace(string(ctx.PostArgs().Peek("password")))

	var dbID int
	var dbHashedPass string
	err := a.db.QueryRow("SELECT id_usr, password FROM usuarios WHERE correo = ?", correo).Scan(&dbID, &dbHashedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			a.renderHTML(ctx, "login.html", map[string]string{"Error": "Correo o contraseña incorrectos"})
		} else {
			log.Printf("Error en query login: %v", err)
			a.renderHTML(ctx, "login.html", map[string]string{"Error": "Error interno del servidor"})
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbHashedPass), []byte(password)); err != nil {
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Correo o contraseña incorrectos"})
		return
	}

	sessionToken, err := generarCookieSegura(32)
	if err != nil {
		log.Printf("Error generando token: %v", err)
		a.renderHTML(ctx, "login.html", map[string]string{"Error": "Error interno del servidor"})
		return
	}

	_, err = a.db.Exec("UPDATE usuarios SET cookie = ? WHERE id_usr = ?", sessionToken, dbID)
	if err != nil {
		log.Printf("Error guardando token en BD: %v", err)
	}

	cookie := fasthttp.AcquireCookie()
	cookie.SetKey("session")
	cookie.SetValue(sessionToken)
	cookie.SetExpire(time.Now().Add(24 * time.Hour))
	cookie.SetPath("/")
	cookie.SetHTTPOnly(true)
	// cookie.SetSecure(true)          // Descomenta cuando tengas HTTPS
	cookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	ctx.Response.Header.SetCookie(cookie)
	fasthttp.ReleaseCookie(cookie)

	log.Printf("Usuario %d logueado correctamente", dbID)
	ctx.Redirect("/", fasthttp.StatusFound)
}
func generarCookieSegura(longitud int) (string, error) {
	bytes := make([]byte, longitud)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	letras := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := range bytes {
		bytes[i] = letras[bytes[i]%byte(len(letras))]
	}
	return string(bytes), nil
}

// Generar token aleatorio
func generarToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// Enviar email
func enviarEmailRecuperacion(destinatario, token string, conf Configuracion) error {
	log.Println("📧 Intentando enviar email a:", destinatario)
	log.Println("📧 Usando SMTP:", conf.EmailSMTP)

	m := gomail.NewMessage()
	m.SetHeader("From", conf.EmailSMTP)
	m.SetHeader("To", destinatario)
	m.SetHeader("Subject", "Recuperación de Contraseña")

	enlace := fmt.Sprintf("http://localhost:8080/resetear-password?token=%s", token)

	cuerpo := fmt.Sprintf(`
        <html>
        <body style="font-family: Arial, sans-serif;">
            <h2>Recuperación de Contraseña</h2>
            <p>Has solicitado restablecer tu contraseña.</p>
            <p>Haz clic en el siguiente enlace para crear una nueva contraseña:</p>
            <p><a href="%s" style="background-color: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 5px; display: inline-block;">Restablecer Contraseña</a></p>
            <p>Este enlace expirará en 1 hora.</p>
            <p>Si no solicitaste este cambio, ignora este correo.</p>
        </body>
        </html>
    `, enlace)

	m.SetBody("text/html", cuerpo)

	d := gomail.NewDialer("smtp.gmail.com", 587, conf.EmailSMTP, conf.PasswordSMTP)

	err := d.DialAndSend(m)
	if err != nil {
		log.Println("❌ Error enviando email:", err)
		return err
	}

	log.Println("✅ Email enviado correctamente a:", destinatario)
	return nil
}

// Mostrar página de recuperación (GET)
func (a *App) mostrarRecuperacion(ctx *fasthttp.RequestCtx) {
	data := map[string]interface{}{}
	a.renderHTML(ctx, "recuperar.html", data)
}

// Procesar solicitud de recuperación (POST)
func (a *App) procesarRecuperacion(ctx *fasthttp.RequestCtx) {
	correo := string(ctx.FormValue("correo"))

	// Verificar si el correo existe
	var idUsr int
	err := a.db.QueryRow("SELECT id_usr FROM usuarios WHERE correo = ?", correo).Scan(&idUsr)

	if err != nil {
		// No revelar si el correo existe o no (seguridad)
		data := map[string]interface{}{
			"Exito": "Si el correo existe, recibirás un enlace de recuperación.",
		}
		a.renderHTML(ctx, "recuperar.html", data)
		return
	}

	// Generar token
	token := generarToken()
	expiracion := time.Now().Add(1 * time.Hour) // Expira en 1 hora

	// Guardar token en BD
	_, err = a.db.Exec(`
        INSERT INTO password_reset_tokens (id_usr, token, expiracion) 
        VALUES (?, ?, ?)
    `, idUsr, token, expiracion)

	if err != nil {
		data := map[string]interface{}{
			"Error": "Error al procesar la solicitud. Intenta nuevamente.",
		}
		a.renderHTML(ctx, "recuperar.html", data)
		return
	}

	// Enviar email
	err = enviarEmailRecuperacion(correo, token, a.Sistema.Configuracion)

	if err != nil {
		log.Println("Error al enviar email:", err)
		data := map[string]interface{}{
			"Error": "Error al enviar el correo. Verifica la configuración.",
		}
		a.renderHTML(ctx, "recuperar.html", data)
		return
	}

	data := map[string]interface{}{
		"Exito": "Si el correo existe, recibirás un enlace de recuperación.",
	}
	a.renderHTML(ctx, "recuperar.html", data)
}

func verificarYCrearTablas(db *sql.DB) {
	// Verificar si existe la tabla password_reset_tokens
	var tableName string
	err := db.QueryRow("SHOW TABLES LIKE 'password_reset_tokens'").Scan(&tableName)

	if err != nil || tableName == "" {
		// La tabla no existe, crearla
		log.Println("⚠️  Tabla password_reset_tokens no encontrada. Creándola...")

		createTable := `
		CREATE TABLE password_reset_tokens (
			id INT AUTO_INCREMENT PRIMARY KEY,
			id_usr INT NOT NULL,
			token VARCHAR(100) NOT NULL UNIQUE,
			expiracion DATETIME NOT NULL,
			usado BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (id_usr) REFERENCES usuarios(id_usr) ON DELETE CASCADE
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
		`

		_, err := db.Exec(createTable)
		if err != nil {
			log.Println("❌ Error al crear tabla password_reset_tokens:", err)
		} else {
			log.Println("✔ Tabla password_reset_tokens creada correctamente")
		}
	} else {
		log.Println("✔ Tabla password_reset_tokens ya existe")
	}
}
func (a *App) limpiarTokensExpirados() {
	ticker := time.NewTicker(24 * time.Hour) // Cada 24 horas
	go func() {
		for range ticker.C {
			_, err := a.db.Exec("DELETE FROM password_reset_tokens WHERE expiracion < NOW()")
			if err != nil {
				log.Println("Error limpiando tokens:", err)
			} else {
				log.Println("Tokens expirados limpiados")
			}
		}
	}()
}
func (a *App) logout(ctx *fasthttp.RequestCtx) {
	// Limpiar cookie session
	cookie := fasthttp.AcquireCookie()
	cookie.SetKey("session")
	cookie.SetValue("")
	cookie.SetPath("/")
	cookie.SetHTTPOnly(true)
	// cookie.SetSecure(true) // cuando HTTPS
	cookie.SetSameSite(fasthttp.CookieSameSiteLaxMode)
	cookie.SetExpire(time.Now().Add(-24 * time.Hour)) // Expira en pasado
	ctx.Response.Header.SetCookie(cookie)
	fasthttp.ReleaseCookie(cookie)

	a.redirect(ctx, "/")
}
func (a *App) redirect(ctx *fasthttp.RequestCtx, url string) {
	ctx.Redirect(url, fasthttp.StatusFound)
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
func (a *App) getUserID(ctx *fasthttp.RequestCtx) (int, bool) {
	cookie := ctx.Request.Header.Cookie("session")
	if cookie == nil {
		return 0, false
	}
	token := string(cookie)

	var userID int
	err := a.db.QueryRow("SELECT id_usr FROM usuarios WHERE cookie = ?", token).Scan(&userID)
	if err != nil {
		return 0, false
	}
	return userID, true
}
func (a *App) authMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		userID, ok := a.getUserID(ctx)
		if !ok {
			ctx.Redirect("/", fasthttp.StatusFound)
			return
		}
		ctx.SetUserValue("userID", userID)
		next(ctx)
	}
}
func (a *App) mostrarHomeOLogin(ctx *fasthttp.RequestCtx) {
	if _, ok := a.getUserID(ctx); ok {
		a.renderHTML(ctx, "admin.html", nil)
		//a.redirect(ctx, "/")
		return
	}
	a.renderHTML(ctx, "login.html", nil)

}

// Mostrar página de reseteo (GET)
func (a *App) mostrarReseteo(ctx *fasthttp.RequestCtx) {
	token := string(ctx.QueryArgs().Peek("token"))

	if token == "" {
		data := map[string]interface{}{
			"Error": "Token no válido o faltante.",
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Verificar que el token existe y no ha expirado
	var idUsr int
	var expiracion time.Time
	var usado bool

	err := a.db.QueryRow(`
        SELECT id_usr, expiracion, usado 
        FROM password_reset_tokens 
        WHERE token = ?
    `, token).Scan(&idUsr, &expiracion, &usado)

	if err != nil {
		data := map[string]interface{}{
			"Error": "El enlace de recuperación no es válido.",
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Verificar si ya fue usado
	if usado {
		data := map[string]interface{}{
			"Error": "Este enlace ya fue utilizado.",
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Verificar si expiró
	if time.Now().After(expiracion) {
		data := map[string]interface{}{
			"Error": "Este enlace ha expirado. Solicita uno nuevo.",
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Token válido, mostrar formulario
	data := map[string]interface{}{
		"Token": token,
	}
	a.renderHTML(ctx, "resetear.html", data)
}

// Procesar reseteo de contraseña (POST)
func (a *App) procesarReseteo(ctx *fasthttp.RequestCtx) {
	token := string(ctx.FormValue("token"))
	password := string(ctx.FormValue("password"))
	passwordConfirm := string(ctx.FormValue("password_confirm"))

	// Validar que las contraseñas coincidan
	if password != passwordConfirm {
		data := map[string]interface{}{
			"Error": "Las contraseñas no coinciden.",
			"Token": token,
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Validar longitud mínima
	if len(password) < 6 {
		data := map[string]interface{}{
			"Error": "La contraseña debe tener al menos 6 caracteres.",
			"Token": token,
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Verificar token nuevamente
	var idUsr int
	var expiracion time.Time
	var usado bool

	err := a.db.QueryRow(`
        SELECT id_usr, expiracion, usado 
        FROM password_reset_tokens 
        WHERE token = ?
    `, token).Scan(&idUsr, &expiracion, &usado)

	if err != nil || usado || time.Now().After(expiracion) {
		data := map[string]interface{}{
			"Error": "El enlace no es válido o ha expirado.",
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Hashear nueva contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		data := map[string]interface{}{
			"Error": "Error al procesar la contraseña.",
			"Token": token,
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Actualizar contraseña en BD
	_, err = a.db.Exec("UPDATE usuarios SET password = ? WHERE id_usr = ?", string(hash), idUsr)
	if err != nil {
		data := map[string]interface{}{
			"Error": "Error al actualizar la contraseña.",
			"Token": token,
		}
		a.renderHTML(ctx, "resetear.html", data)
		return
	}

	// Marcar token como usado
	_, err = a.db.Exec("UPDATE password_reset_tokens SET usado = TRUE WHERE token = ?", token)
	if err != nil {
		log.Println("Error al marcar token como usado:", err)
	}

	// Éxito
	data := map[string]interface{}{
		"Exito": "✔ Tu contraseña ha sido actualizada correctamente.",
	}
	a.renderHTML(ctx, "resetear.html", data)
	a.renderHTML(ctx, "resetear.html", nil)

}

func (a *App) renderHTML(ctx *fasthttp.RequestCtx, name string, data any) {
	ctx.SetContentType("text/html; charset=utf-8")
	if templates == nil {
		fmt.Fprint(ctx, "<h1>Error fatal: Plantillas no cargadas</h1>")
		return
	}
	err := templates.ExecuteTemplate(ctx, name, data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		fmt.Fprintf(ctx, "<h1>Error al renderizar la plantilla: %s</h1><pre>%v</pre>", name, err)
		log.Printf("ERROR RENDER %s: %v", name, err)
		return
	}
}
func Connect(user, pass string) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_Super?parseTime=true", user, pass)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error conectando a MySQL:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping falló:", err)
	}
	log.Println("✓ Base de datos conectada")
	return db
}

func Connect2(user, pass string) *sql.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/?multiStatements=true", user, pass)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	if err = db.Ping(); err != nil {
		log.Fatal("Ping falló:", err)
	}
	log.Println("✓ Base de datos conectada")
	return db

}

func validarDB(user, pass string) bool {

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/", user, pass)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return false
	}
	defer db.Close()

	return db.Ping() == nil

}
func getID(x interface{}) int {
	s, ok := x.(string)
	if !ok {
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
func getUint(s string) uint64 {
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return n
}
