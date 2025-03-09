package initializer

func Init() {
	// 顺序初始化
	initConfig()
	initDatabase()

	// 异步初始化
	go initJudge0()
	go initYuki()
	go initNeko()

	initServer()
}
