package models

import (
	"testing" 
)


func TestPostUser(t *testing.T) {
    
	got := PostUser()
	want := "post success"

	if got != want {
		t.Errorf("Test failed")
	}
	
}


func TestGetUser(t *testing.T) {
    
	got := GetUser()
	want := "List of Users"

	if got != want {
		t.Errorf("Test failed")
	}
	
}