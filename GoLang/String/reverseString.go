package main

import "fmt"


func main(){

	fmt.Println(ReverseString("I am  back"))
}

func ReverseString(s string)string{
	byteString := []byte(s)
	j:= (len(s) -1)
	for i := 0; i < len(byteString) /2; i++{
		byteString[i], byteString[j] = byteString[j], byteString[i] 
		j--
	}
	return string(byteString)
}