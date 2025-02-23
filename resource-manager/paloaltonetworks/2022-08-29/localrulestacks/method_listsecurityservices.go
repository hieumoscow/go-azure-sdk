package localrulestacks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSecurityServicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SecurityServicesResponse
}

type ListSecurityServicesOperationOptions struct {
	Skip *string
	Top  *int64
	Type *SecurityServicesTypeEnum
}

func DefaultListSecurityServicesOperationOptions() ListSecurityServicesOperationOptions {
	return ListSecurityServicesOperationOptions{}
}

func (o ListSecurityServicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSecurityServicesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListSecurityServicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Skip != nil {
		out.Append("skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("top", fmt.Sprintf("%v", *o.Top))
	}
	if o.Type != nil {
		out.Append("type", fmt.Sprintf("%v", *o.Type))
	}
	return &out
}

// ListSecurityServices ...
func (c LocalRuleStacksClient) ListSecurityServices(ctx context.Context, id LocalRuleStackId, options ListSecurityServicesOperationOptions) (result ListSecurityServicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/listSecurityServices", id.ID()),
		OptionsObject: options,
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
