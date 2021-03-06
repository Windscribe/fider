package handlers

import (
	"github.com/Windscribe/fider/app/pkg/backup"
	"github.com/Windscribe/fider/app/pkg/web"
)

// ExportBackupZip returns a Zip file with all content
func ExportBackupZip() web.HandlerFunc {
	return func(c *web.Context) error {

		file, err := backup.Create(c)
		if err != nil {
			return c.Failure(err)
		}

		return c.Attachment("backup.zip", "application/zip", file.Bytes())
	}
}
