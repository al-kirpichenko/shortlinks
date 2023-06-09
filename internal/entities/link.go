package entities

type Link struct {
	ID       int    `json:"-"`
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
}
