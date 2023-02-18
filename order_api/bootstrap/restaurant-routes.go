package bootstrap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapp/order-api/database"
	"restapp/order-api/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/validator.v2"
)

// needs uuid
func (a *application) getRestaurant(w http.ResponseWriter, r *http.Request, qp httprouter.Params) {
	rUuid := qp.ByName("uuid")

	if rUuid == "" {
		http.Error(w, "restaurant uuid cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := database.GetRestaurant(a.db, rUuid)
	if err != nil {
		http.Error(w, "restaurant not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&res)
}

func (a *application) getRestaurants(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues, err := getDefaultPagingConfig(r.URL.Query())

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rs, err := database.GetRestaurants(a.db, *queryValues)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&rs)
}

func (a *application) createRestaurant(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var res models.Restaurant

	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := validator.Validate(res); err != nil {
		fmt.Println(res, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.UUID = uuid.New().String()

	if err := database.InsertRestaurant(a.db, res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&res)
}
