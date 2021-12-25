package main

import "fmt"
import "./other"

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	other.Other()
}
