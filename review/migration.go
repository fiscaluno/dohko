package review

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	db = db.Set("gorm:auto_preload", true)
	defer db.Close()

	db.LogMode(true)

	entity := new(Entity)

	// Migrate the schema
	db.AutoMigrate(entity)

	entity.CourseID = 1
	entity.StudentID = 1
	entity.InstitutionID = 1
	entity.Rate = 5.0
	entity.Title = "Melhor lugar onde eu poderia estudar"
	entity.Pros = "Tem professores muito bons, os alunos são de um nivel bem alto, tem varios eventos legais"
	entity.Cons = "A infraestrutura tem que melhorar um pouco para as algumas aulas fluirem melhor"
	entity.Suggestion = "Colocarem computadores melhores nos laboratórios"

	// Create
	// db.Create(&entity)

	// db.Preload("DetailedReviews").Find(&entity)
	// Read
	db.Find(&entity)

	// log.Println(entity)
	// log.Println(entity.DetailedReviews[0].DetailedReviewType)

	// Delete - delete entity
	// db.Delete(&entity)
}
