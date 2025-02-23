package roleeligibilityschedulerequests

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleRequestsClient struct {
	Client *resourcemanager.Client
}

func NewRoleEligibilityScheduleRequestsClientWithBaseURI(api environments.Api) (*RoleEligibilityScheduleRequestsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "roleeligibilityschedulerequests", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoleEligibilityScheduleRequestsClient: %+v", err)
	}

	return &RoleEligibilityScheduleRequestsClient{
		Client: client,
	}, nil
}
