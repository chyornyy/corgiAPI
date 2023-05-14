package model

type Corgi struct {
	ID          string `json:"id"`
	Author      string `json:"author"`
	Name        string `json:"name"`
	Breed       string `json:"breed"`
	Color       string `json:"color"`
	Age         int    `json:"age"`
	Kennel      string `json:"kennel"`
	Description string `json:"description"`
	Photo       string `json:"photo_url"`
	Likes       int    `json:"total_likes"`
	//	Comments     []Comment `json:"comments"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
}
