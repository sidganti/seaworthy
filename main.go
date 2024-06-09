package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println("Booting up")

	log.Println("Loading environment")
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	us_trade := os.Getenv("EODHD_US_TRADE_SERVER")
	us_quote := os.Getenv("EODHD_US_QUOTE_SERVER")
	forex := os.Getenv("EODHD_FOREX_SERVER")
	crypto := os.Getenv("EODHD_CRYPTO_SERVER")
	demo_token := os.Getenv("EODHD_DEMO_TOKEN")

	if port == "" {
		log.Fatal("PORT is not defined in .env")
	}
	if us_trade == "" {
		log.Fatal("EODHD_US_TRADE_SERVER is not defined in .env")
	}
	if us_quote == "" {
		log.Fatal("EODHD_US_QUOTE_SERVER is not defined in .env")
	}
	if forex == "" {
		log.Fatal("EODHD_FOREX_SERVER is not defined in .env")
	}
	if crypto == "" {
		log.Fatal("EODHD_CRYPTO_SERVER is not defined in .env")
	}
	if demo_token == "" {
		log.Fatal("EODHD_DEMO_TOKEN is not defined in .env")
	}
	log.Println("Environment loaded")

	log.Println("Server spinning up on port", port)
}
