package cognitiveservicescommitmentplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CommitmentPlanId{}

// CommitmentPlanId is a struct representing the Resource ID for a Commitment Plan
type CommitmentPlanId struct {
	SubscriptionId     string
	ResourceGroupName  string
	CommitmentPlanName string
}

// NewCommitmentPlanID returns a new CommitmentPlanId struct
func NewCommitmentPlanID(subscriptionId string, resourceGroupName string, commitmentPlanName string) CommitmentPlanId {
	return CommitmentPlanId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		CommitmentPlanName: commitmentPlanName,
	}
}

// ParseCommitmentPlanID parses 'input' into a CommitmentPlanId
func ParseCommitmentPlanID(input string) (*CommitmentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommitmentPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommitmentPlanId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commitmentPlanName", *parsed)
	}

	return &id, nil
}

// ParseCommitmentPlanIDInsensitively parses 'input' case-insensitively into a CommitmentPlanId
// note: this method should only be used for API response data and not user input
func ParseCommitmentPlanIDInsensitively(input string) (*CommitmentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommitmentPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommitmentPlanId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commitmentPlanName", *parsed)
	}

	return &id, nil
}

// ValidateCommitmentPlanID checks that 'input' can be parsed as a Commitment Plan ID
func ValidateCommitmentPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommitmentPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Commitment Plan ID
func (id CommitmentPlanId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/commitmentPlans/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CommitmentPlanName)
}

// Segments returns a slice of Resource ID Segments which comprise this Commitment Plan ID
func (id CommitmentPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticCommitmentPlans", "commitmentPlans", "commitmentPlans"),
		resourceids.UserSpecifiedSegment("commitmentPlanName", "commitmentPlanValue"),
	}
}

// String returns a human-readable description of this Commitment Plan ID
func (id CommitmentPlanId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Commitment Plan Name: %q", id.CommitmentPlanName),
	}
	return fmt.Sprintf("Commitment Plan (%s)", strings.Join(components, "\n"))
}
