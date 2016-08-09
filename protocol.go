package pitcher

import "time"

// User in Versatile
type User struct {
	ID         string `json:"_key"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	OwnerToken string `json:"ownerToken"`
	Password   string `json:"password"`
	Role       string `json:"role"`
}

// SigninRequest .
type SigninRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Session .
type Session struct {
	Agent   string    `json:"agent,omitempty"`
	Token   string    `json:"token,omitempty"`
	Role    string    `json:"role,omitempty"`
	Created time.Time `json:"created,omitempty"`
	ValidTo time.Time `json:"validTo,omitempty"`
	// OwnerToken string `json:"ownerToken,omitempty"`
	// Payload  string    `json:"payload,omitempty"`
	// Policies []string  `json:"policies,omitempty"`
}

// Data .
type Data struct {
	ID        string `json:"_key"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	LastValue string `json:"lastValue"`
}

// ErrorResponse .
type ErrorResponse struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	ErrorCode   string `json:"errorCode"`
}
