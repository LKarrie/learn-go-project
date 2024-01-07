package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/LKarrie/learn-go-project/api"
	db "github.com/LKarrie/learn-go-project/db/sqlc"
	_ "github.com/LKarrie/learn-go-project/doc/statik"
	"github.com/LKarrie/learn-go-project/gapi"
	"github.com/LKarrie/learn-go-project/pb"
	"github.com/LKarrie/learn-go-project/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
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

	// run db migration
	runDBMigration(config.MigrationURL, config.DBSource)

	stroe := db.NewStore(conn)
	go runGatewayServer(config, stroe)
	runGrpcServer(config, stroe)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
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

func runGatewayServer(config util.Config, stroe db.Store) {
	server, err := gapi.NewServer(config, stroe)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pb.RegisterLearnGoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("cannot register handler server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := fs.NewWithNamespace("swagger")
	if err != nil {
		log.Fatal("cannot create statik fs:", err)
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("start HTTP gateway server at %s", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot create HTTP gateway server:", err)
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
