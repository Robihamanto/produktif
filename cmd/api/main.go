package main

import (
	"github.com/Robihamanto/produktif/cmd/api/config"
	"github.com/Robihamanto/produktif/cmd/api/jwt"
	"github.com/Robihamanto/produktif/cmd/api/rbac"
	"github.com/Robihamanto/produktif/cmd/api/server"
	"github.com/Robihamanto/produktif/cmd/api/service"
	"github.com/Robihamanto/produktif/internal/auth"
	"github.com/Robihamanto/produktif/internal/platform/mysql"
	"github.com/Robihamanto/produktif/internal/todolist"
	"github.com/Robihamanto/produktif/internal/user"
	"github.com/casbin/casbin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	_ "github.com/spf13/cobra"
)

func main() {
	config, err := config.Load("staging")
	checkErr(err)

	e := server.New()

	db, err := mysql.New(config.DB.PSN)
	checkErr(err)

	addService(config, db, e)

	server.Start(e)
}

func addService(config *config.Configuration, db *gorm.DB, e *echo.Echo) {
	casbinService := casbin.NewEnforcer(config.Casbin.Model, config.Casbin.Policy)
	rbacService := rbac.New(casbinService)

	userDB := mysql.NewUserDB(db)
	todolistDB := mysql.NewTodolistDB(db)

	jwtService := jwt.New(config.JWT, rbacService)
	jwtMiddleware := jwtService.MWFunc()

	authSvc := auth.New(userDB, jwtService)

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
		jwtMiddleware,
	)

	todolistRouter := e.Group("/todolist")
	service.NewTodolist(
		todolistSvc,
		todolistRouter,
		jwtMiddleware,
	)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
