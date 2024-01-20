package main

import (
	"github.com/labstack/gommon/log"
	"keyword-generator/src/config"
	"keyword-generator/src/infrastructure/wathcer"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	wathcer.RunMoveWatcher()
}
