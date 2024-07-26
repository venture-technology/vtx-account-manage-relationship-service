package models

type School struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	CNPJ       string `json:"cnpj"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZIP        string `json:"zip"`
}
