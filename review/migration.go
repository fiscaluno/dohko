package review

import (
	"fmt"

	"github.com/fiscaluno/pandorabox/db"
)

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	db.LogMode(true)

	entity := new(Review)
	detailedReview := new(DetailedReview)

	// Migrate the schema
	db.AutoMigrate(&entity, &detailedReview)
	entity.Title = "katarina linda"
	entity.DetailedReviews = []DetailedReview{
		DetailedReview{
			ReviewID:       1,
			NameReviewType: "jc",
			Rate:           6.7,
		},
	}
	db.Model(&entity).Related(&detailedReview)

	// Create
	db.Create(entity)

	// Read
	// var entity Entity
	// db.First(&entity, 1) // find entity with id 1
	// db.First(&entity, "name = ?", "JC") // find entity with name JC

	// Update - update entity's Name to SI
	// db.Model(&entity).Update("Title", "SI")

	db.Model(&entity).Related(&detailedReview)

	fmt.Println(detailedReview)

	// Delete - delete entity
	// db.Delete(&entity)
}
