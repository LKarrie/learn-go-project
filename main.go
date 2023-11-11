package main

import (
	"database/sql"
	"log"

	"github.com/LKarrie/learn-go-project/api"
	db "github.com/LKarrie/learn-go-project/db/sqlc"
	"github.com/LKarrie/learn-go-project/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

func main() {
  config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:",err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	stroe := db.NewStore(conn)
	server := api.NewServer(stroe)
	
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannt start server:",err)
	}
}