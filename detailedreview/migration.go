package detailedreview

import (
	"log"

	"github.com/fiscaluno/pandorabox/db"
)

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	db = db.Set("gorm:auto_preload", true)
	defer db.Close()

	// db.LogMode(true)

	entity := new(Entity)

	// Migrate the schema
	db.AutoMigrate(&entity)

	entity.CourseID = 1
	entity.StudentID = 1
	entity.InstitutionID = 1
	entity.ReviewID = 1
	entity.DetailedReviewTypeID = 1
	entity.Rate = 5.0

	// Create - create entity
	db.Create(&entity)

	// Find - find entity
	db.Find(&entity)

	log.Println(entity)

	// Delete - delete entity
	// db.Delete(&entity)

}
