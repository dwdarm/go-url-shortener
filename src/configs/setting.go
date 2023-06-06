package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Setting struct {
	Env         string
	Port        string
	BaseDomain  string
	UseMongodb  bool
	MongodbURI  string
	MongodbName string
}

func InitSetting() *Setting {
	setting := &Setting{
		Env:         "development",
		Port:        "8080",
		BaseDomain:  "http://localhost:8080",
		UseMongodb:  false,
		MongodbURI:  "mongodb://localhost:27017/",
		MongodbName: "link_shortener",
	}

	err := godotenv.Load()
	if err != nil {
		log.Printf("cannot load .env")
	}

	if os.Getenv("ENV") != "" {
		setting.Env = os.Getenv("ENV")
	}

	if os.Getenv("PORT") != "" {
		setting.Port = os.Getenv("PORT")
	}

	if os.Getenv("BASE_DOMAIN") != "" {
		setting.BaseDomain = os.Getenv("BASE_DOMAIN")
	}

	if os.Getenv("USE_MONGODB") != "" {
		setting.UseMongodb = true
	}

	if os.Getenv("MONGODB_URI") != "" {
		setting.MongodbURI = os.Getenv("MONGODB_URI")
	}

	if os.Getenv("MONGODB_NAME") != "" {
		setting.MongodbName = os.Getenv("MONGODB_NAME")
	}

	return setting
}
