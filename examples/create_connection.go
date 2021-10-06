package main

import (
	"context"
	"fmt"

	"github.com/gsoultan/dataX/builder"
	"github.com/gsoultan/dataX/factory"
)

func main() {
	ctx := context.Background()

	bc := builder.NewConfig()
	bc.WithHost("localhost").WithPort(54432).WithDatabase("postgres").WithProvider("postgres").WithUserName("postgres").WithPassword("postgres")

	cb := factory.Create(ctx, bc)
	if err := cb.Ping(); err != nil {
		fmt.Println("error")
	}
}
