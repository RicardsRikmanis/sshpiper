// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package storage

import original "github.com/Azure/azure-sdk-for-go/services/preview/storage/mgmt/2018-03-01-preview/storage"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessTier = original.AccessTier

const (
	Cool AccessTier = original.Cool
	Hot  AccessTier = original.Hot
)

type AccountStatus = original.AccountStatus

const (
	Available   AccountStatus = original.Available
	Unavailable AccountStatus = original.Unavailable
)

type Action = original.Action

const (
	Allow Action = original.Allow
)

type Bypass = original.Bypass

const (
	AzureServices Bypass = original.AzureServices
	Logging       Bypass = original.Logging
	Metrics       Bypass = original.Metrics
	None          Bypass = original.None
)

type DefaultAction = original.DefaultAction

const (
	DefaultActionAllow DefaultAction = original.DefaultActionAllow
	DefaultActionDeny  DefaultAction = original.DefaultActionDeny
)

type HTTPProtocol = original.HTTPProtocol

const (
	HTTPS     HTTPProtocol = original.HTTPS
	Httpshttp HTTPProtocol = original.Httpshttp
)

type ImmutabilityPolicyState = original.ImmutabilityPolicyState

const (
	Locked   ImmutabilityPolicyState = original.Locked
	Unlocked ImmutabilityPolicyState = original.Unlocked
)

type ImmutabilityPolicyUpdateType = original.ImmutabilityPolicyUpdateType

const (
	Extend ImmutabilityPolicyUpdateType = original.Extend
	Lock   ImmutabilityPolicyUpdateType = original.Lock
	Put    ImmutabilityPolicyUpdateType = original.Put
)

type KeyPermission = original.KeyPermission

const (
	Full KeyPermission = original.Full
	Read KeyPermission = original.Read
)

type KeySource = original.KeySource

const (
	MicrosoftKeyvault KeySource = original.MicrosoftKeyvault
	MicrosoftStorage  KeySource = original.MicrosoftStorage
)

type Kind = original.Kind

const (
	BlobStorage Kind = original.BlobStorage
	Storage     Kind = original.Storage
	StorageV2   Kind = original.StorageV2
)

type LeaseDuration = original.LeaseDuration

const (
	Fixed    LeaseDuration = original.Fixed
	Infinite LeaseDuration = original.Infinite
)

type LeaseState = original.LeaseState

const (
	LeaseStateAvailable LeaseState = original.LeaseStateAvailable
	LeaseStateBreaking  LeaseState = original.LeaseStateBreaking
	LeaseStateBroken    LeaseState = original.LeaseStateBroken
	LeaseStateExpired   LeaseState = original.LeaseStateExpired
	LeaseStateLeased    LeaseState = original.LeaseStateLeased
)

type LeaseStatus = original.LeaseStatus

const (
	LeaseStatusLocked   LeaseStatus = original.LeaseStatusLocked
	LeaseStatusUnlocked LeaseStatus = original.LeaseStatusUnlocked
)

type Permissions = original.Permissions

const (
	A Permissions = original.A
	C Permissions = original.C
	D Permissions = original.D
	L Permissions = original.L
	P Permissions = original.P
	R Permissions = original.R
	U Permissions = original.U
	W Permissions = original.W
)

type ProvisioningState = original.ProvisioningState

const (
	Creating     ProvisioningState = original.Creating
	ResolvingDNS ProvisioningState = original.ResolvingDNS
	Succeeded    ProvisioningState = original.Succeeded
)

type PublicAccess = original.PublicAccess

const (
	PublicAccessBlob      PublicAccess = original.PublicAccessBlob
	PublicAccessContainer PublicAccess = original.PublicAccessContainer
	PublicAccessNone      PublicAccess = original.PublicAccessNone
)

type Reason = original.Reason

const (
	AccountNameInvalid Reason = original.AccountNameInvalid
	AlreadyExists      Reason = original.AlreadyExists
)

type ReasonCode = original.ReasonCode

