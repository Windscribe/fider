package query

import "github.com/Windscribe/fider/app/models"

type GetAttachments struct {
	Post    *models.Post
	Comment *models.Comment

	Result []string
}
