package managedinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstance
}

type ListByResourceGroupCompleteResult struct {
	Items []ManagedInstance
}

type ListByResourceGroupOperationOptions struct {
	Expand *string
}

func DefaultListByResourceGroupOperationOptions() ListByResourceGroupOperationOptions {
	return ListByResourceGroupOperationOptions{}
}

func (o ListByResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	return &out
}

// ListByResourceGroup ...
func (c ManagedInstancesClient) ListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ListByResourceGroupOperationOptions) (result ListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Sql/managedInstances", id.ID()),
		OptionsObject: options,
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
		Values *[]ManagedInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByResourceGroupComplete retrieves all the results into a single object
func (c ManagedInstancesClient) ListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options ListByResourceGroupOperationOptions) (ListByResourceGroupCompleteResult, error) {
	return c.ListByResourceGroupCompleteMatchingPredicate(ctx, id, options, ManagedInstanceOperationPredicate{})
}

// ListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedInstancesClient) ListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ListByResourceGroupOperationOptions, predicate ManagedInstanceOperationPredicate) (result ListByResourceGroupCompleteResult, err error) {
	items := make([]ManagedInstance, 0)

	resp, err := c.ListByResourceGroup(ctx, id, options)
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

	result = ListByResourceGroupCompleteResult{
		Items: items,
	}
	return
}
