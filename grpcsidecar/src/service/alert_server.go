package service

import (
	"context"
	"grsidecar/pbgen"
	"grsidecar/serialize_wrappers"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AlertRequestServer struct {
	pbgen.UnimplementedCaterAlertRequestServer
	AlertsMem *AlertStore
}

func NewAlertCaterServer(AS *AlertStore) *AlertRequestServer { //if assoicated with AlertStore struct, mutex is passed by value
	AS.ReadyAlertStore()
	return &AlertRequestServer{AlertsMem: AS}
}

func (s *AlertRequestServer) CaterAlert(ctx context.Context, req *pbgen.AlertRequest,
) (*pbgen.AlertResponse, error) {

	if len(req.Cid) > 0 {
		_, err := uuid.Parse(req.Cid)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "not a valid uuid %v", err)
		}
	} else {
		return nil, status.Errorf(codes.NotFound, "please provide a CID ")
	}
	alert := req.GetAlertsbatch()
	if alert == nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"Sorry, you need to provide the parameters! Please call DescribeAlert for info!")
	} else {
		log.Infof("Received request to service an alert with cid : %s", req.Cid)
		err := s.AlertsMem.StoreAlert(alert)
		if err != nil {
			log.Debugf("%v", err)
		} else {
			log.Printf("Alert is Stored : %v/n", s.AlertsMem.Store)
		}
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline exceeded!")
		return nil, status.Error(codes.DeadlineExceeded, "Deadline exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request was canceled!")
		return nil, status.Error(codes.Canceled, "Context canceled")
	}
	//pass to relevant processes to handle
	return &pbgen.AlertResponse{Seen: "Alert will be serviced soon"}, nil
}

func (s *AlertRequestServer) DescribeAlert(ctx context.Context, request *pbgen.DescriptionRequest) (*pbgen.AlertSchema, error) {
	schemafor := request.Giveme
	if schemafor == "Alerts" || schemafor == "alerts" {
		schema := serialize_wrappers.ProtobuftoJson(&pbgen.AlertRequest{
			Alertsbatch: &pbgen.AlertList{
				Alerts: []*pbgen.AlertDetail{{},
					{},
				},
			},
		})
		if schema == "" {
			log.Println("Something went wrong")
			log.Debugf("Schema produced %v\n", schema)
			return &pbgen.AlertSchema{Schema: "oops"}, status.Error(codes.Internal, "Error converting to string")
		} else {
			return &pbgen.AlertSchema{Schema: schema}, nil
		}
	}
	return &pbgen.AlertSchema{Schema: "Describe What?"}, nil
}

func (s *AlertRequestServer) RetrieveAlert(ctx context.Context, requested *pbgen.RetrieveAlertRequest) (list *pbgen.AlertList, err error) {
	found, err := s.AlertsMem.Search(requested)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pbgen.AlertList{Alerts: []*pbgen.AlertDetail{found}}, nil
}
