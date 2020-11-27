package main

import "github.com/lucaspichi06/dependency-injection/src/initial/service"

func main() {
	printService := &service.PrintService{}
	printService.Print()
}
