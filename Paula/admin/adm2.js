function LocalStorage(n){
    var res = JSON.parse(window.localStorage.getItem(n));
    if (res === null ){
        if (n == "local_cats_01"){
            res = { Idmax: 0, Lista: [] };
        }
        if (n == "local_user_01"){
            res = {};
        }
        if (n == "local_lang_01"){
            res = {};
        }
    }
    return res;
}
function SaveLocalStorage(n, v){
    localStorage.setItem(n, JSON.stringify(v));
}
function GC(c, n = 0){
    return document.getElementsByClassName(c)[n];
}
function GI(i){
    return document.getElementById(i);
}
function POP(){
    GC("p").classList.toggle("po");
}
function TM(){
    if (GC("c").classList.contains('menu_largo')){
        HideMenu();
    }else {
        ShowMenu();
    }
}
function ShowMenu(){
    var x = GC("c").classList;
    x.add('menu_largo');
    x.remove('menu_corto');
}
function HideMenu(){
    var x = GC("c").classList;
    x.add('menu_corto');
    x.remove('menu_largo');
}
function HE(){
    GC("c").classList.toggle('ce');
}
function ajustarClaseSegunAncho(){
    if (window.innerWidth < 600) {
        HideMenu();
    } else {
        ShowMenu();
    }
}

function SendRequestBytes(url) {
    return fetch(url).then(response => {
        if (!response.ok) {
            throw new Error('Error en la respuesta: ' + response.status);
        }
        return response.arrayBuffer();
    }).then(buffer => {
        const bytes = new Uint8Array(buffer);
        return bytes; // Retornar los bytes para el .then externo
    });
}
function SendRequestJson(url) {
    return fetch(url).then(response => {
        if (!response.ok) {
            throw new Error('Error en la respuesta: ' + response.status);
        }
        return response.json(); // Convertir la respuesta a JSON
    });
}
function Parents(e, c) {
    while ((e = e.parentElement) && e.className != c);
    return e;
}
function SplitParam(x){
    return { id: parseInt(x.split("-")[0]), pid: parseInt(x.split("-")[1]) }
}
function DecodeUserEmp(b){
    var res = {};

    let i = 0;

    for (;;) {

        res.Terminos = b[0]/2;
        res.Admin = b[0]%2;
        res.Idioma = b[1]*256 + b[2];
        res.Pais = b[3];
        
        i = i + 3

        if (i >= b.length) {
            break; // salir del bucle cuando i sea mayor que 5
        }
        i++;
    }

    return res;
}
window.addEventListener('resize', ajustarClaseSegunAncho);