package main

import (
	"todo/webserver"
)

func main() {
	app := webserver.CreateWebServer()
	app.Run("localhost:8003")
}
