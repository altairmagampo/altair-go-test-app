package main

import (
	"log"

	"github.com/altairmagampo/altair-go-test-app/bd"
	"github.com/altairmagampo/altair-go-test-app/handlers"
)

func main() {

	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a la base de datos.")
		return
	}

	handlers.Manejadores()

}
