package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppEnvs struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	PORT        string
}

func ConfigureEnvs() AppEnvs {
	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not configured")
	}

	DB_HOST := os.Getenv("DB_HOST")
	if DB_HOST == "" {
		log.Fatal("DB_HOST not configured")
	}

	DB_USER := os.Getenv("DB_USER")
	if DB_USER == "" {
		log.Fatal("DB_USER not configured")
	}

	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	if DB_PASSWORD == "" {
		log.Fatal("DB_PASSWORD not configured")
	}

	DB_NAME := os.Getenv("DB_NAME")
	if DB_NAME == "" {
		log.Fatal("DB_NAME not configured")
	}

	DB_PORT := os.Getenv("DB_PORT")
	if DB_PORT == "" {
		log.Fatal("DB_PORT not configured")
	}

	return AppEnvs{
		DB_HOST:     DB_HOST,
		DB_USER:     DB_USER,
		DB_PASSWORD: DB_PASSWORD,
		DB_NAME:     DB_NAME,
		DB_PORT:     DB_PORT,
		PORT:        PORT,
	}
}
