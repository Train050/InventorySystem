package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	//"io"
	"log"
	"net/http"

	//for use with json files
	//initializers "inventory-system/initializers"

	//"github.com/mattn/go-sqlite3"
	"github.com/bxcodec/faker/v4"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User Database
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique; not null"`
	Password     string `gorm:"not null"`
	Email        string `gorm:"not null"`
	PhoneNumber  string `gorm:"not null"`
	HashPassword string
}

// Inventory Database
type Inventory struct {
	ID            uint   `gorm:"primaryKey"`
	ProductName   string `gorm:"not null"`
	DateAcquired  string `gorm:"not null"`
	ProductAmount uint   `gorm:"not null"`
}

var db *gorm.DB
var err error

// function to seed the database with users
/*func userSeeder(database *gorm.DB, entries int) error {
	//creates users with random information based on the number of entries specified
	for i := 0; i < entries; i++ {
		user := User{
			Username: faker.Username(), Password: faker.Password(), Email: faker.Email(), PhoneNumber: faker.Phonenumber(), HashPassword: faker.Password(),
		}
		//creates the user in the database
		err := db.Create(&user).Error

		//if there is an error, return the error
		if err != nil {
			return err
		}
	}
	return nil
}

// function to seed the database with items
func inventorySeeder(database *gorm.DB, entries int) error {
	//creates items with random information based on the number of entries specified
	for i := 0; i < entries; i++ {
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
*/
// Checks the authorization of users requesting information
func userAuthenticator(w http.ResponseWriter, r *http.Request) {

	//first checks for authorization within the json header
	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	//decodes the authorization
	encodeAuth := strings.TrimPrefix(auth, "Basic ")
	decodeAuth, err := base64.StdEncoding.DecodeString(encodeAuth)

	//if there is an error decoding the authorization, return an error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization failed"))
		return
	}

	//splits the authorization into username and password
	authArray := strings.Split(string(decodeAuth), ":")
	if len(authArray) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Authorization failed"))
		return
	}

	//assigns the username and password to variables
	username := authArray[0]
	password := authArray[1]

	//if the authorization is not empty, then it checks the database for the user
	var user User
	err = db.Where("Username = ?", username).First(&user).Error

	//if the user is not found, then it returns an error
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username or Password not found"))
		return
	}

	//if the user is found, then it checks the password hash
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Username or Password not found"))
		return
	}

	//Creating JWT token for the user (lasts for 12 hours upon being granted)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"password":   user.Password,
		"email":      user.Email,
		"phone":      user.PhoneNumber,
		"expiration": time.Now().Add(time.Hour * 12).Unix(),
	})

	//signing the token with the secret key
	tokenString, err := jwtToken.SignedString([]byte("VerySecretKey"))

	//if there is an error signing the token, return an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error signing the token"))
		return
	}

	//sending the token to the user
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

	//if there is an error sending the token, return an error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error sending the token in json"))
		return
	}

	//if the user is authenticated, then it returns a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User authenticated"))
}

// Function to check the validity of the token and return the token
func checkToken(inputToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(inputToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {

		//checks if the token is valid
		verified := token.Method.(*jwt.SigningMethodHMAC)
		if (verified == nil) || (verified.Alg() != "HS256") {
			return nil, fmt.Errorf("Error in token")
		}

		//secret key to sign the token
		passKey := []byte("VerySecretKey")
		return passKey, nil
	})

	//if there is an error parsing the token, return an error
	if err != nil {
		return nil, err
	}

	//if the token is not valid, return an error
	if !token.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}

	return token, nil
}

/*
	The authentication middleware checks the validity of the token and returns the token if it is valid.
	It is used to check the validity of the token before allowing the user to access the information. This is
	used to prevent unauthorized access of the information. The token is valid for 12 hours. The password
	is hashed and compared to the password in the database. If the password is correct, then the user is
	authenticated and a token is created for the user. The token is then sent to the user. After, the token is
	checked for validity before allowing the user to access the information.
*/

