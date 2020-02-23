package main

import (
	"fmt"
	"server"
	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println("Hello Produktif")
	fmt.Println(randomdata.SillyName())

	server.Show()
	e := server.New()
	e.Start(":1818")
}
