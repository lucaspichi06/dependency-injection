package service

import "fmt"

type PrintService struct{}

func (p *PrintService) Print() {
	fmt.Println("sending the document to printer")
	fmt.Println("printing the document in pdf format")
}
