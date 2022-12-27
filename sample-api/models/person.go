package models

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	City      string `json:"city,omitempty"`
	Age       string `json:"age,omitempty"`
}
