package mapper

type Mapper struct {
	UserMapper        UserMapping
	PhotoMapper       PhotoMapping
	CommentMapper     CommentMapping
	SocialMediaMapper SocialMediaMapping
}

func NewMapper() *Mapper {
	return &Mapper{
		UserMapper:        NewUserMapper(),
		PhotoMapper:       NewPhotoMapper(),
		CommentMapper:     NewCommentMapper(),
		SocialMediaMapper: NewSocialMediaMapper(),
	}
}
