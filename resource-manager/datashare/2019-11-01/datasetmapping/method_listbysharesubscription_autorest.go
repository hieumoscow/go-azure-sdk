package datasetmapping

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]DataSetMapping

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByShareSubscriptionOperationResponse, error)
}

type ListByShareSubscriptionCompleteResult struct {
	Items []DataSetMapping
}

func (r ListByShareSubscriptionOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByShareSubscriptionOperationResponse) LoadMore(ctx context.Context) (resp ListByShareSubscriptionOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByShareSubscriptionOperationOptions struct {
	Filter  *string
	Orderby *string
}

func DefaultListByShareSubscriptionOperationOptions() ListByShareSubscriptionOperationOptions {
	return ListByShareSubscriptionOperationOptions{}
}

func (o ListByShareSubscriptionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByShareSubscriptionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	return out
}

// ListByShareSubscription ...
func (c DataSetMappingClient) ListByShareSubscription(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions) (resp ListByShareSubscriptionOperationResponse, err error) {
	req, err := c.preparerForListByShareSubscription(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByShareSubscription(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByShareSubscription prepares the ListByShareSubscription request.
func (c DataSetMappingClient) preparerForListByShareSubscription(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/dataSetMappings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByShareSubscriptionWithNextLink prepares the ListByShareSubscription request with the given nextLink token.
func (c DataSetMappingClient) preparerForListByShareSubscriptionWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByShareSubscription handles the response to the ListByShareSubscription request. The method always
// closes the http.Response Body.
func (c DataSetMappingClient) responderForListByShareSubscription(resp *http.Response) (result ListByShareSubscriptionOperationResponse, err error) {
	type page struct {
		Values   []json.RawMessage `json:"value"`
		NextLink *string           `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	temp := make([]DataSetMapping, 0)
	for i, v := range respObj.Values {
		val, err := unmarshalDataSetMappingImplementation(v)
		if err != nil {
			err = fmt.Errorf("unmarshalling item %d for DataSetMapping (%q): %+v", i, v, err)
			return result, err
		}
		temp = append(temp, val)
	}
	result.Model = &temp
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByShareSubscriptionOperationResponse, err error) {
			req, err := c.preparerForListByShareSubscriptionWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByShareSubscription(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "datasetmapping.DataSetMappingClient", "ListByShareSubscription", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByShareSubscriptionComplete retrieves all of the results into a single object
func (c DataSetMappingClient) ListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions) (ListByShareSubscriptionCompleteResult, error) {
	return c.ListByShareSubscriptionCompleteMatchingPredicate(ctx, id, options, DataSetMappingOperationPredicate{})
}

// ListByShareSubscriptionCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c DataSetMappingClient) ListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions, predicate DataSetMappingOperationPredicate) (resp ListByShareSubscriptionCompleteResult, err error) {
	items := make([]DataSetMapping, 0)

	page, err := c.ListByShareSubscription(ctx, id, options)
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

	out := ListByShareSubscriptionCompleteResult{
		Items: items,
	}
	return out, nil
}
