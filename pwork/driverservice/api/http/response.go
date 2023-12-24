package http

import "github.com/google/uuid"

type ResponseGet struct {
	Record []RequstAdd `json:"order"`
}

type ResponseDriverGet struct {
	Record []RequstDriverGet `json:"drivers"`
}

type RequstDriverGet struct {
	DriverName string    `json:"name"`
	Uuid       uuid.UUID `json:"uuid"`
}
