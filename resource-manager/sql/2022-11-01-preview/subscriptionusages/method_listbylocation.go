package subscriptionusages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SubscriptionUsage
}

type ListByLocationCompleteResult struct {
	Items []SubscriptionUsage
}

// ListByLocation ...
func (c SubscriptionUsagesClient) ListByLocation(ctx context.Context, id LocationId) (result ListByLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/usages", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]SubscriptionUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByLocationComplete retrieves all the results into a single object
func (c SubscriptionUsagesClient) ListByLocationComplete(ctx context.Context, id LocationId) (ListByLocationCompleteResult, error) {
	return c.ListByLocationCompleteMatchingPredicate(ctx, id, SubscriptionUsageOperationPredicate{})
}

// ListByLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SubscriptionUsagesClient) ListByLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate SubscriptionUsageOperationPredicate) (result ListByLocationCompleteResult, err error) {
	items := make([]SubscriptionUsage, 0)

	resp, err := c.ListByLocation(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByLocationCompleteResult{
		Items: items,
	}
	return
}
