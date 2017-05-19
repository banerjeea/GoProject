package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dogwalkingserver/models"
	"github.com/julienschmidt/httprouter"
)

//move it to env later
const secret = "hssgt12345"

type Key int

const MyKey Key = 0

//SetToken :
func SetToken(username string) string {

	//Token lives for an hour
	expireToken := time.Now().Add(time.Hour * 1).Unix()

	//assign claim
	claims := models.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:3000",
		},
	}

	//create token using claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign token using secret
	signedToken, _ := token.SignedString([]byte(secret))

	return signedToken
}

//Validate :
func Validate(controller httprouter.Handle) httprouter.Handle {

	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		authToken := r.Header.Get("Authorization")
		if authToken == "" {
			http.NotFound(w, r)
			return
		}

		token, err := jwt.ParseWithClaims(authToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

				return nil, fmt.Errorf("Unexpected signing method")
			}

			return []byte(secret), nil
		})

		if err != nil {
			http.NotFound(w, r)
			return
		}

		if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), MyKey, *claims)

			tst := ctx.Value(MyKey).(models.Claims)
			fmt.Println("Context prints fine right before passing..")
			fmt.Println(tst)
			controller(w, r.WithContext(ctx), nil)

		} else {
			http.NotFound(w, r)
			return
		}

	})

}
