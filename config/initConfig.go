package config

func Init() {
	loadVariables()
	dbConfig()
	syncDatabase()
}
