package storageaccounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountsClient struct {
	Client *resourcemanager.Client
}

func NewStorageAccountsClientWithBaseURI(api environments.Api) (*StorageAccountsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "storageaccounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StorageAccountsClient: %+v", err)
	}

	return &StorageAccountsClient{
		Client: client,
	}, nil
}
