package main

import "github.com/lucaspichi06/dependency-injection/src/dependency-injection/service"

func main() {
	printService := service.NewPrintService()
	printService.Print()
}
