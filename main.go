package main

import (
	"log"

	"github.com/AxLShnitzel/twittor/bd"
	"github.com/AxLShnitzel/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		handlers.Manejadores()
	}
}
