package main

import (
	"eracenergy/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router()
	config, err := utils.LoadConfig("config.json")
	utils.PanicOnError(err)

	/* db, err := utils.DBCon(config)
	utils.PanicOnError(err)
	err = dbMigrate(db)
	utils.PanicOnError(err)

	defer db.Close()*/

	// Launching the server
	fmt.Println("Serving on port", config.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
