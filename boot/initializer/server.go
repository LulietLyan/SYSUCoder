package initializer

import (
	"SYSUCODER/server"
	"log"
)

func initServer() {
	err := server.InitServer()
	if err != nil {
		log.Println("Initialize server failed!")
		panic(err)
	}
	log.Println("Initialize server successed!")
}
