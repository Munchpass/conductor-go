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

// checks if the QuickbooksDesktopInvoicesPostRequestLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopInvoicesPostRequestLinesInner{}

// QuickbooksDesktopInvoicesPostRequestLinesInner struct for QuickbooksDesktopInvoicesPostRequestLinesInner
type QuickbooksDesktopInvoicesPostRequestLinesInner struct {
	// The item associated with this invoice line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item.
	ItemId *string `json:"itemId,omitempty"`
	// A description of this invoice line.
	Description *string `json:"description,omitempty"`
	// The quantity of the item associated with this invoice line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item.
	Quantity *float32 `json:"quantity,omitempty"`
	// The unit-of-measure used for the `quantity` in this invoice line. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// The price per unit for this invoice line. If both `rate` and `amount` are specified, `rate` will be ignored. If both `quantity` and `amount` are specified but not `rate`, QuickBooks will use them to calculate `rate`. Represented as a decimal string. This field cannot be cleared.
	Rate *string `json:"rate,omitempty"`
	// The price of this invoice line expressed as a percentage. Typically used for discount or markup items.
	RatePercent *string `json:"ratePercent,omitempty"`
	// The price level applied to this invoice line. This overrides any price level set on the corresponding customer. The resulting invoice line will not show this price level, only the final `rate` calculated from it.
	PriceLevelId *string `json:"priceLevelId,omitempty"`
	// The invoice line's class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all invoice lines unless overridden here, at the transaction line level.
	ClassId *string `json:"classId,omitempty"`
	// The monetary amount of this invoice line, represented as a decimal string. If both `quantity` and `rate` are specified but not `amount`, QuickBooks will use them to calculate `amount`. If `amount`, `rate`, and `quantity` are all unspecified, then QuickBooks will calculate `amount` based on a `quantity` of `1` and the suggested `rate`. This field cannot be cleared.
	Amount *string `json:"amount,omitempty"`
	// Specifies how to resolve price rule conflicts when adding or modifying this invoice line.
	PriceRuleConflictStrategy *string `json:"priceRuleConflictStrategy,omitempty"`
	// The site location where inventory for the item associated with this invoice line is stored.
	InventorySiteId *string `json:"inventorySiteId,omitempty"`
	// The specific location (e.g., bin or shelf) within the inventory site where the item associated with this invoice line is stored.
	InventorySiteLocationId *string `json:"inventorySiteLocationId,omitempty"`
	// The serial number of the item associated with this invoice line. This is used for tracking individual units of serialized inventory items.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// The lot number of the item associated with this invoice line. Used for tracking groups of inventory items that are purchased or manufactured together.
	LotNumber *string `json:"lotNumber,omitempty"`
	// The date on which the service for this invoice line was or will be performed, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for service items.
	ServiceDate *string `json:"serviceDate,omitempty"`
	// The sales-tax code for this invoice line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \"Non\" (non-taxable) and \"Tax\" (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \"Do You Charge Sales Tax?\" preference), it will assign the default non-taxable code to all sales.
	SalesTaxCodeId *string `json:"salesTaxCodeId,omitempty"`
	// The account to use for this invoice line, overriding the default account associated with the item.
	OverrideItemAccountId *string `json:"overrideItemAccountId,omitempty"`
	// A built-in custom field for additional information specific to this invoice line. Unlike the user-defined fields in the `customFields` array, this is a standard QuickBooks field that exists for all invoice lines for convenience. Developers often use this field for tracking information that doesn't fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI.
	OtherCustomField1 *string `json:"otherCustomField1,omitempty"`
	// A second built-in custom field for additional information specific to this invoice line. Unlike the user-defined fields in the `customFields` array, this is a standard QuickBooks field that exists for all invoice lines for convenience. Like `otherCustomField1`, developers often use this field for tracking information that doesn't fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI.
	OtherCustomField2 *string `json:"otherCustomField2,omitempty"`
	LinkToTransactionLine *QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine `json:"linkToTransactionLine,omitempty"`
	// The custom fields for the invoice line object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner `json:"customFields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopInvoicesPostRequestLinesInner QuickbooksDesktopInvoicesPostRequestLinesInner

// NewQuickbooksDesktopInvoicesPostRequestLinesInner instantiates a new QuickbooksDesktopInvoicesPostRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopInvoicesPostRequestLinesInner() *QuickbooksDesktopInvoicesPostRequestLinesInner {
	this := QuickbooksDesktopInvoicesPostRequestLinesInner{}
	return &this
}

// NewQuickbooksDesktopInvoicesPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopInvoicesPostRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopInvoicesPostRequestLinesInnerWithDefaults() *QuickbooksDesktopInvoicesPostRequestLinesInner {
	this := QuickbooksDesktopInvoicesPostRequestLinesInner{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetItemId(v string) {
	o.ItemId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetRate(v string) {
	o.Rate = &v
}

// GetRatePercent returns the RatePercent field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetRatePercent() string {
	if o == nil || IsNil(o.RatePercent) {
		var ret string
		return ret
	}
	return *o.RatePercent
}

// GetRatePercentOk returns a tuple with the RatePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetRatePercentOk() (*string, bool) {
	if o == nil || IsNil(o.RatePercent) {
		return nil, false
	}
	return o.RatePercent, true
}

// HasRatePercent returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasRatePercent() bool {
	if o != nil && !IsNil(o.RatePercent) {
		return true
	}

	return false
}

// SetRatePercent gets a reference to the given string and assigns it to the RatePercent field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetRatePercent(v string) {
	o.RatePercent = &v
}

// GetPriceLevelId returns the PriceLevelId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetPriceLevelId() string {
	if o == nil || IsNil(o.PriceLevelId) {
		var ret string
		return ret
	}
	return *o.PriceLevelId
}

// GetPriceLevelIdOk returns a tuple with the PriceLevelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PriceLevelId) {
		return nil, false
	}
	return o.PriceLevelId, true
}

// HasPriceLevelId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasPriceLevelId() bool {
	if o != nil && !IsNil(o.PriceLevelId) {
		return true
	}

	return false
}

// SetPriceLevelId gets a reference to the given string and assigns it to the PriceLevelId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetPriceLevelId(v string) {
	o.PriceLevelId = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetClassId(v string) {
	o.ClassId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetAmount(v string) {
	o.Amount = &v
}

// GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetPriceRuleConflictStrategy() string {
	if o == nil || IsNil(o.PriceRuleConflictStrategy) {
		var ret string
		return ret
	}
	return *o.PriceRuleConflictStrategy
}

// GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRuleConflictStrategy) {
		return nil, false
	}
	return o.PriceRuleConflictStrategy, true
}

// HasPriceRuleConflictStrategy returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasPriceRuleConflictStrategy() bool {
	if o != nil && !IsNil(o.PriceRuleConflictStrategy) {
		return true
	}

	return false
}

// SetPriceRuleConflictStrategy gets a reference to the given string and assigns it to the PriceRuleConflictStrategy field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetPriceRuleConflictStrategy(v string) {
	o.PriceRuleConflictStrategy = &v
}

// GetInventorySiteId returns the InventorySiteId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetInventorySiteId() string {
	if o == nil || IsNil(o.InventorySiteId) {
		var ret string
		return ret
	}
	return *o.InventorySiteId
}

// GetInventorySiteIdOk returns a tuple with the InventorySiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteId) {
		return nil, false
	}
	return o.InventorySiteId, true
}

// HasInventorySiteId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasInventorySiteId() bool {
	if o != nil && !IsNil(o.InventorySiteId) {
		return true
	}

	return false
}

// SetInventorySiteId gets a reference to the given string and assigns it to the InventorySiteId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetInventorySiteId(v string) {
	o.InventorySiteId = &v
}

// GetInventorySiteLocationId returns the InventorySiteLocationId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetInventorySiteLocationId() string {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		var ret string
		return ret
	}
	return *o.InventorySiteLocationId
}

// GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		return nil, false
	}
	return o.InventorySiteLocationId, true
}

// HasInventorySiteLocationId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasInventorySiteLocationId() bool {
	if o != nil && !IsNil(o.InventorySiteLocationId) {
		return true
	}

	return false
}

// SetInventorySiteLocationId gets a reference to the given string and assigns it to the InventorySiteLocationId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetInventorySiteLocationId(v string) {
	o.InventorySiteLocationId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber) {
		var ret string
		return ret
	}
	return *o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetLotNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LotNumber) {
		return nil, false
	}
	return o.LotNumber, true
}

// HasLotNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasLotNumber() bool {
	if o != nil && !IsNil(o.LotNumber) {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given string and assigns it to the LotNumber field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetLotNumber(v string) {
	o.LotNumber = &v
}

// GetServiceDate returns the ServiceDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetServiceDate() string {
	if o == nil || IsNil(o.ServiceDate) {
		var ret string
		return ret
	}
	return *o.ServiceDate
}

// GetServiceDateOk returns a tuple with the ServiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetServiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDate) {
		return nil, false
	}
	return o.ServiceDate, true
}

// HasServiceDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasServiceDate() bool {
	if o != nil && !IsNil(o.ServiceDate) {
		return true
	}

	return false
}

// SetServiceDate gets a reference to the given string and assigns it to the ServiceDate field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetServiceDate(v string) {
	o.ServiceDate = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetOverrideItemAccountId returns the OverrideItemAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOverrideItemAccountId() string {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		var ret string
		return ret
	}
	return *o.OverrideItemAccountId
}

// GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOverrideItemAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		return nil, false
	}
	return o.OverrideItemAccountId, true
}

// HasOverrideItemAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasOverrideItemAccountId() bool {
	if o != nil && !IsNil(o.OverrideItemAccountId) {
		return true
	}

	return false
}

// SetOverrideItemAccountId gets a reference to the given string and assigns it to the OverrideItemAccountId field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetOverrideItemAccountId(v string) {
	o.OverrideItemAccountId = &v
}

// GetOtherCustomField1 returns the OtherCustomField1 field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOtherCustomField1() string {
	if o == nil || IsNil(o.OtherCustomField1) {
		var ret string
		return ret
	}
	return *o.OtherCustomField1
}

// GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool) {
	if o == nil || IsNil(o.OtherCustomField1) {
		return nil, false
	}
	return o.OtherCustomField1, true
}

// HasOtherCustomField1 returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasOtherCustomField1() bool {
	if o != nil && !IsNil(o.OtherCustomField1) {
		return true
	}

	return false
}

// SetOtherCustomField1 gets a reference to the given string and assigns it to the OtherCustomField1 field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetOtherCustomField1(v string) {
	o.OtherCustomField1 = &v
}

// GetOtherCustomField2 returns the OtherCustomField2 field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOtherCustomField2() string {
	if o == nil || IsNil(o.OtherCustomField2) {
		var ret string
		return ret
	}
	return *o.OtherCustomField2
}

// GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool) {
	if o == nil || IsNil(o.OtherCustomField2) {
		return nil, false
	}
	return o.OtherCustomField2, true
}

// HasOtherCustomField2 returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasOtherCustomField2() bool {
	if o != nil && !IsNil(o.OtherCustomField2) {
		return true
	}

	return false
}

// SetOtherCustomField2 gets a reference to the given string and assigns it to the OtherCustomField2 field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetOtherCustomField2(v string) {
	o.OtherCustomField2 = &v
}

// GetLinkToTransactionLine returns the LinkToTransactionLine field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetLinkToTransactionLine() QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine {
	if o == nil || IsNil(o.LinkToTransactionLine) {
		var ret QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine
		return ret
	}
	return *o.LinkToTransactionLine
}

// GetLinkToTransactionLineOk returns a tuple with the LinkToTransactionLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetLinkToTransactionLineOk() (*QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine, bool) {
	if o == nil || IsNil(o.LinkToTransactionLine) {
		return nil, false
	}
	return o.LinkToTransactionLine, true
}

