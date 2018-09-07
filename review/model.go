package review

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Review is a Entity
type Review struct {
	CourseID      uint    `json:"course_id"`
	StudentID     uint    `json:"student_id"`
	InstitutionID uint    `json:"institution_id"`
	Rate          float64 `json:"rate"`
	Title         string  `json:"title"`
	Pros          string  `json:"pros"`
	Cons          string  `json:"cons"`
	Suggestion    string  `json:"suggestion"`
}

// Entity is a review
type Entity struct {
	Review
	// DetailedReviews []detailedreview.Entity `json:"detailed_reviews" gorm:"ForeignKey:ReviewID"`
	db.CommonModelFields
}

// TableName for review
func (Entity) TableName() string {
	return "review"
}

// Entities is Entity slice
type Entities []Entity

// GetAll Entities
func GetAll() Entities {
	db := db.Conn()
	// db = db.Set("gorm:auto_preload", true)
	defer db.Close()
	var entities Entities
	// db.Preload("DetailedReviews").Find(&entities)
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
