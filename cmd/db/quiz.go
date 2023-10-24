package db

import (
	"log"

	"github.com/jackc/pgx/v5/pgtype"
)

func (Quiz) TableName() string {
	return "quiz"
}

type Quiz struct {
	ID      uint   `gorm:"primaryKey;default:auto_random()"`
	Content string `gorm:"type:text"`
	Answers pgtype.Array[int]
}

// CREATE
func createQuiz() {
	quiz := Quiz{Content: "Jinzhu", Answers: [...]int{50, 45, 30, 49, 38}}
	DB.Model(&quiz)

	result := DB.Create(&quiz) // pass pointer of data to Create
	if result.Error == nil {
		log.Println("[ERROR] creating quiz:", quiz.ID)
	} else {
		log.Println("Successfully inserted quiz:", result)
	}
}

// READ
func readQuiz() {
	quizes := []Quiz{}
	DB.AutoMigrate(&Quiz{})
	DB.Table("quiz")
	// Get all records
	result := DB.Find(&quizes)
	// SELECT * FROM users;
	if result.Error != nil {
		log.Println("[ERROR] couldn't read all quizes:", result.Error)
		return
	}
	log.Println(quizes)
}

func readQuizSpecific(answers int) {

	readProduct := &Quiz{}
	result := DB.First(&readProduct, "answers = ?", answers)
	if result.Error != nil {
		log.Println("[ERROR] error getting quiz with answers", answers)
		return
	}
	log.Println(readProduct)
}
