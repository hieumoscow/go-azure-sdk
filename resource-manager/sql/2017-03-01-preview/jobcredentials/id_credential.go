package jobcredentials

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CredentialId{}

// CredentialId is a struct representing the Resource ID for a Credential
type CredentialId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	JobAgentName      string
	CredentialName    string
}

// NewCredentialID returns a new CredentialId struct
func NewCredentialID(subscriptionId string, resourceGroupName string, serverName string, jobAgentName string, credentialName string) CredentialId {
	return CredentialId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		JobAgentName:      jobAgentName,
		CredentialName:    credentialName,
	}
}

// ParseCredentialID parses 'input' into a CredentialId
func ParseCredentialID(input string) (*CredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(CredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CredentialId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.CredentialName, ok = parsed.Parsed["credentialName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "credentialName", *parsed)
	}

	return &id, nil
}

// ParseCredentialIDInsensitively parses 'input' case-insensitively into a CredentialId
// note: this method should only be used for API response data and not user input
func ParseCredentialIDInsensitively(input string) (*CredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(CredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CredentialId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.JobAgentName, ok = parsed.Parsed["jobAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jobAgentName", *parsed)
	}

	if id.CredentialName, ok = parsed.Parsed["credentialName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "credentialName", *parsed)
	}

	return &id, nil
}

// ValidateCredentialID checks that 'input' can be parsed as a Credential ID
func ValidateCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Credential ID
func (id CredentialId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/jobAgents/%s/credentials/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.JobAgentName, id.CredentialName)
}

// Segments returns a slice of Resource ID Segments which comprise this Credential ID
func (id CredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticJobAgents", "jobAgents", "jobAgents"),
		resourceids.UserSpecifiedSegment("jobAgentName", "jobAgentValue"),
		resourceids.StaticSegment("staticCredentials", "credentials", "credentials"),
		resourceids.UserSpecifiedSegment("credentialName", "credentialValue"),
	}
}

// String returns a human-readable description of this Credential ID
func (id CredentialId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Job Agent Name: %q", id.JobAgentName),
		fmt.Sprintf("Credential Name: %q", id.CredentialName),
	}
	return fmt.Sprintf("Credential (%s)", strings.Join(components, "\n"))
}
