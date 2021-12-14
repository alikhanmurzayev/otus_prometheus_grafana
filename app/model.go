package main

type User struct {
	ID        *int64  `json:"id" db:"id"`
	Username  *string `json:"username" db:"username"`
	FirstName *string `json:"firstName" db:"first_name"`
	LastName  *string `json:"lastName" db:"last_name"`
	Email     *string `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
}

/*
docker run --rm -d -p 5432:5432 --name postgres --env POSTGRES_DB=mydb --env POSTGRES_USER=myuser --env POSTGRES_PASSWORD=mypassword postgres:latest


psql -h localhost -p 5432 -U myuser -W mydb

create table users (id bigserial primary key, username varchar, first_name varchar, last_name varchar, email varchar, phone varchar);
*/
