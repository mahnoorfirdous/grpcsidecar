package main

import (
	"context"
	"flag"
	"grsidecar/pbgen"
	"grsidecar/sample"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	dialaddress := flag.String("address", "", "we are dialing this address")
	cid := flag.String("param", "", "egg")
	flag.Parse()
	log.Infof("Dialing server %s", *dialaddress)
	connection, err := grpc.Dial(*dialaddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Cannot dial server! Error: ", err)
	}

	AlertClient := pbgen.NewCaterAlertRequestClient(connection)
	sample.Names.Initstore(3)
	alert := sample.NewAlert() //we only have 3 alert objects for readibility right now

	response, err := AlertClient.CaterAlert(context.Background(), alert)
	if err != nil {
		log.Fatalln("Cannot receive response from server! Error: ", err)
	}
	log.Println(response.Seen, err)

	describeresponse, err := AlertClient.DescribeAlert(context.Background(), &pbgen.DescriptionRequest{Giveme: "alerts"})
	if err != nil {
		log.Fatalln("Cannot receive response from server! Error: ", err)
	}
	log.Println(describeresponse.Schema)

	list, err := AlertClient.RetrieveAlert(context.Background(), &pbgen.RetrieveAlertRequest{Cid: *cid})
	if err != nil {
		log.Fatalln("Cannot receive response from server! Error: ", err)
	}
	log.Println(list)
}
