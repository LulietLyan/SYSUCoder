package initializer

import (
	"SYSUCODER/tools/yuki"
	"log"
)

func initYuki() {
	err := yuki.InitYukiImage()
	if err != nil {
		log.Println(err)
		log.Println("------------------------Initialize yuki-image failed!------------------------")
		return
	}
	log.Println("------------------------Initialize yuki-image succeeded!------------------------")
}
