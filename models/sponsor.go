package models

import (
	"github.com/google/uuid"
)

type Sponsor struct {
	ID          uuid.UUID   `json:"id"`
	Driver      Driver      `json:"driver"`
	School      School      `json:"school"`
	Responsible Responsible `json:"responsible"`
}
