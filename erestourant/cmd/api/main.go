package main

import (
	"log"

	"github.com/coleYab/erestourant/config"
	"github.com/coleYab/erestourant/internal/db"
	"github.com/coleYab/erestourant/internal/db/repository"
	"github.com/coleYab/erestourant/internal/server"
)

func main() {
	conn, err := db.ConnectToDb(config.Cfg.DbConnString)
	if err != nil {
		log.Fatalf("Unable to connect to the database %v", err.Error())
	}
	qry := repository.New(conn)
	srv := server.New(config.Cfg.Port, qry)
	srv.RegisterRoutes()
	srv.Run()
}
