package models

type Message struct {
	ID          int    `json:"id,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Message     string `json:"message,omitempty"`
	Status      string `json:"status,omitempty"`
	CreatedBy   string `json:"createdBy,omitempty"`
	CreatedTime int64  `json:"createdTime,omitempty"`
}
