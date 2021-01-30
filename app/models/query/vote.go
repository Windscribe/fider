package query

import "github.com/Windscribe/fider/app/models"

type ListPostVotes struct {
	PostID int
	Limit  int

	Result []*models.Vote
}
