package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
}

//database used map
var database = make(map[string]Product)

func SetJSONResp(res http.ResponseWriter, message []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func main() {
	//init db
	database["001"] = Product{ID: "001", Name: "Samsung Galaxy S1", Quantity: "10"}
	database["002"] = Product{ID: "002", Name: "Realme Pro 2", Quantity: "10"}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		message := []byte(`{"message": "Server is ready"}`)
		SetJSONResp(res, message, http.StatusOK)
	})

	//get list of product
	http.HandleFunc("/get-products", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "GET" {
			message := []byte(`{"message": "Invalid http method"}`)
			SetJSONResp(res, message, http.StatusMethodNotAllowed)
			return
		}

		//append product
		var products []Product
		for _, product := range database {
			products = append(products, product)
		}

		productJSON, err := json.Marshal(&products)
		if err != nil {
			message := []byte(`{"message": "Error while marshalling data"}`)
			//status 500
			SetJSONResp(res, message, http.StatusInternalServerError)
			return
		}
		SetJSONResp(res, productJSON, http.StatusOK)
	})

	//add product
	http.HandleFunc("/add-product", func(res http.ResponseWriter, req *http.Request) {

		if req.Method != "POST" {
			message := []byte(`{"message": "Invalid http method"}`)
			SetJSONResp(res, message, http.StatusMethodNotAllowed)
			return
		}

		//append product
		var products []Product
		for _, product := range database {
			products = append(products, product)
		}

		productJSON, err := json.Marshal(&products)
		if err != nil {
			message := []byte(`{"message": "Error while marshalling data"}`)
			//status 500
			SetJSONResp(res, message, http.StatusInternalServerError)
			return
		}
		SetJSONResp(res, productJSON, http.StatusOK)
	})

	//start server
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Printf("Server error, cause:%+v\n", err)
		os.Exit(1)
	}
}
