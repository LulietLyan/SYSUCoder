package initializer

import (
	"SYSUCODER/boot/configuration"
	"log"
)

func initConfig() {
	err := configuration.InitConfig()
	if err != nil {
		log.Println("------------------------Initialize configurations failed!------------------------")
		panic(err)
	}
	log.Println("------------------------Initialize configurations successed!------------------------")
}
