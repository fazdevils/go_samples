/**
  Simple REST implementation using the 'gorest' Go library.
*/
package main
import (
    "code.google.com/p/gorest"
    "net/http"
    "time"
    "encoding/json"
    "fmt"
)
func main() {
    gorest.RegisterService(new(HelloService))
    
    http.Handle("/",gorest.Handle())    
    http.ListenAndServe(":8786",nil)
}

//Service Definition
type HelloService struct {
    gorest.RestService `root:"/tutorial/"`
    helloWorld     gorest.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
    sayHello       gorest.EndPoint `method:"GET" path:"/hello/{name:string}" output:"string"`
    printStruct    gorest.EndPoint `method:"GET" path:"/hello/struct/{name:string}" output:"SimpleStruct"`
}

func(serv HelloService) HelloWorld() string{
    return "Hello World"
}

func(serv HelloService) SayHello(name string) string{
    return "Hello " + name
}

func(serv HelloService) PrintStruct(name string) SimpleStruct  {
	newStruct := newSimpleStruct(name)
	b,e := json.Marshal(newStruct)
	if e != nil {
		fmt.Println("error:", e)
	}	
	fmt.Println(newStruct)
	fmt.Printf("%s\n", b)
    return *newStruct
}

type SimpleStruct struct {
	Name string
	Now time.Time
} 

func newSimpleStruct(name string) *SimpleStruct {
     return &SimpleStruct{name, time.Now()}
}