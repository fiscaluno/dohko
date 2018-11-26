package detailedreview

import (
	"time"

	"github.com/fiscaluno/dohko/detailedreviewtype"
	"github.com/fiscaluno/pandorabox/db"
)

// DetailedReview ...
type DetailedReview struct {
	ID                   uint                                  `gorm:"primary_key" json:"id"`
	CourseID             uint                                  `json:"course_id"`
	StudentID            uint                                  `json:"student_id"`
	InstitutionID        uint                                  `json:"institution_id"`
	ReviewID             uint                                  `json:"review_id"`
	DetailedReviewTypeID uint                                  `json:"review_type_id" `
	Rate                 float64                               `json:"rate"`
	DetailedReviewType   detailedreviewtype.DetailedReviewType `json:"review_type" `
	CreatedAt            time.Time                             `json:"created_at"`
	UpdatedAt            time.Time                             `json:"updated_at"`
	DeletedAt            *time.Time                            `json:"deleted_at"`
}

// TableName for detailedreview
func (DetailedReview) TableName() string {
	return "detailed_reviews"
}

// GetAll []DetailedReview
func GetAll() []DetailedReview {
	db := db.Conn()
	db = db.Set("gorm:auto_preload", true)
	defer db.Close()
	var entities []DetailedReview
	db.Find(&entities)
	return entities
}

// Save a DetailedReview
func (entity DetailedReview) Save() (DetailedReview, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a DetailedReview
func GetByID(id int) DetailedReview {
	db := db.Conn()
	defer db.Close()

	var entity DetailedReview

	db.Find(&entity, id)

	return entity
}

// GetByQuery a DetailedReview
func GetByQuery(query string, value interface{}) []DetailedReview {
	db := db.Conn()
	defer db.Close()

	var entities []DetailedReview

	db.Find(&entities, query, value)
	return entities
}
