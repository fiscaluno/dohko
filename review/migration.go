package review

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Migrate migration review BD
func Migrate() {
	db := db.Conn()
	db = db.Set("gorm:auto_preload", true)
	defer db.Close()

	// db.LogMode(true)

	// review := new(Review)

	// Migrate the schema
	db.AutoMigrate(&Review{})

	// review.StudentID = 1
	// review.InstitutionID = 1
	// review.CourseID = 1
	// review.Rate = 5.0
	// review.Title = "Melhor lugar onde eu poderia estudar"
	// review.Pros = "Tem professores muito bons, os alunos são de um nivel bem alto, tem varios eventos legais"
	// review.Cons = "A infraestrutura tem que melhorar um pouco para as algumas aulas fluirem melhor"
	// review.Suggestion = "Colocarem computadores melhores nos laboratórios"

	// review.CourseInfo = Course{
	// 	ID:                  1,
	// 	Name:                "Information Systems",
	// 	Type:                "Bachelor",
	// 	Period:              "nightly",
	// 	Semester:            8,
	// 	MonthlyPaymentValue: 1000,
	// }

	// Create
	// db.Create(&review)

	// db.Preload("DetailedReviews").Find(&review)
	// Read
	// db.Find(&review)

	// log.Println(entity)
	// log.Println(entity.DetailedReviews[0].DetailedReviewType)

	// Delete - delete entity
	// db.Delete(&review)
}
