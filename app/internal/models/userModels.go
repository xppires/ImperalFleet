package models

// User represents a user in the system
// It contains fields for ID, username, password, email, and role.
type User struct {
    UID        	string       `json:"id"`
    Username    string    `json:"username"` 
	Password  	string    `json:"password"`
	Email	 	string    `json:"email"`
	Role 		string    `json:"role"`
}





// type ValidLogin  interface{
// 	Authenticate(username, password string) (bool, userId string, role string, err error)
// }

// func (u *User) Authenticate() (bool, string, string, error) {
// 	for _, user := range Users {
// 		if user.Username == u.Username && user.Password == u.Password {
// 			return true, user.UID, user.Role, nil
// 		}	
		
// 	}
// 	return false, "", "", nil
// }