package main

import (
//    "fmt"
   "log"
   "net/http"
   "math/rand"
   "strconv"
   "github.com/gorilla/mux"
   "encoding/json"
)

type Fruta struct {
    ID    string   `json:"id"`
    Nome  string   `json:"nome"`
    Tipo  string   `json:"tipo"`
    Valor float64  `json:"valor"`
}

var frutas []Fruta

// primeiro eu declaro uma variavel importando funções newrouter e uma sub-função HandleFunc
// isso serve como ponte entre o Client e o Server onde pode ser feito requests
// depois aqui ocorre uma especie de interpolação entre a extremidade  / com alguma da função
// por ultimo eu coloco um "Listener" para essas requisições serem ouvidas atraves da porta 8080
func main() {

    frutas = append(frutas, Fruta{ID: "1", Nome: "uva", Tipo: "roxa", Valor: 5.8})
    frutas = append(frutas, Fruta{ID: "2", Nome: "maça", Tipo: "pequena", Valor: 3.6})
    frutas = append(frutas, Fruta{ID: "3", Nome: "banana", Tipo: "da terra", Valor: 7.0})
    frutas = append(frutas, Fruta{ID: "4", Nome: "tomate", Tipo: "pequen", Valor: 4.5})

    r := mux.NewRouter()
    r.HandleFunc("/api/fruta/{id}", getFruta).Methods("GET")
    r.HandleFunc("/api/fruta/", createFruta).Methods("POST")
    r.HandleFunc("/api/fruta/{id}", updateFruta).Methods("PUT")
    r.HandleFunc("/api/fruta/{id}", deleteFruta).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// função Index onde o lado "Server" responde com alguma coisa
func getFruta( w http.ResponseWriter, r *http.Request ) { 
  w.Header().Set("Content-Type", "application/json") 
  params := mux.Vars(r)
  
  for _, item := range(frutas){
      if item.ID == params["id"]{
        json.NewEncoder(w).Encode(item)
        return    
      }
  }
  json.NewEncoder(w).Encode(&Fruta{})
}

func createFruta( w http.ResponseWriter, r *http.Request ) {    
    w.Header().Set("Content-Type", "application/json")
	var fruta Fruta
	_ = json.NewDecoder(r.Body).Decode(&fruta)
	fruta.ID = strconv.Itoa(rand.Intn(1000)) // Mock ID - not safe
	frutas = append(frutas, fruta)
	json.NewEncoder(w).Encode(fruta)
}

func updateFruta( w http.ResponseWriter, r *http.Request ) {    
    params := mux.Vars(r)
	for index, item := range frutas {
		if item.ID == params["id"] {
			frutas = append(frutas[:index], frutas[index+1:]...)
			var fruta Fruta
			_ = json.NewDecoder(r.Body).Decode(&fruta)
			fruta.ID = params["id"]
			frutas = append(frutas, fruta)
			json.NewEncoder(w).Encode(fruta)
			return
		}
	}
}

func deleteFruta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range frutas {
		if item.ID == params["id"] {
			frutas = append(frutas[:index], frutas[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(frutas)
}
