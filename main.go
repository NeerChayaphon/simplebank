package main

import (
	"database/sql"
	"log"

	"github.com/NeerChayaphon/simplebank/api"
	db "github.com/NeerChayaphon/simplebank/db/sqlc"
	"github.com/NeerChayaphon/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("canno start server:", err)
	}
}
