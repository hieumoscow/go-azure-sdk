package privateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionProperties struct {
	PrivateEndpoint                   *PrivateEndpointProperty                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `json:"privateLinkServiceConnectionState,omitempty"`
	ProvisioningState                 *PrivateEndpointProvisioningState          `json:"provisioningState,omitempty"`
}
