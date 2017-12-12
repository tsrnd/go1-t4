package models

import (
	// "fmt"
	// "github.com/fatih/structs"
	"time"
)

type Model struct {
	ID				uint			`schema:"id"`
	CreatedAt	time.Time	`schema:"created_at"`
	UpdatedAt time.Time `schema:"updated_at"`
	DeletedAt time.Time `schema:"deleted_at"`
}

type ModelExtend interface {
	TableName()		string
	GetSchema()		map[string]interface{}
}
