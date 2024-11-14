package main

import (
	"github.com/diki-haryadi/go-micro-template/app"
	"github.com/diki-haryadi/go-micro-template/pkg/logger"
)

func main() {
	err := app.New().Run()
	if err != nil {
		logger.Zap.Sugar().Fatal(err)
	}
}