// Still need to implement the authenticaiton middleware into the routing to check creditentials before allowing access and
// modification of the information
func main() {
	router := mux.NewRouter()
	//authRouter := router.PathPrefix("/api").Subrouter()

	//opens the SQLite3 database inventory (or creates it if it doesn't exist)
	db, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})

	//in the event that the database can't be opened
	if err != nil {
		log.Fatal("Failed to connect to the database.")
	}

	//seeds the database with users and items
	userSeeder(db, 100)
	inventorySeeder(db, 100)

	//create the tables in inventory if they don't already exist
	db.AutoMigrate(&User{}, &Inventory{})

	/*
		In order to use the routing, be it a GET, PUT, POST, or DELETE action,
		you must go through the router designated after the slash (/). In the front end,
		you will use the url to identify the router you are looking to send the information,
		say /login, and then include the attribute, sent through JSON, that you are
		looking to input/create/edit, like {ID}. For creating users, the router is sent
		the entire JSON entity containing all information so no specific attribute is specified.

		Upon the implmentation of the authentication middleware, the user will have their password
		checked and a token will be created for them. The token will be sent to the user and they
		will have to send the token back to the server to be checked for validity. This should not
		affect the routes for front end.
	*/

	//Creating route definitions for login page

	//routes for getting the information of the user
	router.HandleFunc("/login/{ID}", getUserWithID).Methods("GET")
	router.HandleFunc("/login/{Username}", getUserWithUsername).Methods("GET")
	router.HandleFunc("/login/{Email}", getUserWithEmail).Methods("GET")
	router.HandleFunc("/login/{PhoneNumber}", getUserWithEmail).Methods("GET")
	router.HandleFunc("/login", getAllUsers).Methods("GET")

	router.HandleFunc("/api/login/{ID}", getUserWithID).Methods("GET")
	router.HandleFunc("/api/login/{Username}", getUserWithUsername).Methods("GET")
	router.HandleFunc("/api/login/{Email}", getUserWithEmail).Methods("GET")
	router.HandleFunc("/api/login/{PhoneNumber}", getUserWithEmail).Methods("GET")
	router.HandleFunc("/api/login", getAllUsers).Methods("GET")

	//routes for updating the user information
	router.HandleFunc("/login/{ID}", updateUserById).Methods("PUT")
	router.HandleFunc("/login/{Username}", updateUserByUsername).Methods("PUT")

	router.HandleFunc("/api/login/{ID}", updateUserById).Methods("PUT")
	router.HandleFunc("/api/login/{Username}", updateUserByUsername).Methods("PUT")

	//routes for deleting the user based on input field (All unique attributes)
	router.HandleFunc("/login/{ID}", removeUserByID).Methods("DELETE")
	router.HandleFunc("/login/{Username}", removeUserByUsername).Methods("DELETE")
	router.HandleFunc("/login/{Email}", removeUserByEmail).Methods("DELETE")
	//router.HandleFunc("/login/removeAll", removeAllUsers).Methods("DELETE")

	router.HandleFunc("/api/login/{ID}", removeUserByID).Methods("DELETE")
	router.HandleFunc("/api/login/{Username}", removeUserByUsername).Methods("DELETE")
	router.HandleFunc("/api/login/{Email}", removeUserByEmail).Methods("DELETE")

	//Creating route definitions for registration page (just creating a new user)
	router.HandleFunc("http://localhost:4200/register", makeUser).Methods("POST")
	router.HandleFunc("http://localhost:8080/register", makeUser).Methods("POST")

	router.HandleFunc("http://localhost:4200/register", makeUser).Methods("POST")
	router.HandleFunc("http://localhost:8080/register", makeUser).Methods("POST")

	//Creating route definitions for inventory page (waiting for front end to send inventory json)
	//route for creating a new item
	router.HandleFunc("http://localhost:4200/inventory", makeItem).Methods("POST")

	//routes for getting the information of items in the inventory
	/*
		router.HandleFunc("/inventory/{ID}", checkToken(getItemWithID)).Methods("GET")
		router.HandleFunc("/inventory/{ProductName}", checkToken(getItemWithName)).Methods("GET")
		router.HandleFunc("/inventory/{DateAcquired}", checkToken(getFirstItemWithDate)).Methods("GET")
		router.HandleFunc("/inventory/{DateAcquired}", checkToken(getItemsWithDate)).Methods("GET")
		router.HandleFunc("/inventory", checkToken(getAllItems)).Methods("GET")
	*/

	//routes for getting the information of items in the inventory
	router.HandleFunc("/inventory/{ID}", getItemWithID).Methods("GET")
	router.HandleFunc("/inventory/{ProductName}", getItemWithName).Methods("GET")
	//router.HandleFunc("/inventory/{DateAcquired}", getFirstItemWithDate).Methods("GET")
	router.HandleFunc("/inventory/{DateAcquired}", getItemsWithDate).Methods("GET")
	router.HandleFunc("/inventory", getAllItems).Methods("GET")

	//routes for updating the information of items in the inventory
	router.HandleFunc("/inventory/{ID}", updateItemById).Methods("PUT")
	router.HandleFunc("/inventory/{ProductName}", updateItemByName).Methods("PUT")

	//routes for deleting items in the inventory
	router.HandleFunc("/inventory/{ID}", removeItemByID).Methods("DELETE")
	router.HandleFunc("/inventory/{ProductName}", removeItemByName).Methods("DELETE")
	//router.HandleFunc("/inventory/removeAll", removeAllItems).Methods("DELETE")

	//Creates the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Routing calls for the User table
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")

}

