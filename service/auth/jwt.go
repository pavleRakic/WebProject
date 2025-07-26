package auth

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pavleRakic/testGoApi/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {

	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := []byte(config.Envs.JWTSecret)

		userID, err := VerifyJWT(tokenStr, secret)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optionally attach user ID to request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}

func VerifyJWT(tokenString string, secret []byte) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secret, nil
	})
	if err != nil {
		return 0, err
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDStr, ok := claims["userID"].(string)
		if !ok {
			return 0, errors.New("userID not found in token claims")
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return 0, errors.New("invalid userID in token claims")
		}

		/*err2 := store.GetUserByID(userID)
		// Optional: check expiration if you want, though jwt.Parse usually does it
		// expiredAt, ok := claims["expiredAt"].(float64)
		// if ok && int64(expiredAt) < time.Now().Unix() {
		//     return 0, errors.New("token expired")
		// }
		if err2 == nil {
			return 0, errors.New("That user doesn't exist")
		}*/

		return userID, nil
	}

	return 0, errors.New("invalid token claims")
}
