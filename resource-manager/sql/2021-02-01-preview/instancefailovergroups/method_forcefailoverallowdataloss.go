package instancefailovergroups

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

type ForceFailoverAllowDataLossOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ForceFailoverAllowDataLoss ...
func (c InstanceFailoverGroupsClient) ForceFailoverAllowDataLoss(ctx context.Context, id InstanceFailoverGroupId) (result ForceFailoverAllowDataLossOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/forceFailoverAllowDataLoss", id.ID()),
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

// ForceFailoverAllowDataLossThenPoll performs ForceFailoverAllowDataLoss then polls until it's completed
func (c InstanceFailoverGroupsClient) ForceFailoverAllowDataLossThenPoll(ctx context.Context, id InstanceFailoverGroupId) error {
	result, err := c.ForceFailoverAllowDataLoss(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ForceFailoverAllowDataLoss: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ForceFailoverAllowDataLoss: %+v", err)
	}

	return nil
}
