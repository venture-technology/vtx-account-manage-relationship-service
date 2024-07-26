package models

import "time"

type Sponsor struct {
	Record    int       `json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	Child     Child     `json:"child"`
	CreatedAt time.Time `json:"created_at"`
}

type Handshake struct {
	Record    int       `json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at"`
}

type EmploymentContract struct {
	Record      int         `json:"record"`
	Driver      Driver      `json:"driver"`
	School      School      `json:"school"`
	Responsible Responsible `json:"responsible"`
	Amount      int64       `json:"amount"`
	Months      int64       `json:"months"`
}
