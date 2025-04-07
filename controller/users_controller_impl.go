package controller

import (
	"net/http"
	"strconv"

	"github.com/L200160149/be-sewa-alat-berat/helper"
	"github.com/L200160149/be-sewa-alat-berat/model/web"
	"github.com/L200160149/be-sewa-alat-berat/service"
	"github.com/julienschmidt/httprouter"
)

type UsersControllerImpl struct {
	UsersService service.UsersService
}

func NewUsersController(usersService service.UsersService) UsersController {
	return &UsersControllerImpl{
		UsersService: usersService,
	}
}

func (controller *UsersControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersResponses := controller.UsersService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersCreateRequest := web.UsersCreateRequest{}
	helper.ReadFromRequestBody(request, &usersCreateRequest)

	usersResponse := controller.UsersService.Create(request.Context(), usersCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "OK",
		Data: usersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	usersUpdateRequest := web.UsersUpdateRequest{}
	helper.ReadFromRequestBody(request, &usersUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	usersUpdateRequest.Id = id

	usersResponse := controller.UsersService.Update(request.Context(), usersUpdateRequest)

	webResponse := web.WebResponse{
		Code: 201,
		Status: "OK",
		Data: usersResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UsersControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UsersService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
