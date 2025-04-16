package main

import (
	"log"

	"test/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// NICETOHAVE: potremmo prevedere un flag "-resetDB=true" che mi cancella il DB e ricrea
	// diamo evidenza di questa cosa nel `readme.md`
	// FIXME: nella folder "/query" ci sono anche dei commands, non solo delle query. Se vogliamo tenere questo approccio separiamo. Oltretutto separiamo i normali commands/queries da quelli di creazione dei SQL objects e seed dei dati.
	router := gin.Default()

	routes.SetupRoutes(router)

	// NICETOHAVE: merge those two lines
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
