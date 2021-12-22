package main

import ( 
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
    fmt.Fprint(w, "Go Handler")
}

func main(){
//    fmt.Println("C Go run")
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    http.Handle("/", r)
    http.ListenAndServe(":8080",r)
}
