package dto

type LoginUserInput struct {
	Email    string
	Password string
}

type LoginUserOutput struct {
	ID           string
	Username     string
	Email        string
	AccessToken  string
	RefreshToken string
}
