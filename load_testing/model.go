package main

type User struct {
	ID        *int64  `json:"id,omitempty" db:"id"`
	Username  *string `json:"username" db:"username"`
	FirstName *string `json:"firstName" db:"first_name"`
	LastName  *string `json:"lastName" db:"last_name"`
	Email     *string `json:"email" db:"email"`
	Phone     *string `json:"phone" db:"phone"`
}
