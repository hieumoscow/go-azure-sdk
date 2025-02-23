package containers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = StorageAccountId{}

// StorageAccountId is a struct representing the Resource ID for a Storage Account
type StorageAccountId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	StorageAccountName    string
}

// NewStorageAccountID returns a new StorageAccountId struct
func NewStorageAccountID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, storageAccountName string) StorageAccountId {
	return StorageAccountId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		StorageAccountName:    storageAccountName,
	}
}

// ParseStorageAccountID parses 'input' into a StorageAccountId
func ParseStorageAccountID(input string) (*StorageAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StorageAccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", *parsed)
	}

	if id.StorageAccountName, ok = parsed.Parsed["storageAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageAccountName", *parsed)
	}

	return &id, nil
}

// ParseStorageAccountIDInsensitively parses 'input' case-insensitively into a StorageAccountId
// note: this method should only be used for API response data and not user input
func ParseStorageAccountIDInsensitively(input string) (*StorageAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(StorageAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := StorageAccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DataBoxEdgeDeviceName, ok = parsed.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", *parsed)
	}

	if id.StorageAccountName, ok = parsed.Parsed["storageAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageAccountName", *parsed)
	}

	return &id, nil
}

// ValidateStorageAccountID checks that 'input' can be parsed as a Storage Account ID
func ValidateStorageAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStorageAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Storage Account ID
func (id StorageAccountId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/storageAccounts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.StorageAccountName)
}

// Segments returns a slice of Resource ID Segments which comprise this Storage Account ID
func (id StorageAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticStorageAccounts", "storageAccounts", "storageAccounts"),
		resourceids.UserSpecifiedSegment("storageAccountName", "storageAccountValue"),
	}
}

// String returns a human-readable description of this Storage Account ID
func (id StorageAccountId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Storage Account Name: %q", id.StorageAccountName),
	}
	return fmt.Sprintf("Storage Account (%s)", strings.Join(components, "\n"))
}
