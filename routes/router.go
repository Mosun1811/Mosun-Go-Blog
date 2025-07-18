package routes

import (
	"net/http"

	"blog-platform/handlers"
	"blog-platform/middleware"

	"github.com/gorilla/mux"
)


func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetMe))).Methods("GET")

	// Posts
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.GetPostByID).Methods("GET")

	protected := r.PathPrefix("/posts").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("", handlers.CreatePost).Methods("POST")
	protected.HandleFunc("/{id}", handlers.UpdatePost).Methods("PUT")
	protected.HandleFunc("/{id}", handlers.DeletePost).Methods("DELETE")

	return r
}