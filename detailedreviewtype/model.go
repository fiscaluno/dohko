package detailedreviewtype

import (
	"github.com/fiscaluno/pandorabox/db"
)

// DetailedReviewType ...
type DetailedReviewType struct {
	Name string `json:"name"`
}

// Entity is a review
type Entity struct {
	DetailedReviewType
	db.CommonModelFields
}

// Entities is Entity slice
type Entities []Entity

// TableName for detailedreviewtype
func (Entity) TableName() string {
	return "detailed_review_types"
}

// GetAll Entities
func GetAll() Entities {
	db := db.Conn()
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
