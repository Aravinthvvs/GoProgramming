package main

import (
	"fmt"
	"io"
	"net/http"
	"math/rand"
	"time"
)

func main(){
	
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(100))
	cli := http.Client{Timeout: time.Duration(1)* time.Second}

	req, err := http.NewRequest("GET","https://api.github.com",nil)
	if err != nil{
		fmt.Println("err while creating request", err)
	}
	req.Header.Add("accept",`application/json`)
	res, err := cli.Do(req)

	if err != nil{
		fmt.Println("err while receiving responses", err)
	}
	defer res.Body.Close()
	//fmt.Println(res.Header)
	readBytes,err := io.ReadAll(res.Body)
	if err != nil{
		fmt.Println("err  while reading responses", err)
	}
	fmt.Printf(" read status %v and content %s",res.Status, string(readBytes))
}