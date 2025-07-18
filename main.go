package main


import (
	"log"
	"net/http"

	"blog-platform/config"
	"blog-platform/models"
	"blog-platform/routes"
)

func main() {
	config.LoadEnv()
	config.InitDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := routes.SetupRouter()

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}