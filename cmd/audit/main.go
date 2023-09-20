package main

import (
	"github.com/samarec1812/audit-service/internal/app"
	"github.com/samarec1812/audit-service/internal/config"
)

func main() {
	cfg := config.MustLoad()

	app.Run(cfg)
}
