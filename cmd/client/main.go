package main

import (
	"context"
	"net/http"

	log "github.com/gothew/l-og"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/karchx/gw-grpc/protogen/output/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	orderServiceAddr := "localhost:50051"
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	if err = orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	addr := "0.0.0.0:8080"
	log.Infof("API gateway server is running on %s\n", addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
