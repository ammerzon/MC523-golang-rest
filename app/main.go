package main

import (
	"ammerzon.com/golang-rest/config"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Info("❓ No .env file found. Default values will be used instead if no environment variables are specified!")
	}
}

func main() {
	a := App{}
	conf := config.GetConfig()

	a.Initialize(conf)
	a.Run(":8010")
}
