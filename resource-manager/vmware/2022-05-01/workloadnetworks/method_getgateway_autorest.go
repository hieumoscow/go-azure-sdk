package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetGatewayOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkGateway
}

// GetGateway ...
func (c WorkloadNetworksClient) GetGateway(ctx context.Context, id GatewayId) (result GetGatewayOperationResponse, err error) {
	req, err := c.preparerForGetGateway(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetGateway", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetGateway", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetGateway(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetGateway", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetGateway prepares the GetGateway request.
func (c WorkloadNetworksClient) preparerForGetGateway(ctx context.Context, id GatewayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetGateway handles the response to the GetGateway request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetGateway(resp *http.Response) (result GetGatewayOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
