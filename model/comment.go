package model

type Comment struct {
	User      string `json:"user"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
