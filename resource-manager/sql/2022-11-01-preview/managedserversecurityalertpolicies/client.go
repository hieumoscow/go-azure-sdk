package managedserversecurityalertpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedServerSecurityAlertPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewManagedServerSecurityAlertPoliciesClientWithBaseURI(api environments.Api) (*ManagedServerSecurityAlertPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedserversecurityalertpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedServerSecurityAlertPoliciesClient: %+v", err)
	}

	return &ManagedServerSecurityAlertPoliciesClient{
		Client: client,
	}, nil
}
