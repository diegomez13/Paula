var idiomabase = {"es-CL": ["Bienvenidos1", "Lorem Ipsum es simplemente el texto de relleno de las imprentas y archivos de texto. Lorem Ipsum ha sido el texto de relleno estándar de las industrias desde el año 1500, cuando un impresor (N. del T. persona que se dedica a la imprenta) desconocido usó una galería de textos y los mezcló de tal manera que logró hacer un libro de textos especimen."]};
var idiomadefaults = ["Bienvenidos2", "Lorem Ipsum es simplemente el texto de relleno de las imprentas y archivos de texto. Lorem Ipsum ha sido el texto de relleno estándar de las industrias desde el año 1500, cuando un impresor (N. del T. persona que se dedica a la imprenta) desconocido usó una galería de textos y los mezcló de tal manera que logró hacer un libro de textos especimen."];
var idiomanum = {"es-CL":0, "en-US": 1}
var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_";

var base;
document.addEventListener('DOMContentLoaded', async function () {
    ajustarClaseSegunAncho();
    base = new Base();
    base.IdiomaBase();

    //var pop = {c:"pi",d:[{c:"pd", d:[{c:"pd0", d:[{c:"pt", h:x[0]}, {c:"pe", h:x[1]}, {c:"pf"}, {c:"pa",d:[{c:"px",h:"iniciando...."},{c:"loading", d:[{c:"barra"}]}]}, {c:"pb",d:[{c:"pz", h:"Empezar"}]}]}]}]};
    //GC("p").replaceChild(render(pop), GC("pi"));
    //GC("pi").style.display = "block";
    //base.Load(version);
});

function cargarContenido(url) {
    const contenedor = document.getElementsByClassName('ch')[0];

    // 1. Llamar a la URL
    fetch(url)
        .then(function(respuesta) {
    console.log("URL pedida:", respuesta.url);
    console.log("Código de estado:", respuesta.status); // Si sale 404, la ruta está mal
    
    if (!respuesta.ok) {
        throw new Error("El servidor respondió con código " + respuesta.status);
    }
    return respuesta.text();
})
        .then(function(html) {
            // 4. Incrustar el HTML en el div
            contenedor.innerHTML = html;
        })
        .catch(function(error) {
            // 5. Manejar errores (CORS, 404, etc.)
            console.error("Hubo un problema:", error);
            contenedor.innerHTML = "Error al cargar el contenido.";
        });
}

function Send(){

    console.log(GC("f"));


    var f = GC("f").querySelectorAll('input');
    var obj = {};
    for(var x of f){
        obj[x.id] = x.value
               // obj[x.id] = { v: x.value, c: x.className }

    }

 fetch('/guardar', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json'            
    },
    body: JSON.stringify(obj)
})
.then(response => {
    if (!response.ok) throw new Error('Error en la respuesta');
    return response.json();
})
.then(data => console.log('Éxito:', data))
.catch(error => console.error('Fallo:', error));
}

