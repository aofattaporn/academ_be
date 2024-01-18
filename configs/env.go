package configs

func EnvMongoURI() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return "mongodb://root:example@localhost:27017/"
}

func EnvMongoDatabase() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// return os.Getenv("MONGODATABASE")
	return "academDB"

}
