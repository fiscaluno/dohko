package detailedreview

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fiscaluno/dohko/detailedreviewtype"
	"github.com/fiscaluno/pandorabox"
	"github.com/fiscaluno/pandorabox/db"

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
		Content: "Not exist this DetailedReview",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// Add a entity
func Add(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity DetailedReview
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
		Content: "New DetailedReview successfully added",
		Status:  "OK",
		Body:    entity,
	}
	pandorabox.RespondWithJSON(w, http.StatusCreated, msg)

}

// DeleteByID find a entity by ID
func DeleteByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity DetailedReview
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
			Content: "Deleted DetailedReview successfully",
			Status:  "OK",
			Body:    entity,
		}

		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this DetailedReview",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

// UpdateByID find a entity by ID
func UpdateByID(w http.ResponseWriter, r *http.Request) {
	db := db.Conn()
	defer db.Close()

	var entity DetailedReview
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
			Content: "Update DetailedReview successfully ",
			Status:  "OK",
			Body:    entity,
		}
		pandorabox.RespondWithJSON(w, http.StatusAccepted, msg)
		return
	}

	msg = pandorabox.Message{
		Content: "Not exist this DetailedReview",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}

func GetDetailedReviewsAverage(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	institutionId, _ := strconv.Atoi(vars["institution_id"][0])

	allInstitutionDetailedReviews := GetByQuery("institution_id = ?", institutionId)

	mappedReviews := make(map[string]float64)
	reviewCount := make(map[string]float64)

	for index, _ := range allInstitutionDetailedReviews {
		detailedReview := allInstitutionDetailedReviews[index]
		detailedReviewType := detailedreviewtype.GetByID(int(detailedReview.DetailedReviewTypeID))

		reviewCount[detailedReviewType.Name] += float64(1)
		mappedReviews[detailedReviewType.Name] += detailedReview.Rate
	}

	type ReviewType struct {
		Description string  `json:"description"`
		Rate        float64 `json:"rate"`
	}

	var response []ReviewType
	for index, _ := range mappedReviews {
		reviewType := ReviewType{
			Description: index,
			Rate:        mappedReviews[index] / reviewCount[index],
		}

		response = append(response, reviewType)
	}

	msg := pandorabox.Message{
		Content: "",
		Status:  "OK",
		Body:    response,
	}

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
		Content: "Not exist this DetailedReview",
		Status:  "ERROR",
		Body:    nil,
	}
	pandorabox.RespondWithJSON(w, http.StatusOK, msg)

}
