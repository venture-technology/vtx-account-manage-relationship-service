package models

type Responsible struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	CPF        string `json:"cpf"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZIP        string `json:"zip"`
}

type Child struct {
	Name        string      `json:"name"`
	RG          string      `json:"rg"`
	Responsible Responsible `json:"responsible"`
	Shift       string      `json:"shift"`
}
