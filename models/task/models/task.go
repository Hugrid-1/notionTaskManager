package models

import "encoding/json"

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	CreatedDate string `json:"created_date"`
	UpdatedDate string `json:"updated_date"`
	Status      string `json:"status"`
}

//unmarshal function
func (t *Task) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}
