package hdinsight

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AsyncOperationState enumerates the values for async operation state.
type AsyncOperationState string

const (
	// Failed ...
	Failed AsyncOperationState = "Failed"
	// InProgress ...
	InProgress AsyncOperationState = "InProgress"
	// Succeeded ...
	Succeeded AsyncOperationState = "Succeeded"
)

// PossibleAsyncOperationStateValues returns an array of possible values for the AsyncOperationState const type.
func PossibleAsyncOperationStateValues() []AsyncOperationState {
	return []AsyncOperationState{Failed, InProgress, Succeeded}
}

// ClusterProvisioningState enumerates the values for cluster provisioning state.
type ClusterProvisioningState string

const (
	// ClusterProvisioningStateCanceled ...
	ClusterProvisioningStateCanceled ClusterProvisioningState = "Canceled"
	// ClusterProvisioningStateDeleting ...
	ClusterProvisioningStateDeleting ClusterProvisioningState = "Deleting"
	// ClusterProvisioningStateFailed ...
	ClusterProvisioningStateFailed ClusterProvisioningState = "Failed"
	// ClusterProvisioningStateInProgress ...
	ClusterProvisioningStateInProgress ClusterProvisioningState = "InProgress"
	// ClusterProvisioningStateSucceeded ...
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
)

// PossibleClusterProvisioningStateValues returns an array of possible values for the ClusterProvisioningState const type.
func PossibleClusterProvisioningStateValues() []ClusterProvisioningState {
	return []ClusterProvisioningState{ClusterProvisioningStateCanceled, ClusterProvisioningStateDeleting, ClusterProvisioningStateFailed, ClusterProvisioningStateInProgress, ClusterProvisioningStateSucceeded}
}

// DaysOfWeek enumerates the values for days of week.
type DaysOfWeek string

const (
	// Friday ...
	Friday DaysOfWeek = "Friday"
	// Monday ...
	Monday DaysOfWeek = "Monday"
	// Saturday ...
	Saturday DaysOfWeek = "Saturday"
	// Sunday ...
	Sunday DaysOfWeek = "Sunday"
	// Thursday ...
	Thursday DaysOfWeek = "Thursday"
	// Tuesday ...
	Tuesday DaysOfWeek = "Tuesday"
	// Wednesday ...
	Wednesday DaysOfWeek = "Wednesday"
)

// PossibleDaysOfWeekValues returns an array of possible values for the DaysOfWeek const type.
func PossibleDaysOfWeekValues() []DaysOfWeek {
	return []DaysOfWeek{Friday, Monday, Saturday, Sunday, Thursday, Tuesday, Wednesday}
}

// DirectoryType enumerates the values for directory type.
type DirectoryType string

const (
	// ActiveDirectory ...
	ActiveDirectory DirectoryType = "ActiveDirectory"
)

// PossibleDirectoryTypeValues returns an array of possible values for the DirectoryType const type.
func PossibleDirectoryTypeValues() []DirectoryType {
	return []DirectoryType{ActiveDirectory}
}

// FilterMode enumerates the values for filter mode.
type FilterMode string

const (
	// Default ...
	Default FilterMode = "Default"
	// Exclude ...
	Exclude FilterMode = "Exclude"
	// Include ...
	Include FilterMode = "Include"
	// Recommend ...
	Recommend FilterMode = "Recommend"
)

// PossibleFilterModeValues returns an array of possible values for the FilterMode const type.
func PossibleFilterModeValues() []FilterMode {
	return []FilterMode{Default, Exclude, Include, Recommend}
}

// JSONWebKeyEncryptionAlgorithm enumerates the values for json web key encryption algorithm.
type JSONWebKeyEncryptionAlgorithm string

const (
	// RSA15 ...
	RSA15 JSONWebKeyEncryptionAlgorithm = "RSA1_5"
	// RSAOAEP ...
	RSAOAEP JSONWebKeyEncryptionAlgorithm = "RSA-OAEP"
	// RSAOAEP256 ...
	RSAOAEP256 JSONWebKeyEncryptionAlgorithm = "RSA-OAEP-256"
)

// PossibleJSONWebKeyEncryptionAlgorithmValues returns an array of possible values for the JSONWebKeyEncryptionAlgorithm const type.
func PossibleJSONWebKeyEncryptionAlgorithmValues() []JSONWebKeyEncryptionAlgorithm {
	return []JSONWebKeyEncryptionAlgorithm{RSA15, RSAOAEP, RSAOAEP256}
}

// OSType enumerates the values for os type.
type OSType string

const (
	// Linux ...
	Linux OSType = "Linux"
	// Windows ...
	Windows OSType = "Windows"
)

// PossibleOSTypeValues returns an array of possible values for the OSType const type.
func PossibleOSTypeValues() []OSType {
	return []OSType{Linux, Windows}
}

// PrivateLink enumerates the values for private link.
type PrivateLink string

const (
	// Disabled ...
	Disabled PrivateLink = "Disabled"
	// Enabled ...
	Enabled PrivateLink = "Enabled"
)

// PossiblePrivateLinkValues returns an array of possible values for the PrivateLink const type.
func PossiblePrivateLinkValues() []PrivateLink {
	return []PrivateLink{Disabled, Enabled}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// None ...
	None ResourceIdentityType = "None"
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
	// SystemAssignedUserAssigned ...
	SystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	// UserAssigned ...
	UserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{None, SystemAssigned, SystemAssignedUserAssigned, UserAssigned}
}

// ResourceProviderConnection enumerates the values for resource provider connection.
type ResourceProviderConnection string

const (
	// Inbound ...
	Inbound ResourceProviderConnection = "Inbound"
	// Outbound ...
	Outbound ResourceProviderConnection = "Outbound"
)

// PossibleResourceProviderConnectionValues returns an array of possible values for the ResourceProviderConnection const type.
func PossibleResourceProviderConnectionValues() []ResourceProviderConnection {
	return []ResourceProviderConnection{Inbound, Outbound}
}

// Tier enumerates the values for tier.
type Tier string

const (
	// Premium ...
	Premium Tier = "Premium"
	// Standard ...
	Standard Tier = "Standard"
)

// PossibleTierValues returns an array of possible values for the Tier const type.
func PossibleTierValues() []Tier {
	return []Tier{Premium, Standard}
}