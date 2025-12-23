package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const DbPath = "./db.sqlite3"
const InitSql = "./init.sql"
const SeedJson = "./seed.json"

var DB *sql.DB

// @title           Aadhar Backend API
// @version         1.0
// @description     This is an Go based backend for Assignment 1.
// @host            localhost:8081
func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Failed to load .env. Error: %#v", err)
	}

	err = generateRandomUsers()

	if err != nil {
		log.Printf("Failed to generate seed data. Seed Data might be already created?. Error: %#v", err)
	}

	err = InitDB(InitSql, DbPath)

	if err != nil {
		log.Fatalf("Failed to run database init. Error: %#v", err)
	}

	db, err := GetDB(DbPath)

	if err != nil {
		log.Fatalf("Failed to connect to db. Error: %#v", err)
	}

	DB = db

	err = SeedDb()

	if err != nil {
		log.Printf("Failed to seed Database. Database might already be seeded?. Error: %#v", err)
	}

	router := gin.Default()
	
	// a strict cors setup for frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Request-Origin"},
		MaxAge:       12 * time.Hour,
	}))

	router.POST("/login", Login)
	router.POST("/register", Register)
	router.POST("/refresh", Refresh)
	
	auth := router.Group("/")
	auth.Use(AuthMiddleware())
	auth.GET("get-data", GetData)
	auth.GET("profile", GetProfile)
	
	// exposing swagger files for openapi specs
	router.StaticFS("/swagger", http.Dir("./docs"))

	router.Run(":8081")
}
