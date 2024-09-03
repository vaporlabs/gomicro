package main

import (
	"log"

	"github.com/vaporlabs/gomicro/order/config"
	"github.com/vaporlabs/gomicro/order/internal/adapters/db"
	"github.com/vaporlabs/gomicro/order/internal/adapters/grpc"
	"github.com/vaporlabs/gomicro/order/internal/adapters/payment"
	"github.com/vaporlabs/gomicro/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error %v", err)
	}
	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
