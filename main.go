package main

import (
   "fmt"
   "log"
   "net/http"
   "github.com/gorilla/mux"
   "encoding/json"
)

type fruta struct {
    nome string
    tipo string
    valor float32
}

type Frutas []fruta

// primeiro eu declaro uma variavel importando funções newrouter e uma sub-função HandleFunc
// isso serve como ponte entre o Client e o Server onde pode ser feito requests
// depois aqui ocorre uma especie de interpolação entre a extremidade  / com alguma da função
// por ultimo eu coloco um "Listener" para essas requisições serem ouvidas atraves da porta 8080
func main() {
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/", Index)
    r.HandleFunc("/abacate", AbacateIndex)
    r.HandleFunc("/xuxu", XuxuIndex)
    r.HandleFunc("/banana/{frutaId}", BananaIndex)
    log.Fatal(http.ListenAndServe(":8080", r))

}

// função Index onde o lado "Server" responde com alguma coisa
func Index(w http.ResponseWriter, r *http.Request) {
    frutas := fruta{
         nome: "banana",
         tipo: "banana da terra",
         valor: 12.0,
    }

    ret, err := json.Marshal(frutas)

    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(ret))
    w.Write(ret)
}


func AbacateIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome Abacate!")
}

func XuxuIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome Xuxu!")
}

// aqui coloquei uma variavel frutaId, com isso eu passo qualquer valor para ela e la em 
// cima na rota ele vai entender o meu parametro e vai alocar dinamicamente minha varuivel. 
func BananaIndex(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    frutaId := vars["frutaId"]
    fmt.Fprintln(w, "Welcome: ", frutaId)
}
