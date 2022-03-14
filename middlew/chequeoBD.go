package middlew

import (
	"net/http"

	"github.com/altairmagampo/altair-go-test-app/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(w, "Sin conexión con la BD.", http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	}
}
