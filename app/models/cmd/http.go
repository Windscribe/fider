package cmd

import (
	"io"

	"github.com/Windscribe/fider/app/models/dto"
)

type HTTPRequest struct {
	URL       string
	Body      io.Reader
	Method    string
	Headers   map[string]string
	BasicAuth *dto.BasicAuth

	ResponseBody       []byte
	ResponseStatusCode int
}
