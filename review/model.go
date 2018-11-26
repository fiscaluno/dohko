package review

import (
	"time"

	"github.com/fiscaluno/pandorabox/db"
)

// Review is a Entity
type Review struct {
	ID            uint       `gorm:"primary_key" json:"id"`
	StudentID     uint       `json:"student_id"`
	InstitutionID uint       `json:"institution_id"`
	CourseID      uint       `json:"course_id"`
	Rate          float64    `json:"rate"`
	Title         string     `json:"title"`
	Pros          string     `json:"pros"`
	Cons          string     `json:"cons"`
	Suggestion    string     `json:"suggestion"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	// CourseInfo    Course     `json:"course_info"`
}

// Course ...
// type Course struct {
// 	ID                  uint   `json:"course_id"`
// 	Type                string `json:"course_type"`
// 	Period              string `json:"period"`
// 	Semester            int    `json:"semester"`
// 	Name                string `json:"course_name"`
// 	MonthlyPaymentValue int    `json:"monthly_payment_value"`
// }

// TableName for review
func (Review) TableName() string {
	return "review"
}

// GetAll Review
func GetAll() []Review {
	db := db.Conn()
	// db = db.Set("gorm:auto_preload", true)
	defer db.Close()
	var entities []Review
	// db.Preload("DetailedReviews").Find(&entities)
	db.Find(&entities)
	return entities
}

// Save a Review
func (entity Review) Save() (Review, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Review
func GetByID(id int) Review {
	db := db.Conn()
	defer db.Close()

	var entity Review

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Review
func GetByQuery(query string, value ...interface{}) []Review {
	db := db.Conn()
	defer db.Close()

	var entities []Review

	db.Where(query, value...).Find(&entities)
	return entities
}
