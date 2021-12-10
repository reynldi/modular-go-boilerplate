package application

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context){
	auth := c.Request.Header.Get("Authorization")

	fmt.Println(auth)
}
//
//import "net/http"
//
//func AuthHandler(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		dbConn := DB.Connect()
//		header := r.Header.Get("Authorization")
//		// if auth is not available then proceed to resolver
//		if header == "" {
//			next.ServeHTTP(w, r)
//		} else {
//			userData, err := Auth.VerifyIDToken(dbConn.Context, header)
//			if err != nil {
//				next.ServeHTTP(w, r)
//			} else {
//				// merge userID into request context
//				ctx := context.WithValue(r.Context(), "userID", userData.UID)
//				next.ServeHTTP(w, r.WithContext(ctx))
//			}
//		}
//	})
//}
//
