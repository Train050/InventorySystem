package main

import (
	"testing"
)

// User tests
func TestInsertUser(t *testing.T) {
	result := 1 + 3

	if result != 4 {
		t.Errorf("Add(1, 3) Failed")
	} else {
		t.Logf("Add(1, 3) PASSED")
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
