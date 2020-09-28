package domain

// Link .
type Link struct {
	Model
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	PostedBy    UUID   `sql:",type:uuid" json:"posted_by"`
}
