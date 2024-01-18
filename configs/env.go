package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	fmt.Println(err)

	if err != nil {
		fmt.Println(err)
	}

	return os.Getenv("MONGOURI")
}

func EnvMongoDatabase() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	return os.Getenv("MONGODATABASE")
}
