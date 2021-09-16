package seed

import (
	"log"

	"github.com/igorariza/Dockerized-Golang_API-MySql-React.js/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Firstname: "Firstname",
		Lastname:  "Lastname",
		Position:  "Position",
		Nation:    "Nation",
		Club:      "Club",
		Page:      "Page",
	},
	models.User{
		Firstname: "Firstname2",
		Lastname:  "Lastname2",
		Position:  "Position2",
		Nation:    "Nation2",
		Club:      "Club2",
		Page:      "Page2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
