package view

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type SignInRequest struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type SignInResponse struct {
	Token string `json:"token,omitempty"`
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

func NewLSignInResponse(token string) *SignInResponse {
	return &SignInResponse{
		Token: token,
	}
}
