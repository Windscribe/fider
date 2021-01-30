package actions

import (
	"context"

	"github.com/Windscribe/fider/app/models"
	"github.com/Windscribe/fider/app/pkg/validate"
)

// Actionable is any action that the user can perform using the web app
type Actionable interface {
	Initialize() interface{}
	IsAuthorized(ctx context.Context, user *models.User) bool
	Validate(ctx context.Context, user *models.User) *validate.Result
}
