package restorabledroppedmanageddatabases

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RestorableDroppedDatabaseId{}

// RestorableDroppedDatabaseId is a struct representing the Resource ID for a Restorable Dropped Database
type RestorableDroppedDatabaseId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ManagedInstanceName         string
	RestorableDroppedDatabaseId string
}

// NewRestorableDroppedDatabaseID returns a new RestorableDroppedDatabaseId struct
func NewRestorableDroppedDatabaseID(subscriptionId string, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseId string) RestorableDroppedDatabaseId {
	return RestorableDroppedDatabaseId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ManagedInstanceName:         managedInstanceName,
		RestorableDroppedDatabaseId: restorableDroppedDatabaseId,
	}
}

// ParseRestorableDroppedDatabaseID parses 'input' into a RestorableDroppedDatabaseId
func ParseRestorableDroppedDatabaseID(input string) (*RestorableDroppedDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorableDroppedDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorableDroppedDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.RestorableDroppedDatabaseId, ok = parsed.Parsed["restorableDroppedDatabaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "restorableDroppedDatabaseId", *parsed)
	}

	return &id, nil
}

// ParseRestorableDroppedDatabaseIDInsensitively parses 'input' case-insensitively into a RestorableDroppedDatabaseId
// note: this method should only be used for API response data and not user input
func ParseRestorableDroppedDatabaseIDInsensitively(input string) (*RestorableDroppedDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorableDroppedDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorableDroppedDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.RestorableDroppedDatabaseId, ok = parsed.Parsed["restorableDroppedDatabaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "restorableDroppedDatabaseId", *parsed)
	}

	return &id, nil
}

// ValidateRestorableDroppedDatabaseID checks that 'input' can be parsed as a Restorable Dropped Database ID
func ValidateRestorableDroppedDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorableDroppedDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restorable Dropped Database ID
func (id RestorableDroppedDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/restorableDroppedDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.RestorableDroppedDatabaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Restorable Dropped Database ID
func (id RestorableDroppedDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticRestorableDroppedDatabases", "restorableDroppedDatabases", "restorableDroppedDatabases"),
		resourceids.UserSpecifiedSegment("restorableDroppedDatabaseId", "restorableDroppedDatabaseIdValue"),
	}
}

// String returns a human-readable description of this Restorable Dropped Database ID
func (id RestorableDroppedDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Restorable Dropped Database: %q", id.RestorableDroppedDatabaseId),
	}
	return fmt.Sprintf("Restorable Dropped Database (%s)", strings.Join(components, "\n"))
}
