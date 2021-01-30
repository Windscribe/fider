package middlewares_test

import (
	"net/http"
	"testing"

	"github.com/Windscribe/fider/app/middlewares"
	"github.com/Windscribe/fider/app/pkg/mock"
	"github.com/Windscribe/fider/app/pkg/web"

	. "github.com/Windscribe/fider/app/pkg/assert"
)

func TestCatchPanic_Success(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	server.Use(middlewares.CatchPanic())
	status, _ := server.Execute(func(c *web.Context) error {
		return c.Ok(web.Map{})
	})

	Expect(status).Equals(http.StatusOK)
}

func TestCatchPanic_Panic(t *testing.T) {
	RegisterT(t)

	server := mock.NewServer()
	server.Use(middlewares.CatchPanic())
	status, _ := server.Execute(func(c *web.Context) error {
		panic("Boom!")
	})

	Expect(status).Equals(http.StatusInternalServerError)
}
