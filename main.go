package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/shashank-mugiwara/meowxy/conf"
	customLogger "github.com/shashank-mugiwara/meowxy/logger"
	"github.com/shashank-mugiwara/meowxy/metrics"
)

func InitalizeAllRoutes() {
	metrics.RegisterRoutes(conf.GetMeowxyEngine(), customLogger.GetMeowxyLogger())
}

func main() {
	conf.InitEngine()
	app := conf.GetMeowxyEngine()
	InitalizeAllRoutes()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(":8080"); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
}
