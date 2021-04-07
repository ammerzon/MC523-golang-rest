package services

import (
	"ammerzon.com/golang-rest/internal/config"
	"ammerzon.com/golang-rest/internal/models"
	"database/sql"
	"encoding/json"
	"fmt"
  "math"
  "net/http"
	"strconv"
  "strings"

  "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type ProductService struct {
	Router *mux.Router
	DB     *sql.DB
}

func (ps *ProductService) Initialize(conf *config.Config, r *mux.Router) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DatabaseUsername, conf.DatabasePassword, conf.DatabaseName)

	var err error
	ps.DB, err = sql.Open("postgres", connectionString)
	ps.Router = r
	if err != nil {
		log.Fatal(err)
	}

	ps.initializeRoutes()
}

func (ps *ProductService) initializeRoutes() {
  ps.Router.HandleFunc("/products", ps.getProducts).Methods("GET")
	ps.Router.HandleFunc("/product", ps.createProduct).Methods("POST")
	ps.Router.HandleFunc("/product/{id:[0-9]+}", ps.getProduct).Methods("GET")
	ps.Router.HandleFunc("/product/{id:[0-9]+}", ps.updateProduct).Methods("PUT")
	ps.Router.HandleFunc("/product/{id:[0-9]+}", ps.deleteProduct).Methods("DELETE")
}

func (ps *ProductService) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.GetProduct(ps.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (ps *ProductService) createProduct(w http.ResponseWriter, r *http.Request) {
	var p models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.CreateProduct(ps.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (ps *ProductService) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var p models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p.ID = id

	if err := p.UpdateProduct(ps.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (ps *ProductService) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	p := models.Product{ID: id}
	if err := p.DeleteProduct(ps.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (ps *ProductService) getProducts(w http.ResponseWriter, r *http.Request) {
  sort := strings.ToLower(r.FormValue("s"))
  lowerBound, lbErr := strconv.ParseFloat(r.FormValue("lb"), 64)
  upperBound, ubErr := strconv.ParseFloat(r.FormValue("ub"), 64)
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if sort == "" || !(sort == "name" || sort == "price") {
	  sort = "id"
  }
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	var products []models.Product
	var err error
	if lbErr == nil || ubErr == nil {
	  if lbErr != nil {
	    lowerBound = 0.0
    } else if ubErr != nil {
      upperBound = math.MaxFloat64
    }
    products, err = models.GetProductsWithBounds(ps.DB, lowerBound, upperBound, sort, start, count)
  } else {
    products, err = models.GetProducts(ps.DB, sort, start, count)
  }

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
