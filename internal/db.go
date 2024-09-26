package internal

import (
	"fmt"

	"github.com/karchx/gw-grpc/protogen/output/orders"
)

type DB struct {
	collection []*orders.Order
}

func NewDB() *DB {
	return &DB{
		collection: make([]*orders.Order, 0),
	}
}

func (d *DB) AddOrder(order *orders.Order) error {
	for _, o := range d.collection {
		if o.OrderId == order.OrderId {
			return fmt.Errorf("duplicate order id %d", order.GetOrderId())
		}
	}

	d.collection = append(d.collection, order)
	return nil
}
