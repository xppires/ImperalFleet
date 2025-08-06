package common

import (
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"	
	"strings" 
	"errors" 
)

// getToken generates a JWT token for the user with the given userId and role.
// It uses a signing key to sign the token and returns the token string.
// The token contains claims for userId and role.
func GetToken(key string, userId ,role string) (string, error) {
	signingKey := []byte(key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"Role": role,
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}
// verifyToken verifies the JWT token string and returns the claims if valid.
// It uses the same signing key to parse the token and returns the claims or an error if verification fails.
func verifyToken(key,tokenString string) (jwt.Claims, error) {
	signingKey := []byte(key)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}

func ExtractUserFromJWT(key string,w http.ResponseWriter,r *http.Request) (string, string, error) {
    token := r.Header.Get("Authorization")
	if len(token) == 0 {	 
		return  "", "", errors.New("Missing Authorization Header")
	}
	tokenString := strings.Replace(token, "Bearer ", "", 1)
    // Validar e extrair claims do JWT
    claims, err := verifyToken(key,tokenString)
	if err != nil {	
		return  "", "", err
		}
    return  claims.(jwt.MapClaims)["userId"].(string), claims.(jwt.MapClaims)["Role"].(string), nil
}