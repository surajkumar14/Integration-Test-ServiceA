package integrationtest

import (
	"context"
	"testing"

	grpc_client "github.com/surajkumar14/Integration-Test-ServiceA/grpcClient"
	"github.com/surajkumar14/Integration-Test-ServiceA/mocks"
	"github.com/surajkumar14/Integration-Test-ServiceA/models/protomodel"
	grpcroutes "github.com/surajkumar14/Integration-Test-ServiceA/router/grpc_routes"

	serviceB "github.com/surajkumar14/ServiceB/models/protomodel"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRatesGrpcUsingGrpc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock Service B gRPC client
	mockServiceBRatesClient := mocks.NewMockGetRatesServiceClient(ctrl)

	// Set up the expected response from Service B
	expectedResponseB := &serviceB.RatesResponse{Rates: "1000"}
	mockServiceBRatesClient.EXPECT().GetRates(gomock.Any(), gomock.Any()).Return(expectedResponseB, nil)

	//set mock client in global variable
	grpc_client.SetServiceBRatesGrpcClient(mockServiceBRatesClient)

	// Service A's gRPC server setup
	serviceAGetRates := &grpcroutes.GetRatesServiceWithGrpcServer{}

	// Call the method in Service A that internally calls Service B
	resp, err := serviceAGetRates.GetRatesGrpc(context.Background(), &protomodel.RatesRequestGrpc{})
	require.NoError(t, err)

	assert.Equal(t, "1000", resp.Rates)
}
