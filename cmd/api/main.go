package main

import (
	"log"

	"github.com/honganji/social/internal/env"
	"github.com/honganji/social/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	store := store.NewStorage(nil)
	app := &applilcation{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal((app.run(mux)))
}
