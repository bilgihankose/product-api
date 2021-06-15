package handlers

import (
	. "../helpers"
	. "../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

//we need an object because I don't add any DB
var productStore = make(map[string]Product)
var id = 0

//HTTP POST - api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request){ //handlerlar public olmasi icin buyuk harfle baslamali
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product) // r icindeki body'i al, decode et ve onu yukarida olusturdugum product nesnesinin hafiza adresine aktar.
	CheckError(err)
	product.CreatedOn = time.Now()
	product.ID = id
	key, _ := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product) //object to json
	CheckError(err)

	w.Header().Set("Content-Type", "application/jspn") //gelen verinin content tipini kur
	w.WriteHeader(http.StatusCreated) //islem basarili
	w.Write(data) // olusturdugum veriyi yaz

}

//HTTP GET - api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request){
	var products []Product
	for _, product := range productStore {
		products = append(products,product)
	}

	data, err := json.Marshal(products) //object to json
	CheckError(err)

	w.Header().Set("Content-Type", "application/jspn") //gelen verinin content tipini kur
	w.WriteHeader(http.StatusOK) //islem basarili
	w.Write(data) // olusturdugum veriyi yaz
}

//HTTP GET -api/products/{id}
func GetProductHandler (w http.ResponseWriter, r *http.Request){
 var product Product //Product turunden tek bir product tanimla
 vars := mux.Vars(r) //request icindeki variablelari al
 key, _ := strconv.Atoi(vars["id"]) //vars icindeki id'yi al

 for _, prd := range productStore{
 	if prd.ID == key {
 		product = prd
	}
 }
	data, err := json.Marshal(product) //object to json
	CheckError(err)

	w.Header().Set("Content-Type", "application/jspn") //gelen verinin content tipini kur
	w.WriteHeader(http.StatusOK) //islem basarili
	w.Write(data) // olusturdugum veriyi yaz


}

//HTTP PUT - api/products/{id]
func PutProductHandler(w http.ResponseWriter, r *http.Request){
var err error

vars := mux.Vars(r)
key := vars["id"] //belli bi unique degere gore silinmesi gerekir.

var productUpdate Product
err = json.NewDecoder(r.Body).Decode(&productUpdate) //request uzerindeki bodyden gelen veriyo decode ettikten sonra o veriyi nesneye donustturuyoruz
CheckError(err)

if _, ok := productStore[key]; ok { //disaridan olusturdugum key degeri productStore icinde varsa ve donen cevap ok ise (map ile alakali kullanim turu) biz bu veriyi guncelleyebiliriz.
	productUpdate.ID, _ = strconv.Atoi(key)
	productUpdate.ChangedOn = time.Now()
	delete(productStore, key)
	productStore[key] = productUpdate
} else {
	log.Println("Deger bulunamadi")
}
	w.WriteHeader(http.StatusOK)

}

//HTTP DELETE - api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request)  {
 vars := mux.Vars(r)
 key := vars["id"]

 if _, ok := productStore[key]; ok {
 	delete(productStore, key)
 } else {
	 log.Println("Deger bulunamadi")
 }
	w.WriteHeader(http.StatusOK)

}
