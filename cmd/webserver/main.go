package main

import (
	"log"
	"net/http"

	"github.com/gtuc/ddd-ct/internal/domainmodeling/application"
)

func main() {
	service := application.DomainModelService{}

	http.HandleFunc("/create-domain-model", func(w http.ResponseWriter, r *http.Request) {
		// Simulate creating a new domain model
		model := service.NewDomainModel("1", "Sample Model", "A sample domain model for demonstration.")
		response := "Created domain model: " + model.Name
		w.Write([]byte(response))
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
