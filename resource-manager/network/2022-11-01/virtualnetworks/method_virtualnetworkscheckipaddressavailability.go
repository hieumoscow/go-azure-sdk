package virtualnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualNetworksCheckIPAddressAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *IPAddressAvailabilityResult
}

type VirtualNetworksCheckIPAddressAvailabilityOperationOptions struct {
	IPAddress *string
}

func DefaultVirtualNetworksCheckIPAddressAvailabilityOperationOptions() VirtualNetworksCheckIPAddressAvailabilityOperationOptions {
	return VirtualNetworksCheckIPAddressAvailabilityOperationOptions{}
}

func (o VirtualNetworksCheckIPAddressAvailabilityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o VirtualNetworksCheckIPAddressAvailabilityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o VirtualNetworksCheckIPAddressAvailabilityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.IPAddress != nil {
		out.Append("ipAddress", fmt.Sprintf("%v", *o.IPAddress))
	}
	return &out
}

// VirtualNetworksCheckIPAddressAvailability ...
func (c VirtualNetworksClient) VirtualNetworksCheckIPAddressAvailability(ctx context.Context, id VirtualNetworkId, options VirtualNetworksCheckIPAddressAvailabilityOperationOptions) (result VirtualNetworksCheckIPAddressAvailabilityOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/checkIPAddressAvailability", id.ID()),
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
