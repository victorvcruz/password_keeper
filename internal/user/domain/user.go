package domain

import "time"

type User struct {
	id             *int
	name           *string
	email          *string
	masterPassword *string
	createdAt      *time.Time
	updatedAt      *time.Time
	deletedAt      *time.Time
}
