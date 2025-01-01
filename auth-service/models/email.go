package models

type EmailType struct {
	Subject  string `json:"subject"`
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}
