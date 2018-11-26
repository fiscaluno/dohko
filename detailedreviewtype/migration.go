package detailedreviewtype

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	// db.LogMode(true)

	// entity := new(Entity)

	// Migrate the schema
	db.AutoMigrate(&DetailedReviewType{})

	// entity.Name = "Infraestrutura"

	// Create - create entity
	// db.Create(&entity)

	// Find - find entity
	// db.Find(&entity)

	// log.Println(entity)

	// Delete - delete entity
	// db.Delete(&entity)

}
