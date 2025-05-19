package commands

import (
	"fmt"
	"log"
	"ocserv/internal/models"
	"ocserv/pkg/database"
	"ocserv/pkg/oc"
)

var tables = []interface{}{
	&models.User{},
	&models.UserToken{},
	&models.Panel{},
	oc.OcservUser{},
}

func Migrate() {
	log.Println("starting migrations...")
	engine := database.Get()
	err := engine.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln(fmt.Sprintf("error sync tables: %v", err))
	}
	log.Println("migrating tables successfully")
}
