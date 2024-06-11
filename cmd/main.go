package main

import (
	"homework/internal/db"
	"homework/internal/handlers"
	"homework/internal/rpc"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	database, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		log.Fatal("DB connection could not be set")
	}
	defer database.Close()

	go rpc.StartRPCServer()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetUser(w, r)
		case http.MethodPost:
			handlers.CreateUser(w, r)
		case http.MethodPut:
			handlers.UpdateUser(w, r)
		case http.MethodDelete:
			handlers.DeleteUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
