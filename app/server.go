package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"sewa-in/app/database/seeders"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSslMode  string
	DBTimeZone string
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.initializeDB(dbConfig)
	server.initializeRoutes()
	seeders.DBSeed(server.DB)
}

func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort, dbConfig.DBSslMode, dbConfig.DBTimeZone)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect")
	}

	for _, model := range RegisterModels() {
		err = server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Database Migration Sucess")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}

	//appConfig.AppName = "Ini Coba"
	//appConfig.AppEnv = "stagging"
	//appConfig.AppPort = "8080"

	//appConfig.AppName = os.Getenv("APP_NAME")
	//appConfig.AppEnv = os.Getenv("APP_ENV")
	//appConfig.AppPort = os.Getenv("APP_PORT")

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

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}
