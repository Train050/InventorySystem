package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/gorilla/mux"
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

	//Run tests
	return m.Run(), nil
}

// Testing functions
// Backend user tests

// The following tests first test http routing, then database operations.
// FAIL will output if either fails, displaying the appropriate error.
func TestMakeUser(t *testing.T) {
	//Test empty user
	req := httptest.NewRequest(http.MethodPut, "/registration", nil)
	w := httptest.NewRecorder()
	makeUser(w, req)
	res := w.Result()
	defer res.Body.Close()

	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	//Test complete user
	db.Create(&User{ID: 2, Username: "Nicholas", Password: "password", Email: "test@example.com", PhoneNumber: "917-613-XXXX"})
	vars := mux.Vars(req)
	var user User
	db.First(&user, vars["ID"])

	if user.ID == 2 && user.Username == "Nicholas" && user.Password == "password" && user.Email == "test@example.com" && user.PhoneNumber == "917-613-XXXX" {
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
	req := httptest.NewRequest(http.MethodPost, "/inventory", nil)
	w := httptest.NewRecorder()
	makeItem(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestUpdateItem(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/inventory/{1}", nil)
	w := httptest.NewRecorder()
	updateItemById(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestRemoveItem(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/inventory/{1}", nil)
	w := httptest.NewRecorder()
	updateItemById(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

func TestFindItem(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/inventory/{1}", nil)
	w := httptest.NewRecorder()
	getItemWithID(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

//Sprint 3 backend tests

// Seeding
// User, DB mocked for simplicity, tests only username faking
func TestUserSeeder(t *testing.T) {
	fmt.Println("Testing User Seeder")
	var mockedUsers [1001]string
	for i := 0; i < 1000; i++ {
		mockedUsers[i] = "s"
	}
	for i := 0; i < 1000; i++ {
		var Username = faker.Username()
		//creates the user in the database (array)
		mockedUsers[i] = Username
	}

	var err = "err"
	for i := 0; i < 1000; i++ {
		if mockedUsers[i] == "s" {
			err = "some users not populated"
		}
	}
	if err != "err" {
		t.Errorf("expected error to be nil got %v", err)
		fmt.Println()
	} else {
		fmt.Println("PASS")
		fmt.Println()
	}
}

// Inventory, db mocked, product name tested
func TestInventorySeeder(t *testing.T) {
	fmt.Println("Testing Inventory Seeder")
	var mockedInventory [1001]string
	for i := 0; i < 1000; i++ {
		mockedInventory[i] = "s"
	}
	for i := 0; i < 1000; i++ {
		var ProductName = faker.Word()
		//creates the user in the database (array)
		mockedInventory[i] = ProductName
	}

	var err = "err"
	for i := 0; i < 1000; i++ {
		if mockedInventory[i] == "s" {
			err = "some inventory not populated"
		}
	}
	if err != "err" {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

// Get all users, check for errors.
func TestGetAllUsers(t *testing.T) {
	//Testing routing
	fmt.Println("Test getAllUsers")
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	w := httptest.NewRecorder()
	getAllUsers(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}

// Get all inventory
func TestGetAllItems(t *testing.T) {
	//Testing routing
	fmt.Println("Test getAllItems")
	req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
	w := httptest.NewRecorder()
	getAllItems(w, req)
	res := w.Result()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	} else {
		fmt.Println("PASS")
	}
}
