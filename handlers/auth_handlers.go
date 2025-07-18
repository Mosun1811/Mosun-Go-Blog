package handlers

import (
	"encoding/json"
	"net/http"

	"blog-platform/config"
	"blog-platform/models"
	"blog-platform/utils"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed)
	user.ID = uuid.New()

	config.DB.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	var user models.User
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid input", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(string)

	var user models.User
	if err := config.DB.First(&user, "id = ?", userID).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
