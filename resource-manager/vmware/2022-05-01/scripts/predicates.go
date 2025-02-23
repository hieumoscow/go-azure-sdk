package scripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptCmdletOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptCmdletOperationPredicate) Matches(input ScriptCmdlet) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type ScriptExecutionOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptExecutionOperationPredicate) Matches(input ScriptExecution) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}

type ScriptPackageOperationPredicate struct {
	Id   *string
	Name *string
	Type *string
}

func (p ScriptPackageOperationPredicate) Matches(input ScriptPackage) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil && *p.Name != *input.Name) {
		return false
	}

	if p.Type != nil && (input.Type == nil && *p.Type != *input.Type) {
		return false
	}

	return true
}
