package cmd

import (
	"testing"

	. "github.com/Windscribe/fider/app/pkg/assert"
	"github.com/Windscribe/fider/app/pkg/web"
)

func TestGetMainEngine(t *testing.T) {
	RegisterT(t)

	r := routes(web.New(nil))
	Expect(r).IsNotNil()
}
