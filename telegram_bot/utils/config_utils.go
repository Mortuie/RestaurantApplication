package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramApiToken string
	Timeout          int
}

func checkEnvVariable(e string) string {
	tempEnv := os.Getenv(e)
	if tempEnv == "" {
		log.Fatalf("error parsing: %s.", e)
	}
	return tempEnv
}

func (c *Config) GetEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c.TelegramApiToken = checkEnvVariable("TELEGRAM_API_TOKEN")
	c.Timeout, _ = strconv.Atoi(checkEnvVariable("TIMEOUT"))
}
