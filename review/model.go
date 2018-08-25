package review

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Review is a Entity
type Review struct {
	CourseID        int             `json:"course_id"`
	StudentID       int             `json:"student_id"`
	InstitutionID   int             `json:"institution_id"`
	Rate            float64         `json:"rate"`
	Title           string          `json:"title"`
	Pros            string          `json:"pros"`
	Cons            string          `json:"cons"`
	Suggestion      string          `json:"suggestion"`
	DetailedReviews DetailedReviews `json:"detailed_reviews"`
}

// DetailedReview ...
type DetailedReview struct {
	CourseID       int     `json:"course_id"`
	StudentID      int     `json:"student_id"`
	InstitutionID  int     `json:"institution_id"`
	ReviewID       int     `json:"review_id"`
	ReviewType     int     `json:"review_type"`
	NameReviewType string  `json:"name_review_type"`
	Rate           float64 `json:"rate"`
}

// DetailedReviewType ...
type DetailedReviewType struct {
	Name string `json:"name"`
	db.CommonModelFields
}

// DetailedReviews ...
type DetailedReviews []DetailedReview

// Entity is a review
type Entity struct {
	Review
	db.CommonModelFields
}

// Entitys is Entity slice
type Entitys []Entity

// GetAll Entitys
func GetAll() Entitys {
	db := db.Conn()
	defer db.Close()
	var entitys Entitys
	db.Find(&entitys)
	return entitys
}

// Save a Entity
func (entity Entity) Save() (Entity, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Entity
func GetByID(id int) Entity {
	db := db.Conn()
	defer db.Close()

	var entity Entity

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Entity
func GetByQuery(query string, value interface{}) Entitys {
	db := db.Conn()
	defer db.Close()

	var entitys Entitys

	db.Find(&entitys, query, value)
	return entitys
}