function LoadCategoria(){
    


    base.LoadCat(0, 0);
}
function ShowCat(){

}
function SetDir(){
    var info = SplitParam(Parents(this, "bt").id);
    base.SetTipo(info.pid, "Directorio");
    base.LoadCat(info.id, info.pid);
}
function SetProd(){
    var info = SplitParam(Parents(this, "bt").id);
    base.SetTipo(info.pid, "Producto");
    base.LoadCat(info.id, info.pid);
}
function Accion1(){
    console.log("ACCION1", this);
}
function Accion2(){
    console.log("ACCION2", this);
}
function AccionItem(){
    var info = SplitParam(Parents(this, "it r").id);
    base.LoadCat(0, info.id);
    //console.log("Accion: "+info.id+"/"+info.pid);
}
function AccionesItems(){
    var info = SplitParam(Parents(this, "ft").id);
    console.log("Accion: "+info.id+"/"+info.pid);
}
function AccionesForm(){
    var info = SplitParam(Parents(this, "ft").id);
    console.log("Accion: "+info.id+"/"+info.pid);
}
var tempauto = {};
function FindAuto(v, n){

    const texto = "Noticias de Deportes";

    if (texto.toLowerCase().includes("deportes")) {
        console.log("Contiene deportes");
    }

    var html = {
        c: "al",
        d: []
    };

    html.d.push({c:"al1", d:[{c:"al2", h:"HOLA"},{c:"al3", h:"Nelson"}]});
    html.d.push({c:"al1", d:[{c:"al2", h:"Diego"},{c:"al3", h:"Gomez"}]});

    for (var i=0; i<v.length; i++){
        if (tempauto[v.substring(0, i+1)] !== undefined){
            console.log(tempauto[v.substring(0, i+1)]);
        } else {
            break;
        }
    }
    
    
    
    //console.log(v);
    //console.log(tempauto);
    return html;
}
function Auto(){

    GC("rauto").replaceChild(render(FindAuto(this.value, 0)), GC("al"));
    if (tempauto[this.value] === undefined){
        SendRequestBytes('/a?d='+this.value).then(data => {
            tempauto[this.value] = data;
            GC("rauto").replaceChild(render(FindAuto(this.value, 1)), GC("al"));
        }).catch(error => {
            console.log("Error Usuario", error);
        });
    }
    
}
function LoadAdmin(){
    /*
    Loding(0, "Iniciando");
    setTimeout(function(){
        Loding(20, "Lenguaje");
    }, 500);
    setTimeout(function(){
        Loding(30, "Informacion del usuario");
    }, 1000);
    setTimeout(function(){
        Loding(40, "Informacion de empresa");
    }, 1500);
    setTimeout(function(){
        GC("pa").style.display = "none";
        GC("pb").style.display = "flex";
    }, 2000);

    GC("h").replaceChild(render(data.Start()), GC("ch"));
    */
}
function Loding(n, s){
    var z = parseInt(GC("barra").style.width || 0) + n;
    GC("barra").style.width = z+"%";
    GC("px").textContent = s;
}
function Empezar(){
    data.SetInicio();
    LoadAdmin();
}
function Terminos(){
    data.SetTerminos();
    LoadAdmin();
}
function ConfBase(){
    var f = Parents(this, "f").querySelectorAll('input');
    console.log(f);
}
function DeleteStorage(){
    localStorage.setItem("local_user_01", null);
    localStorage.setItem("local_lang_01", null);
    localStorage.setItem("local_cats_01", null);
}

class Htmls {
    
