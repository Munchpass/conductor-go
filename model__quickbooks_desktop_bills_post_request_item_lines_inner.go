/*
Conductor API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package conductor

import (
	"encoding/json"
)

// checks if the QuickbooksDesktopBillsPostRequestItemLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopBillsPostRequestItemLinesInner{}

// QuickbooksDesktopBillsPostRequestItemLinesInner struct for QuickbooksDesktopBillsPostRequestItemLinesInner
type QuickbooksDesktopBillsPostRequestItemLinesInner struct {
	// The item associated with this item line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item.
	ItemId *string `json:"itemId,omitempty"`
	// The site location where inventory for the item associated with this item line is stored.
	InventorySiteId *string `json:"inventorySiteId,omitempty"`
	// The specific location (e.g., bin or shelf) within the inventory site where the item associated with this item line is stored.
	InventorySiteLocationId *string `json:"inventorySiteLocationId,omitempty"`
	// The serial number of the item associated with this item line. This is used for tracking individual units of serialized inventory items.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The lot number of the item associated with this item line. Used for tracking groups of inventory items that are purchased or manufactured together.
	LotNumber *string `json:"lotNumber,omitempty"`
	// The expiration date for the serial number or lot number of the item associated with this item line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// A description of this item line.
	Description *string `json:"description,omitempty"`
	// The quantity of the item associated with this item line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item.
	Quantity *float32 `json:"quantity,omitempty"`
	// The unit-of-measure used for the `quantity` in this item line. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// The cost of this item line, represented as a decimal string. If both `quantity` and `amount` are specified but not `cost`, QuickBooks will use them to calculate `cost`.
	Cost *string `json:"cost,omitempty"`
	// The monetary amount of this item line, represented as a decimal string. If both `quantity` and `cost` are specified but not `amount`, QuickBooks will use them to calculate `amount`. If `amount`, `cost`, and `quantity` are all unspecified, then QuickBooks will calculate `amount` based on a `quantity` of `1` and the suggested `cost`. This field cannot be cleared.
	Amount *string `json:"amount,omitempty"`
	// The customer or customer-job associated with this item line.
	CustomerId *string `json:"customerId,omitempty"`
	// The item line's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all item lines unless overridden here, at the transaction line level.
	ClassId *string `json:"classId,omitempty"`
	// The sales-tax code for this item line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The billing status of this item line.
	BillingStatus *string `json:"billingStatus,omitempty"`
	// The account to use for this item line, overriding the default account associated with the item.
	OverrideItemAccountId *string `json:"overrideItemAccountId,omitempty"`
	LinkToTransactionLine *QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine `json:"linkToTransactionLine,omitempty"`
	// The item line's sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks.
	SalesRepresentativeId *string `json:"salesRepresentativeId,omitempty"`
	// The custom fields for the item line object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner `json:"customFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopBillsPostRequestItemLinesInner QuickbooksDesktopBillsPostRequestItemLinesInner

// NewQuickbooksDesktopBillsPostRequestItemLinesInner instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopBillsPostRequestItemLinesInner() *QuickbooksDesktopBillsPostRequestItemLinesInner {
	this := QuickbooksDesktopBillsPostRequestItemLinesInner{}
	var billingStatus string = "billable"
	this.BillingStatus = &billingStatus
	return &this
}

// NewQuickbooksDesktopBillsPostRequestItemLinesInnerWithDefaults instantiates a new QuickbooksDesktopBillsPostRequestItemLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopBillsPostRequestItemLinesInnerWithDefaults() *QuickbooksDesktopBillsPostRequestItemLinesInner {
	this := QuickbooksDesktopBillsPostRequestItemLinesInner{}
	var billingStatus string = "billable"
	this.BillingStatus = &billingStatus
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetItemId(v string) {
	o.ItemId = &v
}

// GetInventorySiteId returns the InventorySiteId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteId() string {
	if o == nil || IsNil(o.InventorySiteId) {
		var ret string
		return ret
	}
	return *o.InventorySiteId
}

// GetInventorySiteIdOk returns a tuple with the InventorySiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteId) {
		return nil, false
	}
	return o.InventorySiteId, true
}

// HasInventorySiteId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasInventorySiteId() bool {
	if o != nil && !IsNil(o.InventorySiteId) {
		return true
	}

	return false
}

// SetInventorySiteId gets a reference to the given string and assigns it to the InventorySiteId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetInventorySiteId(v string) {
	o.InventorySiteId = &v
}

// GetInventorySiteLocationId returns the InventorySiteLocationId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteLocationId() string {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		var ret string
		return ret
	}
	return *o.InventorySiteLocationId
}

// GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetInventorySiteLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		return nil, false
	}
	return o.InventorySiteLocationId, true
}

// HasInventorySiteLocationId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasInventorySiteLocationId() bool {
	if o != nil && !IsNil(o.InventorySiteLocationId) {
		return true
	}

	return false
}

// SetInventorySiteLocationId gets a reference to the given string and assigns it to the InventorySiteLocationId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetInventorySiteLocationId(v string) {
	o.InventorySiteLocationId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber) {
		var ret string
		return ret
	}
	return *o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLotNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LotNumber) {
		return nil, false
	}
	return o.LotNumber, true
}

// HasLotNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasLotNumber() bool {
	if o != nil && !IsNil(o.LotNumber) {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given string and assigns it to the LotNumber field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetLotNumber(v string) {
	o.LotNumber = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCost() string {
	if o == nil || IsNil(o.Cost) {
		var ret string
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCostOk() (*string, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given string and assigns it to the Cost field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCost(v string) {
	o.Cost = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetAmount(v string) {
	o.Amount = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomerId() string {
	if o == nil || IsNil(o.CustomerId) {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerId) {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCustomerId() bool {
	if o != nil && !IsNil(o.CustomerId) {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCustomerId(v string) {
	o.CustomerId = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetClassId(v string) {
	o.ClassId = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetBillingStatus returns the BillingStatus field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetBillingStatus() string {
	if o == nil || IsNil(o.BillingStatus) {
		var ret string
		return ret
	}
	return *o.BillingStatus
}

// GetBillingStatusOk returns a tuple with the BillingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetBillingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.BillingStatus) {
		return nil, false
	}
	return o.BillingStatus, true
}

// HasBillingStatus returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasBillingStatus() bool {
	if o != nil && !IsNil(o.BillingStatus) {
		return true
	}

	return false
}

// SetBillingStatus gets a reference to the given string and assigns it to the BillingStatus field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetBillingStatus(v string) {
	o.BillingStatus = &v
}

// GetOverrideItemAccountId returns the OverrideItemAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetOverrideItemAccountId() string {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		var ret string
		return ret
	}
	return *o.OverrideItemAccountId
}

// GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetOverrideItemAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		return nil, false
	}
	return o.OverrideItemAccountId, true
}

// HasOverrideItemAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasOverrideItemAccountId() bool {
	if o != nil && !IsNil(o.OverrideItemAccountId) {
		return true
	}

	return false
}

// SetOverrideItemAccountId gets a reference to the given string and assigns it to the OverrideItemAccountId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetOverrideItemAccountId(v string) {
	o.OverrideItemAccountId = &v
}

// GetLinkToTransactionLine returns the LinkToTransactionLine field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLinkToTransactionLine() QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine {
	if o == nil || IsNil(o.LinkToTransactionLine) {
		var ret QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine
		return ret
	}
	return *o.LinkToTransactionLine
}

// GetLinkToTransactionLineOk returns a tuple with the LinkToTransactionLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetLinkToTransactionLineOk() (*QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine, bool) {
	if o == nil || IsNil(o.LinkToTransactionLine) {
		return nil, false
	}
	return o.LinkToTransactionLine, true
}

// HasLinkToTransactionLine returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasLinkToTransactionLine() bool {
	if o != nil && !IsNil(o.LinkToTransactionLine) {
		return true
	}

	return false
}

// SetLinkToTransactionLine gets a reference to the given QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine and assigns it to the LinkToTransactionLine field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetLinkToTransactionLine(v QuickbooksDesktopBillsPostRequestItemLinesInnerLinkToTransactionLine) {
	o.LinkToTransactionLine = &v
}

// GetSalesRepresentativeId returns the SalesRepresentativeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesRepresentativeId() string {
	if o == nil || IsNil(o.SalesRepresentativeId) {
		var ret string
		return ret
	}
	return *o.SalesRepresentativeId
}

// GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetSalesRepresentativeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesRepresentativeId) {
		return nil, false
	}
	return o.SalesRepresentativeId, true
}

// HasSalesRepresentativeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasSalesRepresentativeId() bool {
	if o != nil && !IsNil(o.SalesRepresentativeId) {
		return true
	}

	return false
}

// SetSalesRepresentativeId gets a reference to the given string and assigns it to the SalesRepresentativeId field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetSalesRepresentativeId(v string) {
	o.SalesRepresentativeId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner {
	if o == nil || IsNil(o.CustomFields) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) GetCustomFieldsOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner and assigns it to the CustomFields field.
func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) {
	o.CustomFields = v
}

func (o QuickbooksDesktopBillsPostRequestItemLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopBillsPostRequestItemLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["itemId"] = o.ItemId
	}
	if !IsNil(o.InventorySiteId) {
		toSerialize["inventorySiteId"] = o.InventorySiteId
	}
	if !IsNil(o.InventorySiteLocationId) {
		toSerialize["inventorySiteLocationId"] = o.InventorySiteLocationId
	}
	if !IsNil(o.SerialNumber) {
		toSerialize["serialNumber"] = o.SerialNumber
	}
	if !IsNil(o.LotNumber) {
		toSerialize["lotNumber"] = o.LotNumber
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.UnitOfMeasure) {
		toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CustomerId) {
		toSerialize["customerId"] = o.CustomerId
	}
	if !IsNil(o.ClassId) {
		toSerialize["classId"] = o.ClassId
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.BillingStatus) {
		toSerialize["billingStatus"] = o.BillingStatus
	}
	if !IsNil(o.OverrideItemAccountId) {
		toSerialize["overrideItemAccountId"] = o.OverrideItemAccountId
	}
	if !IsNil(o.LinkToTransactionLine) {
		toSerialize["linkToTransactionLine"] = o.LinkToTransactionLine
	}
	if !IsNil(o.SalesRepresentativeId) {
		toSerialize["salesRepresentativeId"] = o.SalesRepresentativeId
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopBillsPostRequestItemLinesInner) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopBillsPostRequestItemLinesInner := _QuickbooksDesktopBillsPostRequestItemLinesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopBillsPostRequestItemLinesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopBillsPostRequestItemLinesInner(varQuickbooksDesktopBillsPostRequestItemLinesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemId")
		delete(additionalProperties, "inventorySiteId")
		delete(additionalProperties, "inventorySiteLocationId")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "lotNumber")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "customerId")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "billingStatus")
		delete(additionalProperties, "overrideItemAccountId")
		delete(additionalProperties, "linkToTransactionLine")
		delete(additionalProperties, "salesRepresentativeId")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopBillsPostRequestItemLinesInner struct {
	value *QuickbooksDesktopBillsPostRequestItemLinesInner
	isSet bool
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInner) Get() *QuickbooksDesktopBillsPostRequestItemLinesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInner) Set(val *QuickbooksDesktopBillsPostRequestItemLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopBillsPostRequestItemLinesInner(val *QuickbooksDesktopBillsPostRequestItemLinesInner) *NullableQuickbooksDesktopBillsPostRequestItemLinesInner {
	return &NullableQuickbooksDesktopBillsPostRequestItemLinesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopBillsPostRequestItemLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopBillsPostRequestItemLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


