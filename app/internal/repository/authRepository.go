package repository


type AuthRepository interface {
	Authenticate(username, password string) (bool, string, string, error)
}