    Inicio(){

        var x = this.UserEmp.data;
        var pop = {c:"pi",d:[{c:"pd", d:[{c:"pd0", d:[{c:"pt", h:x[0]}, {c:"pe", h:x[1]}, {c:"pf"}, {c:"pa",d:[{c:"px",h:"iniciando...."},{c:"loading", d:[{c:"barra"}]}]}, {c:"pb",d:[{c:"pz", h:"Empezar"}]}]}]}]};
        GC("p").replaceChild(render(pop), GC("pi"));
        GC("pi").style.display = "block";

    }
    LoadCat(id, pid){
        
        var Data = this.SqlCategoria(id, pid);
        var html;

        if (pid == 0 || Data.p.Tipo == "Directorio"){

            var form = [
                { t: 1, i: "accion", v: "suser" },
                { t: 1, i: "id", v: id },
                { t: 1, i: "pid", v: pid },
                { t: 2, h: "Nombre", i: "n", v: Data.e.Nombre },
                { t: 3, h: "Enviar", c: "enviar", a: "click", f: Send }
            ];

            var acciones_titulo = [{c:"i i1", f:Accion1}, {c:"i i2", f:Accion2}];
            var acciones_form = [{c:"i i3", f:AccionesForm}, {c:"i i4", f:AccionesForm}];
            var acciones_items = [{c:"i i5", f:AccionesItems}, {c:"i i6", f:AccionesItems}];
            var acciones_lista = [{c:"i i5", f:AccionItem}, {c:"i i6", f:AccionItem}];

            html = {
                c:"ch",d:[
                    this.TituloSimple("Categorias", "Completa los Campos", acciones_titulo),
                    this.Form("HOLA", "BUENA", id, pid, form, acciones_form),
                    this.Items("HOLA", "BUENA", id, pid, Data.l, acciones_items, acciones_lista)
                ]
            };

        } else if (Data.p.Tipo == undefined) {

            var acciones_titulo = [{c:"i i1", f:Accion1}, {c:"i i2", f:Accion2}];
            html = {
                c:"ch",d:[
                    this.TituloSimple("Categorias", "Completa los Campos", acciones_titulo),
                    this.PreguntaCategoria(id, pid),
                ]
            };

        } else if (Data.p.Tipo == "Producto") {

            if(Data.p.Cats === undefined){

                var acciones_titulo = [{c:"i i1", f:Accion1}, {c:"i i2", f:Accion2}];
                var acciones_cat = [{c:"i i1", f:Accion1}, {c:"i i2", f:Accion2}];

                html = {
                    c:"ch",d:[
                        this.TituloSimple("Seleccionar Categoria", "Completa los Campos", acciones_titulo),
                        this.CategoriaAutoComplete("HOLA", "BUENA", id, pid, acciones_cat)
                    ]
                };

            } else {

                var form = [
                    { t: 1, i: "accion", v: "suser" },
                    { t: 1, i: "id", v: id },
                    { t: 1, i: "pid", v: pid },
                    { t: 2, h: "Nombre", i: "n", v: Data.e.Nombre },
                    { t: 3, h: "Enviar", c: "enviar", a: "click", f: Send }
                ];

                var acciones_titulo = [{c:"i i1", f:Accion1}, {c:"i i2", f:Accion2}];
                var acciones_form = [{c:"i i3", f:AccionesForm}, {c:"i i4", f:AccionesForm}];
                var acciones_items = [{c:"i i5", f:AccionesItems}, {c:"i i6", f:AccionesItems}];
                var acciones_lista = [{c:"i i5", f:AccionItem}, {c:"i i6", f:AccionItem}];

                html = {
                    c:"ch",d:[
                        this.TituloSimple("Productos", "Completa los Campos", acciones_titulo),
                        this.Form("HOLA", "BUENA", id, pid, form, acciones_form),
                        this.Items("HOLA", "BUENA", id, pid, Data.l, acciones_items, acciones_lista)
                    ]
                };

            }

        } else {

        }
        
        GC("h").replaceChild(render(html), GC("ch"));
    }
    TituloSimple(ti, st, a = []){
        return {
            c: "ht r",
            d: [
                { t:6, h: ti },
                { t:9, h: st },
                { c: "clearfix va ri", d: this.RenderAcciones(a) }
            ]
        }
    }
    Form(Titulo, SubTitulo, id, pid, Inputs, a = []){
        return {
            c: "hf bt",
            d: [
                { c: "ft", i: id+"-"+pid, d: [
                    { c:"r", d:[
                        { t:8, h: Titulo }, 
                        { c: "clearfix va ri", d: this.RenderAcciones(a)
                        }] 
                    }, 
                    { t:9, h: SubTitulo },
                    { c: "f", d: this.AddInputs(Inputs) }
                    ]
                }
            ]
        }
    }
    Items(Titulo, Subtitulo, id, pid, Lista, a = [], b = []){
        if (Lista.length > 0) {
            return {
                c: "hf ab",
                d: [
                    { c: "ft", i: id+"-"+pid, d: [
                        { c:"r", d:[
                            { t:8, h: Titulo },
                            { c: "clearfix va ri", d: this.RenderAcciones(a)
                            }]
                        },
                        { t:9, h: Subtitulo },
                        { n: "list", c: "lis", d: this.RenderLista(Lista, ShowCat, b) }
                        ]
                    }
                ]
            }
        } else {
            return {}
        }
    }
    AddInputs(Inputs){
        var res = [];
        for (var x of Inputs){
            if (x.t == 1){
                res.push({t:3,y:3,i:x.i,v:x.v});
            }else if(x.t == 2){
                res.push({t:10,d:[{t:2,h:x.h},{t:3,y:0,i:x.i,v:x.v},{c:"msg"}]});
            }else{
                res.push({c:"s", d:[{h:x.h,c:x.c,f:{m:x.a,n:x.f}}]});
            }
        }
        return res;
    }
    RenderLista(Lista, Func, a = []){
        var res = [];
        for (var x of Lista){
            res.push({c:"it r",i:x.i+"-"+x.p,d:[{c:"im",h:x.n, f:{m:"click",n:Func}},{c:"clearfix va ri",d: this.RenderAcciones(a)}]});
        }
        return res;
    }
    RenderAcciones(Acciones){
        var res = [];
        for (var x of Acciones){
            res.push({t:7,c:x.c,f:{m:"click",n:x.f}})
        }
        return res;
    }
    PreguntaCategoria(id, pid){
        return {
            c: "bt", 
            i:id+"-"+pid,
            d: [{c:"rt", f:{m:"click",n:SetDir},d:[{c:"f1", h:"Subdirectorio"},{c:"f2",h:"Seleccionar para ingresar subdirectorios"}]}, {c:"rt", f:{m:"click",n:SetProd},d:[{c:"f1", h:"Producto"},{c:"f2", h:"Seleccionar para ingresar productos"}]}]
        }
    }
    CategoriaAutoComplete(Titulo, SubTitulo, id, pid, a){
        return {
            c: "hf bt",
            d: [
                { c: "ft", i: id+"-"+pid, d: [
                    { c:"r", d:[
                        { t:8, h: Titulo }, 
                        { c: "clearfix va ri", d: this.RenderAcciones(a)
                        }] 
                    }, 
                    { t:9, h: SubTitulo },
                    { c: "f", d: [{t:10,d:[{t:2,h:"AutoComplete"},{t:3,y:0,v:"",f:{m:"input",n:Auto}},{c:"rauto", d:[{c:"al"}]}]}] }
                    ]
                }
            ]
        }
    }
    LoadMenu(){
        var x = this.UserEmp.data;
        var menu = {
            c:"lc",
            d:[
            {
                c:"lt",
                d:[{c:"lo"},{c:"ld",h:"Administracion"}]
            },{
                c:"li",
                d:[{c:"lo"},{c:"ld",h:"Ingresar Categoria",i:1,f:{m:"click",n:LoadCategoria}}]
            },{
                c:"li",
                d:[{c:"lo"},{c:"ld",h:"Ingresar Productos"}]
            },{
                c:"li",
                d:[{c:"lo"},{c:"ld",h:"Configuracion"}]
            }]
        };
        
        GC("l").replaceChild(render(menu), GC("lc"));
    }
    TituloAcciones(a){
        var aux = {c:"ha",d:[]}
        for (x in a){
            aux.d.push({});
        }
        return aux;
    }
}


