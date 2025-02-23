package recoverablemanageddatabases

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoverableManagedDatabasesClient struct {
	Client *resourcemanager.Client
}

func NewRecoverableManagedDatabasesClientWithBaseURI(api environments.Api) (*RecoverableManagedDatabasesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "recoverablemanageddatabases", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoverableManagedDatabasesClient: %+v", err)
	}

	return &RecoverableManagedDatabasesClient{
		Client: client,
	}, nil
}
