package domain

type SignInRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Username        string `json:"username"`
	University      string `json:"university,omitempty"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
