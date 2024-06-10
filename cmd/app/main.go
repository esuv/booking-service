package main

import (
	"booking-service/internal/app"
)

const configsDir = "configs"

// @title Booking service
// @version 1.0
// @description Booking API server.

// @contact.name Evgenii Suvorov
// @contact.email eo.suvorov@gmail.com
func main() {
	app.Run(configsDir)
}
