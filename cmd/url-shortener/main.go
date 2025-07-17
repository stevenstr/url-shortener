package main

import (
	"github.com/stevenstr/url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()

	// TODO: init logger: slog (log/slog)

	// TODO: init  storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server

}
