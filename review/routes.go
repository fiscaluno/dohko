package review

import (
	"github.com/gorilla/mux"
)

// SetRoutes add routes from course
func SetRoutes(r *mux.Router) {
	r.HandleFunc("", Add).Methods("POST")
	r.HandleFunc("", FindAll).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", FindByID).Methods("GET")
	r.HandleFunc("/{id:[0-9]+}", DeleteByID).Methods("DELETE")
	r.HandleFunc("/{id:[0-9]+}", UpdateByID).Methods("PUT")
	r.HandleFunc("/details/average", TestAverage).Methods("PUT")
}
