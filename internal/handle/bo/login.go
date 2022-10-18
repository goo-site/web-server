package bo

type LoginRequest struct {
	UserId   string `json:"user_id"`
	PassWord string `json:"password"`
}

type RegisterRequest struct {
	UserId           string `json:"user_id"`
	PassWord         string `json:"password"`
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}
