package main

import "fmt"



func main(){
	s := make([]string, 0)
	fmt.Println("s:", s)
	s = append(s, "a")
	fmt.Println("s:", s)
	s = append(s, "b")
	fmt.Println("s:", s)

}