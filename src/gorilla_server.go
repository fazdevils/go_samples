/**
  Simple REST implementation using the 'gorilla' Go library.
*/
package main

import (
    "encoding/json"
    "encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
    "net/http"
    "time"    
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{name}.{format}", HelloWorld)
	http.Handle("/", router)
	http.ListenAndServe(":8787",nil)
}

func HelloWorld(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	newStruct := NewSimpleStruct(vars["name"])
	
	var b []byte
	var e error
	
	switch vars["format"] {
		case "xml":
			b,e = xml.MarshalIndent(newStruct, "", "   ")
		default:
			b,e = json.Marshal(newStruct)	
	}
	if e != nil {
		fmt.Println("error:", e)
	}
	response.Write(b)	
}


// http://golang.org/pkg/encoding/json/
// http://golang.org/pkg/encoding/xml/

type SimpleStruct struct {
	Name string `json:"name" xml:"Name>Full"`
	CreatedAt time.Time `json:"created_at"`
} 

func NewSimpleStruct(name string) *SimpleStruct {
	 s := &SimpleStruct{Name:name, CreatedAt:time.Now()}
     return s
}

