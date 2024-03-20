package comment

import "github.com/go-playground/validator/v10"

type UpdateCommentRequest struct {
	Message string `json:"message" validate:"required"`
}

func (u *UpdateCommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
