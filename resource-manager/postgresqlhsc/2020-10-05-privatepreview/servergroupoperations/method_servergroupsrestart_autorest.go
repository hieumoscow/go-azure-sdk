package servergroupoperations

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

type ServerGroupsRestartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServerGroupsRestart ...
func (c ServerGroupOperationsClient) ServerGroupsRestart(ctx context.Context, id ServerGroupsv2Id) (result ServerGroupsRestartOperationResponse, err error) {
	req, err := c.preparerForServerGroupsRestart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsRestart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServerGroupsRestart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servergroupoperations.ServerGroupOperationsClient", "ServerGroupsRestart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServerGroupsRestartThenPoll performs ServerGroupsRestart then polls until it's completed
func (c ServerGroupOperationsClient) ServerGroupsRestartThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ServerGroupsRestart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServerGroupsRestart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServerGroupsRestart: %+v", err)
	}

	return nil
}

// preparerForServerGroupsRestart prepares the ServerGroupsRestart request.
func (c ServerGroupOperationsClient) preparerForServerGroupsRestart(ctx context.Context, id ServerGroupsv2Id) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServerGroupsRestart sends the ServerGroupsRestart request. The method will close the
// http.Response Body if it receives an error.
func (c ServerGroupOperationsClient) senderForServerGroupsRestart(ctx context.Context, req *http.Request) (future ServerGroupsRestartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
