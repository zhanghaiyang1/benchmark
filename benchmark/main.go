package main

import (
	"fmt"
	"time"
)



func main(){
	mill := time.Now().UnixNano() / 1e6
	fmt.Println("mill:", mill)
}