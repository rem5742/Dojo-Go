package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)
///Users/usuario/go/src/github.com/gorilla/mux/
func HomeHandler(w http.ResponseWriter, r *http.Request){
	// response, _ := json.Marshal(payload)
	// w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write([]byte("blabla\n"))
}
func TestHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["testArg"])

	// w.WriteHeader(200)
    // w.Write([]byte("blabla\n"))
}
func main(){
	fmt.Println("start server")
	r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/test/{testArg}", TestHandler).Methods("GET")
    // r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", r)
	http.ListenAndServe("localhost:4000", nil)
}
