package cmd

import (
	"github.com/Windscribe/fider/app/models"
	"github.com/Windscribe/fider/app/models/dto"
)

type SaveCustomOAuthConfig struct {
	Config *models.CreateEditOAuthConfig
}

type ParseOAuthRawProfile struct {
	Provider string
	Body     string

	Result *dto.OAuthUserProfile
}
