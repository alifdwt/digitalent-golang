package socialmedia

import "github.com/go-playground/validator/v10"

type CreateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required,url"`
}

func (c *CreateSocialMediaRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
