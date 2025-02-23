package environmentcontainer

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryEnvironmentContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]EnvironmentContainerResource

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (RegistryEnvironmentContainersListOperationResponse, error)
}

type RegistryEnvironmentContainersListCompleteResult struct {
	Items []EnvironmentContainerResource
}

func (r RegistryEnvironmentContainersListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r RegistryEnvironmentContainersListOperationResponse) LoadMore(ctx context.Context) (resp RegistryEnvironmentContainersListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type RegistryEnvironmentContainersListOperationOptions struct {
	ListViewType *ListViewType
	Skip         *string
}

func DefaultRegistryEnvironmentContainersListOperationOptions() RegistryEnvironmentContainersListOperationOptions {
	return RegistryEnvironmentContainersListOperationOptions{}
}

func (o RegistryEnvironmentContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegistryEnvironmentContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ListViewType != nil {
		out["listViewType"] = *o.ListViewType
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	return out
}

// RegistryEnvironmentContainersList ...
func (c EnvironmentContainerClient) RegistryEnvironmentContainersList(ctx context.Context, id RegistryId, options RegistryEnvironmentContainersListOperationOptions) (resp RegistryEnvironmentContainersListOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForRegistryEnvironmentContainersList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForRegistryEnvironmentContainersList prepares the RegistryEnvironmentContainersList request.
func (c EnvironmentContainerClient) preparerForRegistryEnvironmentContainersList(ctx context.Context, id RegistryId, options RegistryEnvironmentContainersListOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/environments", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForRegistryEnvironmentContainersListWithNextLink prepares the RegistryEnvironmentContainersList request with the given nextLink token.
func (c EnvironmentContainerClient) preparerForRegistryEnvironmentContainersListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegistryEnvironmentContainersList handles the response to the RegistryEnvironmentContainersList request. The method always
// closes the http.Response Body.
func (c EnvironmentContainerClient) responderForRegistryEnvironmentContainersList(resp *http.Response) (result RegistryEnvironmentContainersListOperationResponse, err error) {
	type page struct {
		Values   []EnvironmentContainerResource `json:"value"`
		NextLink *string                        `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result RegistryEnvironmentContainersListOperationResponse, err error) {
			req, err := c.preparerForRegistryEnvironmentContainersListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForRegistryEnvironmentContainersList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// RegistryEnvironmentContainersListComplete retrieves all of the results into a single object
func (c EnvironmentContainerClient) RegistryEnvironmentContainersListComplete(ctx context.Context, id RegistryId, options RegistryEnvironmentContainersListOperationOptions) (RegistryEnvironmentContainersListCompleteResult, error) {
	return c.RegistryEnvironmentContainersListCompleteMatchingPredicate(ctx, id, options, EnvironmentContainerResourceOperationPredicate{})
}

// RegistryEnvironmentContainersListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c EnvironmentContainerClient) RegistryEnvironmentContainersListCompleteMatchingPredicate(ctx context.Context, id RegistryId, options RegistryEnvironmentContainersListOperationOptions, predicate EnvironmentContainerResourceOperationPredicate) (resp RegistryEnvironmentContainersListCompleteResult, err error) {
	items := make([]EnvironmentContainerResource, 0)

	page, err := c.RegistryEnvironmentContainersList(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := RegistryEnvironmentContainersListCompleteResult{
		Items: items,
	}
	return out, nil
}
