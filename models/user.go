package models

import "time"

// User model to save information
type User struct {
	ID        string    `bson:"_id" json:"id" structs:"id,omitempty"`
	Name      string    `bson:"name" json:"name" structs:"name,omitempty"`
	Email     string    `bson:"email" json:"email" structs:"email,omitempty"`
	Address   string    `bson:"address" json:"address" structs:"address,omitempty"`
	CreatedAt time.Time `bson:"created_at" json:"created_at" structs:"created_at,omitempty"`
}
