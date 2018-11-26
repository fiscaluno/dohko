package detailedreviewtype

import (
	"time"

	"github.com/fiscaluno/pandorabox/db"
)

// DetailedReviewType ...
type DetailedReviewType struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// TableName for detailedreviewtype
func (DetailedReviewType) TableName() string {
	return "detailed_review_types"
}

// GetAll []DetailedReviewType
func GetAll() []DetailedReviewType {
	db := db.Conn()
	defer db.Close()
	var entities []DetailedReviewType
	db.Find(&entities)
	return entities
}

// Save a DetailedReviewType
func (entity DetailedReviewType) Save() (DetailedReviewType, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a DetailedReviewType
func GetByID(id int) DetailedReviewType {
	db := db.Conn()
	defer db.Close()

	var entity DetailedReviewType

	db.Find(&entity, id)

	return entity
}

// GetByQuery a DetailedReviewType
func GetByQuery(query string, value interface{}) []DetailedReviewType {
	db := db.Conn()
	defer db.Close()

	var entities []DetailedReviewType

	db.Find(&entities, query, value)
	return entities
}
