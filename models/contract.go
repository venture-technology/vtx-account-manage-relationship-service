package models

import (
	"time"

	"github.com/google/uuid"
)

type Contract struct {
	Record    uuid.UUID `json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	Child     Child     `json:"child"`
	CreatedAt time.Time `json:"created_at"`
	ExpireAt  time.Time `json:"expire_at"`
	Amount    int64     `json:"amount"`
	Months    int64     `json:"months"`
}

type Handshake struct {
	Record    int       `json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at"`
}
