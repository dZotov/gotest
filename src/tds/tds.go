package main

import (
//    "encoding/json"
    "log"
    "net/http"
   "./tour"
    "io"
//    "expvar"
)



func handler(w http.ResponseWriter, req *http.Request){

	io.WriteString(w,"!!!!")

}


func main(){

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
    	log.Fatal("ListenAndServe: ", err)
  	}

}


// import (
// 	"io"
// 	"net/http"

// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "Hello world!")

// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	http.ListenAndServe(":3000", mux)
// }
