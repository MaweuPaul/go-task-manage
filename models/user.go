package models

type User struct {
	ID        string `json:"id"`
	NameFirst string `json:"nameFirst"`
	NameLast  string `json:"nameLast"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}

type AuthResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
