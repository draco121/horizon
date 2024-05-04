package models

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
