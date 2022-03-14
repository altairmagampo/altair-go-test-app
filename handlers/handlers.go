package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/altairmagampo/altair-go-test-app/middlew"
	"github.com/altairmagampo/altair-go-test-app/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escuhar al servidor*/
func Manejadores() {
	//Aqui registramos las rutas
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")

	PORT := os.Getenv("PORT") //Se le da preferencia a las variables de entorno del sistema, para el caso de herouku o donde se despliegue
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router) //Le damos permisos a todos para que desde caulquier lugar se pueda acceder
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
