package domain

// User .
type User struct {
	Model
	Name         string  `json:"name,omitempty"`
	Email        string  `json:"email,omitempty"`
	PasswordHash string  `json:"-"`
	Links        []*Link `json:"links,omitempty"`
}
