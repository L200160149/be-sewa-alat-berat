package controller

import (
	"net/http"

	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
	"github.com/L200160149/be-sewa-alat-berat/service"
	"github.com/julienschmidt/httprouter"
)


type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	authRequest := web.AuthRequest{}
	helper.ReadFromRequestBody(request, &authRequest)

	authResponse := controller.AuthService.Login(request.Context(), authRequest)

	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   authResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
