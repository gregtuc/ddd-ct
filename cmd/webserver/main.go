package main

import (
	"log"
	"net/http"

	"github.com/gregtuc/ddd-ct/internal/domain_modeling/application"
	"github.com/gregtuc/ddd-ct/internal/domain_modeling/infrastructure"
)

func main() {
	db, err := infrastructure.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repo := infrastructure.NewDomainModelRepository(db)
	service := application.NewDomainModelService(repo)

	projectRepo := infrastructure.NewProjectRepository(db)
	projectService := application.NewProjectService(projectRepo)

	http.HandleFunc("/create-domain-model", func(w http.ResponseWriter, r *http.Request) {
		// Simulate creating a new domain model
		model, err := service.CreateDomainModel("Sample Model", "A sample domain model for demonstration.")
		if err != nil {
			http.Error(w, "Failed to create domain model", http.StatusInternalServerError)
			return
		}
		response := "Created domain model: " + model.Name
		w.Write([]byte(response))
	})

	http.HandleFunc("/projects/create", func(w http.ResponseWriter, r *http.Request) {
		// Simplified example: In a real application, you should handle request parsing and error checking.
		name := r.URL.Query().Get("name")
		description := r.URL.Query().Get("description")

		project, err := projectService.CreateProject(name, description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := "Created project: " + project.Name
		w.Write([]byte(response))
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
