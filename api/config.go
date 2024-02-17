package api

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func envACCOUNTSID() string {
	println(godotenv.Unmarshal("cmd/.env"))
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal(err)
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func envAUTHTOKEN() string {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}

func envSERVICESID() string {
	err := godotenv.Load("cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}