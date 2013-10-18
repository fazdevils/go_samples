/**
  Sample "Hello World" type go application.
*/
package main 

import (
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Admin %s!</h1>", r.URL.Path[1:])
}


func main() {
    http.HandleFunc("/vinnie", defaultHandler)
    //http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/admin", adminHandler)
    http.ListenAndServe(":8080", nil)
}

