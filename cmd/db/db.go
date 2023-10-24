package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

var DB *gorm.DB

func Init() {
	initDBVariables()
	var err error
	connectionString := fmt.Sprintf(`host=%s user=%s 
	password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`, PGGlobal["PG_HOST"],
		PGGlobal["PG_USER"], PGGlobal["PG_PASSWORD"], PGGlobal["PG_DBNAME"], PGGlobal["PG_PORT"])
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	createQuiz()
	readQuiz()
}
