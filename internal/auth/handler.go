package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/request"
	"go/adv-demo/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}
func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		response.Json(w, data, 200)
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		fmt.Println(body)
	}
}
