package appplatform

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

type ServiceRegistriesCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ServiceRegistriesCreateOrUpdate ...
func (c AppPlatformClient) ServiceRegistriesCreateOrUpdate(ctx context.Context, id ServiceRegistryId) (result ServiceRegistriesCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

// ServiceRegistriesCreateOrUpdateThenPoll performs ServiceRegistriesCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) ServiceRegistriesCreateOrUpdateThenPoll(ctx context.Context, id ServiceRegistryId) error {
	result, err := c.ServiceRegistriesCreateOrUpdate(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServiceRegistriesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ServiceRegistriesCreateOrUpdate: %+v", err)
	}

	return nil
}
