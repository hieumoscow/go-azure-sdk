package restorepoints

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RestorePointId{}

// RestorePointId is a struct representing the Resource ID for a Restore Point
type RestorePointId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	RestorePointName  string
}

// NewRestorePointID returns a new RestorePointId struct
func NewRestorePointID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, restorePointName string) RestorePointId {
	return RestorePointId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		RestorePointName:  restorePointName,
	}
}

// ParseRestorePointID parses 'input' into a RestorePointId
func ParseRestorePointID(input string) (*RestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.RestorePointName, ok = parsed.Parsed["restorePointName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "restorePointName", *parsed)
	}

	return &id, nil
}

// ParseRestorePointIDInsensitively parses 'input' case-insensitively into a RestorePointId
// note: this method should only be used for API response data and not user input
func ParseRestorePointIDInsensitively(input string) (*RestorePointId, error) {
	parser := resourceids.NewParserFromResourceIdType(RestorePointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RestorePointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.RestorePointName, ok = parsed.Parsed["restorePointName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "restorePointName", *parsed)
	}

	return &id, nil
}

// ValidateRestorePointID checks that 'input' can be parsed as a Restore Point ID
func ValidateRestorePointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRestorePointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Restore Point ID
func (id RestorePointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/restorePoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.RestorePointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Restore Point ID
func (id RestorePointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticRestorePoints", "restorePoints", "restorePoints"),
		resourceids.UserSpecifiedSegment("restorePointName", "restorePointValue"),
	}
}

// String returns a human-readable description of this Restore Point ID
func (id RestorePointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Restore Point Name: %q", id.RestorePointName),
	}
	return fmt.Sprintf("Restore Point (%s)", strings.Join(components, "\n"))
}
