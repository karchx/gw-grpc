package main

import (
	"net"

	log "github.com/gothew/l-og"
	"github.com/karchx/gw-grpc/internal"
	"github.com/karchx/gw-grpc/protogen/output/orders"
	"google.golang.org/grpc"
)

func main() {
	const addr = "0.0.0.0:50051"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	orders.RegisterOrdersServer(server, &orderService)

	log.Infof("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
