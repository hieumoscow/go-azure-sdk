package expressroutecrossconnectionroutetablesummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteCrossConnectionRouteTableSummaryClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteCrossConnectionRouteTableSummaryClientWithBaseURI(api environments.Api) (*ExpressRouteCrossConnectionRouteTableSummaryClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutecrossconnectionroutetablesummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteCrossConnectionRouteTableSummaryClient: %+v", err)
	}

	return &ExpressRouteCrossConnectionRouteTableSummaryClient{
		Client: client,
	}, nil
}
