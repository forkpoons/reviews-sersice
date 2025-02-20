package api

import (
	"github.com/fasthttp/router"
	"github.com/rs/zerolog"
)

type service struct {
	r         *router.Router
	log       zerolog.Logger
	appSecret string
}

func NewService(log zerolog.Logger, appSecret string) *service {
	s := service{
		log:       log,
		appSecret: appSecret,
	}

	return &s
}
