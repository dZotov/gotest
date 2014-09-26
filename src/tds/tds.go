package main

import (
//    "encoding/json"
    "log"
    "net/http"
    "fmt"
    "gopkg.in/mgo.v2"
//    "expvar"
)



func handler(w http.ResponseWriter, req *http.Request){

	fmt.Fprintf(w, "!!!!\n")

	//io.WriteStrning(writer,"!!!!")

}


func main(){

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":2000", nil); err != nil {
    	log.Fatal("ListenAndServe: ", err)
  	}

}






