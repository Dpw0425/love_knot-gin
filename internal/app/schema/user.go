package schema

type UserRegister struct {
	NickName string
	Password string
	Avatar   string
	Gender   int
	Email    string
}

type UserLogin struct {
	Email    string
	Password string
}
