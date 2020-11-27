package main

import "fmt"

type printService struct{}

func (p *printService) imprimir() {
	fmt.Println("sending the document to printer")
	fmt.Println("printing the document in pdf format")
}
