/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
	"fmt"
)

// checks if the QuickbooksDesktopEstimatesIdPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopEstimatesIdPostRequest{}

// QuickbooksDesktopEstimatesIdPostRequest struct for QuickbooksDesktopEstimatesIdPostRequest
type QuickbooksDesktopEstimatesIdPostRequest struct {
	// The current QuickBooks-assigned revision number of the estimate object you are updating, which you can get by fetching the object first. Provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The customer or customer-job associated with this estimate.
	CustomerId *string `json:"customerId,omitempty"`
	// The estimate's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this estimate's line items unless overridden at the line item level.
	ClassId *string `json:"classId,omitempty"`
	// The predefined template in QuickBooks that determines the layout and formatting for this estimate when printed or displayed.
	DocumentTemplateId *string `json:"documentTemplateId,omitempty"`
	// The date of this estimate, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate *string `json:"transactionDate,omitempty"`
	// The case-sensitive user-defined reference number for this estimate, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.
	RefNumber *string `json:"refNumber,omitempty"`
	BillingAddress *QuickbooksDesktopEstimatesPostRequestBillingAddress `json:"billingAddress,omitempty"`
	ShippingAddress *QuickbooksDesktopEstimatesPostRequestShippingAddress `json:"shippingAddress,omitempty"`
	// Indicates whether this estimate is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive *bool `json:"isActive,omitempty"`
	// When `true`, creates a \"change order\" that appears in this estimate's description field in QuickBooks's estimate form, specifying exactly what changed in this update request, the dollar amount of each change, and the net dollar change to this estimate.
	CreateChangeOrder *bool `json:"createChangeOrder,omitempty"`
	// The customer's Purchase Order (PO) number associated with this estimate. This field is often used to cross-reference the estimate with the customer's purchasing system.
	PurchaseOrderNumber *string `json:"purchaseOrderNumber,omitempty"`
	// The estimate's payment terms, defining when payment is due and any applicable discounts.
	TermsId *string `json:"termsId,omitempty"`
	// The date by which this estimate must be paid, in ISO 8601 format (YYYY-MM-DD).
	DueDate *string `json:"dueDate,omitempty"`
	// The estimate's sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks.
	SalesRepresentativeId *string `json:"salesRepresentativeId,omitempty"`
	// The origin location from where the product associated with this estimate is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \"FOB\" for this field, which stands for \"freight on board\". This field is informational and has no accounting implications.
	ShipmentOrigin *string `json:"shipmentOrigin,omitempty"`
	// The sales-tax item used to calculate the actual tax amount for this estimate's transactions by applying a specific tax rate collected for a single tax agency. Unlike `salesTaxCode`, which only indicates general taxability, this field drives the actual tax calculation and reporting.
	SalesTaxItemId *string `json:"salesTaxItemId,omitempty"`
	// A memo or note for this estimate that appears in reports, but not on the estimate. Use `customerMessage` to add a note to this estimate.
	Memo *string `json:"memo,omitempty"`
	// The message to display to the customer on the estimate.
	CustomerMessageId *string `json:"customerMessageId,omitempty"`
	// Indicates whether this estimate is included in the queue of documents for QuickBooks to email to the customer.
	IsQueuedForEmail *bool `json:"isQueuedForEmail,omitempty"`
	// The sales-tax code for this estimate, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// A built-in custom field for additional information specific to this estimate. Unlike the user-defined fields in the `customFields` array, this is a standard QuickBooks field that exists for all estimates for convenience. Developers often use this field for tracking information that doesn't fit into other standard QuickBooks fields. Unlike `otherCustomField1` and `otherCustomField2`, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI.
	OtherCustomField *string `json:"otherCustomField,omitempty"`
	// The market exchange rate between this estimate's currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR = 1.2345 USD if USD is the home currency).
	ExchangeRate *float32 `json:"exchangeRate,omitempty"`
	// The estimate's line items, each representing a single product or service quoted.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the estimate with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the `id` field set to `-1`.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged.
	Lines []QuickbooksDesktopEstimatesIdPostRequestLinesInner `json:"lines,omitempty"`
	// The estimate's line item groups, each representing a predefined set of related items.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line item groups for the estimate with this array. To keep any existing line item groups, you must include them in this array even if they have not changed. **Any line item groups not included will be removed.**  2. To add a new line item group, include it here with the `id` field set to `-1`.  3. If you do not wish to modify any line item groups, omit this field entirely to keep them unchanged.
	LineGroups []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner `json:"lineGroups,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopEstimatesIdPostRequest QuickbooksDesktopEstimatesIdPostRequest

// NewQuickbooksDesktopEstimatesIdPostRequest instantiates a new QuickbooksDesktopEstimatesIdPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopEstimatesIdPostRequest(revisionNumber string) *QuickbooksDesktopEstimatesIdPostRequest {
	this := QuickbooksDesktopEstimatesIdPostRequest{}
	this.RevisionNumber = revisionNumber
	return &this
}

// NewQuickbooksDesktopEstimatesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopEstimatesIdPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopEstimatesIdPostRequestWithDefaults() *QuickbooksDesktopEstimatesIdPostRequest {
	this := QuickbooksDesktopEstimatesIdPostRequest{}
	return &this
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetClassId(v string) {
	o.ClassId = &v
}

// GetDocumentTemplateId returns the DocumentTemplateId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDocumentTemplateId() string {
	if o == nil || IsNil(o.DocumentTemplateId) {
		var ret string
		return ret
	}
	return *o.DocumentTemplateId
}

// GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDocumentTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentTemplateId) {
		return nil, false
	}
	return o.DocumentTemplateId, true
}

// HasDocumentTemplateId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasDocumentTemplateId() bool {
	if o != nil && !IsNil(o.DocumentTemplateId) {
		return true
	}

	return false
}

// SetDocumentTemplateId gets a reference to the given string and assigns it to the DocumentTemplateId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetDocumentTemplateId(v string) {
	o.DocumentTemplateId = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetRefNumber returns the RefNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRefNumber() string {
	if o == nil || IsNil(o.RefNumber) {
		var ret string
		return ret
	}
	return *o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRefNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RefNumber) {
		return nil, false
	}
	return o.RefNumber, true
}

// HasRefNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasRefNumber() bool {
	if o != nil && !IsNil(o.RefNumber) {
		return true
	}

	return false
}

// SetRefNumber gets a reference to the given string and assigns it to the RefNumber field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetRefNumber(v string) {
	o.RefNumber = &v
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetBillingAddress() QuickbooksDesktopEstimatesPostRequestBillingAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret QuickbooksDesktopEstimatesPostRequestBillingAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetBillingAddressOk() (*QuickbooksDesktopEstimatesPostRequestBillingAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given QuickbooksDesktopEstimatesPostRequestBillingAddress and assigns it to the BillingAddress field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetBillingAddress(v QuickbooksDesktopEstimatesPostRequestBillingAddress) {
	o.BillingAddress = &v
}

// GetShippingAddress returns the ShippingAddress field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShippingAddress() QuickbooksDesktopEstimatesPostRequestShippingAddress {
	if o == nil || IsNil(o.ShippingAddress) {
		var ret QuickbooksDesktopEstimatesPostRequestShippingAddress
		return ret
	}
	return *o.ShippingAddress
}

// GetShippingAddressOk returns a tuple with the ShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopEstimatesPostRequestShippingAddress, bool) {
	if o == nil || IsNil(o.ShippingAddress) {
		return nil, false
	}
	return o.ShippingAddress, true
}

// HasShippingAddress returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasShippingAddress() bool {
	if o != nil && !IsNil(o.ShippingAddress) {
		return true
	}

	return false
}

// SetShippingAddress gets a reference to the given QuickbooksDesktopEstimatesPostRequestShippingAddress and assigns it to the ShippingAddress field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetShippingAddress(v QuickbooksDesktopEstimatesPostRequestShippingAddress) {
	o.ShippingAddress = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetCreateChangeOrder returns the CreateChangeOrder field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCreateChangeOrder() bool {
	if o == nil || IsNil(o.CreateChangeOrder) {
		var ret bool
		return ret
	}
	return *o.CreateChangeOrder
}

// GetCreateChangeOrderOk returns a tuple with the CreateChangeOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCreateChangeOrderOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateChangeOrder) {
		return nil, false
	}
	return o.CreateChangeOrder, true
}

// HasCreateChangeOrder returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCreateChangeOrder() bool {
	if o != nil && !IsNil(o.CreateChangeOrder) {
		return true
	}

	return false
}

// SetCreateChangeOrder gets a reference to the given bool and assigns it to the CreateChangeOrder field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCreateChangeOrder(v bool) {
	o.CreateChangeOrder = &v
}

// GetPurchaseOrderNumber returns the PurchaseOrderNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetPurchaseOrderNumber() string {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderNumber
}

// GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetPurchaseOrderNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderNumber) {
		return nil, false
	}
	return o.PurchaseOrderNumber, true
}

// HasPurchaseOrderNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasPurchaseOrderNumber() bool {
	if o != nil && !IsNil(o.PurchaseOrderNumber) {
		return true
	}

	return false
}

// SetPurchaseOrderNumber gets a reference to the given string and assigns it to the PurchaseOrderNumber field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetPurchaseOrderNumber(v string) {
	o.PurchaseOrderNumber = &v
}

// GetTermsId returns the TermsId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTermsId() string {
	if o == nil || IsNil(o.TermsId) {
		var ret string
		return ret
	}
	return *o.TermsId
}

// GetTermsIdOk returns a tuple with the TermsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTermsIdOk() (*string, bool) {
	if o == nil || IsNil(o.TermsId) {
		return nil, false
	}
	return o.TermsId, true
}

// HasTermsId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasTermsId() bool {
	if o != nil && !IsNil(o.TermsId) {
		return true
	}

	return false
}

// SetTermsId gets a reference to the given string and assigns it to the TermsId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetTermsId(v string) {
	o.TermsId = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetSalesRepresentativeId returns the SalesRepresentativeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesRepresentativeId() string {
	if o == nil || IsNil(o.SalesRepresentativeId) {
		var ret string
		return ret
	}
	return *o.SalesRepresentativeId
}

// GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesRepresentativeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesRepresentativeId) {
		return nil, false
	}
	return o.SalesRepresentativeId, true
}

// HasSalesRepresentativeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesRepresentativeId() bool {
	if o != nil && !IsNil(o.SalesRepresentativeId) {
		return true
	}

	return false
}

// SetSalesRepresentativeId gets a reference to the given string and assigns it to the SalesRepresentativeId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesRepresentativeId(v string) {
	o.SalesRepresentativeId = &v
}

// GetShipmentOrigin returns the ShipmentOrigin field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShipmentOrigin() string {
	if o == nil || IsNil(o.ShipmentOrigin) {
		var ret string
		return ret
	}
	return *o.ShipmentOrigin
}

// GetShipmentOriginOk returns a tuple with the ShipmentOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShipmentOriginOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentOrigin) {
		return nil, false
	}
	return o.ShipmentOrigin, true
}

// HasShipmentOrigin returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasShipmentOrigin() bool {
	if o != nil && !IsNil(o.ShipmentOrigin) {
		return true
	}

	return false
}

// SetShipmentOrigin gets a reference to the given string and assigns it to the ShipmentOrigin field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetShipmentOrigin(v string) {
	o.ShipmentOrigin = &v
}

// GetSalesTaxItemId returns the SalesTaxItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxItemId() string {
	if o == nil || IsNil(o.SalesTaxItemId) {
		var ret string
		return ret
	}
	return *o.SalesTaxItemId
}

// GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxItemId) {
		return nil, false
	}
	return o.SalesTaxItemId, true
}

// HasSalesTaxItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesTaxItemId() bool {
	if o != nil && !IsNil(o.SalesTaxItemId) {
		return true
	}

	return false
}

// SetSalesTaxItemId gets a reference to the given string and assigns it to the SalesTaxItemId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesTaxItemId(v string) {
	o.SalesTaxItemId = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetMemo(v string) {
	o.Memo = &v
}

// GetCustomerMessageId returns the CustomerMessageId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerMessageId() string {
	if o == nil || IsNil(o.CustomerMessageId) {
		var ret string
		return ret
	}
	return *o.CustomerMessageId
}

// GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerMessageId) {
		return nil, false
	}
	return o.CustomerMessageId, true
}

// HasCustomerMessageId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCustomerMessageId() bool {
	if o != nil && !IsNil(o.CustomerMessageId) {
		return true
	}

	return false
}

// SetCustomerMessageId gets a reference to the given string and assigns it to the CustomerMessageId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCustomerMessageId(v string) {
	o.CustomerMessageId = &v
}

// GetIsQueuedForEmail returns the IsQueuedForEmail field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsQueuedForEmail() bool {
	if o == nil || IsNil(o.IsQueuedForEmail) {
		var ret bool
		return ret
	}
	return *o.IsQueuedForEmail
}

// GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsQueuedForEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.IsQueuedForEmail) {
		return nil, false
	}
	return o.IsQueuedForEmail, true
}

// HasIsQueuedForEmail returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasIsQueuedForEmail() bool {
	if o != nil && !IsNil(o.IsQueuedForEmail) {
		return true
	}

	return false
}

// SetIsQueuedForEmail gets a reference to the given bool and assigns it to the IsQueuedForEmail field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetIsQueuedForEmail(v bool) {
	o.IsQueuedForEmail = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetOtherCustomField returns the OtherCustomField field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetOtherCustomField() string {
	if o == nil || IsNil(o.OtherCustomField) {
		var ret string
		return ret
	}
	return *o.OtherCustomField
}

// GetOtherCustomFieldOk returns a tuple with the OtherCustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetOtherCustomFieldOk() (*string, bool) {
	if o == nil || IsNil(o.OtherCustomField) {
		return nil, false
	}
	return o.OtherCustomField, true
}

// HasOtherCustomField returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasOtherCustomField() bool {
	if o != nil && !IsNil(o.OtherCustomField) {
		return true
	}

	return false
}

// SetOtherCustomField gets a reference to the given string and assigns it to the OtherCustomField field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetOtherCustomField(v string) {
	o.OtherCustomField = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetLines returns the Lines field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLines() []QuickbooksDesktopEstimatesIdPostRequestLinesInner {
	if o == nil || IsNil(o.Lines) {
		var ret []QuickbooksDesktopEstimatesIdPostRequestLinesInner
		return ret
	}
	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLinesOk() ([]QuickbooksDesktopEstimatesIdPostRequestLinesInner, bool) {
	if o == nil || IsNil(o.Lines) {
		return nil, false
	}
	return o.Lines, true
}

// HasLines returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasLines() bool {
	if o != nil && !IsNil(o.Lines) {
		return true
	}

	return false
}

// SetLines gets a reference to the given []QuickbooksDesktopEstimatesIdPostRequestLinesInner and assigns it to the Lines field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetLines(v []QuickbooksDesktopEstimatesIdPostRequestLinesInner) {
	o.Lines = v
}

// GetLineGroups returns the LineGroups field value if set, zero value otherwise.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLineGroups() []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner {
	if o == nil || IsNil(o.LineGroups) {
		var ret []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner
		return ret
	}
	return o.LineGroups
}

// GetLineGroupsOk returns a tuple with the LineGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLineGroupsOk() ([]QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner, bool) {
	if o == nil || IsNil(o.LineGroups) {
		return nil, false
	}
	return o.LineGroups, true
}

// HasLineGroups returns a boolean if a field has been set.
func (o *QuickbooksDesktopEstimatesIdPostRequest) HasLineGroups() bool {
	if o != nil && !IsNil(o.LineGroups) {
		return true
	}

	return false
}

// SetLineGroups gets a reference to the given []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner and assigns it to the LineGroups field.
func (o *QuickbooksDesktopEstimatesIdPostRequest) SetLineGroups(v []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner) {
	o.LineGroups = v
}

func (o QuickbooksDesktopEstimatesIdPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopEstimatesIdPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["revisionNumber"] = o.RevisionNumber
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.ClassId) {
		toSerialize["classId"] = o.ClassId
	}
	if !IsNil(o.DocumentTemplateId) {
		toSerialize["documentTemplateId"] = o.DocumentTemplateId
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !IsNil(o.RefNumber) {
		toSerialize["refNumber"] = o.RefNumber
	}
	if !IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !IsNil(o.ShippingAddress) {
		toSerialize["shippingAddress"] = o.ShippingAddress
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.CreateChangeOrder) {
		toSerialize["createChangeOrder"] = o.CreateChangeOrder
	}
	if !IsNil(o.PurchaseOrderNumber) {
		toSerialize["purchaseOrderNumber"] = o.PurchaseOrderNumber
	}
	if !IsNil(o.TermsId) {
		toSerialize["termsId"] = o.TermsId
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.SalesRepresentativeId) {
		toSerialize["salesRepresentativeId"] = o.SalesRepresentativeId
	}
	if !IsNil(o.ShipmentOrigin) {
		toSerialize["shipmentOrigin"] = o.ShipmentOrigin
	}
	if !IsNil(o.SalesTaxItemId) {
		toSerialize["salesTaxItemId"] = o.SalesTaxItemId
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.CustomerMessageId) {
		toSerialize["customerMessageId"] = o.CustomerMessageId
	}
	if !IsNil(o.IsQueuedForEmail) {
		toSerialize["isQueuedForEmail"] = o.IsQueuedForEmail
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.OtherCustomField) {
		toSerialize["otherCustomField"] = o.OtherCustomField
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchangeRate"] = o.ExchangeRate
	}
	if !IsNil(o.Lines) {
		toSerialize["lines"] = o.Lines
	}
	if !IsNil(o.LineGroups) {
		toSerialize["lineGroups"] = o.LineGroups
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopEstimatesIdPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"revisionNumber",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQuickbooksDesktopEstimatesIdPostRequest := _QuickbooksDesktopEstimatesIdPostRequest{}

	err = json.Unmarshal(data, &varQuickbooksDesktopEstimatesIdPostRequest)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopEstimatesIdPostRequest(varQuickbooksDesktopEstimatesIdPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "documentTemplateId")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "billingAddress")
		delete(additionalProperties, "shippingAddress")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "createChangeOrder")
		delete(additionalProperties, "purchaseOrderNumber")
		delete(additionalProperties, "termsId")
		delete(additionalProperties, "dueDate")
		delete(additionalProperties, "salesRepresentativeId")
		delete(additionalProperties, "shipmentOrigin")
		delete(additionalProperties, "salesTaxItemId")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "customerMessageId")
		delete(additionalProperties, "isQueuedForEmail")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "otherCustomField")
		delete(additionalProperties, "exchangeRate")
		delete(additionalProperties, "lines")
		delete(additionalProperties, "lineGroups")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopEstimatesIdPostRequest struct {
	value *QuickbooksDesktopEstimatesIdPostRequest
	isSet bool
}

func (v NullableQuickbooksDesktopEstimatesIdPostRequest) Get() *QuickbooksDesktopEstimatesIdPostRequest {
	return v.value
}

func (v *NullableQuickbooksDesktopEstimatesIdPostRequest) Set(val *QuickbooksDesktopEstimatesIdPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopEstimatesIdPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopEstimatesIdPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopEstimatesIdPostRequest(val *QuickbooksDesktopEstimatesIdPostRequest) *NullableQuickbooksDesktopEstimatesIdPostRequest {
	return &NullableQuickbooksDesktopEstimatesIdPostRequest{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopEstimatesIdPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopEstimatesIdPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


