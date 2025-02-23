package localrules

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalRulesClient struct {
	Client *resourcemanager.Client
}

func NewLocalRulesClientWithBaseURI(api environments.Api) (*LocalRulesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "localrules", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocalRulesClient: %+v", err)
	}

	return &LocalRulesClient{
		Client: client,
	}, nil
}
