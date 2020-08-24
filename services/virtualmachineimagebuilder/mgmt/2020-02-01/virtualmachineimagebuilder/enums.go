package virtualmachineimagebuilder

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ProvisioningErrorCode enumerates the values for provisioning error code.
type ProvisioningErrorCode string

const (
	// BadCustomizerType ...
	BadCustomizerType ProvisioningErrorCode = "BadCustomizerType"
	// BadDistributeType ...
	BadDistributeType ProvisioningErrorCode = "BadDistributeType"
	// BadManagedImageSource ...
	BadManagedImageSource ProvisioningErrorCode = "BadManagedImageSource"
	// BadPIRSource ...
	BadPIRSource ProvisioningErrorCode = "BadPIRSource"
	// BadSharedImageDistribute ...
	BadSharedImageDistribute ProvisioningErrorCode = "BadSharedImageDistribute"
	// BadSharedImageVersionSource ...
	BadSharedImageVersionSource ProvisioningErrorCode = "BadSharedImageVersionSource"
	// BadSourceType ...
	BadSourceType ProvisioningErrorCode = "BadSourceType"
	// NoCustomizerScript ...
	NoCustomizerScript ProvisioningErrorCode = "NoCustomizerScript"
	// Other ...
	Other ProvisioningErrorCode = "Other"
	// ServerError ...
	ServerError ProvisioningErrorCode = "ServerError"
	// UnsupportedCustomizerType ...
	UnsupportedCustomizerType ProvisioningErrorCode = "UnsupportedCustomizerType"
)

// PossibleProvisioningErrorCodeValues returns an array of possible values for the ProvisioningErrorCode const type.
func PossibleProvisioningErrorCodeValues() []ProvisioningErrorCode {
	return []ProvisioningErrorCode{BadCustomizerType, BadDistributeType, BadManagedImageSource, BadPIRSource, BadSharedImageDistribute, BadSharedImageVersionSource, BadSourceType, NoCustomizerScript, Other, ServerError, UnsupportedCustomizerType}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Deleting, Failed, Succeeded, Updating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, UserAssigned}
}

// RunState enumerates the values for run state.
type RunState string

const (
	// RunStateCanceled ...
	RunStateCanceled RunState = "Canceled"
	// RunStateCanceling ...
	RunStateCanceling RunState = "Canceling"
	// RunStateFailed ...
	RunStateFailed RunState = "Failed"
	// RunStatePartiallySucceeded ...
	RunStatePartiallySucceeded RunState = "PartiallySucceeded"
	// RunStateRunning ...
	RunStateRunning RunState = "Running"
	// RunStateSucceeded ...
	RunStateSucceeded RunState = "Succeeded"
)

// PossibleRunStateValues returns an array of possible values for the RunState const type.
func PossibleRunStateValues() []RunState {
	return []RunState{RunStateCanceled, RunStateCanceling, RunStateFailed, RunStatePartiallySucceeded, RunStateRunning, RunStateSucceeded}
}

// RunSubState enumerates the values for run sub state.
type RunSubState string

const (
	// Building ...
	Building RunSubState = "Building"
	// Customizing ...
	Customizing RunSubState = "Customizing"
	// Distributing ...
	Distributing RunSubState = "Distributing"
	// Queued ...
	Queued RunSubState = "Queued"
)

// PossibleRunSubStateValues returns an array of possible values for the RunSubState const type.
func PossibleRunSubStateValues() []RunSubState {
	return []RunSubState{Building, Customizing, Distributing, Queued}
}

// SharedImageStorageAccountType enumerates the values for shared image storage account type.
type SharedImageStorageAccountType string

const (
	// StandardLRS ...
	StandardLRS SharedImageStorageAccountType = "Standard_LRS"
	// StandardZRS ...
	StandardZRS SharedImageStorageAccountType = "Standard_ZRS"
)

// PossibleSharedImageStorageAccountTypeValues returns an array of possible values for the SharedImageStorageAccountType const type.
func PossibleSharedImageStorageAccountTypeValues() []SharedImageStorageAccountType {
	return []SharedImageStorageAccountType{StandardLRS, StandardZRS}
}

// Type enumerates the values for type.
type Type string

const (
	// TypeImageTemplateSource ...
	TypeImageTemplateSource Type = "ImageTemplateSource"
	// TypeManagedImage ...
	TypeManagedImage Type = "ManagedImage"
	// TypePlatformImage ...
	TypePlatformImage Type = "PlatformImage"
	// TypeSharedImageVersion ...
	TypeSharedImageVersion Type = "SharedImageVersion"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{TypeImageTemplateSource, TypeManagedImage, TypePlatformImage, TypeSharedImageVersion}
}

// TypeBasicImageTemplateCustomizer enumerates the values for type basic image template customizer.
type TypeBasicImageTemplateCustomizer string

const (
	// TypeFile ...
	TypeFile TypeBasicImageTemplateCustomizer = "File"
	// TypeImageTemplateCustomizer ...
	TypeImageTemplateCustomizer TypeBasicImageTemplateCustomizer = "ImageTemplateCustomizer"
	// TypePowerShell ...
	TypePowerShell TypeBasicImageTemplateCustomizer = "PowerShell"
	// TypeShell ...
	TypeShell TypeBasicImageTemplateCustomizer = "Shell"
	// TypeWindowsRestart ...
	TypeWindowsRestart TypeBasicImageTemplateCustomizer = "WindowsRestart"
	// TypeWindowsUpdate ...
	TypeWindowsUpdate TypeBasicImageTemplateCustomizer = "WindowsUpdate"
)

// PossibleTypeBasicImageTemplateCustomizerValues returns an array of possible values for the TypeBasicImageTemplateCustomizer const type.
func PossibleTypeBasicImageTemplateCustomizerValues() []TypeBasicImageTemplateCustomizer {
	return []TypeBasicImageTemplateCustomizer{TypeFile, TypeImageTemplateCustomizer, TypePowerShell, TypeShell, TypeWindowsRestart, TypeWindowsUpdate}
}

// TypeBasicImageTemplateDistributor enumerates the values for type basic image template distributor.
type TypeBasicImageTemplateDistributor string

const (
	// TypeBasicImageTemplateDistributorTypeImageTemplateDistributor ...
	TypeBasicImageTemplateDistributorTypeImageTemplateDistributor TypeBasicImageTemplateDistributor = "ImageTemplateDistributor"
	// TypeBasicImageTemplateDistributorTypeManagedImage ...
	TypeBasicImageTemplateDistributorTypeManagedImage TypeBasicImageTemplateDistributor = "ManagedImage"
	// TypeBasicImageTemplateDistributorTypeSharedImage ...
	TypeBasicImageTemplateDistributorTypeSharedImage TypeBasicImageTemplateDistributor = "SharedImage"
	// TypeBasicImageTemplateDistributorTypeVHD ...
	TypeBasicImageTemplateDistributorTypeVHD TypeBasicImageTemplateDistributor = "VHD"
)

// PossibleTypeBasicImageTemplateDistributorValues returns an array of possible values for the TypeBasicImageTemplateDistributor const type.
func PossibleTypeBasicImageTemplateDistributorValues() []TypeBasicImageTemplateDistributor {
	return []TypeBasicImageTemplateDistributor{TypeBasicImageTemplateDistributorTypeImageTemplateDistributor, TypeBasicImageTemplateDistributorTypeManagedImage, TypeBasicImageTemplateDistributorTypeSharedImage, TypeBasicImageTemplateDistributorTypeVHD}
}