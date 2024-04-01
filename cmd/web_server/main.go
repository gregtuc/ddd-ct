package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gregtuc/ddd-ct/internal/domain_modeling/application"
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	db, err := infrastructure.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	repo := infrastructure.NewDomainModelRepository(db)
	service := application.NewDomainModelService(repo)
	projectRepo := infrastructure.NewProjectRepository(db)
	projectService := application.NewProjectService(projectRepo)
	userRepo := infrastructure.NewUserRepository(db)
	userService := application.NewUserService(userRepo)

	http.HandleFunc("/create-domain-model", func(w http.ResponseWriter, r *http.Request) {
		// Simulate creating a new domain model
		model, err := service.CreateDomainModel("Sample Model", "A sample domain model for demonstration.")
		if err != nil {
			http.Error(w, "Failed to create domain model", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Created domain model: " + model.Name))
	})

	http.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		description := r.URL.Query().Get("description")

		project, err := projectService.CreateProject(name, description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Created project: " + project.Name))
	})

	http.HandleFunc("/users/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		user, err := userService.RegisterUser(req.Name, req.Email, string(hashedPassword))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User ID: " + fmt.Sprint(user.ID)))
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
