package main

import (
	"flag"
	"fmt"
	"log"
	_ "ocserv/docs"
	"ocserv/internal/commands"
	"ocserv/pkg/config"
	"ocserv/pkg/database"
	"ocserv/pkg/routing"
	"os"
	"os/signal"
	"syscall"
)

// @title Ocserv User management Example Api
// @version 1.0
// @description This is a sample Ocserv User management Api server.
// @BasePath /api
func main() {
	var (
		debug   bool
		migrate bool
		drop    bool
	)

	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&migrate, "migrate", false, "migrate models to database")
	flag.BoolVar(&drop, "drop", false, "drop models table from database")
	flag.Parse()

	config.Init(debug)
	database.Connect(debug)
	defer database.Close()

	if migrate {
		commands.Migrate()
	} else if drop && debug {
		commands.Drop()
	} else {

		go func() {
			routing.Serve(debug)
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		sig := <-quit

		fmt.Println()
		log.Printf("signal %v recieved\n", sig)
		log.Println("initiating shutdown...")

		routing.Shutdown()
	}
}
