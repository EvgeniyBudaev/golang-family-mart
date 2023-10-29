package schemas

type RegisterRequest struct {
	/* Схема данных при регистрации */
	Username  string
	Password1 string
	Password2 string
	Email     string
}
