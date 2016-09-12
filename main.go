package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

var count int

func echoString(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.path))
}

func counter(w http.ResponseWriter, r *http.Request){
    mu.Lock()
    fmt.Fprintf(w, "Count %d\n", count)
    mu.Unlock()
}

func main() {
    http.HandleFunc("/", echoString)

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
