package firewalls

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallsClient struct {
	Client *resourcemanager.Client
}

func NewFirewallsClientWithBaseURI(api environments.Api) (*FirewallsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "firewalls", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FirewallsClient: %+v", err)
	}

	return &FirewallsClient{
		Client: client,
	}, nil
}
