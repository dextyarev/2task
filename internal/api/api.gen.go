// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Message defines model for Message.
type Message struct {
	Id      *uint   `json:"id,omitempty"`
	Message *string `json:"message,omitempty"`
}

// MessageID defines model for MessageID.
type MessageID struct {
	Id *uint `json:"id,omitempty"`
}

// NewMessage defines model for NewMessage.
type NewMessage struct {
	Message *string `json:"message,omitempty"`
}

// NewUser defines model for NewUser.
type NewUser struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UpdateUser defines model for UpdateUser.
type UpdateUser struct {
	union json.RawMessage
}

// User defines model for User.
type User struct {
	Email    *string `json:"email,omitempty"`
	Id       *uint   `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UserEmail defines model for UserEmail.
type UserEmail struct {
	Email *string `json:"email,omitempty"`
	Id    *uint   `json:"id,omitempty"`
}

// UserID defines model for UserID.
type UserID struct {
	Id *uint `json:"id,omitempty"`
}

// UserPassword defines model for UserPassword.
type UserPassword struct {
	Id       *uint   `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}

// DeleteMessagesJSONRequestBody defines body for DeleteMessages for application/json ContentType.
type DeleteMessagesJSONRequestBody = MessageID

// PatchMessagesJSONRequestBody defines body for PatchMessages for application/json ContentType.
type PatchMessagesJSONRequestBody = Message

// PostMessagesJSONRequestBody defines body for PostMessages for application/json ContentType.
type PostMessagesJSONRequestBody = NewMessage

// DeleteUsersJSONRequestBody defines body for DeleteUsers for application/json ContentType.
type DeleteUsersJSONRequestBody = UserID

// PatchUsersJSONRequestBody defines body for PatchUsers for application/json ContentType.
type PatchUsersJSONRequestBody = UpdateUser

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = NewUser

// AsUserEmail returns the union data inside the UpdateUser as a UserEmail
func (t UpdateUser) AsUserEmail() (UserEmail, error) {
	var body UserEmail
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUserEmail overwrites any union data inside the UpdateUser as the provided UserEmail
func (t *UpdateUser) FromUserEmail(v UserEmail) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUserEmail performs a merge with any union data inside the UpdateUser, using the provided UserEmail
func (t *UpdateUser) MergeUserEmail(v UserEmail) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsUserPassword returns the union data inside the UpdateUser as a UserPassword
func (t UpdateUser) AsUserPassword() (UserPassword, error) {
	var body UserPassword
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUserPassword overwrites any union data inside the UpdateUser as the provided UserPassword
func (t *UpdateUser) FromUserPassword(v UserPassword) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUserPassword performs a merge with any union data inside the UpdateUser, using the provided UserPassword
func (t *UpdateUser) MergeUserPassword(v UserPassword) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsUser returns the union data inside the UpdateUser as a User
func (t UpdateUser) AsUser() (User, error) {
	var body User
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUser overwrites any union data inside the UpdateUser as the provided User
func (t *UpdateUser) FromUser(v User) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUser performs a merge with any union data inside the UpdateUser, using the provided User
func (t *UpdateUser) MergeUser(v User) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t UpdateUser) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *UpdateUser) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete message by id
	// (DELETE /messages)
	DeleteMessages(ctx echo.Context) error
	// Get all messages
	// (GET /messages)
	GetMessages(ctx echo.Context) error
	// Update message by id
	// (PATCH /messages)
	PatchMessages(ctx echo.Context) error
	// Create a new message
	// (POST /messages)
	PostMessages(ctx echo.Context) error
	// Delete user by id
	// (DELETE /users)
	DeleteUsers(ctx echo.Context) error
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Update user by id
	// (PATCH /users)
	PatchUsers(ctx echo.Context) error
	// Create a new user
	// (POST /users)
	PostUsers(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteMessages converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteMessages(ctx)
	return err
}

// GetMessages converts echo context to params.
func (w *ServerInterfaceWrapper) GetMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetMessages(ctx)
	return err
}

// PatchMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PatchMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchMessages(ctx)
	return err
}

// PostMessages converts echo context to params.
func (w *ServerInterfaceWrapper) PostMessages(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostMessages(ctx)
	return err
}

// DeleteUsers converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsers(ctx)
	return err
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PatchUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/messages", wrapper.DeleteMessages)
	router.GET(baseURL+"/messages", wrapper.GetMessages)
	router.PATCH(baseURL+"/messages", wrapper.PatchMessages)
	router.POST(baseURL+"/messages", wrapper.PostMessages)
	router.DELETE(baseURL+"/users", wrapper.DeleteUsers)
	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.PATCH(baseURL+"/users", wrapper.PatchUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)

}

type DeleteMessagesRequestObject struct {
	Body *DeleteMessagesJSONRequestBody
}

type DeleteMessagesResponseObject interface {
	VisitDeleteMessagesResponse(w http.ResponseWriter) error
}

type DeleteMessages200TextResponse string

func (response DeleteMessages200TextResponse) VisitDeleteMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type DeleteMessages400TextResponse string