// creates the user based on the input information
func makeUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user User
	decode := json.NewDecoder(r.Body).Decode(&user)
	if decode != nil {
		http.Error(w, decode.Error(), http.StatusBadRequest)
		return
	}
	create := db.Create(&user)
	if create.Error != nil {
		http.Error(w, create.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// returns the specific user based on the ID
func getUserWithID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var user User
	err := db.Where("ID = ?", vars["ID"]).First(&user)
	if err != nil {
		log.Fatalf("No with that ID found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(user)
	if otherErr != nil {
		log.Fatalf("Couldn't encode user")
		return
	}
}

// returns the specific user based on username
func getUserWithUsername(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var user User
	err := db.Where("Username = ?", vars["Username"]).First(&user)
	if err != nil {
		log.Fatalf("No with that username found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(user)
	if otherErr != nil {
		log.Fatalf("Couldn't encode user.")
		return
	}

	//fmt.Println(user)
}

// returns the specific user based on phone number
func getUserWithPhoneNumber(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var user User
	err := db.Where("PhoneNumber = ?", vars["PhoneNumber"]).First(&user)
	if err != nil {
		log.Fatalf("No with that phone number found")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(user)
	if otherErr != nil {
		log.Fatalf("Couldn't encode user")
		return
	}

	//fmt.Println(user)
}

// returns the specific user based on email
func getUserWithEmail(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var user User
	err := db.Where("Email = ?", vars["Email"]).First(&user)
	if err != nil {
		log.Fatalf("No user with that email found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(user)
	if otherErr != nil {
		log.Fatalf("Couldn't encode user.")
		return
	}
}

// returns all of the users in the database
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var allUsers []User
	db.Find(&allUsers)
	fmt.Println(allUsers)

	// Setup the backend response so we can return the items in a JSON object:
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := allUsers
	jsonResp, err := json.Marshal(resp)

	if err != nil {
		log.Fatalf("getAllUsers failed to JSON marshal. Error: %s", err)
	}
	w.Write(jsonResp)
}

// function to remove the information of the user by ID
func removeUserByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	err := db.Where("ID = ?", vars["ID"]).Delete(&User{})

	if err.Error != nil {
		fmt.Println("User wasn't deleted")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removed the user")
}

// function to remove the information of the user by Email
func removeUserByEmail(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	err := db.Where("Email = ?", vars["Email"]).Delete(&User{})

	if err.Error != nil {
		fmt.Println("User wasn't deleted")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removed the user")
}

// function to remove the information of the user by Username
func removeUserByUsername(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	err := db.Where("Username = ?", vars["Username"]).Delete(&User{})

	if err.Error != nil {
		fmt.Println("User wasn't deleted")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removed the user")
}

// function to update the information of the user by ID
func updateUserById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
	vars := mux.Vars(r)
	var user User

	//find the user in the database
	db.First(&user, vars["Username"])

	//updates the user in the database
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)

	fmt.Println(user)
}

func removeAllUsers(db *gorm.DB) {
	//removes all users, regardless of their role (used for redoing the database)
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&User{})
}

//the routing for the inventory table

// function creates a new item in the Inventory table
func makeItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var item Inventory
	json.NewDecoder(r.Body).Decode(&item)
	db.Create(&item)
	fmt.Printf("Created Item: %v\n", item)
}

// fuction retrieves the information of an item based on its ID
func getItemWithID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	err := db.Where("ID = ?", vars["ID"]).First(&item)
	if err != nil {
		log.Fatalf("No item with that ID found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(item)
	if otherErr != nil {
		log.Fatalf("Couldn't encode the Item")
		return
	}
}

// function retrieves the information of an item based on its name
func getItemWithName(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	err := db.Where("ProductName = ?", vars["ProductName"]).First(&item)
	if err != nil {
		log.Fatalf("No item with that name found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherErr := json.NewEncoder(w).Encode(item)
	if otherErr != nil {
		log.Fatalf("Couldn't encode the Item.")
		return
	}
}

// function retrieves multiple item information based on its date (since it isn't unique)
func getItemsWithDate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	err := db.Where("DateAcquired = ?", vars["DateAcquired"]).First(&item)
	if err != nil {
		log.Fatalf("No item with that date found.")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	otherError := json.NewEncoder(w).Encode(item)
	if otherError != nil {
		log.Fatalf("Couldn't encode the Item.")
		return
	}
}

// function retrieves the first item information based on its date
func getFirstItemWithDate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["DateAcquired"])
	fmt.Println(item)
}

// function gets the information of all items in the Inventory table
func getAllItems(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var items []Inventory
	db.Find(&items) // select * from inventory;

	fmt.Println("getAllItems: ")
	fmt.Println(items)

	// Setup the backend response so we can return the items in a JSON object:
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := items
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("getAllItems failed to JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

// function removes the tuple that contains the input ID
func removeItemByID(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	err := db.Where("ID = ?", vars["ID"]).Delete(&item)

	if err.Error != nil {
		fmt.Println("Item wasn't deleted")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removed the Item")
}

// function removes the tuple by the product name
func removeItemByName(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	err := db.Where("Name = ?", vars["Name"]).Delete(&item)

	if err.Error != nil {
		fmt.Println("Item wasn't deleted")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Removed the Item")
}

// function updates the item based on the ID
func updateItemById(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ID"])
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	fmt.Printf("Updated Item: %v\n", item)
}

// function updates the item based on the item Name
func updateItemByName(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	vars := mux.Vars(r)
	var item Inventory
	db.First(&item, vars["ProductName"])
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	fmt.Println(item)
}

func removeAllItems(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Inventory{})
}
