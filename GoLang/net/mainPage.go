package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	//mux := http.NewServeMux()
	cli := http.Client{Timeout: time.Duration(0)*time.Microsecond}

	resp , err := cli.Get("https://www.google.com")

	if err!=nil{
		fmt.Println(err)
	}
	
	fmt.Println(resp)
	//http.HandleFunc("/",HelloServer)
	//http.ListenAndServe(":8080",mux)
}

func HelloServer(w http.ResponseWriter, r* http.Request){
	log.Print(r.URL.Path)
	fmt.Fprintf(w," Hello, %v",r.URL.Path[1:])
}