package main

import (
	"ammerzon.com/golang-rest/internal/config"
	"ammerzon.com/golang-rest/internal/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Info("❓ No .env file found. Default values will be used instead if no environment variables are specified!")
	}
}

func main() {
	r := mux.NewRouter()
	ps := services.ProductService{}
  ss := services.SearchService{}
	conf := config.GetConfig()

	ps.Initialize(conf, r)
  ss.Initialize(conf, r)
	log.Fatal(http.ListenAndServe(":8010", r))
}
