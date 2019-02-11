package main

import (
	"flag"
	"github.com/matgarcia/edcomments/migration"
	"log"
)

func main() {

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migraciónn a la BD")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Comenzó la migración")
		migration.Migrate()
		log.Println("Finalizó la migración")
	}
}
