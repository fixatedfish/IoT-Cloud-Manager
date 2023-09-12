package middleware

import (
	"IoTManager/utils"
	"database/sql"
	"log"
	"net/http"
)

//var deviceIdRegex = "([0-9a-zA-Z]{8}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{4}-[0-9a-zA-Z]{11})"

// Authenticator
// An authenticated URL is a URL that has a device ID in the URL params
// and a secret key in the Authorization header that must be validated
type Authenticator struct {
	unauthenticatedUris []string
	db                  *sql.DB
}

func NewAuthenticator(db *sql.DB, authUris []string) *Authenticator {
	auth := new(Authenticator)
	auth.unauthenticatedUris = authUris
	auth.db = db
	return auth
}

func (authenticator *Authenticator) Auth(handler http.Handler) http.Handler {
	handlerFn := func(w http.ResponseWriter, r *http.Request) {
		if utils.Contains(authenticator.unauthenticatedUris, r.RequestURI) {
			handler.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if checkIfSecretExists(authenticator.db, authHeader) {
			handler.ServeHTTP(w, r)
		} else {
			log.Println("Authentication Failed.")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized\n"))
		}

		log.Print("END - ", r.RequestURI, "  -  ", r.Method)
	}

	return http.HandlerFunc(handlerFn)
}

func checkIfSecretExists(db *sql.DB, secret string) bool {
	row := db.QueryRow("select 1 from device where secret_key = ?", secret)
	var found bool
	err := row.Scan(&found)
	if err != nil {
		return false
	}

	return found
}
