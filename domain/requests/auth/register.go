package auth

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Age             uint   `json:"age" validate:"required,gte=8"`
	Password        string `json:"password" validate:"required,min=6"`
	ProfileImageURL string `json:"profile_image_url" validate:"http_url"`
}

func (l *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
