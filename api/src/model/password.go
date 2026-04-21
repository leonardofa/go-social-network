package model

// Password represents a user's password change request.
type Password struct {
	Actual string `json:"actual,omitempty"`
	New    string `json:"new,omitempty"`
}
