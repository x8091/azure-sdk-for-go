package commitmentplans

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

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go//services/preview/machinelearning/mgmt/2016-05-01-preview/commitmentplans"

// ResourceSkuRestrictionsReasonCode enumerates the values for resource sku restrictions reason code.
type ResourceSkuRestrictionsReasonCode string

const (
	// NotAvailableForSubscription ...
	NotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	// QuotaID ...
	QuotaID ResourceSkuRestrictionsReasonCode = "QuotaId"
)

// PossibleResourceSkuRestrictionsReasonCodeValues returns an array of possible values for the ResourceSkuRestrictionsReasonCode const type.
func PossibleResourceSkuRestrictionsReasonCodeValues() []ResourceSkuRestrictionsReasonCode {
	return []ResourceSkuRestrictionsReasonCode{NotAvailableForSubscription, QuotaID}
}

// ResourceSkuRestrictionsType enumerates the values for resource sku restrictions type.
type ResourceSkuRestrictionsType string

const (
	// Location ...
	Location ResourceSkuRestrictionsType = "location"
	// Zone ...
	Zone ResourceSkuRestrictionsType = "zone"
)

// PossibleResourceSkuRestrictionsTypeValues returns an array of possible values for the ResourceSkuRestrictionsType const type.
func PossibleResourceSkuRestrictionsTypeValues() []ResourceSkuRestrictionsType {
	return []ResourceSkuRestrictionsType{Location, Zone}
}

// SkuCapacityScaleType enumerates the values for sku capacity scale type.
type SkuCapacityScaleType string

const (
	// Automatic ...
	Automatic SkuCapacityScaleType = "Automatic"
	// Manual ...
	Manual SkuCapacityScaleType = "Manual"
	// None ...
	None SkuCapacityScaleType = "None"
)

// PossibleSkuCapacityScaleTypeValues returns an array of possible values for the SkuCapacityScaleType const type.
func PossibleSkuCapacityScaleTypeValues() []SkuCapacityScaleType {
	return []SkuCapacityScaleType{Automatic, Manual, None}
}

// CatalogSku details of a commitment plan SKU.
type CatalogSku struct {
	// ResourceType - Resource type name
	ResourceType *string `json:"resourceType,omitempty"`
	// Name - SKU name
	Name *string `json:"name,omitempty"`
	// Tier - SKU tier
	Tier *string `json:"tier,omitempty"`
	// Locations - Regions where the SKU is available.
	Locations *[]string `json:"locations,omitempty"`
	// Capacity - SKU scaling information
	Capacity *SkuCapacity `json:"capacity,omitempty"`
	// Capabilities - The capability information for the specified SKU.
	Capabilities *[]SkuCapability `json:"capabilities,omitempty"`
	// Costs - The cost information for the specified SKU.
	Costs *[]SkuCost `json:"costs,omitempty"`
	// Restrictions - Restrictions which would prevent a SKU from being used. This is empty if there are no restrictions.
	Restrictions *[]SkuRestrictions `json:"restrictions,omitempty"`
}

