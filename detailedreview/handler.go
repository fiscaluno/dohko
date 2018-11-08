package detailedreview

import (
	"log"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fiscaluno/pandorabox"
	"github.com/fiscaluno/pandorabox/db"
	"github.com/fiscaluno/dohko/detailedreviewtype"

	"github.com/gorilla/mux"
)

// FindAll entitys
func FindAll(w http.ResponseWriter, r *http.Request) {
	entitys := GetAll()
	pandorabox.RespondWithJSON(w, http.StatusOK, entitys)
}

// FindByID find a entity by ID
func FindByID(w http.ResponseWriter, r *http.Request) {

	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	entity := GetByID(id)

	if entity.ID != 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, entity)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Course",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a entity
func Add(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	db.Create(&entity)

	msg = pandorabox.Message{
		Content: "New Course successfully added",
		Status:  "OK",
		Body:    entity,
	}
	pandorabox.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a entity by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	db.Find(&entity, id)
	if entity.ID != 0 {
		db.Delete(&entity)
		msg = pandorabox.Message{
			Content: "Deleted Course successfully",
			Status:  "OK",
			Body:    entity,
		}

		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Course",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a entity by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity Entity
	var msg pandorabox.Message

	msg = pandorabox.Message{
		Content: "Invalid ID, not a int",
		Status:  "ERROR",
		Body:    nil,
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		pandorabox.RespondWithJSON(w, http.StatusOK, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Invalid request payload",
		Status:  "ERROR",
		Body:    nil,
	}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	if id >= 0 {
		entity.ID = uint(id)
		db.Save(&entity)

		msg = pandorabox.Message{
			Content: "Update Course successfully ",
			Status:  "OK",
			Body:    entity,
		}
		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this Course",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

func GetDetailedReviewsAverage(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	institutionId, err := strconv.Atoi(vars["institution_id"][0])

	if err != nil {
		msg := pandorabox.Message {
			Content: "Institution id must be defined",
			Status: "ERROR",
			Body: nil,
		}
		pandorabox.RespondWithJSON(w, http.StatusBadRequest, msg)
		return
	}

	allInstitutionDetailedReviews := GetByQuery("institution_id = ?", institutionId)

	mappedReviews := make(map[string] float64)
	reviewCount := make(map[string] int)

	for index, _ := range allInstitutionDetailedReviews {
		detailedReview := allInstitutionDetailedReviews[index]
		detailedReviewType := detailedreviewtype.GetByID(int(detailedReview.DetailedReviewTypeID))
		mappedReviews[detailedReviewType.Name] += detailedReview.Rate / reviewCount[detailedReviewType.Name]
		reviewCount[detailedReviewType.Name] += 1
	}

	msg := pandorabox.Message {
		Content: "",
		Status: "OK",
		Body: mappedReviews,
	}

	log.Println(mappedReviews)

	pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
}

// FindByFacebookID find a entity by FacebookID
func FindByFacebookID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	FacebookID := vars["id"]
	entitys := GetByQuery("facebook_id = ?", FacebookID)

	if len(entitys) >= 0 {
		pandorabox.RespondWithJSON(w, http.StatusOK, entitys)
		return
	}

	msg := pandorabox.Message{
		Content: "Not exist this Course",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}
