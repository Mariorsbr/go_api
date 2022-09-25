package main

import (
	"testing" 
)

type user struct {
    ID     string  `json:"id"`
    Username  string  `json:"username"`
    Password string  `json:"password"`
}

func TestGetUser(t *testing.T) {
    
	got := getuser()
	want := "List of Users"

	if got != want {
		t.Errorf("Test failed")
	}
}