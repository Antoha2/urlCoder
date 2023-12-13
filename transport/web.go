package transport

import (
	"context"
	"net/http"

	urlService "github.com/antoha2/urlCoder/service"
)

type webImpl struct {
	Transport
	service urlService.UrlService
	server  *http.Server
}

func NewHTTP(buyService urlService.UrlService) *webImpl {
	return &webImpl{
		service: buyService,
	}
}

type Transport interface {
}

func (wImpl *webImpl) Stop() {

	if err := wImpl.server.Shutdown(context.TODO()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
