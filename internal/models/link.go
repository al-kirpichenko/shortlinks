package models

type Link struct {
	ID       int    `json:"-"`
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
	UserID   string `json:"-"`
	Deleted  bool   `json:"-"`
}
