package longtermretentionmanagedinstancebackups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LocationLongTermRetentionManagedInstanceId{}

// LocationLongTermRetentionManagedInstanceId is a struct representing the Resource ID for a Location Long Term Retention Managed Instance
type LocationLongTermRetentionManagedInstanceId struct {
	SubscriptionId                       string
	ResourceGroupName                    string
	LocationName                         string
	LongTermRetentionManagedInstanceName string
}

// NewLocationLongTermRetentionManagedInstanceID returns a new LocationLongTermRetentionManagedInstanceId struct
func NewLocationLongTermRetentionManagedInstanceID(subscriptionId string, resourceGroupName string, locationName string, longTermRetentionManagedInstanceName string) LocationLongTermRetentionManagedInstanceId {
	return LocationLongTermRetentionManagedInstanceId{
		SubscriptionId:                       subscriptionId,
		ResourceGroupName:                    resourceGroupName,
		LocationName:                         locationName,
		LongTermRetentionManagedInstanceName: longTermRetentionManagedInstanceName,
	}
}

// ParseLocationLongTermRetentionManagedInstanceID parses 'input' into a LocationLongTermRetentionManagedInstanceId
func ParseLocationLongTermRetentionManagedInstanceID(input string) (*LocationLongTermRetentionManagedInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocationLongTermRetentionManagedInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocationLongTermRetentionManagedInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceName, ok = parsed.Parsed["longTermRetentionManagedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceName", *parsed)
	}

	return &id, nil
}

// ParseLocationLongTermRetentionManagedInstanceIDInsensitively parses 'input' case-insensitively into a LocationLongTermRetentionManagedInstanceId
// note: this method should only be used for API response data and not user input
func ParseLocationLongTermRetentionManagedInstanceIDInsensitively(input string) (*LocationLongTermRetentionManagedInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocationLongTermRetentionManagedInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocationLongTermRetentionManagedInstanceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceName, ok = parsed.Parsed["longTermRetentionManagedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceName", *parsed)
	}

	return &id, nil
}

// ValidateLocationLongTermRetentionManagedInstanceID checks that 'input' can be parsed as a Location Long Term Retention Managed Instance ID
func ValidateLocationLongTermRetentionManagedInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationLongTermRetentionManagedInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location Long Term Retention Managed Instance ID
func (id LocationLongTermRetentionManagedInstanceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.LongTermRetentionManagedInstanceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Location Long Term Retention Managed Instance ID
func (id LocationLongTermRetentionManagedInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "longTermRetentionManagedInstanceValue"),
	}
}

// String returns a human-readable description of this Location Long Term Retention Managed Instance ID
func (id LocationLongTermRetentionManagedInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
	}
	return fmt.Sprintf("Location Long Term Retention Managed Instance (%s)", strings.Join(components, "\n"))
}
