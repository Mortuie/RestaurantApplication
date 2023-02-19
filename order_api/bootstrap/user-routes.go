package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapp/order-api/database"
	"restapp/order-api/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (a *application) createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = a.v.Struct(u); err != nil {
		fmt.Println("MEMES, validation err", err)
		return
	}

	u.UUID = uuid.NewString()
	err = u.HashPassword()
	if err != nil {
		fmt.Println("rip failure", err)
		return
	}

	err = database.InsertUser(a.db, u)

	if err != nil {
		fmt.Println("rip failure", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&models.UserResponse{UUID: u.UUID, Username: u.Username})
}

func (a *application) getUser(w http.ResponseWriter, r *http.Request, qp httprouter.Params) {
	userUuid := qp.ByName("uuid")
	if userUuid == "" {
		http.Error(w, "user uuid cannot be empty", http.StatusBadRequest)
		return
	}

	user, err := database.GetUser(a.db, userUuid)
	if err != nil {
		fmt.Println("rip failure", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)

}

func (a *application) getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues, err := getDefaultPagingConfig(r.URL.Query())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := database.GetUsers(a.db, *queryValues)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}
