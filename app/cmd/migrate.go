package cmd

import (
	"context"

	"github.com/Windscribe/fider/app/models/dto"
	"github.com/Windscribe/fider/app/pkg/bus"
	"github.com/Windscribe/fider/app/pkg/dbx"
	"github.com/Windscribe/fider/app/pkg/log"
	"github.com/Windscribe/fider/app/pkg/rand"
	_ "github.com/Windscribe/fider/app/services/log/console"
)

// RunMigrate run all pending migrations on current DATABASE_URL
// Returns an exitcode, 0 for OK and 1 for ERROR
func RunMigrate() int {
	bus.Init()

	ctx := log.WithProperties(context.Background(), dto.Props{
		log.PropertyKeyTag:       "MIGRATE",
		log.PropertyKeyContextID: rand.String(32),
	})

	err := dbx.Migrate(ctx, "/migrations")
	if err != nil {
		log.Error(ctx, err)
		return 1
	}
	return 0
}
