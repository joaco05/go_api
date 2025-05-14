package middleware

import (
	"errors"
	"net/http"

	"go_api/api"
	"go_api/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Usuario Invalido.")

func Autorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var username string = r.URL.Query().Get("username")
		var err error

		if username == ""{
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w,  UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		var LoginDetails *tools.LoginDetails
		LoginDetails = (*database).GetUserLoginDetails(username)
	
		if (LoginDetails == nil || (username != (*LoginDetails).UsernameVerify)){
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}