package serverendpointresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerEndpointsCreateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ServerEndpointsCreate ...
func (c ServerEndpointResourceClient) ServerEndpointsCreate(ctx context.Context, id ServerEndpointId, input ServerEndpointCreateParameters) (result ServerEndpointsCreateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// ServerEndpointsCreateThenPoll performs ServerEndpointsCreate then polls until it's completed
func (c ServerEndpointResourceClient) ServerEndpointsCreateThenPoll(ctx context.Context, id ServerEndpointId, input ServerEndpointCreateParameters) error {
	result, err := c.ServerEndpointsCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServerEndpointsCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ServerEndpointsCreate: %+v", err)
	}

	return nil
}
