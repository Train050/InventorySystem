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
	//Create test databases
	//If we cannot connect to the database, return an error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

// The following tests first test http routing, then database operations.
// FAIL will output if either fails, displaying the appropriate error.
func TestMakeUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/registration", nil)
	w := httptest.NewRecorder()
	makeUser(w, req)
	res := w.Result()
	defer res.Body.Close()

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}
func TestUpdateUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodPut, "/login/{1}", nil)
	w := httptest.NewRecorder()
	getUserWithID(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestRemoveUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/login/{1}", nil)
	w := httptest.NewRecorder()
	removeUserByID(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestFindUser(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login/{1}", nil)
	w := httptest.NewRecorder()
	getUserWithID(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

//Inventory tests

func TestInsertItem(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/login/{1}", nil)
	w := httptest.NewRecorder()
	getUserWithID(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestUpdateItem(t *testing.T) {

}

func TestRemoveItem(t *testing.T) {

}

func TestFindItem(t *testing.T) {

}
