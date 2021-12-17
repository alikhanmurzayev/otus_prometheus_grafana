package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func makeRequests(ctx context.Context, processID int) error {
	log.Printf("process %d started", processID)
	defer log.Printf("process %d finished", processID)

	for ctx.Err() == nil {
		user := getNewUser()
		user, err := createUser(ctx, user)
		if err != nil {
			return fmt.Errorf("createUser: %w", err)
		}
		//log.Printf("process %d created user with name %d", processID, *user.ID)
		for i := 0; i < 3; i++ {
			user, err = getUser(ctx, *user.ID)
			if err != nil {
				return fmt.Errorf("getUser: %w", err)
			}
			//log.Printf("process %d got user with id %d", processID, *user.ID)
			user, err = updateUser(ctx, user)
			if err != nil {
				return fmt.Errorf("updateUser: %w", err)
			}
			//log.Printf("process %d updated user with id %d", processID, *user.ID)
		}
		err = deleteUser(ctx, *user.ID)
		if err != nil {
			return fmt.Errorf("deleteUser: %w", err)
		}
		//log.Printf("process %d deleted user with id %d", processID, *user.ID)
		_, err = getUser(ctx, *user.ID)
		if err == nil {
			return fmt.Errorf("got user, but not expected")
		}
		//log.Printf("process %d could not get user with id %d", processID, *user.ID)

		err = panicRequest(ctx)
		if err == nil {
			return fmt.Errorf("panicRequest: expected error, but got nil")
		}
		//log.Printf("process %d got panic: %s", processID, err)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}

	return nil
}

func createUser(ctx context.Context, user *User) (*User, error) {
	return user, makeRequest(ctx, user, user, http.MethodPost, "/user")
}

func getUser(ctx context.Context, id int64) (*User, error) {
	response := new(User)
	return response, makeRequest(ctx, nil, response, http.MethodGet, fmt.Sprintf("/user/%d", id))
}

func updateUser(ctx context.Context, user *User) (*User, error) {
	return user, makeRequest(ctx, user, user, http.MethodPut, fmt.Sprintf("/user/%d", *user.ID))
}

func deleteUser(ctx context.Context, id int64) error {
	return makeRequest(ctx, nil, nil, http.MethodDelete, fmt.Sprintf("/user/%d", id))
}

func panicRequest(ctx context.Context) error {
	return makeRequest(ctx, nil, nil, http.MethodGet, "/panic")
}
