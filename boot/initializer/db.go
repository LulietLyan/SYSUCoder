package initializer

import (
	"SYSUCODER/boot/database"
	"log"
)

func initDatabase() {
	err := database.InitDatabase()
	if err != nil {
		log.Println("------------------------Initialize database failed!------------------------")
		panic(err)
	}
	log.Println("------------------------Initialize database succeeded!------------------------")
}
