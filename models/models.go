package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type MyTestModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Define other fields
}

// Define functions to interact with the database (CRUD operations)
