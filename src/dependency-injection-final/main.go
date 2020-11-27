package main

import (
	"github.com/lucaspichi06/dependency-injection/src/dependency-injection-final/service"
)

type printService interface {
	HandlePrint()
}

type services struct {
	printer printService
}

func main() {
	services := loadDependencies()
	services.printer.HandlePrint()
}

func loadDependencies() *services {
	sendService := service.NewSendService()
	pdfService := service.NewPdfService()
	//docService := service.NewDocService()
	printService := service.NewPrintService(sendService, pdfService)
	//printService := service.NewPrintService(sendService, docService)

	return &services{
		printer: printService,
	}
}
