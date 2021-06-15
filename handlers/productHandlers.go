package handlers

import (
	. "../helpers"
	. "../models"
	"encoding/json"
	"net/http"
	"time"
)

//we need an object because I don't add any DB
var productStore = make(map[int]Product)
var id int = 0

//HTTP POST - api/products
func postProductHandler(w http.ResponseWriter, r *http.Request){
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product) // r icindeki body'i al, decode et ve onu yukarida olusturdugum product nesnesinin hafiza adresine aktar.
	CheckError(err)
	product.CreatedOn = time.Now()
	product.ID = id
	productStore[id] = product

	data, err := json.Marshal(product) //object to json
	CheckError(err)

	w.Header().Set("Content-Type", "application/jspn") //gelen verinin content tipini kur
	w.WriteHeader(http.StatusCreated) //islem basarili
	w.Write(data) // olusturdugum veriyi yaz

}

//HTTP GET - api/products
func getProductsHandler(w http.ResponseWriter, r *http.Request){

}

//HTTP GET -api/products/{id}
func getProductHandler (w http.ResponseWriter, r *http.Request){

}

//HTTP PUT - api/products/{id]
func putProductHandler(w http.ResponseWriter, r *http.Request){

}

//HTTP DELETE - api/products/{id}
func deleteProductHandler(w http.ResponseWriter, r *http.Request)  {

}
