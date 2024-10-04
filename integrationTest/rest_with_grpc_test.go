package integrationtest

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/surajkumar14/Integration-Test-ServiceA/controller"
	"github.com/surajkumar14/Integration-Test-ServiceA/mocks"
	"github.com/surajkumar14/Integration-Test-ServiceA/router"

	serviceB "github.com/surajkumar14/ServiceB/models/protomodel"

	grpc_client "github.com/surajkumar14/Integration-Test-ServiceA/grpcClient"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRatesHTTPApiWithGrpc(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockServiceBRatesClient := mocks.NewMockGetRatesServiceClient(ctrl)

	// Expected response from Service B
	expectedResponse := &serviceB.RatesResponse{Rates: "1000"}
	mockServiceBRatesClient.EXPECT().GetRates(gomock.Any(), gomock.Any()).Return(expectedResponse, nil)

	//set mock client in global variable
	grpc_client.SetServiceBRatesGrpcClient(mockServiceBRatesClient)

	// Set Gin to test mode
	gin.SetMode(gin.TestMode)
	r := router.SetupHttpServer()

	resp, body, err := performRequest(r, "GET", "/service_a/getrates_grpc")
	require.NoError(t, err)

	var ratesResponse controller.RatesResponse
	require.NoError(t, json.Unmarshal(body, &ratesResponse))

	assert.Equal(t, "1000", ratesResponse.Rates)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