// CommitmentAssociation represents the association between a commitment plan and some other resource, such
// as a Machine Learning web service.
type CommitmentAssociation struct {
	autorest.Response `json:"-"`
	// Etag - An entity tag used to enforce optimistic concurrency.
	Etag *string `json:"etag,omitempty"`
	// Properties - The properties of the commitment association resource.
	Properties *CommitmentAssociationProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - User-defined tags for the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for CommitmentAssociation.
func (ca CommitmentAssociation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ca.Etag != nil {
		objectMap["etag"] = ca.Etag
	}
	if ca.Properties != nil {
		objectMap["properties"] = ca.Properties
	}
	if ca.ID != nil {
		objectMap["id"] = ca.ID
	}
	if ca.Name != nil {
		objectMap["name"] = ca.Name
	}
	if ca.Location != nil {
		objectMap["location"] = ca.Location
	}
	if ca.Type != nil {
		objectMap["type"] = ca.Type
	}
	if ca.Tags != nil {
		objectMap["tags"] = ca.Tags
	}
	return json.Marshal(objectMap)
}

// CommitmentAssociationListResult a page of commitment association resources.
type CommitmentAssociationListResult struct {
	autorest.Response `json:"-"`
	// NextLink - A URI to retrieve the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The set of results for this page.
	Value *[]CommitmentAssociation `json:"value,omitempty"`
}

// CommitmentAssociationListResultIterator provides access to a complete listing of CommitmentAssociation
// values.
type CommitmentAssociationListResultIterator struct {
	i    int
	page CommitmentAssociationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *CommitmentAssociationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentAssociationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *CommitmentAssociationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter CommitmentAssociationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter CommitmentAssociationListResultIterator) Response() CommitmentAssociationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter CommitmentAssociationListResultIterator) Value() CommitmentAssociation {
	if !iter.page.NotDone() {
		return CommitmentAssociation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (calr CommitmentAssociationListResult) IsEmpty() bool {
	return calr.Value == nil || len(*calr.Value) == 0
}

// commitmentAssociationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (calr CommitmentAssociationListResult) commitmentAssociationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if calr.NextLink == nil || len(to.String(calr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(calr.NextLink)))
}

// CommitmentAssociationListResultPage contains a page of CommitmentAssociation values.
type CommitmentAssociationListResultPage struct {
	fn   func(context.Context, CommitmentAssociationListResult) (CommitmentAssociationListResult, error)
	calr CommitmentAssociationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *CommitmentAssociationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CommitmentAssociationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.calr)
	if err != nil {
		return err
	}
	page.calr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *CommitmentAssociationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page CommitmentAssociationListResultPage) NotDone() bool {
	return !page.calr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page CommitmentAssociationListResultPage) Response() CommitmentAssociationListResult {
	return page.calr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page CommitmentAssociationListResultPage) Values() []CommitmentAssociation {
	if page.calr.IsEmpty() {
		return nil
	}
	return *page.calr.Value
}

// CommitmentAssociationProperties properties of an Azure ML commitment association.
type CommitmentAssociationProperties struct {
	// AssociatedResourceID - The ID of the resource this association points to, such as the ARM ID of an Azure ML web service.
	AssociatedResourceID *string `json:"associatedResourceId,omitempty"`
	// CommitmentPlanID - The ARM ID of the parent Azure ML commitment plan.
	CommitmentPlanID *string `json:"commitmentPlanId,omitempty"`
	// CreationDate - The date at which this commitment association was created, in ISO 8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
}

