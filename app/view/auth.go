package view

import (
	"encoding/json"
	"io"
	"regexp"
	"unicode/utf8"

	"github.com/go-playground/validator"
)

type SignInRequest struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required"`
}

type SignInResponse struct {
	Token  string `json:"token,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

type SignUpRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,securePassword"`
}

type SignUpResponse struct {
	Token  string `json:"token,omitempty"`
	UserID int    `json:"user_id,omitempty"`
}

func NewSignInRequest(body io.ReadCloser) (*SignInRequest, error) {
	signInRequest := &SignInRequest{}
	if err := json.NewDecoder(body).Decode(signInRequest); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(signInRequest)
	if err != nil {
		return nil, err
	}

	return signInRequest, nil
}

func NewLSignInResponse(token string, userID int) *SignInResponse {
	return &SignInResponse{
		Token:  token,
		UserID: userID,
	}
}

func NewSignUpRequest(body io.ReadCloser) (*SignUpRequest, error) {
	signUpRequest := &SignUpRequest{}
	if err := json.NewDecoder(body).Decode(signUpRequest); err != nil {
		return nil, err
	}

	validate := validator.New()
	validate.RegisterValidation("securePassword", securePassword)
	if err := validate.Struct(signUpRequest); err != nil {
		return nil, err
	}

	return signUpRequest, nil
}

func NewLSignUpResponse(token string, userID int) *SignUpResponse {
	return &SignUpResponse{
		Token:  token,
		UserID: userID,
	}
}

func securePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if utf8.RuneCountInString(password) < 10 {
		return false
	}
	r1, _ := regexp.Compile(`([A-Z]+)`)
	r2, _ := regexp.Compile(`([a-z]+)`)
	r3, _ := regexp.Compile(`([0-9]+)`)

	l1 := len(r1.FindAllString(password, -1)) > 0
	l2 := len(r2.FindAllString(password, -1)) > 0
	l3 := len(r3.FindAllString(password, -1)) > 0

	return l1 && l2 && l3
}
