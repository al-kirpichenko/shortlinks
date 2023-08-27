package models

// Link - структура, содержащая сведения об url и их владельце
// (оригинальный урл, сокращенный, ID  пользователя, а также статус удаления)
type Link struct {
	ID       int    `json:"-"`
	Short    string `json:"short_url"`
	Original string `json:"original_url"`
	UserID   string `json:"-"`
	Deleted  bool   `json:"-"`
}
