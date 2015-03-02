package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/ruprict/loccasions-go/api"
	"log"
)

func init() {
	log.Println("** Migrating DB")
	var err error
	db, err := gorm.Open("postgres", "user=postgres dbname=loccasions_development sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&api.User{})
	db.AutoMigrate(&api.Loccasion{})
}