class Base extends Htmls{
    constructor(){
        super();
        this.UserEmp = LocalStorage('local_user_01');
        this.Idioma = LocalStorage('local_lang_01');
        this.Categorias = LocalStorage('local_cats_01');
    }
    GetIdCat(){
        this.Categorias.Idmax++;
        return this.Categorias.Idmax;
    }
    AddCategoria(obj){

        if (obj.id.v == 0){
            // AGREGAR
            this.Categorias.Lista.push({ Nombre: obj.n.v, id: this.GetIdCat(), pid: parseInt(obj.pid.v)});
        } else {
            // MODIFICAR
            for (var i=0; i<this.Categorias.Lista.length; i++){
                if (this.Categorias.Lista[i].id == obj.id.v) {
                    this.Categorias.Lista[i].Nombre = obj.n.v;
                }
            }
        }
        this.SaveCat();
    }
    SqlCategoria(id, pid){
        var res = {l:[], e:{ Nombre:"" }, p:{} };
        for (var x of this.Categorias.Lista){
            if (x.id == id){
                res.e = x;
            }
            if (x.pid == pid){
                res.l.push({i:x.id, p:x.pid, n: x.Nombre});
            }
            if (x.id == pid){
                res.p = x;
            }
        }
        return res;
    }
    SetTipo(id, tipo){
        for (var i=0; i<this.Categorias.Lista.length; i++){
            if (this.Categorias.Lista[i].id == id){
                this.Categorias.Lista[i].Tipo = tipo;
            }
        }
    }
    IdiomaBase(){

        super.LoadMenu();
        /*
        if (this.UserEmp === undefined || this.UserEmp.version != version) {
            var user_idioma = this.UserEmp.idioma || navigator.language || navigator.languages[0];
            SendRequestJson("/i?i="+user_idioma+"&v="+version).then(data => {
                //Loding(50, "Datos Idioma...");
                this.UserEmp = { version: version, idioma: user_idioma, data: data };
                super.Inicio();
                this.SaveUser();
            }).catch(error => {
                console.log(error);
            });
        } else {
            super.Inicio();
        }

        SendRequestBytes('/u').then(data => {
            //Loding(50, "Datos Usuarios...");
            this.UserEmp = DecodeUserEmp(data);
            console.log("/u", data);
        }).catch(error => {
            console.log("Error Usuario", error);
        });
        */
    }
    Inicio(){

    }
    Load(){

       
    }
    SetInicio(){
        this.UserEmp.inicio = 1;
        this.SaveUser();
    }
    SetTerminos(){
        this.UserEmp.terminos = 1;
        this.SaveUser();
    }
    SaveUser(){
        localStorage.setItem("local_user_01", JSON.stringify(this.UserEmp));
    }
    SaveIdioma(){
        localStorage.setItem("local_lang_01", JSON.stringify(this.Idioma));
    }
    SaveCat(){
        localStorage.setItem("local_cats_01", JSON.stringify(this.Categorias));
    }
    
