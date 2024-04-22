package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dev-patmazo/user-management/controllers"
	"github.com/dev-patmazo/user-management/middlewares"
	"github.com/dev-patmazo/user-management/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {

	// Load environment file for database connection
	// Note: Environment file should be in the root directory of the project
	log.Println("Loading environment file...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	log.Println("Connecting to database...")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbName)
	models.DbClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, err := models.DbClient.DB()
	if err != nil {
		panic("failed to connect database")
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetConnMaxLifetime(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	models.DbClient.AutoMigrate(&models.User{})
	log.Println("Database connected successfully.")

}

func main() {
	// Initialize router and routes to be used as api endpoints.
	r := mux.NewRouter()
	r.HandleFunc("/", healthChecker).Methods("GET")
	r.Use(middlewares.AuthenticationRBAC)
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":8080", r)
}

// Health checker function to check if the server is running
func healthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
