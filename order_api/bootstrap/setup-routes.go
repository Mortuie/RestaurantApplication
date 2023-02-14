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

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pong\n")
}

func ServeHttp(app *application) error {
	router := httprouter.New()

	router.GET("/ping", index)
	router.GET("/users/:uuid", app.getUser)
	router.POST("/users", app.createUser)

	fmt.Printf("Serving http requests on port: %d\n", 8080)
	return http.ListenAndServe(":8080", router)
}

func (a *application) createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}
