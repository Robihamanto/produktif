package main

import (
	"github.com/Robihamanto/produktif/cmd/migration/cmd"
	"github.com/Robihamanto/produktif/cmd/migration/schema"
	"github.com/Robihamanto/produktif/internal/platform/mysql"
)

func main() {
	db, err := mysql.New()
	checkErr(err)

	s := schema.New(db)
	cmd.Execute(s)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
