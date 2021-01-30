package query

import "github.com/Windscribe/fider/app/models"

type GetTagBySlug struct {
	Slug string

	Result *models.Tag
}

type GetAssignedTags struct {
	Post *models.Post

	Result []*models.Tag
}

type GetAllTags struct {
	Result []*models.Tag
}
