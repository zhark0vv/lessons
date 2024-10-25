package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	original := Person{Name: "Alice", Age: 30}

	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(original)
	if err != nil {
		log.Printf("encode error: %v", err)
		return
	}

	var decoded Person
	decoder := gob.NewDecoder(&buffer)
	err = decoder.Decode(&decoded)
	if err != nil {
		log.Printf("decode error: %v", err)
		return
	}

	fmt.Println("Original:", original)
	fmt.Println("Decoded:", decoded)
}
