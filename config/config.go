package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port                string
	TokenDuration		int
	Host				string
	FEPort              string
	DbUsername          string
	DbPassword          string
	DbHost              string
	DbPort              string
	DbName              string
	Issuer              string
	SendEmailIdentity   string
	SendEmailUsername   string
	SendEmailPassword   string
	SendEmailHost       string
	SendEmailPort       string
	VerifSecret         string
	AccessSecret        string
	RefreshSecret       string
	ResetPasswordSecret string
	RajaOngkirApiKey    string
	HashCost            int
	GracefulPeriod      int
	DbDsn               string  
}

func Init(log *logrus.Logger) *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	hashCost, err := strconv.Atoi(os.Getenv("HASH_COST"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": "HASH_COST must be integer",
		}).Fatal("error loading .env file")
	}

	gracefulPeriod, err := strconv.Atoi(os.Getenv("GRACEFUL_PERIOD"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": "GRACEFUL_PERIOD must be integer",
		}).Fatal("error loading .env file")
	}

	tokenDUration, err := strconv.Atoi(os.Getenv("TOKEN_DURATION"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": "TOKEN_DURATION must be integer",
		}).Fatal("error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USERNAME"),    // MySQL username
		os.Getenv("DB_PASSWORD"),    // MySQL password
		os.Getenv("DB_HOST"),        // MySQL host
		os.Getenv("DB_PORT"),        // MySQL port
		os.Getenv("DB_NAME"),        // MySQL database name
	)

	return &Config{
		Port:                os.Getenv("BE_PORT"),
		TokenDuration:		 tokenDUration,
		Host:				 os.Getenv("HOST"),
		FEPort:              os.Getenv("FE_PORT"),
		DbUsername:          os.Getenv("DB_USERNAME"),
		DbPassword:          os.Getenv("DB_PASSWORD"),
		DbHost:              os.Getenv("DB_HOST"),
		DbPort:              os.Getenv("DB_PORT"),
		DbName:              os.Getenv("DB_NAME"),
		Issuer:              os.Getenv("ISSUER"),
		SendEmailIdentity:   os.Getenv("SEND_EMAIL_IDENTITY"),
		SendEmailUsername:   os.Getenv("SEND_EMAIL_USERNAME"),
		SendEmailPassword:   os.Getenv("SEND_EMAIL_PASSWORD"),
		SendEmailHost:       os.Getenv("SEND_EMAIL_HOST"),
		SendEmailPort:       os.Getenv("SEND_EMAIL_PORT"),
		VerifSecret:         os.Getenv("VERIFICATION_CODE_SECRET_KEY"),
		AccessSecret:        os.Getenv("ACCESS_TOKEN_SECRET_KEY"),
		RefreshSecret:       os.Getenv("REFRESH_TOKEN_SECRET_KEY"),
		ResetPasswordSecret: os.Getenv("RESET_PASSWORD_SECRET_KEY"),
		RajaOngkirApiKey:    os.Getenv("RAJA_ONGKIR_API_KEY"),
		HashCost:            hashCost,
		GracefulPeriod:      gracefulPeriod,
		DbDsn:               dsn,  // Store the DSN in the Config struct
	}
}
