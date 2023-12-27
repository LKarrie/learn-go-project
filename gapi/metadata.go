package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

const (
	grpcgatewayUserAgentHeader = "grpcgateway-user-agent"
	userAgentHeader            = "user-agent"
	xForwardedForHeader        = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	clientIP  string
}

func (server *Server) extracMetadata(ctx context.Context) *Metadata {
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("md: %+v\n", md)
		// fot http request
		if userAgents := md.Get(grpcgatewayUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		// fot grpc request
		if userAgents := md.Get(userAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
		// fot http request
		if clientIPs := md.Get(xForwardedForHeader); len(clientIPs) > 0 {
			mtdt.clientIP = clientIPs[0]
		}
	}

	// fot grpc request
	if p, ok := peer.FromContext(ctx); ok {
		mtdt.clientIP = p.Addr.String()
	}

	return mtdt
}
