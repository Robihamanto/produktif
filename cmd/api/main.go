package main

import (
	"github.com/Robihamanto/produktif/cmd/api/server"
	//"github.com/Robihamanto/produktif/internal"
	"github.com/jinzhu/gorm"

)

func main() {
	e := server.New()
	server.Start(e)
}


func addService(db *gorm.DB) {
	
}



func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
