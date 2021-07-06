package main

import (
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })


    fmt.Printf("Starting server at port 80\n")
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal(err)
    }
}
