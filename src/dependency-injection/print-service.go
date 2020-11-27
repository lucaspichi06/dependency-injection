package main

import "fmt"

type printService struct {
	send *sendService
	pdf  *pdfService
}

func NewPrintService() *printService {
	return &printService{
		send: &sendService{},
		pdf:  &pdfService{},
	}
}

func (p *printService) imprimir() {
	p.send.send()
	p.pdf.pdf()
}

type pdfService struct{}

func (p *pdfService) pdf() {
	fmt.Println("printing the document in pdf format")
}

type sendService struct{}

func (s *sendService) send() {
	fmt.Println("sending the document to printer")
}
