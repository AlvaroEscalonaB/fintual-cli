package models

type BankAttribute struct {
	Name string `json:"name"`
}

type Bank struct {
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Attributes BankAttribute `json:"attributes"`
}

type BankResponse struct {
	Data []Bank `json:"data"`
}
