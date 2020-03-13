package main

import (
	"github.com/Robihamanto/produktif/cmd/api/server"
	"github.com/Robihamanto/produktif/cmd/api/service"
	"github.com/Robihamanto/produktif/internal/auth"
	"github.com/Robihamanto/produktif/internal/platform/mysql"
	"github.com/Robihamanto/produktif/internal/todolist"
	"github.com/Robihamanto/produktif/internal/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/spf13/cobra"
)

func main() {
	e := server.New()

	db, err := mysql.New()
	checkErr(err)

	addService(db, e)

	server.Start(e)
}

func addService(db *gorm.DB, e *echo.Echo) {
	userDB := mysql.NewUserDB(db)
	todolistDB := mysql.NewTodolistDB(db)

	authSvc := auth.New(userDB)

	userSvc := user.New(userDB)

	todolistSvc := todolist.New(
		todolistDB,
		userDB,
	)

	//Bind app service to http service
	service.NewAuth(authSvc, e)

	userRouter := e.Group("/users")
	service.NewUser(
		userSvc,
		userRouter,
	)

	todolistRouter := e.Group("/todolist")
	//service.NewTodolist(todolistSvc, todolistRouter, )

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
