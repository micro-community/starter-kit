package handler

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v3/server"
)

func RegisterHandler(server server.Server) {
	if err := registerAccount(server); err != nil {
		log.Error(err)
	}
}
