package dataflows

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowsClient struct {
	Client *resourcemanager.Client
}

func NewDataFlowsClientWithBaseURI(api environments.Api) (*DataFlowsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "dataflows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataFlowsClient: %+v", err)
	}

	return &DataFlowsClient{
		Client: client,
	}, nil
}
