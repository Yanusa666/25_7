package main

import (
	"fmt"
	"log"
)

func main() {
	var n float32
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&n)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Printf("Ваше число: %f\n", n)
}