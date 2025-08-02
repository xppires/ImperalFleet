package repository


type UmsRepository interface {
	Authenticate(username, password string) (bool, string, string, error)
}

