package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Model struct {
	ID				uint								`schema:"id"`
	CreatedAt	timestamp.Timestamp	`schema:"created_at"`
	UpdatedAt timestamp.Timestamp `schema:"updated_at"`
	DeletedAt timestamp.Timestamp `schema:"deleted_at"`
}
