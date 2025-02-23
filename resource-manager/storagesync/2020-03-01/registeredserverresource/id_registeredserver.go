package registeredserverresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegisteredServerId{}

// RegisteredServerId is a struct representing the Resource ID for a Registered Server
type RegisteredServerId struct {
	SubscriptionId         string
	ResourceGroupName      string
	StorageSyncServiceName string
	ServerId               string
}

// NewRegisteredServerID returns a new RegisteredServerId struct
func NewRegisteredServerID(subscriptionId string, resourceGroupName string, storageSyncServiceName string, serverId string) RegisteredServerId {
	return RegisteredServerId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		StorageSyncServiceName: storageSyncServiceName,
		ServerId:               serverId,
	}
}

// ParseRegisteredServerID parses 'input' into a RegisteredServerId
func ParseRegisteredServerID(input string) (*RegisteredServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegisteredServerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegisteredServerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageSyncServiceName", *parsed)
	}

	if id.ServerId, ok = parsed.Parsed["serverId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverId", *parsed)
	}

	return &id, nil
}

// ParseRegisteredServerIDInsensitively parses 'input' case-insensitively into a RegisteredServerId
// note: this method should only be used for API response data and not user input
func ParseRegisteredServerIDInsensitively(input string) (*RegisteredServerId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegisteredServerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegisteredServerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageSyncServiceName", *parsed)
	}

	if id.ServerId, ok = parsed.Parsed["serverId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverId", *parsed)
	}

	return &id, nil
}

// ValidateRegisteredServerID checks that 'input' can be parsed as a Registered Server ID
func ValidateRegisteredServerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegisteredServerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registered Server ID
func (id RegisteredServerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageSync/storageSyncServices/%s/registeredServers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageSyncServiceName, id.ServerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Registered Server ID
func (id RegisteredServerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorageSync", "Microsoft.StorageSync", "Microsoft.StorageSync"),
		resourceids.StaticSegment("staticStorageSyncServices", "storageSyncServices", "storageSyncServices"),
		resourceids.UserSpecifiedSegment("storageSyncServiceName", "storageSyncServiceValue"),
		resourceids.StaticSegment("staticRegisteredServers", "registeredServers", "registeredServers"),
		resourceids.UserSpecifiedSegment("serverId", "serverIdValue"),
	}
}

// String returns a human-readable description of this Registered Server ID
func (id RegisteredServerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Sync Service Name: %q", id.StorageSyncServiceName),
		fmt.Sprintf("Server: %q", id.ServerId),
	}
	return fmt.Sprintf("Registered Server (%s)", strings.Join(components, "\n"))
}
