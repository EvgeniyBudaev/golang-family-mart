package schemas

type RegisterRequest struct {
	/* Схема данных при регистрации */
	username  string
	password1 string
	password2 string
	email     string
}
