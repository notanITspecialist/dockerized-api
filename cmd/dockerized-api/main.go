package main

import "dockerized-api/internal/app"

func main() {
	application := app.NewApplication()
	application.Run()
}
