package main

import (
	"encoding/json"
	"fmt"

	//	"io"
	"log"
	"net/http"

	//for use with json files
	//initializers "inventory-system/initializers"

	//routing and database libraries
	//"github.com/mattn/go-sqlite3"
	"github.com/bxcodec/faker/v4"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//
)

type User struct {
	ID          uint   `gorm:"primaryKey"`
	Username    string `gorm:"unique"`
	Password    string
	Email       string `gorm:"unique"`
	PhoneNumber string `gorm:"unique"`
}

type Inventory struct {
	ID            uint   `gorm:"primaryKey"`
	ProductName   string `gorm:"unique"`
	DateAcquired  string
	ProductAmount uint
}

var db *gorm.DB
var err error

// function to seed the database with users
func userSeeder(database *gorm.DB) error {
	//creates 1000 users with random information
	for i := 0; i < 1000; i++ {
		user := User{
			Username: faker.Username(), Password: faker.Password(), Email: faker.Email(), PhoneNumber: faker.Phonenumber(),
		}
		//creates the user in the database
		err := db.Create(&user).Error

		if err != nil {
			return err
		}
	}
	return nil
}

// function to seed the database with items
func inventorySeeder(database *gorm.DB) error {
	//creates 1000 items with random information
	for i := 0; i < 1000; i++ {
		item := Inventory{
			ProductName: faker.Word(), DateAcquired: faker.Date(), ProductAmount: uint(faker.RandomUnixTime()),
		}

		//creates the item in the database
		err := db.Create(&item).Error
		//if there is an error, return the error
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	router := mux.NewRouter()

	//opens the SQLite3 database inventory (or creates it if it doesn't exist)
	db, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})

	//in the event that the database can't be opened
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}

	//create the tables in inventory if they don't already exist
	db.AutoMigrate(&User{}, &Inventory{})

	//Seeding program to insert random data into the database for users and inventory
	err := userSeeder(db)

	if err != nil {
		panic("User seeding failed.")
	}

	//Seeding the inventory table with 1000 tuples of random data
	err = inventorySeeder(db)

	if err != nil {
		panic("Inventory seeding failed.")
	}

	/*
		In order to use the routing, be it a GET, PUT, POST, or DELETE action,
		you must go through the router designated after the slash (/). In the front end,
		you will use the url to identify the router you are looking to send the information,
		say /login, and then include the attribute, sent through JSON, that you are
		looking to input/create/edit, like {ID}. For creating users, the router is sent
		the entire JSON entity containing all information so no specific attribute is specified.

	*/

	//Creating route definitions for login page
	//routes for getting the information of the user
	router.HandleFunc("/login/{ID}", getUserWithID).Methods("GET")
	router.HandleFunc("/login/{Username}", getUserWithUsername).Methods("GET")
	router.HandleFunc("/login/{Email}", getUserWithEmail).Methods("GET")
	router.HandleFunc("/login", getAllUsers).Methods("GET")

	//routes for updating the user information
	router.HandleFunc("/login/{ID}", updateUserById).Methods("PUT")
	router.HandleFunc("/login/{Username}", updateUserByUsername).Methods("PUT")

	//routes for deleting the user based on input field (All unique attributes)
	router.HandleFunc("/login/{ID}", removeUserByID).Methods("DELETE")
	router.HandleFunc("/login/{Username}", removeUserByUsername).Methods("DELETE")
	router.HandleFunc("/login/{Email}", removeUserByEmail).Methods("DELETE")

	//Creating route definitions for registration page (just creating a new user)
	router.HandleFunc("/registration", makeUser).Methods("POST")

	//Creating route definitions for inventory page (waiting for front end to send inventory json)
	//route for creating a new item
	router.HandleFunc("/inventory", makeItem).Methods("POST")

	//routes for getting the information of items in the inventory
	router.HandleFunc("/inventory/{ID}", getItemWithID).Methods("GET")
	router.HandleFunc("/inventory/{ProductName}", getItemWithName).Methods("GET")
	router.HandleFunc("/inventory/{DateAcquired}", getFirstItemWithDate).Methods("GET")
	router.HandleFunc("/inventory/{DateAcquired}", getItemsWithDate).Methods("GET")
	router.HandleFunc("/inventory", getAllItems).Methods("GET")

	//routes for updating the information of items in the inventory
	router.HandleFunc("/inventory/{ID}", updateItemById).Methods("PUT")
	router.HandleFunc("/inventory/{ProductName}", updateItemByName).Methods("PUT")

	//routes for deleting items in the inventory
	router.HandleFunc("/inventory/{ID}", removeItemByID).Methods("DELETE")
	router.HandleFunc("/inventory/{ProductName}", removeItemByName).Methods("DELETE")

	//Creates the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Routing calls for the User table

// creates the user based on the input information
func makeUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	fmt.Printf("Created User %v\n", user)
}

// returns the specific user based on the ID
func getUserWithID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.First(&user, vars["ID"])
	fmt.Printf("Got User: %v\n", user)
}

// returns the specific user based on username
func getUserWithUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.First(&user, vars["Username"])
	fmt.Println(user)
}

// returns the specific user based on email
func getUserWithEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.First(&user, vars["Email"])
	fmt.Println(user)
}

// returns all of the users in the database
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println(allUsers)
}

// function to remove the information of the user by ID
func removeUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.Delete(&user, vars["ID"])
	fmt.Printf("Removed User: %v\n", user)
}

// function to remove the information of the user by Email
func removeUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.Delete(&user, vars["Email"])
	fmt.Println(user)
}

// function to remove the information of the user by Username
func removeUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User
	db.Delete(&user, vars["ID"])
	fmt.Println(user)
}

// function to update the information of the user by ID
func updateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User

	//find the user in the database
	db.First(&user, vars["ID"])

	//updates the user in the database
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)

	fmt.Printf("Updated User: %v\n", user)
}

// function to update the information of the user by
func updateUserByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user User

	//find the user in the database
	db.First(&user, vars["Username"])

	//updates the user in the database
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)

	fmt.Println(user)
}

//the routing for the inventory table

// function creates a new item in the Inventory table
func makeItem(w http.ResponseWriter, r *http.Request) {
	var item Inventory
	json.NewDecoder(r.Body).Decode(&item)
	db.Create(&item)
	fmt.Printf("Created Item: %v\n", item)
}

// fuction retrieves the information of an item based on its ID
func getItemWithID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ID"])
	fmt.Printf("Got Item: %v\n", item)
}

// function retrieves the information of an item based on its name
func getItemWithName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ProductName"])
	fmt.Println(item)
}

// function retrieves multiple item information based on its date (since it isn't unique)
func getItemsWithDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var items []Inventory
	db.First(&items, vars["DateAcquired"])
	fmt.Println(items)
}

// function retrieves the first item information based on its date
func getFirstItemWithDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["DateAcquired"])
	fmt.Println(item)
}

// function gets the information of all items in the Inventory table
func getAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []Inventory
	db.Find(&items)
	// fmt.Printf("test")
	json.NewEncoder(w).Encode(items)

}

// function removes the tuple that contains the input ID
func removeItemByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.Delete(&item, vars["ID"])
	fmt.Printf("Removed Item: %v\n", item)
}

// function removes the tuple by the product name
func removeItemByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.Delete(&item, vars["ProductName"])
	fmt.Println(item)
}

// function updates the item based on the ID
func updateItemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ID"])
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	fmt.Printf("Updated Item: %v\n", item)
}

// function updates the item based on the item Name
func updateItemByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ProductName"])
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	fmt.Println(item)
}
