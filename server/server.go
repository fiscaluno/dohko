package server

import (
	"log"
	"net/http"

	"github.com/fiscaluno/dohko/detailedreview"
	"github.com/fiscaluno/dohko/detailedreviewtype"
	"github.com/fiscaluno/dohko/review"
	"github.com/fiscaluno/pandorabox"
	"github.com/fiscaluno/pandorabox/logs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var name string

func handlerHi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola, Seja bem vindo ao @" + name + " !!"))
}

// Listen init a http server
func Listen() {
	port := pandorabox.GetOSEnvironment("PORT", "5001")

	name = pandorabox.GetOSEnvironment("NAME", "Dohko")

	r := mux.NewRouter()
	r.Use(logs.LoggingMiddleware)

	review.SetRoutes(r.PathPrefix("/reviews").Subrouter())
	detailedreview.SetRoutes(r.PathPrefix("/detailedreviews").Subrouter())
	detailedreviewtype.SetRoutes(r.PathPrefix("/detailedreviewtypes").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Client-ID", "X-User-Token", "Content-Type", "X-Requested-With"})
	methodsOk := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})

	log.Println("Listen on port: " + port)
	if err := http.ListenAndServe(":"+port, handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
		log.Fatal(err)
	}
}
