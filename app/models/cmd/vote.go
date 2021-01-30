package cmd

import "github.com/Windscribe/fider/app/models"

type AddVote struct {
	Post *models.Post
	User *models.User
}

type RemoveVote struct {
	Post *models.Post
	User *models.User
}

type MarkPostAsDuplicate struct {
	Post     *models.Post
	Original *models.Post
}
