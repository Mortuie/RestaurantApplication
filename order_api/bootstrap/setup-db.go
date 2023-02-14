package bootstrap

import (
	"log"
	"os"
	"restapp/order-api/database"
)

type application struct {
	db     *database.DB
	logger *log.Logger
}

func Setup() error {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	db, err := database.New("order.db")

	if err != nil {
		return err
	}

	app := &application{db: db, logger: logger}

	ServeHttp(app)

	return nil
}
