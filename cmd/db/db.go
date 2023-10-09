package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	initDBVariables()
	connectionString := fmt.Sprintf(`host=%s user=%s 
	password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai`, PGGlobal["PG_HOST"],
		PGGlobal["PG_USER"], PGGlobal["PG_PASSWORD"], PGGlobal["PG_DBNAME"], PGGlobal["PG_PORT"])
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Quiz{})

	readProduct := &Quiz{}
	db.First(&readProduct, "answers = ?", "23") // find product with code D42
	log.Println(readProduct)
}
