package expressroutelinks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpressRouteLinksClient struct {
	Client *resourcemanager.Client
}

func NewExpressRouteLinksClientWithBaseURI(api environments.Api) (*ExpressRouteLinksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "expressroutelinks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExpressRouteLinksClient: %+v", err)
	}

	return &ExpressRouteLinksClient{
		Client: client,
	}, nil
}
