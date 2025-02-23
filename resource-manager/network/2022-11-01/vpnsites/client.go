package vpnsites

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnSitesClient struct {
	Client *resourcemanager.Client
}

func NewVpnSitesClientWithBaseURI(api environments.Api) (*VpnSitesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "vpnsites", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VpnSitesClient: %+v", err)
	}

	return &VpnSitesClient{
		Client: client,
	}, nil
}
