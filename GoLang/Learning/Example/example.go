package main

import (
	"fmt"
	"net/http"
	"runtime"
)


func main(){
	fmt.Println(" runtime :: ", runtime.NumCPU())
	res, err := http.Get("https://www.google.com")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf(" response status ::%v  ",res.StatusCode)
}