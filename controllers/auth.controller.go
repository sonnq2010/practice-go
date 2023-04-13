package controllers

import (
	"net/http"
	"practice-go/services"
)

type IAuthController interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
}

type AuthController struct {
	authService services.IAuthService
}

func NewAuthController(s services.IAuthService) IAuthController {
	return AuthController{
		authService: s,
	}
}

func (c AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate

	_, err := c.authService.SignUp()

	if err != nil {
		w.Write([]byte("SignUp error"))
	}

	w.Write([]byte("SignUp ok"))
}

func (c AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: Validate

	_, err := c.authService.SignIn()

	if err != nil {
		w.Write([]byte("SignIn error"))
	}

	w.Write([]byte("SignIn ok"))
}
