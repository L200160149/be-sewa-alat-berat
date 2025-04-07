package app

import (
	"github.com/L200160149/be-sewa-alat-berat/controller"
	"github.com/L200160149/be-sewa-alat-berat/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(usersController controller.UsersController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/v1/users", usersController.FindAll)
	router.POST("/api/v1/users", usersController.Create)
	router.PUT("/api/v1/users/:userId", usersController.Update)
	router.DELETE("/api/v1/users/:userId", usersController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}