package testjobstream

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

type ListByTestJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]JobStream

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListByTestJobOperationResponse, error)
}

type ListByTestJobCompleteResult struct {
	Items []JobStream
}

func (r ListByTestJobOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListByTestJobOperationResponse) LoadMore(ctx context.Context) (resp ListByTestJobOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

type ListByTestJobOperationOptions struct {
	Filter *string
}

func DefaultListByTestJobOperationOptions() ListByTestJobOperationOptions {
	return ListByTestJobOperationOptions{}
}

func (o ListByTestJobOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByTestJobOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByTestJob ...
func (c TestJobStreamClient) ListByTestJob(ctx context.Context, id RunbookId, options ListByTestJobOperationOptions) (resp ListByTestJobOperationResponse, err error) {
	req, err := c.preparerForListByTestJob(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForListByTestJob(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForListByTestJob prepares the ListByTestJob request.
func (c TestJobStreamClient) preparerForListByTestJob(ctx context.Context, id RunbookId, options ListByTestJobOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/draft/testJob/streams", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListByTestJobWithNextLink prepares the ListByTestJob request with the given nextLink token.
func (c TestJobStreamClient) preparerForListByTestJobWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
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

// responderForListByTestJob handles the response to the ListByTestJob request. The method always
// closes the http.Response Body.
func (c TestJobStreamClient) responderForListByTestJob(resp *http.Response) (result ListByTestJobOperationResponse, err error) {
	type page struct {
		Values   []JobStream `json:"value"`
		NextLink *string     `json:"nextLink"`
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
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListByTestJobOperationResponse, err error) {
			req, err := c.preparerForListByTestJobWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForListByTestJob(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "testjobstream.TestJobStreamClient", "ListByTestJob", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListByTestJobComplete retrieves all of the results into a single object
func (c TestJobStreamClient) ListByTestJobComplete(ctx context.Context, id RunbookId, options ListByTestJobOperationOptions) (ListByTestJobCompleteResult, error) {
	return c.ListByTestJobCompleteMatchingPredicate(ctx, id, options, JobStreamOperationPredicate{})
}

// ListByTestJobCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c TestJobStreamClient) ListByTestJobCompleteMatchingPredicate(ctx context.Context, id RunbookId, options ListByTestJobOperationOptions, predicate JobStreamOperationPredicate) (resp ListByTestJobCompleteResult, err error) {
	items := make([]JobStream, 0)

	page, err := c.ListByTestJob(ctx, id, options)
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

	out := ListByTestJobCompleteResult{
		Items: items,
	}
	return out, nil
}
