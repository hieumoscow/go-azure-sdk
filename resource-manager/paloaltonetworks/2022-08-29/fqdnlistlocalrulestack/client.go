package fqdnlistlocalrulestack

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FqdnListLocalRulestackClient struct {
	Client *resourcemanager.Client
}

func NewFqdnListLocalRulestackClientWithBaseURI(api environments.Api) (*FqdnListLocalRulestackClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "fqdnlistlocalrulestack", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FqdnListLocalRulestackClient: %+v", err)
	}

	return &FqdnListLocalRulestackClient{
		Client: client,
	}, nil
}
