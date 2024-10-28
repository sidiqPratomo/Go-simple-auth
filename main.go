package main

import (
	"github.com/sidiqPratomo/DJKI-Pengaduan/server"
)

func main() {
	// Initialize the logger
	// log := logrus.New()
	// log.Out = os.Stdout

	// // Initialize configuration
	// conf := config.Init(log)

	// // Create router with database connection
	// router := server.CreateRouter(log, conf)

	// // Run the server
	// log.Fatal(router.Run(":8080"))
	server.Init()
}
