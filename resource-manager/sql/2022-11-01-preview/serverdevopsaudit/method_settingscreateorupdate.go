package serverdevopsaudit

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

type SettingsCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// SettingsCreateOrUpdate ...
func (c ServerDevOpsAuditClient) SettingsCreateOrUpdate(ctx context.Context, id ServerId, input ServerDevOpsAuditingSettings) (result SettingsCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/devOpsAuditingSettings/default", id.ID()),
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

// SettingsCreateOrUpdateThenPoll performs SettingsCreateOrUpdate then polls until it's completed
func (c ServerDevOpsAuditClient) SettingsCreateOrUpdateThenPoll(ctx context.Context, id ServerId, input ServerDevOpsAuditingSettings) error {
	result, err := c.SettingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SettingsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SettingsCreateOrUpdate: %+v", err)
	}

	return nil
}
