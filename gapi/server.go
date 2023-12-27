package gapi

import (
	"fmt"

	db "github.com/LKarrie/learn-go-project/db/sqlc"
	"github.com/LKarrie/learn-go-project/pb"
	"github.com/LKarrie/learn-go-project/token"
	"github.com/LKarrie/learn-go-project/util"
)

// Server servers gRPC requests for my service.
type Server struct {
	pb.UnimplementedLearnGoServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// tokenMaker,err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
