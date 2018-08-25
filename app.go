package main

import (
	"github.com/fiscaluno/dohko/review"
	"github.com/fiscaluno/dohko/server"
)

func main() {
	review.Migrate()
	server.Listen()
}