func (response DeleteMessages400TextResponse) VisitDeleteMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)

	_, err := w.Write([]byte(response))
	return err
}

type GetMessagesRequestObject struct {
}

type GetMessagesResponseObject interface {
	VisitGetMessagesResponse(w http.ResponseWriter) error
}

type GetMessages200JSONResponse []Message

func (response GetMessages200JSONResponse) VisitGetMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchMessagesRequestObject struct {
	Body *PatchMessagesJSONRequestBody
}

type PatchMessagesResponseObject interface {
	VisitPatchMessagesResponse(w http.ResponseWriter) error
}

type PatchMessages200JSONResponse Message

func (response PatchMessages200JSONResponse) VisitPatchMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostMessagesRequestObject struct {
	Body *PostMessagesJSONRequestBody
}

type PostMessagesResponseObject interface {
	VisitPostMessagesResponse(w http.ResponseWriter) error
}

type PostMessages200JSONResponse Message

func (response PostMessages200JSONResponse) VisitPostMessagesResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersRequestObject struct {
	Body *DeleteUsersJSONRequestBody
}

type DeleteUsersResponseObject interface {
	VisitDeleteUsersResponse(w http.ResponseWriter) error
}

type DeleteUsers200TextResponse string

func (response DeleteUsers200TextResponse) VisitDeleteUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)

	_, err := w.Write([]byte(response))
	return err
}

type DeleteUsers400TextResponse string

func (response DeleteUsers400TextResponse) VisitDeleteUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)

	_, err := w.Write([]byte(response))
	return err
}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchUsersRequestObject struct {
	Body *PatchUsersJSONRequestBody
}

type PatchUsersResponseObject interface {
	VisitPatchUsersResponse(w http.ResponseWriter) error
}

type PatchUsers200JSONResponse User

func (response PatchUsers200JSONResponse) VisitPatchUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PatchUsers400TextResponse string

func (response PatchUsers400TextResponse) VisitPatchUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)

	_, err := w.Write([]byte(response))
	return err
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers200JSONResponse User

func (response PostUsers200JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsers400TextResponse string

func (response PostUsers400TextResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(400)

	_, err := w.Write([]byte(response))
	return err
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Delete message by id
	// (DELETE /messages)
	DeleteMessages(ctx context.Context, request DeleteMessagesRequestObject) (DeleteMessagesResponseObject, error)
	// Get all messages
	// (GET /messages)
	GetMessages(ctx context.Context, request GetMessagesRequestObject) (GetMessagesResponseObject, error)
	// Update message by id
	// (PATCH /messages)
	PatchMessages(ctx context.Context, request PatchMessagesRequestObject) (PatchMessagesResponseObject, error)
	// Create a new message
	// (POST /messages)
	PostMessages(ctx context.Context, request PostMessagesRequestObject) (PostMessagesResponseObject, error)
	// Delete user by id
	// (DELETE /users)
	DeleteUsers(ctx context.Context, request DeleteUsersRequestObject) (DeleteUsersResponseObject, error)
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Update user by id
	// (PATCH /users)
	PatchUsers(ctx context.Context, request PatchUsersRequestObject) (PatchUsersResponseObject, error)
	// Create a new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// DeleteMessages operation middleware
func (sh *strictHandler) DeleteMessages(ctx echo.Context) error {
	var request DeleteMessagesRequestObject

	var body DeleteMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteMessages(ctx.Request().Context(), request.(DeleteMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteMessagesResponseObject); ok {
		return validResponse.VisitDeleteMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetMessages operation middleware
func (sh *strictHandler) GetMessages(ctx echo.Context) error {
	var request GetMessagesRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetMessages(ctx.Request().Context(), request.(GetMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetMessagesResponseObject); ok {
		return validResponse.VisitGetMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchMessages operation middleware
func (sh *strictHandler) PatchMessages(ctx echo.Context) error {
	var request PatchMessagesRequestObject

	var body PatchMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchMessages(ctx.Request().Context(), request.(PatchMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchMessagesResponseObject); ok {
		return validResponse.VisitPatchMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostMessages operation middleware
func (sh *strictHandler) PostMessages(ctx echo.Context) error {
	var request PostMessagesRequestObject

	var body PostMessagesJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostMessages(ctx.Request().Context(), request.(PostMessagesRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostMessages")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostMessagesResponseObject); ok {
		return validResponse.VisitPostMessagesResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsers operation middleware
func (sh *strictHandler) DeleteUsers(ctx echo.Context) error {
	var request DeleteUsersRequestObject

	var body DeleteUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsers(ctx.Request().Context(), request.(DeleteUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersResponseObject); ok {
		return validResponse.VisitDeleteUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsers operation middleware
func (sh *strictHandler) PatchUsers(ctx echo.Context) error {
	var request PatchUsersRequestObject

	var body PatchUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsers(ctx.Request().Context(), request.(PatchUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersResponseObject); ok {
		return validResponse.VisitPatchUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx echo.Context) error {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx.Request().Context(), request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		return validResponse.VisitPostUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
