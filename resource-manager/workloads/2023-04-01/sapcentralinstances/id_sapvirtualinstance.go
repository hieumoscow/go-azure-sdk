package sapcentralinstances

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SapVirtualInstanceId{}

// SapVirtualInstanceId is a struct representing the Resource ID for a Sap Virtual Instance
type SapVirtualInstanceId struct {
	SubscriptionId         string
	ResourceGroupName      string
	SapVirtualInstanceName string
}

// NewSapVirtualInstanceID returns a new SapVirtualInstanceId struct
func NewSapVirtualInstanceID(subscriptionId string, resourceGroupName string, sapVirtualInstanceName string) SapVirtualInstanceId {
	return SapVirtualInstanceId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		SapVirtualInstanceName: sapVirtualInstanceName,
	}
}

// ParseSapVirtualInstanceID parses 'input' into a SapVirtualInstanceId
func ParseSapVirtualInstanceID(input string) (*SapVirtualInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SapVirtualInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SapVirtualInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SapVirtualInstanceName, ok = parsed.Parsed["sapVirtualInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sapVirtualInstanceName", *parsed)
	}

	return &id, nil
}

// ParseSapVirtualInstanceIDInsensitively parses 'input' case-insensitively into a SapVirtualInstanceId
// note: this method should only be used for API response data and not user input
func ParseSapVirtualInstanceIDInsensitively(input string) (*SapVirtualInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SapVirtualInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SapVirtualInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SapVirtualInstanceName, ok = parsed.Parsed["sapVirtualInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sapVirtualInstanceName", *parsed)
	}

	return &id, nil
}

// ValidateSapVirtualInstanceID checks that 'input' can be parsed as a Sap Virtual Instance ID
func ValidateSapVirtualInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSapVirtualInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sap Virtual Instance ID
func (id SapVirtualInstanceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Workloads/sapVirtualInstances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SapVirtualInstanceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sap Virtual Instance ID
func (id SapVirtualInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWorkloads", "Microsoft.Workloads", "Microsoft.Workloads"),
		resourceids.StaticSegment("staticSapVirtualInstances", "sapVirtualInstances", "sapVirtualInstances"),
		resourceids.UserSpecifiedSegment("sapVirtualInstanceName", "sapVirtualInstanceValue"),
	}
}

// String returns a human-readable description of this Sap Virtual Instance ID
func (id SapVirtualInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Sap Virtual Instance Name: %q", id.SapVirtualInstanceName),
	}
	return fmt.Sprintf("Sap Virtual Instance (%s)", strings.Join(components, "\n"))
}
