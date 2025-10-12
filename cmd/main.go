package main

import (
	"github.com/jinpain/go-auth-service/internal/app"
	"github.com/jinpain/go-auth-service/internal/config"
	"github.com/jinpain/go-auth-service/pkg/configloader"
)

func main() {
	cfg, err := configloader.Load[config.Config]()
	if err != nil {
		panic(err)
	}

	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
