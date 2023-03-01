package main

import (
	"fmt"
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

	//Run tests
	return m.Run(), nil
}

// Testing functions
// User tests
func TestInsertUser(t *testing.T) {
	result := 1 + 3
	if result != 4 {
		t.Errorf("Add failed")
	} else {
		t.Logf("Add passed")
	}
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
