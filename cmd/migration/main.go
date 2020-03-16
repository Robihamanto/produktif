package main

import (
	"github.com/Robihamanto/produktif/cmd/api/config"
	"github.com/Robihamanto/produktif/cmd/migration/cmd"
	"github.com/Robihamanto/produktif/cmd/migration/schema"
	"github.com/Robihamanto/produktif/internal/platform/mysql"
)

func main() {
	cfg, err := config.Load("staging")
	checkErr(err)

	db, err := mysql.New(cfg.DB.PSN)
	checkErr(err)

	s := schema.New(db)
	cmd.Execute(s)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
