package service_test

import (
	"context"
	"grsidecar/pbgen"
	"grsidecar/sample"
	"grsidecar/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartTestAlertServer(t *testing.T, alertstore *service.AlertStore) string {
	alertserver := service.NewAlertCaterServer(alertstore)
	grpcServe := grpc.NewServer()
	pbgen.RegisterCaterAlertRequestServer(grpcServe, alertserver)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	go grpcServe.Serve(listener)

	return listener.Addr().String()

}

func MakeTestAlertClient(t *testing.T, serverAdd string) pbgen.CaterAlertRequestClient {
	connection, err := grpc.Dial(serverAdd, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return pbgen.NewCaterAlertRequestClient(connection)
}

func StartTestAlertClient(t *testing.T) {
	t.Parallel()

	alertss := &service.AlertStore{}
	alertss.ReadyAlertStore()
	serverAdd := StartTestAlertServer(t, alertss)
	alertclient := MakeTestAlertClient(t, serverAdd)

	sample.Names.Initstore(10)
	alertsample := sample.NewAlert()
	response, err := alertclient.CaterAlert(context.Background(), alertsample)
	require.NotNil(t, response)
	require.NoError(t, err)

}
