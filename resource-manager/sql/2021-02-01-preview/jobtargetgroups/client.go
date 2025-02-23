package jobtargetgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobTargetGroupsClient struct {
	Client *resourcemanager.Client
}

func NewJobTargetGroupsClientWithBaseURI(api environments.Api) (*JobTargetGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobtargetgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobTargetGroupsClient: %+v", err)
	}

	return &JobTargetGroupsClient{
		Client: client,
	}, nil
}
