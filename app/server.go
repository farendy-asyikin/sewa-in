package app

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sewa-in/app/controllers"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Sewa.no")
	appConfig.AppEnv = getEnv("APP_ENV", "Stagging")
	appConfig.AppPort = getEnv("APP_PORT", "7070")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "farendfal")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "sewa.in")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	dbConfig.DBSslMode = getEnv("DB_SSL_MODE", "disable")
	dbConfig.DBTimeZone = getEnv("DB_TIME_ZONE", "Asia/Jakarta")

	flag.Parse()
	arg := flag.Arg(0)
	if arg != "" {
		server.InitComands(appConfig, dbConfig)
	} else {
		server.Initialize(appConfig, dbConfig)
		server.Run(":" + appConfig.AppPort)
	}
	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
