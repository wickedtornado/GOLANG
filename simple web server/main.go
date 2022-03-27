package main

import (
	"fmt"
	"log"
	"net/http"
)
func hellohandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello"{
		http.Error(w, "404",http.StatusNotFound)
		return
	}
if r.Method!="GET"{
	http.Error(w, "method not supported",http.StatusNotFound )
	return
}
fmt.Fprintln(w, "hello")
}

func formhandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!=nil{
		fmt.Fprintf(w, "parserform() err : %v",err)
		return
	}
	fmt.Fprintf(w, "Post request success")
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w, "Name = %s", name)
	fmt.Fprintf(w, "Address = %s", address)
} 
 
func main(){
	 fileserver := http.FileServer(http.Dir("./static"))
	 http.Handle("/", fileserver)
	 http.HandleFunc("/form",formhandler)
	 http.HandleFunc("/hello",hellohandler)

	 fmt.Println("Starting server at port 8080")
	 if err := http.ListenAndServe(":8080",nil);err!=nil{
		 log.Fatal(err)
	 }
	 
}
