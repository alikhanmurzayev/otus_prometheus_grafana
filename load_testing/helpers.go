package main

import (
	"fmt"
	"math/rand"
)

func strToPtr(s string) *string {
	return &s
}

func getNewUser() *User {
	num := rand.Int()
	return &User{
		Username:  strToPtr(fmt.Sprintf("username_%d", num)),
		FirstName: strToPtr(fmt.Sprintf("firstname_%d", num)),
		LastName:  strToPtr(fmt.Sprintf("lastname_%d", num)),
		Email:     strToPtr(fmt.Sprintf("email_%d@example.com", num)),
		Phone:     strToPtr(fmt.Sprintf("phone%d", num)),
	}
}