const (
	NotAvailableForSubscription ReasonCode = original.NotAvailableForSubscription
	QuotaID                     ReasonCode = original.QuotaID
)

type Services = original.Services

const (
	B Services = original.B
	F Services = original.F
	Q Services = original.Q
	T Services = original.T
)

type SignedResource = original.SignedResource

const (
	SignedResourceB SignedResource = original.SignedResourceB
	SignedResourceC SignedResource = original.SignedResourceC
	SignedResourceF SignedResource = original.SignedResourceF
	SignedResourceS SignedResource = original.SignedResourceS
)

type SignedResourceTypes = original.SignedResourceTypes

const (
	SignedResourceTypesC SignedResourceTypes = original.SignedResourceTypesC
	SignedResourceTypesO SignedResourceTypes = original.SignedResourceTypesO
	SignedResourceTypesS SignedResourceTypes = original.SignedResourceTypesS
)

type SkuName = original.SkuName

const (
	PremiumLRS    SkuName = original.PremiumLRS
	StandardGRS   SkuName = original.StandardGRS
	StandardLRS   SkuName = original.StandardLRS
	StandardRAGRS SkuName = original.StandardRAGRS
	StandardZRS   SkuName = original.StandardZRS
)

type SkuTier = original.SkuTier

const (
	Premium  SkuTier = original.Premium
	Standard SkuTier = original.Standard
)

type State = original.State

const (
	StateDeprovisioning       State = original.StateDeprovisioning
	StateFailed               State = original.StateFailed
	StateNetworkSourceDeleted State = original.StateNetworkSourceDeleted
	StateProvisioning         State = original.StateProvisioning
	StateSucceeded            State = original.StateSucceeded
)

type UsageUnit = original.UsageUnit

const (
	Bytes           UsageUnit = original.Bytes
	BytesPerSecond  UsageUnit = original.BytesPerSecond
	Count           UsageUnit = original.Count
	CountsPerSecond UsageUnit = original.CountsPerSecond
	Percent         UsageUnit = original.Percent
	Seconds         UsageUnit = original.Seconds
)