// CommitmentPlan an Azure ML commitment plan resource.
type CommitmentPlan struct {
	autorest.Response `json:"-"`
	// Etag - An entity tag used to enforce optimistic concurrency.
	Etag *string `json:"etag,omitempty"`
	// Properties - The commitment plan properties.
	Properties *Properties `json:"properties,omitempty"`
	// Sku - The commitment plan SKU.
	Sku *ResourceSku `json:"sku,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - User-defined tags for the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for CommitmentPlan.
func (cp CommitmentPlan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.Etag != nil {
		objectMap["etag"] = cp.Etag
	}
	if cp.Properties != nil {
		objectMap["properties"] = cp.Properties
	}
	if cp.Sku != nil {
		objectMap["sku"] = cp.Sku
	}
	if cp.ID != nil {
		objectMap["id"] = cp.ID
	}
	if cp.Name != nil {
		objectMap["name"] = cp.Name
	}
	if cp.Location != nil {
		objectMap["location"] = cp.Location
	}
	if cp.Type != nil {
		objectMap["type"] = cp.Type
	}
	if cp.Tags != nil {
		objectMap["tags"] = cp.Tags
	}
	return json.Marshal(objectMap)
}

// ListResult a page of commitment plan resources.
type ListResult struct {
	autorest.Response `json:"-"`
	// NextLink - A URI to retrieve the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The set of results for this page.
	Value *[]CommitmentPlan `json:"value,omitempty"`
}

// ListResultIterator provides access to a complete listing of CommitmentPlan values.
type ListResultIterator struct {
	i    int
	page ListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *ListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *ListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter ListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter ListResultIterator) Response() ListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter ListResultIterator) Value() CommitmentPlan {
	if !iter.page.NotDone() {
		return CommitmentPlan{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (lr ListResult) IsEmpty() bool {
	return lr.Value == nil || len(*lr.Value) == 0
}

// listResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (lr ListResult) listResultPreparer(ctx context.Context) (*http.Request, error) {
	if lr.NextLink == nil || len(to.String(lr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(lr.NextLink)))
}

// ListResultPage contains a page of CommitmentPlan values.
type ListResultPage struct {
	fn func(context.Context, ListResult) (ListResult, error)
	lr ListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *ListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.lr)
	if err != nil {
		return err
	}
	page.lr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *ListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page ListResultPage) NotDone() bool {
	return !page.lr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page ListResultPage) Response() ListResult {
	return page.lr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page ListResultPage) Values() []CommitmentPlan {
	if page.lr.IsEmpty() {
		return nil
	}
	return *page.lr.Value
}

// MoveCommitmentAssociationRequest specifies the destination Azure ML commitment plan for a move
// operation.
type MoveCommitmentAssociationRequest struct {
	// DestinationPlanID - The ARM ID of the commitment plan to re-parent the commitment association to.
	DestinationPlanID *string `json:"destinationPlanId,omitempty"`
}

// PatchPayload the properties of a commitment plan which may be updated via PATCH.
type PatchPayload struct {
	// Tags - User-defined tags for the commitment plan.
	Tags map[string]*string `json:"tags"`
	// Sku - The commitment plan SKU.
	Sku *ResourceSku `json:"sku,omitempty"`
}

// MarshalJSON is the custom marshaler for PatchPayload.
func (pp PatchPayload) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pp.Tags != nil {
		objectMap["tags"] = pp.Tags
	}
	if pp.Sku != nil {
		objectMap["sku"] = pp.Sku
	}
	return json.Marshal(objectMap)
}

// PlanQuantity represents the quantity a commitment plan provides of a metered resource.
type PlanQuantity struct {
	// Allowance - The quantity added to the commitment plan at an interval specified by its allowance frequency.
	Allowance *float64 `json:"allowance,omitempty"`
	// Amount - The quantity available to the plan the last time usage was calculated.
	Amount *float64 `json:"amount,omitempty"`
	// IncludedQuantityMeter - The Azure meter for usage against included quantities.
	IncludedQuantityMeter *string `json:"includedQuantityMeter,omitempty"`
	// OverageMeter - The Azure meter for usage which exceeds included quantities.
	OverageMeter *string `json:"overageMeter,omitempty"`
}

// PlanUsageHistory represents historical information about usage of the Azure resources associated with a
// commitment plan.
type PlanUsageHistory struct {
	// PlanDeletionOverage - Overage incurred as a result of deleting a commitment plan.
	PlanDeletionOverage map[string]*float64 `json:"planDeletionOverage"`
	// PlanMigrationOverage - Overage incurred as a result of migrating a commitment plan from one SKU to another.
	PlanMigrationOverage map[string]*float64 `json:"planMigrationOverage"`
	// PlanQuantitiesAfterUsage - Included quantities remaining after usage against the commitment plan's associated resources was calculated.
	PlanQuantitiesAfterUsage map[string]*float64 `json:"planQuantitiesAfterUsage"`
	// PlanQuantitiesBeforeUsage - Included quantities remaining before usage against the commitment plan's associated resources was calculated.
	PlanQuantitiesBeforeUsage map[string]*float64 `json:"planQuantitiesBeforeUsage"`
	// PlanUsageOverage - Usage against the commitment plan's associated resources which was not covered by included quantities and is therefore overage.
	PlanUsageOverage map[string]*float64 `json:"planUsageOverage"`
	// Usage - Usage against the commitment plan's associated resources.
	Usage map[string]*float64 `json:"usage"`
	// UsageDate - The date of usage, in ISO 8601 format.
	UsageDate *date.Time `json:"usageDate,omitempty"`
}

// MarshalJSON is the custom marshaler for PlanUsageHistory.
func (puh PlanUsageHistory) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if puh.PlanDeletionOverage != nil {
		objectMap["planDeletionOverage"] = puh.PlanDeletionOverage
	}
	if puh.PlanMigrationOverage != nil {
		objectMap["planMigrationOverage"] = puh.PlanMigrationOverage
	}
	if puh.PlanQuantitiesAfterUsage != nil {
		objectMap["planQuantitiesAfterUsage"] = puh.PlanQuantitiesAfterUsage
	}
	if puh.PlanQuantitiesBeforeUsage != nil {
		objectMap["planQuantitiesBeforeUsage"] = puh.PlanQuantitiesBeforeUsage
	}
	if puh.PlanUsageOverage != nil {
		objectMap["planUsageOverage"] = puh.PlanUsageOverage
	}
	if puh.Usage != nil {
		objectMap["usage"] = puh.Usage
	}
	if puh.UsageDate != nil {
		objectMap["usageDate"] = puh.UsageDate
	}
	return json.Marshal(objectMap)
}

// PlanUsageHistoryListResult a page of usage history.
type PlanUsageHistoryListResult struct {
	autorest.Response `json:"-"`
	// NextLink - A URI to retrieve the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
	// Value - The set of results for this page.
	Value *[]PlanUsageHistory `json:"value,omitempty"`
}

// PlanUsageHistoryListResultIterator provides access to a complete listing of PlanUsageHistory values.
type PlanUsageHistoryListResultIterator struct {
	i    int
	page PlanUsageHistoryListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PlanUsageHistoryListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PlanUsageHistoryListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *PlanUsageHistoryListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PlanUsageHistoryListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PlanUsageHistoryListResultIterator) Response() PlanUsageHistoryListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PlanUsageHistoryListResultIterator) Value() PlanUsageHistory {
	if !iter.page.NotDone() {
		return PlanUsageHistory{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (puhlr PlanUsageHistoryListResult) IsEmpty() bool {
	return puhlr.Value == nil || len(*puhlr.Value) == 0
}

// planUsageHistoryListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (puhlr PlanUsageHistoryListResult) planUsageHistoryListResultPreparer(ctx context.Context) (*http.Request, error) {
	if puhlr.NextLink == nil || len(to.String(puhlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(puhlr.NextLink)))
}

// PlanUsageHistoryListResultPage contains a page of PlanUsageHistory values.
type PlanUsageHistoryListResultPage struct {
	fn    func(context.Context, PlanUsageHistoryListResult) (PlanUsageHistoryListResult, error)
	puhlr PlanUsageHistoryListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PlanUsageHistoryListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PlanUsageHistoryListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.puhlr)
	if err != nil {
		return err
	}
	page.puhlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PlanUsageHistoryListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PlanUsageHistoryListResultPage) NotDone() bool {
	return !page.puhlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PlanUsageHistoryListResultPage) Response() PlanUsageHistoryListResult {
	return page.puhlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PlanUsageHistoryListResultPage) Values() []PlanUsageHistory {
	if page.puhlr.IsEmpty() {
		return nil
	}
	return *page.puhlr.Value
}

// Properties properties of an Azure ML commitment plan.
type Properties struct {
	// ChargeForOverage - Indicates whether usage beyond the commitment plan's included quantities will be charged.
	ChargeForOverage *bool `json:"chargeForOverage,omitempty"`
	// ChargeForPlan - Indicates whether the commitment plan will incur a charge.
	ChargeForPlan *bool `json:"chargeForPlan,omitempty"`
	// CreationDate - The date at which this commitment plan was created, in ISO 8601 format.
	CreationDate *date.Time `json:"creationDate,omitempty"`
	// IncludedQuantities - The included resource quantities this plan gives you.
	IncludedQuantities map[string]*PlanQuantity `json:"includedQuantities"`
	// MaxAssociationLimit - The maximum number of commitment associations that can be children of this commitment plan.
	MaxAssociationLimit *int32 `json:"maxAssociationLimit,omitempty"`
	// MaxCapacityLimit - The maximum scale-out capacity for this commitment plan.
	MaxCapacityLimit *int32 `json:"maxCapacityLimit,omitempty"`
	// MinCapacityLimit - The minimum scale-out capacity for this commitment plan.
	MinCapacityLimit *int32 `json:"minCapacityLimit,omitempty"`
	// PlanMeter - The Azure meter which will be used to charge for this commitment plan.
	PlanMeter *string `json:"planMeter,omitempty"`
	// RefillFrequencyInDays - The frequency at which this commitment plan's included quantities are refilled.
	RefillFrequencyInDays *int32 `json:"refillFrequencyInDays,omitempty"`
	// SuspendPlanOnOverage - Indicates whether this commitment plan will be moved into a suspended state if usage goes beyond the commitment plan's included quantities.
	SuspendPlanOnOverage *bool `json:"suspendPlanOnOverage,omitempty"`
}

// MarshalJSON is the custom marshaler for Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.ChargeForOverage != nil {
		objectMap["chargeForOverage"] = p.ChargeForOverage
	}
	if p.ChargeForPlan != nil {
		objectMap["chargeForPlan"] = p.ChargeForPlan
	}
	if p.CreationDate != nil {
		objectMap["creationDate"] = p.CreationDate
	}
	if p.IncludedQuantities != nil {
		objectMap["includedQuantities"] = p.IncludedQuantities
	}
	if p.MaxAssociationLimit != nil {
		objectMap["maxAssociationLimit"] = p.MaxAssociationLimit
	}
	if p.MaxCapacityLimit != nil {
		objectMap["maxCapacityLimit"] = p.MaxCapacityLimit
	}
	if p.MinCapacityLimit != nil {
		objectMap["minCapacityLimit"] = p.MinCapacityLimit
	}
	if p.PlanMeter != nil {
		objectMap["planMeter"] = p.PlanMeter
	}
	if p.RefillFrequencyInDays != nil {
		objectMap["refillFrequencyInDays"] = p.RefillFrequencyInDays
	}
	if p.SuspendPlanOnOverage != nil {
		objectMap["suspendPlanOnOverage"] = p.SuspendPlanOnOverage
	}
	return json.Marshal(objectMap)
}

// Resource common properties of an ARM resource.
type Resource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Location - Resource location.
	Location *string `json:"location,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
	// Tags - User-defined tags for the resource.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceSku the SKU of a resource.
type ResourceSku struct {
	// Capacity - The scale-out capacity of the resource. 1 is 1x, 2 is 2x, etc. This impacts the quantities and cost of any commitment plan resource.
	Capacity *int32 `json:"capacity,omitempty"`
	// Name - The SKU name. Along with tier, uniquely identifies the SKU.
	Name *string `json:"name,omitempty"`
	// Tier - The SKU tier. Along with name, uniquely identifies the SKU.
	Tier *string `json:"tier,omitempty"`
}

// SkuCapability describes The SKU capabilites object.
type SkuCapability struct {
	// Name - The capability name.
	Name *string `json:"name,omitempty"`
	// Value - The capability value.
	Value *string `json:"value,omitempty"`
}

// SkuCapacity describes scaling information of a SKU.
type SkuCapacity struct {
	// Minimum - The minimum capacity.
	Minimum *int64 `json:"minimum,omitempty"`
	// Maximum - The maximum capacity that can be set.
	Maximum *int64 `json:"maximum,omitempty"`
	// Default - The default capacity.
	Default *int64 `json:"default,omitempty"`
	// ScaleType - The scale type applicable to the sku. Possible values include: 'Automatic', 'Manual', 'None'
	ScaleType SkuCapacityScaleType `json:"scaleType,omitempty"`
}

// SkuCost describes metadata for SKU cost info.
type SkuCost struct {
	// MeterID - The meter used for this part of a SKU's cost.
	MeterID *string `json:"meterID,omitempty"`
	// Quantity - The multiplier for the meter ID.
	Quantity *int64 `json:"quantity,omitempty"`
	// ExtendedUnit - The overall duration represented by the quantity.
	ExtendedUnit *string `json:"extendedUnit,omitempty"`
}

// SkuListResult the list of commitment plan SKUs.
type SkuListResult struct {
	autorest.Response `json:"-"`
	Value             *[]CatalogSku `json:"value,omitempty"`
}

// SkuRestrictions describes restrictions which would prevent a SKU from being used.
type SkuRestrictions struct {
	// Type - The type of restrictions. Possible values include: 'Location', 'Zone'
	Type ResourceSkuRestrictionsType `json:"type,omitempty"`
	// Values - The value of restrictions. If the restriction type is set to location. This would be different locations where the SKU is restricted.
	Values *[]string `json:"values,omitempty"`
	// ReasonCode - The reason for restriction. Possible values include: 'QuotaID', 'NotAvailableForSubscription'
	ReasonCode ResourceSkuRestrictionsReasonCode `json:"reasonCode,omitempty"`
}
