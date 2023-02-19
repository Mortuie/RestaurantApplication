package bootstrap

import (
	"fmt"
	"net/http"
	"net/url"
	"restapp/order-api/models"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "pong\n")
}

func ServeHttp(app *application) error {
	router := httprouter.New()

	// health check
	router.GET("/ping", index)

	// users
	router.GET("/users", app.getUsers)
	router.GET("/users/:uuid", app.getUser)
	router.POST("/users", app.createUser)

	// restaurants
	router.GET("/restaurants", app.getRestaurants)
	router.GET("/restaurants/:uuid", app.getRestaurant)
	router.POST("/restaurants", app.createRestaurant)

	// menus
	router.GET("/menus", app.getMenus)
	router.GET("/menus/:uuid", app.getMenu)
	router.POST("/menus", app.createMenu)

	fmt.Printf("Serving http requests on port: %d\n", 8080)
	return http.ListenAndServe(":8080", router)
}

func getDefaultPagingConfig(u url.Values) (*models.PagingConfig, error) {
	offset := u.Get("offset")
	var pc models.PagingConfig

	if offset == "" {
		pc.Offset = 0
	} else {
		val, err := strconv.Atoi(offset)

		if err != nil {
			return &pc, err
		}

		pc.Offset = val
	}

	pageSize := u.Get("pageSize")

	if pageSize == "" {
		pc.PageSize = 10
	} else {
		val, err := strconv.Atoi(pageSize)

		if err != nil {
			return &pc, err
		}

		pc.PageSize = val
	}

	return &pc, nil
}
