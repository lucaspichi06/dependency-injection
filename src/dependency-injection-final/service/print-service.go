package service

import "fmt"

type printService struct {
	send   *sendService
	format formatService
}

func NewPrintService(send *sendService, format formatService) *printService {
	return &printService{
		send:   send,
		format: format,
	}
}

func (p *printService) Print() {
	p.send.send()
	p.format.print()
}

type sendService struct{}

func NewSendService() *sendService {
	return &sendService{}
}

func (s *sendService) send() {
	fmt.Println("sending the document to printer")
}

type formatService interface {
	print()
}

type pdfService struct{}

func NewPdfService() *pdfService {
	return &pdfService{}
}

func (p *pdfService) print() {
	fmt.Println("printing the document in pdf format")
}

type docService struct{}

func NewDocService() *docService {
	return &docService{}
}

func (d *docService) print() {
	fmt.Println("printing the document in doc format")
}
