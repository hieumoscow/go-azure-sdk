package tenantbackfill

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantBackfillClient struct {
	Client *resourcemanager.Client
}

func NewTenantBackfillClientWithBaseURI(api environments.Api) (*TenantBackfillClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "tenantbackfill", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TenantBackfillClient: %+v", err)
	}

	return &TenantBackfillClient{
		Client: client,
	}, nil
}
