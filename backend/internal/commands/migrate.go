package commands

import (
	"fmt"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/database"
)

var tables = []interface{}{
	&models.User{},
	&models.UserToken{},
	&models.Panel{},
}

func Migrate() {
	engine := database.Get()
	err := engine.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln(fmt.Sprintf("error sync tables: %v", err))
	}
	log.Println("migrating tables successfully")
}
