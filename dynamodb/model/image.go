package model

type Image struct {
	ID     string `json:"id,omitempty" dynamo:"id,omitempty"`
	URL    string `json:"url,omitempty" dynamo:"url,omitempty"`
	UserID string `json:"user_id,omitempty" dynamo:"user_id,omitempty"`
}
