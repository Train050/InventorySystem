package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//for use with json files
	//initializers "inventory-system/initializers"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model 
	ID uint
	UserName string
	password string 
	email string
	phoneNumber string
}

type Inventory struct {
	gorm.Model
	ID uint 
	productName string
	dateAcquired string
	productAmount uint
}

/*
func addUser(id uint, username string, password string, email string, phoneNumber string) {

}


func addInventory(id uint, product string, date string, amount uint) {

}
*/

func handleInventoryGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	product := vars["productName"]
	date := vars["dateAquired"]
	amount := vars["productAmount"]

	fmt.Fprintf(w, "Requested info: %s id, %s product, %s acquisition date, %s amount of product.", id, product, date, amount)
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	password := vars["password"]
	email := vars["email"]
	phoneNumber := vars["phone number"]

	fmt.Fprintf(w, "Requested info: %s username with %s password, %s email, %s phone number\n", username, password, email, phoneNumber)
}

func main() {
	router := mux.NewRouter()

	var err error
	var DB *gorm.DB

	DB, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}
	
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Inventory{})

	//this is the route for creating a new user into the database
	router.HandleFunc("/login-page", func(w http.ResponseWriter, r *http.Request) {
		//creating a new user variable
		var user User
		
		//creating the new user in the database
		newUser := DB.Create(&user)

		if newUser.Error != nil{
			http.Error(w, newUser.Error.Error(), http.StatusInternalServerError)
			return
		}

		//returning the created user to frontend
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(&user)

		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

//this is the route for creating a new user into the database
	router.HandleFunc("/inventory-home-page", func(w http.ResponseWriter, r *http.Request) {
		//creating a new user variable
		var inventory Inventory
		
		//creating the new user in the database
		newInventory := DB.Create(&inventory)

		if newInventory.Error != nil{
			http.Error(w, newInventory.Error.Error(), http.StatusInternalServerError)
			return
		}

		//returning the created user to frontend
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(&newInventory)

		if (err != nil) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
