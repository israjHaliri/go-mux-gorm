package util

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"go-mux-gorm/model"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//List of endpoints that doesn't require auth
		notAuth := []string{"/login"}

		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//Grab the token from the header
		tokenHeader := r.Header.Get("Authorization")

		//Token is missing, returns with error code 403 Unauthorized
		if tokenHeader == "" {
			ResponseWithJson(w, http.StatusForbidden, "Missing auth token")

			return
		}

		//The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			ResponseWithJson(w, http.StatusForbidden, "Invalid/Malformed auth token")

			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in

		tk := &model.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			ResponseWithJson(w, http.StatusForbidden, "Malformed authentication token")

			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			ResponseWithJson(w, http.StatusForbidden, "Token is not valid")

			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		logs.Info("User ", tk.UserId) //Useful for monitoring

		ctx := context.WithValue(r.Context(), "user", tk.UserId)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
