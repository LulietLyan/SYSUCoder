package initializer

import (
	"SYSUCODER/tools/neko"
	"log"
)

func initNeko() {
	err := neko.InitNekoAcm()
	if err != nil {
		log.Println(err)
		log.Println("------------------------Initialize NekoACM failed！------------------------")
		return
	}
	log.Println("------------------------Initialize NekoACM succeeded!------------------------")
}
