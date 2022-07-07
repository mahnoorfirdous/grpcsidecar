package service_test

import (
	"context"
	"grsidecar/sample"
	"grsidecar/service"
	"testing"

	"grsidecar/pbgen"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateAlertStore(tst *testing.T) {
	tst.Parallel()
	sample.Names.Initstore(100)
	alertWithoutCID := sample.NewAlert()
	alertWithoutCID.Cid = ""

	alertWithInvalidUUID := sample.NewAlert()
	alertWithInvalidUUID.Cid = "2439"

	testCases := []struct {
		testname string
		store    *service.AlertStore
		alert    *pbgen.AlertRequest
		errcode  codes.Code
	}{
		{
			testname: "No UUID",
			store:    &service.AlertStore{},
			alert:    alertWithoutCID,
			errcode:  codes.NotFound,
		},

		{
			testname: "Invalid UUID",
			store:    &service.AlertStore{},
			alert:    alertWithInvalidUUID,
			errcode:  codes.InvalidArgument,
		},

		{
			testname: "Normal Alert",
			store:    &service.AlertStore{},
			alert:    sample.NewAlert(),
			errcode:  codes.OK,
		},
	}

	for t := range testCases {
		tc := testCases[t]

		tst.Run(tc.testname, func(t *testing.T) {
			t.Parallel()
			tc.store.ReadyAlertStore()
			server := service.NewAlertCaterServer(tc.store)
			response, err := server.CaterAlert(context.Background(), tc.alert)

			if tc.errcode != codes.OK {
				require.Error(t, err)
				require.Nil(t, response)
				status, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.errcode, status.Code())
			} else {
				require.NoError(t, err)
				require.NotNil(t, response)
				require.NotEmpty(t, response.Seen)
			}
		})

	}
}
