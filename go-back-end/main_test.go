package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup test database
func TestMain(m *testing.M) {
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	//Create test database
	//If we cannot connect to the database, return an error
	db, err = gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		return -1, fmt.Errorf("could not connect to the database: %w", err)
	}
	db.AutoMigrate(&User{}, &Inventory{})

	//db.Create(&User{ID: 1, Username: "Nicholas", Password: "password", Email: "test@example.com", PhoneNumber: "917-613-XXXX"})

	//Run tests
	return m.Run(), nil
}

// Testing functions
// Backend user tests

// Add a user
func TestMakeUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/registration", nil)
	w := httptest.NewRecorder()
	makeUser(w, req)
	res := w.Result()
	defer res.Body.Close()
	/**
	data, err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(data))

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ABC" {
		fmt.Println(string(data))
		t.Errorf("expected ABC got %v", string(data))
	}
	*/
}
func TestUpdateUser(t *testing.T) {

}

func TestRemoveUser(t *testing.T) {

}

func TestFindUser(t *testing.T) {

}

//Inventory tests

func TestInsertItem(t *testing.T) {

}

func TestUpdateItem(t *testing.T) {

}

func TestRemoveItem(t *testing.T) {

}

func TestFindItem(t *testing.T) {

}
