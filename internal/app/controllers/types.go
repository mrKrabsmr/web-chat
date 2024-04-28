package controllers

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}
