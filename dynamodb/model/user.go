package model

// User .
type User struct {
	ID       string `json:"id,omitempty" dynamo:"id,omitempty"`
	Username string `json:"username,omitempty" dynamo:"username,omitempty"`
	Fullname string `json:"fullname,omitempty" dynamo:"fullname,omitempty"`
	ImageID  string `json:"image_id,omitempty" dynamo:"image_id,omitempty"`
}

func (User) TableName() string {
	return "users"
}
