package main

import (
	"log"
	"os"
	"restapp/order-api/internal/database"
	"runtime/debug"
)

type application struct {
	db *database.DB
}

func run(logger *log.Logger) error {

	db, err := database.New("order.db")
	if err != nil {
		return err
	}
	//TODO: db.close?

	app := &application{db: db}

	_ = app
	return nil
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)

	err := run(logger)
	if err != nil {
		trace := debug.Stack()
		logger.Fatalf("%s\n%s", err, trace)
	}
}
