package util 


import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

var hmacSampleSecret []byte

func JWTInit() {
	hmacSampleSecret = []byte(GetEnv("JWT_SECRET", "fail"))
}

func CreateJWT(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		// "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		// do something could not parse string
		return ""
	}
	return tokenString
}

func ParseJWT(tokenString string) bool {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	parsedToken := strings.Replace(tokenString," ", "", -1)
	token, err := jwt.Parse(parsedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["username"])
		return true
	} else {
		fmt.Println(err)
		return false
	}
}
