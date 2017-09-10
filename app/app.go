// package main

// import (
// 	"fmt"
// 	"github.com/julienschmidt/httprouter"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprintf(w, "This is the indeeeex")
// }

// func main() {
// 	router := httprouter.New()
// 	router.GET("/", indexHandler)

// 	// print env
// 	env := os.Getenv("APP_ENV")
// 	if env == "production" {
// 		log.Println("Running api server in production mode")
// 	} else {
// 		log.Println("Running api server in dev mode")
// 	}

// 	http.ListenAndServe(":8080", router)
// }
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
    w.Write([]byte("THIS IS THE INDEX\n"))
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
	http.ListenAndServe("0.0.0.0:8080", nil)
}
