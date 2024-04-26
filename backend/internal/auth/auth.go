package auth

type User struct {
	isAdmin  bool
	Username string `json:"username"`
	Password string `json:"password"`
}

type Storage interface {
	GetPassword(string) (string, error)
	RegisterUser(*User) error
	IsAdmin(string) bool
}

type Auth struct {
}
