package main

import (
	"os"

	"github.com/sidiqPratomo/DJKI-Pengaduan/config"
	"github.com/sidiqPratomo/DJKI-Pengaduan/server"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize the logger
	log := logrus.New()
	log.Out = os.Stdout

	// Initialize configuration
	conf := config.Init(log)

	// Create router with database connection
	router := server.CreateRouter(log, conf)

	// Run the server
	log.Fatal(router.Run(":8080"))
}
