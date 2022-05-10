package handlers

import (
	"context"

	productservice "github.com/evgeniy-dammer/objectsgrpc/productservice/proto"
)

type ProductServiceServer struct {
	productservice.UnimplementedProductServiceServer
}

func (*ProductServiceServer) Find(ctx context.Context, in *productservice.FindRequest) (*productservice.FindResponse, error) {
	return &productservice.FindResponse{
		Product: &productservice.Product{
			Id:       "p01",
			Name:     "name 1",
			Quantity: 27,
			Price:    3,
			Status:   true,
		},
	}, nil
}
