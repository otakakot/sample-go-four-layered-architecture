package main

import (
	"os"

	"github.com/otakakot/sample-go-four-layered-architecture/internal/adapter/controller"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/adapter/gateway"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/application/interactor"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/driver/postgres"
	"github.com/otakakot/sample-go-four-layered-architecture/internal/driver/server"
	"github.com/otakakot/sample-go-four-layered-architecture/pkg/api"
)

func main() {
	db, err := postgres.New(os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	gw := gateway.NewSample(db)

	uc := interactor.NewSample(gw)

	hdl, err := api.NewServer(&controller.Controller{
		Sample: controller.NewSample(uc),
	})
	if err != nil {
		panic(err)
	}

	srv := server.NewServer("8080", hdl)

	srv.Run()
}
