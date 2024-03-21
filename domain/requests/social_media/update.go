package socialmedia

import "github.com/go-playground/validator/v10"

type UpdateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

func (u *UpdateSocialMediaRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
