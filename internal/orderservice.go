package internal

import (
	"context"

	log "github.com/gothew/l-og"
	"github.com/karchx/gw-grpc/protogen/output/orders"
)

// OrderService should implement the OrdersServer interface generated from grpc
type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

// NewOrderService create new OrderService
func NewOrderService(db *DB) OrderService {
	return OrderService{db: db}
}

// AddOrder
func (o *OrderService) AddOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Info("Received ad add-order request")
	err := o.db.AddOrder(req.GetOrder())

	return &orders.Empty{}, err
}
