package main

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"io/ioutil"
)
///Users/usuario/go/src/github.com/gorilla/mux/
func HomeHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "./static/index.html")
}
func getAllPosts(w http.ResponseWriter, r *http.Request){
	resp,errReques := http.Get("https://jsonplaceholder.typicode.com/posts")
	if errReques != nil {
		w.WriteHeader(405)
		w.Write([]byte("error API"))
		return 
	}
	responseData, errParse := ioutil.ReadAll(resp.Body)
	if errParse != nil {
		w.WriteHeader(405)
		w.Write([]byte("error parsing"))
		return 
	}
	responseString := string(responseData)
	w.WriteHeader(200)
	w.Write([]byte(responseString))
}
func getPost(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	postId := vars["postId"]
	url := "https://jsonplaceholder.typicode.com/posts/"+postId
	resp,errReques := http.Get(url)
	if errReques != nil {
		w.WriteHeader(405)
		w.Write([]byte("error API"))
		return 
	}
	responseData, errParse := ioutil.ReadAll(resp.Body)
	if errParse != nil {
		w.WriteHeader(405)
		w.Write([]byte("error parsing"))
		return 
	}
	responseString := string(responseData)
    w.WriteHeader(200)
	w.Write([]byte(responseString))
}
func TestHandler(w http.ResponseWriter, r *http.Request){
	// response, _ := json.Marshal(payload)
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(200)
    // w.Write([]byte("blabla\n"))
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["testArg"])

}
func main(){
	fmt.Println("start server")
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    r.HandleFunc("/", HomeHandler).Methods("GET")
    r.HandleFunc("/get-all-posts", getAllPosts).Methods("GET")
    r.HandleFunc("/get-post/{postId}", getPost).Methods("GET")
    http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", nil)
}
