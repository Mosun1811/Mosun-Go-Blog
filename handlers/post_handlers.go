package handlers

import (
	"encoding/json"
	"net/http"

	"blog-platform/config"
	"blog-platform/models"

	"github.com/gorilla/mux"
	"github.com/google/uuid"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	config.DB.Preload("Author").Find(&posts)
	json.NewEncoder(w).Encode(posts)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r) ["id"]
	var post models.Post
	if err := config.DB.Preload("Author").First(&post, "id = ?", id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	userID, _ := uuid.Parse(r.Context().Value("user_id").(string))
	post.AuthorID = userID
	config.DB.Create(&post)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := r.Context().Value("user_id").(string)

	var post models.Post
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if post.AuthorID.String() != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var input models.Post
	json.NewDecoder(r.Body).Decode(&input)
	post.Title = input.Title
	post.Content = input.Content
	config.DB.Save(&post)
	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	userID := r.Context().Value("user_id").(string)

	var post models.Post
	if err := config.DB.First(&post, "id = ?", id).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	if post.AuthorID.String() != userID {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	config.DB.Delete(&post)
	w.WriteHeader(http.StatusNoContent)
}

