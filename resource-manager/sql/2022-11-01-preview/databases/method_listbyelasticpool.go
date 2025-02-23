package databases

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByElasticPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Database
}

type ListByElasticPoolCompleteResult struct {
	Items []Database
}

// ListByElasticPool ...
func (c DatabasesClient) ListByElasticPool(ctx context.Context, id ElasticPoolId) (result ListByElasticPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/databases", id.ID()),
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
		Values *[]Database `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByElasticPoolComplete retrieves all the results into a single object
func (c DatabasesClient) ListByElasticPoolComplete(ctx context.Context, id ElasticPoolId) (ListByElasticPoolCompleteResult, error) {
	return c.ListByElasticPoolCompleteMatchingPredicate(ctx, id, DatabaseOperationPredicate{})
}

// ListByElasticPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DatabasesClient) ListByElasticPoolCompleteMatchingPredicate(ctx context.Context, id ElasticPoolId, predicate DatabaseOperationPredicate) (result ListByElasticPoolCompleteResult, err error) {
	items := make([]Database, 0)

	resp, err := c.ListByElasticPool(ctx, id)
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

	result = ListByElasticPoolCompleteResult{
		Items: items,
	}
	return
}
