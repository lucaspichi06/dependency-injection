package main

import "fmt"

type printService struct{}

func (p *printService) imprimir() {
	fmt.Println("enviando el documento a imprimir")
	fmt.Println("imprimiendo el documento en formato pdf")
}
