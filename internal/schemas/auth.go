package schemas

type AuthRequest struct {
	/* Схема данных при авторизации */
	username string
	password string
}

type AuthResponse struct {
	/* Ответ на удачную авторизацию */
	UserId   string
	JwtToken string
}
