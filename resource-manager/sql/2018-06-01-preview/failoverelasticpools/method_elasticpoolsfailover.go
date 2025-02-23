package failoverelasticpools

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

type ElasticPoolsFailoverOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ElasticPoolsFailover ...
func (c FailoverElasticPoolsClient) ElasticPoolsFailover(ctx context.Context, id ElasticPoolId) (result ElasticPoolsFailoverOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/failover", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

// ElasticPoolsFailoverThenPoll performs ElasticPoolsFailover then polls until it's completed
func (c FailoverElasticPoolsClient) ElasticPoolsFailoverThenPoll(ctx context.Context, id ElasticPoolId) error {
	result, err := c.ElasticPoolsFailover(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ElasticPoolsFailover: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ElasticPoolsFailover: %+v", err)
	}

	return nil
}
