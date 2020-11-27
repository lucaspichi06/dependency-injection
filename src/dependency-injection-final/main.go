package main

import "github.com/lucaspichi06/dependency-injection/src/dependency-injection-final/service"

func main() {
	sendService := service.NewSendService()
	pdfService := service.NewPdfService()
	docService := service.NewDocService()
	printService := service.NewPrintService(sendService, pdfService)
	printService.Print()
	printService = service.NewPrintService(sendService, docService)
	printService.Print()
}