    Start(){

        if (this.UserEmp.pais === undefined || this.UserEmp.idioma === undefined) {
            return {
                c:"ch", 
                d:[
                    {
                        c:"hh",
                        d: super.TituloSimple("HOLA", "NELSON")
                    },
                    {
                        c:"hc"
                    }
                ]
            };
        }
        if (this.UserEmp.inicio === undefined) {
            return {
                c:"ch", 
                d:[{
                    c:"hh",
                    d: this.Titulo(idioma[0], idioma[1])
                },{
                    c:"hc", d:[{c:"hc0", h:"INICIO", f:{m:"click",n:Empezar}}]
                }]
            };
        }
        if (this.UserEmp.terminos === undefined) {
            return {
                c:"ch", 
                d:[{
                    c:"hh",
                    d: this.Titulo(idioma[0], idioma[1])
                },{
                    c:"hc", d:[{c:"hc0", h:"TERMINOS Y CONDICIONES", f:{m:"click",n:Terminos}}]
                }]
            };
        }
    }
}


function loadJS(url) {
    return new Promise((resolve, reject) => {
        var script = cE("script");
        script.type = "text/javascript";
        script.src = url;
        script.onload = () => resolve(true);
        script.onerror = () => reject(false);
        document.body.appendChild(script);
    });
}

function cE(t){
    return document.createElement(t);
}
function render(t){
    const d = cE(divTipo(t.t));
    if (t.c !== undefined) {
        d.className = t.c;
    }
    if (t.i !== undefined) {
        d.id = t.i;
    }
    if (t.f !== undefined) {
        if (typeof t.f === 'object'){
            d.addEventListener(t.f.m, t.f.n);
        }
        if (Array.isArray(t.f)){
            for(var x of t.f){
                d.addEventListener(x.m, x.n);
            }
        }
    }
    if (t.y !== undefined) {
        d.type = inputType(t.y);
    }
    if (t.v !== undefined) {
        d.value = t.v;
    }
    if (t.h !== undefined) {
        d.textContent = t.h;
    }
    if (t.s !== undefined && t.s == true) {
        d.selected = true;
    }
    
    if (t.d !== undefined) {
        if (Array.isArray(t.d)){
            t.d.forEach((e) => {
                d.appendChild(render(e)) 
            });
        }else{
            d.appendChild(render(t.d)) 
        }
    }
    return d;
}

// TYPES
function divTipo(n){
    switch(n) {
        case 1:
            return "video";
        case 2:
            return "span";
        case 3:
            return "input";
        case 4:
            return "ul";
        case 5:
            return "li";
        case 6:
            return "h1";
        case 7:
            return "a";
        case 8:
            return "h2";
        case 9:
            return "h3";
        case 10:
            return "label";
        case 11:
            return "img";
        case 12:
            return "select";
        case 13:
            return "option";
        case 14:
            return "fieldset";
        case 15:
            return "legend";
        default:
            return "div";
    }
}
function inputType(n){
    switch(n) {
        case 1:
            return "checkbox";
        case 2:
            return "email";
        case 3:
            return "hidden";
        case 4:
            return "file";
        case 5:
            return "range";
        default:
            return "text";
    }
}
function encodeBase64Custom(num, largo) {
    let out = "";
    do {
        out = chars[num % 64] + out;
        num = Math.floor(num / 64);
    } while (num > 0);

    while (out.length < largo) {
        out = "a" + out;
    }
    return out;
}

function decodeBase64Custom(str) {
    let num = 0;
    for (let c of str) {
        num = num * 64 + chars.indexOf(c);
    }
    return num;
}