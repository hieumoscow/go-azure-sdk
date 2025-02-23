package modelcontainer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryModelContainersDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryModelContainersDelete ...
func (c ModelContainerClient) RegistryModelContainersDelete(ctx context.Context, id RegistryModelId) (result RegistryModelContainersDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryModelContainersDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryModelContainersDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryModelContainersDeleteThenPoll performs RegistryModelContainersDelete then polls until it's completed
func (c ModelContainerClient) RegistryModelContainersDeleteThenPoll(ctx context.Context, id RegistryModelId) error {
	result, err := c.RegistryModelContainersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryModelContainersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryModelContainersDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryModelContainersDelete prepares the RegistryModelContainersDelete request.
func (c ModelContainerClient) preparerForRegistryModelContainersDelete(ctx context.Context, id RegistryModelId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRegistryModelContainersDelete sends the RegistryModelContainersDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ModelContainerClient) senderForRegistryModelContainersDelete(ctx context.Context, req *http.Request) (future RegistryModelContainersDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
