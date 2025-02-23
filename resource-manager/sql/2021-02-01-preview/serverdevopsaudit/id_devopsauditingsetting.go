package serverdevopsaudit

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DevOpsAuditingSettingId{}

// DevOpsAuditingSettingId is a struct representing the Resource ID for a Dev Ops Auditing Setting
type DevOpsAuditingSettingId struct {
	SubscriptionId            string
	ResourceGroupName         string
	ServerName                string
	DevOpsAuditingSettingName string
}

// NewDevOpsAuditingSettingID returns a new DevOpsAuditingSettingId struct
func NewDevOpsAuditingSettingID(subscriptionId string, resourceGroupName string, serverName string, devOpsAuditingSettingName string) DevOpsAuditingSettingId {
	return DevOpsAuditingSettingId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		ServerName:                serverName,
		DevOpsAuditingSettingName: devOpsAuditingSettingName,
	}
}

// ParseDevOpsAuditingSettingID parses 'input' into a DevOpsAuditingSettingId
func ParseDevOpsAuditingSettingID(input string) (*DevOpsAuditingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DevOpsAuditingSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DevOpsAuditingSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DevOpsAuditingSettingName, ok = parsed.Parsed["devOpsAuditingSettingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "devOpsAuditingSettingName", *parsed)
	}

	return &id, nil
}

// ParseDevOpsAuditingSettingIDInsensitively parses 'input' case-insensitively into a DevOpsAuditingSettingId
// note: this method should only be used for API response data and not user input
func ParseDevOpsAuditingSettingIDInsensitively(input string) (*DevOpsAuditingSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(DevOpsAuditingSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DevOpsAuditingSettingId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DevOpsAuditingSettingName, ok = parsed.Parsed["devOpsAuditingSettingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "devOpsAuditingSettingName", *parsed)
	}

	return &id, nil
}

// ValidateDevOpsAuditingSettingID checks that 'input' can be parsed as a Dev Ops Auditing Setting ID
func ValidateDevOpsAuditingSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDevOpsAuditingSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dev Ops Auditing Setting ID
func (id DevOpsAuditingSettingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/devOpsAuditingSettings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DevOpsAuditingSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dev Ops Auditing Setting ID
func (id DevOpsAuditingSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDevOpsAuditingSettings", "devOpsAuditingSettings", "devOpsAuditingSettings"),
		resourceids.UserSpecifiedSegment("devOpsAuditingSettingName", "devOpsAuditingSettingValue"),
	}
}

// String returns a human-readable description of this Dev Ops Auditing Setting ID
func (id DevOpsAuditingSettingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Dev Ops Auditing Setting Name: %q", id.DevOpsAuditingSettingName),
	}
	return fmt.Sprintf("Dev Ops Auditing Setting (%s)", strings.Join(components, "\n"))
}