// HasLinkToTransactionLine returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasLinkToTransactionLine() bool {
	if o != nil && !IsNil(o.LinkToTransactionLine) {
		return true
	}

	return false
}

// SetLinkToTransactionLine gets a reference to the given QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine and assigns it to the LinkToTransactionLine field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetLinkToTransactionLine(v QuickbooksDesktopInvoicesPostRequestLinesInnerLinkToTransactionLine) {
	o.LinkToTransactionLine = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetCustomFields() []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner {
	if o == nil || IsNil(o.CustomFields) {
		var ret []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) GetCustomFieldsOk() ([]QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner and assigns it to the CustomFields field.
func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) SetCustomFields(v []QuickbooksDesktopBillsPostRequestExpenseLinesInnerCustomFieldsInner) {
	o.CustomFields = v
}

func (o QuickbooksDesktopInvoicesPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopInvoicesPostRequestLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["itemId"] = o.ItemId
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
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.RatePercent) {
		toSerialize["ratePercent"] = o.RatePercent
	}
	if !IsNil(o.PriceLevelId) {
		toSerialize["priceLevelId"] = o.PriceLevelId
	}
	if !IsNil(o.ClassId) {
		toSerialize["classId"] = o.ClassId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.PriceRuleConflictStrategy) {
		toSerialize["priceRuleConflictStrategy"] = o.PriceRuleConflictStrategy
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
	if !IsNil(o.ServiceDate) {
		toSerialize["serviceDate"] = o.ServiceDate
	}
	if !IsNil(o.SalesTaxCodeId) {
		toSerialize["salesTaxCodeId"] = o.SalesTaxCodeId
	}
	if !IsNil(o.OverrideItemAccountId) {
		toSerialize["overrideItemAccountId"] = o.OverrideItemAccountId
	}
	if !IsNil(o.OtherCustomField1) {
		toSerialize["otherCustomField1"] = o.OtherCustomField1
	}
	if !IsNil(o.OtherCustomField2) {
		toSerialize["otherCustomField2"] = o.OtherCustomField2
	}
	if !IsNil(o.LinkToTransactionLine) {
		toSerialize["linkToTransactionLine"] = o.LinkToTransactionLine
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopInvoicesPostRequestLinesInner) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopInvoicesPostRequestLinesInner := _QuickbooksDesktopInvoicesPostRequestLinesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopInvoicesPostRequestLinesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopInvoicesPostRequestLinesInner(varQuickbooksDesktopInvoicesPostRequestLinesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "itemId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "ratePercent")
		delete(additionalProperties, "priceLevelId")
		delete(additionalProperties, "classId")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "priceRuleConflictStrategy")
		delete(additionalProperties, "inventorySiteId")
		delete(additionalProperties, "inventorySiteLocationId")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "lotNumber")
		delete(additionalProperties, "serviceDate")
		delete(additionalProperties, "salesTaxCodeId")
		delete(additionalProperties, "overrideItemAccountId")
		delete(additionalProperties, "otherCustomField1")
		delete(additionalProperties, "otherCustomField2")
		delete(additionalProperties, "linkToTransactionLine")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopInvoicesPostRequestLinesInner struct {
	value *QuickbooksDesktopInvoicesPostRequestLinesInner
	isSet bool
}

func (v NullableQuickbooksDesktopInvoicesPostRequestLinesInner) Get() *QuickbooksDesktopInvoicesPostRequestLinesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopInvoicesPostRequestLinesInner) Set(val *QuickbooksDesktopInvoicesPostRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopInvoicesPostRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopInvoicesPostRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopInvoicesPostRequestLinesInner(val *QuickbooksDesktopInvoicesPostRequestLinesInner) *NullableQuickbooksDesktopInvoicesPostRequestLinesInner {
	return &NullableQuickbooksDesktopInvoicesPostRequestLinesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopInvoicesPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopInvoicesPostRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


