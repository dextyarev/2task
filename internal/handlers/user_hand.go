package handlers

import (
	"2task/internal/api"
	"2task/internal/repositories"
	"2task/internal/services"
	"context"
	"fmt"
	"gorm.io/gorm"
)

func (h *Handler) DeleteUsers(_ context.Context, request api.DeleteUsersRequestObject) (api.DeleteUsersResponseObject, error) {
	userId := request.Body.Id
	err := h.UserService.DeleteUserByID(int(*userId))
	var res api.DeleteUsers400TextResponse

	if err != nil {
		res = "err"
		fmt.Println("[DB] Ошибка обновления записи ", err)
		return res, err
	}

	res = "deleted"
	return res, nil
}

func (h *Handler) GetUsers(_ context.Context, request api.GetUsersRequestObject) (api.GetUsersResponseObject, error) {
	allUsers, err := h.UserService.GetAllUsers()
	if err != nil {
		return nil, err
	}

	res := api.GetUsers200JSONResponse{}

	for _, user := range allUsers {
		User := api.User{
			Email:    &user.Email,
			Id:       &user.ID,
			Password: &user.Password,
		}

		res = append(res, User)
	}
	return res, nil
}

func (h *Handler) PatchUsers(_ context.Context, request api.PatchUsersRequestObject) (api.PatchUsersResponseObject, error) {
	var resErr api.PatchUsers400TextResponse = "err"
	user, asUserErr := request.Body.AsUser()

	if asUserErr != nil {
		return resErr, fmt.Errorf("invalid input data")
	}

	User := repositories.User{Model: gorm.Model{ID: uint(*user.Id)}}

	if user.Email != nil {
		User.Email = *user.Email
	}

	if user.Password != nil {
		User.Password = *user.Password
	}

	updated, err := h.UserService.UpdateUserByID(int(*user.Id), User)
	if err != nil {
		return resErr, err
	}

	return api.PatchUsers200JSONResponse{
		Id:       user.Id,
		Email:    &updated.Email,
		Password: &updated.Password,
	}, nil
}

func (h *Handler) PostUsers(_ context.Context, request api.PostUsersRequestObject) (api.PostUsersResponseObject, error) {
	user := services.User{Email: *request.Body.Email, Password: *request.Body.Password}
	var resErr api.PostUsers400TextResponse

	created, err := h.UserService.CreateUser(user)

	if err != nil {
		resErr = "err"
		fmt.Println("[DB] Ошибка обновления записи ", err)
		return resErr, err
	}

	var res = api.PostUsers200JSONResponse{Email: &created.Email, Id: &created.ID, Password: &created.Password}

	return res, nil
}
