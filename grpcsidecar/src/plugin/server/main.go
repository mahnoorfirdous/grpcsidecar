package main

import (
	"grsidecar/pbgen"
	"grsidecar/service"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartAlertServer() error {
	alertServer := service.NewAlertCaterServer(&service.AlertStore{})

	grpcServer := grpc.NewServer()
	pbgen.RegisterCaterAlertRequestServer(grpcServer, alertServer)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":8050")
	if err != nil {
		return err
	}
	log.Info("Starting GRPC Listener...")
	log.Fatal(grpcServer.Serve(listener))
	return nil
}

func main() {

	StartAlertServer()
}
