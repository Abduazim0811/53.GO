package handlers

import (
	"encoding/json"
	"homework/internal/models"
	"net/http"
	"github.com/ybbus/jsonrpc/v2"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	client := jsonrpc.NewClient("http://localhost:8081/rpc")
	response, err := client.Call("RPCService.GetUser", map[string]string{"ID": idStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user models.User
	err = response.GetObject(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := jsonrpc.NewClient("http://localhost:8081/rpc")
	response, err := client.Call("RPCService.CreateUser", map[string]models.User{"User": user})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var createdUser models.User
	err = response.GetObject(&createdUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := jsonrpc.NewClient("http://localhost:8081/rpc")
	response, err := client.Call("RPCService.UpdateUser", map[string]models.User{"User": user})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedUser models.User
	err = response.GetObject(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	client := jsonrpc.NewClient("http://localhost:8081/rpc")
	response, err := client.Call("RPCService.DeleteUser", map[string]string{"ID": idStr})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result string
	err = response.GetObject(&result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
