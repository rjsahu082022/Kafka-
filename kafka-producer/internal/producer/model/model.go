package model

type Message struct {
	From    string `json:"from"`
	Message string `json:"message"`
	TO      string `json:"to"`
}
