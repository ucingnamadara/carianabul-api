package middlewares

import (
	errors "dana/anabul-rest-api/src/templates/error"
	"dana/anabul-rest-api/src/templates/response"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKay = []byte(os.Getenv("SECRET_KEY"))

func VerifyJWT(endpointHandler func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		JwtToken := strings.Replace(r.Header["Authorization"][0], fmt.Sprintf("%s ", "Bearer"), "", 1)
		if JwtToken == "" {
			response.Json(w, http.StatusUnauthorized, "Token is Required", nil)
			return
		}
		token, err := jwt.Parse(JwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.Unauthorization("Invalid Token")
			}
			return secretKay, nil
		})

		if err != nil || token == nil {
			response.Json(w, http.StatusUnauthorized, "Invalid Token", nil)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.Json(w, http.StatusUnauthorized, "Invalid Claims", nil)
			return
		}

		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Local().Unix() {
			response.Json(w, http.StatusUnauthorized, "Token Expired", nil)
			return
		}

		r.Header.Set("id", claims["id"].(string))
		r.Header.Set("email", claims["email"].(string))
		r.Header.Set("phoneNumber", claims["phoneNumber"].(string))
		r.Header.Set("fullName", claims["fullName"].(string))
		endpointHandler(w, r)
	})
}
