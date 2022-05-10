package main

import (
	"context"
	"fmt"

	productservice "github.com/evgeniy-dammer/objectsgrpc/productservice/proto"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:1111", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err)
	}

	defer connection.Close()

	productServ := productservice.NewProductServiceClient(connection)

	response, err := productServ.Find(context.Background(), &productservice.FindRequest{})

	if err != nil {
		fmt.Println(err)
	} else {
		product := response.Product

		fmt.Printf(
			"Id: %s\n Name: %s\n Price: %f\n Quantity: %d\n Status: %v",
			product.Id,
			product.Name,
			product.Price,
			product.Quantity,
			product.Status)
	}
}
