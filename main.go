package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/LKarrie/learn-go-project/api"
	db "github.com/LKarrie/learn-go-project/db/sqlc"
	"github.com/LKarrie/learn-go-project/gapi"
	"github.com/LKarrie/learn-go-project/pb"
	"github.com/LKarrie/learn-go-project/util"
	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	stroe := db.NewStore(conn)

	runGrpcServer(config, stroe)
}

func runGrpcServer(config util.Config, stroe db.Store) {
	server, err := gapi.NewServer(config, stroe)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLearnGoServer(grpcServer, server)
	// allow client know what RPC api can execute
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create gRPC server:", err)
	}
}

func runGinServer(config util.Config, stroe db.Store) {
	server, err := api.NewServer(config, stroe)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannt start server:", err)
	}
}
