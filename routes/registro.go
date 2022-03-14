package routes

import (
	"encoding/json"
	"net/http"

	"github.com/altairmagampo/altair-go-test-app/bd"
	"github.com/altairmagampo/altair-go-test-app/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "Email no informado.", http.StatusBadRequest)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "El password debe contener por lo menos 6 caracteres.", http.StatusBadRequest)
		return
	}
	_, encontrado, _ := bd.BuscarUsuarioPorEmail(t.Email)
	if encontrado {
		http.Error(w, "Usuario ya registrado.", http.StatusBadRequest)
		return
	}
	_, status, err := bd.Registrar(t)
	if err != nil {
		http.Error(w, "Ocurrrió un error al intentar registrar: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
