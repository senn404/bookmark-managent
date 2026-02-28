package main

import "github.com/senn404/bookmark-managent/internal/api"

func main() {
	app := api.New()
	app.Start()
}
