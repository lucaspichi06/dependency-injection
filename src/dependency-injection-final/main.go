package main

func main() {
	sendService := NewSendService()
	pdfService := NewPdfService()
	docService := NewDocService()
	printService := NewPrintService(sendService, pdfService)
	printService.imprimir()
	printService = NewPrintService(sendService, docService)
	printService.imprimir()
}
