package domain

type ErrorMessage struct {
	Header      string `json:"header,omitempty"`
	Description string `json:"description,omitempty"`
}
