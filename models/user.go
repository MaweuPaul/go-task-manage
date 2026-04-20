package models

type CreateUserInput struct {
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Used for responses - password never exposed
type User struct {
	ID        string `json:"id,omitempty"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"-"` // never in response
}

type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
