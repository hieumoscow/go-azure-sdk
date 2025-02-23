package services

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServiceId{}

// ServiceId is a struct representing the Resource ID for a Service
type ServiceId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ManagedClusterName string
	ApplicationName    string
	ServiceName        string
}

// NewServiceID returns a new ServiceId struct
func NewServiceID(subscriptionId string, resourceGroupName string, managedClusterName string, applicationName string, serviceName string) ServiceId {
	return ServiceId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ManagedClusterName: managedClusterName,
		ApplicationName:    applicationName,
		ServiceName:        serviceName,
	}
}

// ParseServiceID parses 'input' into a ServiceId
func ParseServiceID(input string) (*ServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedClusterName, ok = parsed.Parsed["managedClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", *parsed)
	}

	if id.ApplicationName, ok = parsed.Parsed["applicationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	return &id, nil
}

// ParseServiceIDInsensitively parses 'input' case-insensitively into a ServiceId
// note: this method should only be used for API response data and not user input
func ParseServiceIDInsensitively(input string) (*ServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedClusterName, ok = parsed.Parsed["managedClusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedClusterName", *parsed)
	}

	if id.ApplicationName, ok = parsed.Parsed["applicationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	return &id, nil
}

// ValidateServiceID checks that 'input' can be parsed as a Service ID
func ValidateServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service ID
func (id ServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ServiceFabric/managedClusters/%s/applications/%s/services/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedClusterName, id.ApplicationName, id.ServiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service ID
func (id ServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftServiceFabric", "Microsoft.ServiceFabric", "Microsoft.ServiceFabric"),
		resourceids.StaticSegment("staticManagedClusters", "managedClusters", "managedClusters"),
		resourceids.UserSpecifiedSegment("managedClusterName", "managedClusterValue"),
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationName", "applicationValue"),
		resourceids.StaticSegment("staticServices", "services", "services"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
	}
}

// String returns a human-readable description of this Service ID
func (id ServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Cluster Name: %q", id.ManagedClusterName),
		fmt.Sprintf("Application Name: %q", id.ApplicationName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
	}
	return fmt.Sprintf("Service (%s)", strings.Join(components, "\n"))
}
