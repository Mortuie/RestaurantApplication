package bootstrap

import (
	"log"
	"os"
	"restapp/order-api/database"

	"github.com/go-playground/validator/v10"
)

type application struct {
	db     *database.DB
	logger *log.Logger
	v      *validator.Validate
}

func Setup() error {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	db, err := database.New("order.db")

	if err != nil {
		return err
	}

	v := validator.New()

	app := &application{db: db, logger: logger, v: v}

	ServeHttp(app)

	return nil
}
