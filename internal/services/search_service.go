package services

import (
  "ammerzon.com/golang-rest/internal/config"
  "ammerzon.com/golang-rest/internal/models"
  "database/sql"
  "fmt"
  "github.com/gorilla/mux"
  log "github.com/sirupsen/logrus"
  "net/http"
  "strconv"
)

type SearchService struct {
  Router *mux.Router
  DB     *sql.DB
}

func (ss *SearchService) Initialize(conf *config.Config, r *mux.Router) {
  connectionString :=
    fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseName)

  var err error
  ss.DB, err = sql.Open("postgres", connectionString)
  ss.Router = r
  if err != nil {
    log.Fatal(err)
  }

  ss.initializeRoutes()
}

func (ss *SearchService) initializeRoutes() {
  ss.Router.HandleFunc("/search/products", ss.searchProducts).Queries("q", "{q}").Methods("GET")
}

func (ss *SearchService) searchProducts(w http.ResponseWriter, r *http.Request) {
  query := "%" + r.FormValue("q") + "%"
  count, _ := strconv.Atoi(r.FormValue("count"))
  start, _ := strconv.Atoi(r.FormValue("start"))

  if count > 10 || count < 1 {
    count = 10
  }
  if start < 0 {
    start = 0
  }

  products, err := models.SearchProducts(ss.DB, query, start, count)
  if err != nil {
    respondWithError(w, http.StatusInternalServerError, err.Error())
    return
  }

  respondWithJSON(w, http.StatusOK, products)
}
