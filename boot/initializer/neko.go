package initializer

import (
	"SYSUCODER/tools/neko"
	"log"
)

func initNeko() {
	err := neko.InitNekoAcm()
	if err != nil {
		log.Println(err)
		log.Println("初始化 NekoACM 失败！")
		return
	}
	log.Println("初始化 NekoACM 成功")
}
