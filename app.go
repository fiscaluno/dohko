package main

import (
	"github.com/fiscaluno/dohko/detailedreview"
	"github.com/fiscaluno/dohko/detailedreviewtype"
	"github.com/fiscaluno/dohko/review"
	"github.com/fiscaluno/dohko/server"
)

func main() {
	detailedreviewtype.Migrate()
	detailedreview.Migrate()
	review.Migrate()

	server.Listen()
}
