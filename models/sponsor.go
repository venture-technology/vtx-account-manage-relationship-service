package models

type Sponsor struct {
	ID          int         `json:"id"`
	Driver      Driver      `json:"driver"`
	School      School      `json:"school"`
	Responsible Responsible `json:"responsible"`
}

type Handshake struct {
	ID     int    `json:"id"`
	Driver Driver `json:"driver"`
	School School `json:"school"`
}
