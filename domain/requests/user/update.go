package user

import "github.com/go-playground/validator/v10"

type UpdateUserRequest struct {
	Username        string `json:"username" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Age             uint   `json:"age" validate:"required,gte=8"`
	ProfileImageURL string `json:"profile_image_url" validate:"http_url"`
}

func (l *UpdateUserRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(l)
	if err != nil {
		return err
	}

	return nil
}
