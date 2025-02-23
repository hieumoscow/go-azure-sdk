package resetcifspassword

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VolumeId{}

// VolumeId is a struct representing the Resource ID for a Volume
type VolumeId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	CapacityPoolName  string
	VolumeName        string
}

// NewVolumeID returns a new VolumeId struct
func NewVolumeID(subscriptionId string, resourceGroupName string, netAppAccountName string, capacityPoolName string, volumeName string) VolumeId {
	return VolumeId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		CapacityPoolName:  capacityPoolName,
		VolumeName:        volumeName,
	}
}

// ParseVolumeID parses 'input' into a VolumeId
func ParseVolumeID(input string) (*VolumeId, error) {
	parser := resourceids.NewParserFromResourceIdType(VolumeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VolumeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", *parsed)
	}

	if id.CapacityPoolName, ok = parsed.Parsed["capacityPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", *parsed)
	}

	if id.VolumeName, ok = parsed.Parsed["volumeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "volumeName", *parsed)
	}

	return &id, nil
}

// ParseVolumeIDInsensitively parses 'input' case-insensitively into a VolumeId
// note: this method should only be used for API response data and not user input
func ParseVolumeIDInsensitively(input string) (*VolumeId, error) {
	parser := resourceids.NewParserFromResourceIdType(VolumeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VolumeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NetAppAccountName, ok = parsed.Parsed["netAppAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", *parsed)
	}

	if id.CapacityPoolName, ok = parsed.Parsed["capacityPoolName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", *parsed)
	}

	if id.VolumeName, ok = parsed.Parsed["volumeName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "volumeName", *parsed)
	}

	return &id, nil
}

// ValidateVolumeID checks that 'input' can be parsed as a Volume ID
func ValidateVolumeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVolumeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Volume ID
func (id VolumeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/capacityPools/%s/volumes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.CapacityPoolName, id.VolumeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Volume ID
func (id VolumeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountValue"),
		resourceids.StaticSegment("staticCapacityPools", "capacityPools", "capacityPools"),
		resourceids.UserSpecifiedSegment("capacityPoolName", "capacityPoolValue"),
		resourceids.StaticSegment("staticVolumes", "volumes", "volumes"),
		resourceids.UserSpecifiedSegment("volumeName", "volumeValue"),
	}
}

// String returns a human-readable description of this Volume ID
func (id VolumeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Capacity Pool Name: %q", id.CapacityPoolName),
		fmt.Sprintf("Volume Name: %q", id.VolumeName),
	}
	return fmt.Sprintf("Volume (%s)", strings.Join(components, "\n"))
}
