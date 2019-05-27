package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello DevOps v2")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}
