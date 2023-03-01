package main

import (
	"github.com/batuhannoz/paribu-case/app"
	"github.com/batuhannoz/paribu-case/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