type Account = original.Account
type AccountCheckNameAvailabilityParameters = original.AccountCheckNameAvailabilityParameters
type AccountCreateParameters = original.AccountCreateParameters
type AccountKey = original.AccountKey
type AccountListKeysResult = original.AccountListKeysResult
type AccountListResult = original.AccountListResult
type AccountManagementPolicies = original.AccountManagementPolicies
type AccountManagementPoliciesRulesProperty = original.AccountManagementPoliciesRulesProperty
type AccountProperties = original.AccountProperties
type AccountPropertiesCreateParameters = original.AccountPropertiesCreateParameters
type AccountPropertiesUpdateParameters = original.AccountPropertiesUpdateParameters
type AccountRegenerateKeyParameters = original.AccountRegenerateKeyParameters
type AccountSasParameters = original.AccountSasParameters
type AccountUpdateParameters = original.AccountUpdateParameters
type AccountsClient = original.AccountsClient
type AccountsCreateFuture = original.AccountsCreateFuture
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type BlobContainer = original.BlobContainer
type BlobContainersClient = original.BlobContainersClient
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ContainerProperties = original.ContainerProperties
type CustomDomain = original.CustomDomain
type Dimension = original.Dimension
type Encryption = original.Encryption
type EncryptionService = original.EncryptionService
type EncryptionServices = original.EncryptionServices
type Endpoints = original.Endpoints
type IPRule = original.IPRule
type Identity = original.Identity
type ImmutabilityPolicy = original.ImmutabilityPolicy
type ImmutabilityPolicyProperties = original.ImmutabilityPolicyProperties
type ImmutabilityPolicyProperty = original.ImmutabilityPolicyProperty
type KeyVaultProperties = original.KeyVaultProperties
type LegalHold = original.LegalHold
type LegalHoldProperties = original.LegalHoldProperties
type ListAccountSasResponse = original.ListAccountSasResponse
type ListContainerItem = original.ListContainerItem
type ListContainerItems = original.ListContainerItems
type ListServiceSasResponse = original.ListServiceSasResponse
type ManagementPoliciesRules = original.ManagementPoliciesRules
type ManagementPoliciesRulesSetParameter = original.ManagementPoliciesRulesSetParameter
type MetricSpecification = original.MetricSpecification
type NetworkRuleSet = original.NetworkRuleSet
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type Restriction = original.Restriction
type SKUCapability = original.SKUCapability
type ServiceSasParameters = original.ServiceSasParameters
type ServiceSpecification = original.ServiceSpecification
type Sku = original.Sku
type SkuListResult = original.SkuListResult
type SkusClient = original.SkusClient
type TagProperty = original.TagProperty
type TrackedResource = original.TrackedResource
type UpdateHistoryProperty = original.UpdateHistoryProperty
type Usage = original.Usage
type UsageListResult = original.UsageListResult
type UsageName = original.UsageName
type UsagesClient = original.UsagesClient
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAccountsClient(subscriptionID string) AccountsClient {
	return original.NewAccountsClient(subscriptionID)
}
func NewAccountsClientWithBaseURI(baseURI string, subscriptionID string) AccountsClient {
	return original.NewAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBlobContainersClient(subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClient(subscriptionID)
}
func NewBlobContainersClientWithBaseURI(baseURI string, subscriptionID string) BlobContainersClient {
	return original.NewBlobContainersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSkusClient(subscriptionID string) SkusClient {
	return original.NewSkusClient(subscriptionID)
}
func NewSkusClientWithBaseURI(baseURI string, subscriptionID string) SkusClient {
	return original.NewSkusClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessTierValues() []AccessTier {
	return original.PossibleAccessTierValues()
}
func PossibleAccountStatusValues() []AccountStatus {
	return original.PossibleAccountStatusValues()
}
func PossibleActionValues() []Action {
	return original.PossibleActionValues()
}
func PossibleBypassValues() []Bypass {
	return original.PossibleBypassValues()
}
func PossibleDefaultActionValues() []DefaultAction {
	return original.PossibleDefaultActionValues()
}
func PossibleHTTPProtocolValues() []HTTPProtocol {
	return original.PossibleHTTPProtocolValues()
}
func PossibleImmutabilityPolicyStateValues() []ImmutabilityPolicyState {
	return original.PossibleImmutabilityPolicyStateValues()
}
func PossibleImmutabilityPolicyUpdateTypeValues() []ImmutabilityPolicyUpdateType {
	return original.PossibleImmutabilityPolicyUpdateTypeValues()
}
func PossibleKeyPermissionValues() []KeyPermission {
	return original.PossibleKeyPermissionValues()
}
func PossibleKeySourceValues() []KeySource {
	return original.PossibleKeySourceValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLeaseDurationValues() []LeaseDuration {
	return original.PossibleLeaseDurationValues()
}
func PossibleLeaseStateValues() []LeaseState {
	return original.PossibleLeaseStateValues()
}
func PossibleLeaseStatusValues() []LeaseStatus {
	return original.PossibleLeaseStatusValues()
}
func PossiblePermissionsValues() []Permissions {
	return original.PossiblePermissionsValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicAccessValues() []PublicAccess {
	return original.PossiblePublicAccessValues()
}
func PossibleReasonCodeValues() []ReasonCode {
	return original.PossibleReasonCodeValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleServicesValues() []Services {
	return original.PossibleServicesValues()
}
func PossibleSignedResourceTypesValues() []SignedResourceTypes {
	return original.PossibleSignedResourceTypesValues()
}
func PossibleSignedResourceValues() []SignedResource {
	return original.PossibleSignedResourceValues()
}
func PossibleSkuNameValues() []SkuName {
	return original.PossibleSkuNameValues()
}
func PossibleSkuTierValues() []SkuTier {
	return original.PossibleSkuTierValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func PossibleUsageUnitValues() []UsageUnit {
	return original.PossibleUsageUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}