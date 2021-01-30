package cmd

import (
	"github.com/Windscribe/fider/app/models"
	"github.com/Windscribe/fider/app/models/enum"
)

type AddNewPost struct {
	Title       string
	Description string

	Result *models.Post
}

type UpdatePost struct {
	Post        *models.Post
	Title       string
	Description string

	Result *models.Post
}

type SetPostResponse struct {
	Post   *models.Post
	Text   string
	Status enum.PostStatus
}
