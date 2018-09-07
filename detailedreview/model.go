package detailedreview

import (
	"github.com/fiscaluno/dohko/detailedreviewtype"
	"github.com/fiscaluno/pandorabox/db"
)

// DetailedReview ...
type DetailedReview struct {
	CourseID             uint                      `json:"course_id"`
	StudentID            uint                      `json:"student_id"`
	InstitutionID        uint                      `json:"institution_id"`
	ReviewID             uint                      `json:"review_id"`
	DetailedReviewTypeID uint                      `json:"review_type_id" `
	Rate                 float64                   `json:"rate"`
	DetailedReviewType   detailedreviewtype.Entity `json:"review_type" `
}

// Entity is a review
type Entity struct {
	DetailedReview
	db.CommonModelFields
}

// Entities is Entity slice
type Entities []Entity

// TableName for detailedreview
func (Entity) TableName() string {
	return "detailed_reviews"
}

// GetAll Entities
func GetAll() Entities {
	db := db.Conn()
	db = db.Set("gorm:auto_preload", true)
	defer db.Close()
	var entities Entities
	db.Find(&entities)
	return entities
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
func GetByQuery(query string, value interface{}) Entities {
	db := db.Conn()
	defer db.Close()

	var entities Entities

	db.Find(&entities, query, value)
	return entities
}
