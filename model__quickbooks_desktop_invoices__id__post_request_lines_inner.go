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

// checks if the QuickbooksDesktopInvoicesIdPostRequestLinesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopInvoicesIdPostRequestLinesInner{}

// QuickbooksDesktopInvoicesIdPostRequestLinesInner struct for QuickbooksDesktopInvoicesIdPostRequestLinesInner
type QuickbooksDesktopInvoicesIdPostRequestLinesInner struct {
	// The QuickBooks-assigned unique identifier of an existing invoice line you wish to retain or update.  **IMPORTANT**: Set this field to `-1` for new invoice lines you wish to add.
	Id string `json:"id"`
	// The item associated with this invoice line. This can refer to any good or service that the business buys or sells, including item types such as a service item, inventory item, or special calculation item like a discount item or sales-tax item.
	ItemId *string `json:"itemId,omitempty"`
	// A description of this invoice line.
	Description *string `json:"description,omitempty"`
	// The quantity of the item associated with this invoice line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item.
	Quantity *float32 `json:"quantity,omitempty"`
	// The unit-of-measure used for the `quantity` in this invoice line. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
	// Specifies an alternative unit-of-measure set when updating this invoice line's `unitOfMeasure` field (e.g., \"pound\" or \"kilogram\"). This allows you to select units from a different set than the item's default unit-of-measure set, which remains unchanged on the item itself. The override applies only to this specific line. For example, you can sell an item typically measured in volume units using weight units in a specific transaction by specifying a different unit-of-measure set with this field.
	OverrideUnitOfMeasureSetId *string `json:"overrideUnitOfMeasureSetId,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopInvoicesIdPostRequestLinesInner QuickbooksDesktopInvoicesIdPostRequestLinesInner

// NewQuickbooksDesktopInvoicesIdPostRequestLinesInner instantiates a new QuickbooksDesktopInvoicesIdPostRequestLinesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopInvoicesIdPostRequestLinesInner(id string) *QuickbooksDesktopInvoicesIdPostRequestLinesInner {
	this := QuickbooksDesktopInvoicesIdPostRequestLinesInner{}
	this.Id = id
	return &this
}

// NewQuickbooksDesktopInvoicesIdPostRequestLinesInnerWithDefaults instantiates a new QuickbooksDesktopInvoicesIdPostRequestLinesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopInvoicesIdPostRequestLinesInnerWithDefaults() *QuickbooksDesktopInvoicesIdPostRequestLinesInner {
	this := QuickbooksDesktopInvoicesIdPostRequestLinesInner{}
	return &this
}

// GetId returns the Id field value
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetId(v string) {
	o.Id = v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetItemId(v string) {
	o.ItemId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetUnitOfMeasure() string {
	if o == nil || IsNil(o.UnitOfMeasure) {
		var ret string
		return ret
	}
	return *o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil || IsNil(o.UnitOfMeasure) {
		return nil, false
	}
	return o.UnitOfMeasure, true
}

// HasUnitOfMeasure returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasUnitOfMeasure() bool {
	if o != nil && !IsNil(o.UnitOfMeasure) {
		return true
	}

	return false
}

// SetUnitOfMeasure gets a reference to the given string and assigns it to the UnitOfMeasure field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = &v
}

// GetOverrideUnitOfMeasureSetId returns the OverrideUnitOfMeasureSetId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOverrideUnitOfMeasureSetId() string {
	if o == nil || IsNil(o.OverrideUnitOfMeasureSetId) {
		var ret string
		return ret
	}
	return *o.OverrideUnitOfMeasureSetId
}

// GetOverrideUnitOfMeasureSetIdOk returns a tuple with the OverrideUnitOfMeasureSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOverrideUnitOfMeasureSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideUnitOfMeasureSetId) {
		return nil, false
	}
	return o.OverrideUnitOfMeasureSetId, true
}

// HasOverrideUnitOfMeasureSetId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasOverrideUnitOfMeasureSetId() bool {
	if o != nil && !IsNil(o.OverrideUnitOfMeasureSetId) {
		return true
	}

	return false
}

// SetOverrideUnitOfMeasureSetId gets a reference to the given string and assigns it to the OverrideUnitOfMeasureSetId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetOverrideUnitOfMeasureSetId(v string) {
	o.OverrideUnitOfMeasureSetId = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetRate(v string) {
	o.Rate = &v
}

// GetRatePercent returns the RatePercent field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetRatePercent() string {
	if o == nil || IsNil(o.RatePercent) {
		var ret string
		return ret
	}
	return *o.RatePercent
}

// GetRatePercentOk returns a tuple with the RatePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetRatePercentOk() (*string, bool) {
	if o == nil || IsNil(o.RatePercent) {
		return nil, false
	}
	return o.RatePercent, true
}

// HasRatePercent returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasRatePercent() bool {
	if o != nil && !IsNil(o.RatePercent) {
		return true
	}

	return false
}

// SetRatePercent gets a reference to the given string and assigns it to the RatePercent field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetRatePercent(v string) {
	o.RatePercent = &v
}

// GetPriceLevelId returns the PriceLevelId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetPriceLevelId() string {
	if o == nil || IsNil(o.PriceLevelId) {
		var ret string
		return ret
	}
	return *o.PriceLevelId
}

// GetPriceLevelIdOk returns a tuple with the PriceLevelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetPriceLevelIdOk() (*string, bool) {
	if o == nil || IsNil(o.PriceLevelId) {
		return nil, false
	}
	return o.PriceLevelId, true
}

// HasPriceLevelId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasPriceLevelId() bool {
	if o != nil && !IsNil(o.PriceLevelId) {
		return true
	}

	return false
}

// SetPriceLevelId gets a reference to the given string and assigns it to the PriceLevelId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetPriceLevelId(v string) {
	o.PriceLevelId = &v
}

// GetClassId returns the ClassId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetClassId() string {
	if o == nil || IsNil(o.ClassId) {
		var ret string
		return ret
	}
	return *o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetClassIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClassId) {
		return nil, false
	}
	return o.ClassId, true
}

// HasClassId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasClassId() bool {
	if o != nil && !IsNil(o.ClassId) {
		return true
	}

	return false
}

// SetClassId gets a reference to the given string and assigns it to the ClassId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetClassId(v string) {
	o.ClassId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetAmount(v string) {
	o.Amount = &v
}

// GetPriceRuleConflictStrategy returns the PriceRuleConflictStrategy field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetPriceRuleConflictStrategy() string {
	if o == nil || IsNil(o.PriceRuleConflictStrategy) {
		var ret string
		return ret
	}
	return *o.PriceRuleConflictStrategy
}

// GetPriceRuleConflictStrategyOk returns a tuple with the PriceRuleConflictStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetPriceRuleConflictStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.PriceRuleConflictStrategy) {
		return nil, false
	}
	return o.PriceRuleConflictStrategy, true
}

// HasPriceRuleConflictStrategy returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasPriceRuleConflictStrategy() bool {
	if o != nil && !IsNil(o.PriceRuleConflictStrategy) {
		return true
	}

	return false
}

// SetPriceRuleConflictStrategy gets a reference to the given string and assigns it to the PriceRuleConflictStrategy field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetPriceRuleConflictStrategy(v string) {
	o.PriceRuleConflictStrategy = &v
}

// GetInventorySiteId returns the InventorySiteId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetInventorySiteId() string {
	if o == nil || IsNil(o.InventorySiteId) {
		var ret string
		return ret
	}
	return *o.InventorySiteId
}

// GetInventorySiteIdOk returns a tuple with the InventorySiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetInventorySiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteId) {
		return nil, false
	}
	return o.InventorySiteId, true
}

// HasInventorySiteId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasInventorySiteId() bool {
	if o != nil && !IsNil(o.InventorySiteId) {
		return true
	}

	return false
}

// SetInventorySiteId gets a reference to the given string and assigns it to the InventorySiteId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetInventorySiteId(v string) {
	o.InventorySiteId = &v
}

// GetInventorySiteLocationId returns the InventorySiteLocationId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetInventorySiteLocationId() string {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		var ret string
		return ret
	}
	return *o.InventorySiteLocationId
}

// GetInventorySiteLocationIdOk returns a tuple with the InventorySiteLocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetInventorySiteLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InventorySiteLocationId) {
		return nil, false
	}
	return o.InventorySiteLocationId, true
}

// HasInventorySiteLocationId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasInventorySiteLocationId() bool {
	if o != nil && !IsNil(o.InventorySiteLocationId) {
		return true
	}

	return false
}

// SetInventorySiteLocationId gets a reference to the given string and assigns it to the InventorySiteLocationId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetInventorySiteLocationId(v string) {
	o.InventorySiteLocationId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetSerialNumber() string {
	if o == nil || IsNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetSerialNumberOk() (*string, bool) {
	if o == nil || IsNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasSerialNumber() bool {
	if o != nil && !IsNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber) {
		var ret string
		return ret
	}
	return *o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetLotNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LotNumber) {
		return nil, false
	}
	return o.LotNumber, true
}

// HasLotNumber returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasLotNumber() bool {
	if o != nil && !IsNil(o.LotNumber) {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given string and assigns it to the LotNumber field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetLotNumber(v string) {
	o.LotNumber = &v
}

// GetServiceDate returns the ServiceDate field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetServiceDate() string {
	if o == nil || IsNil(o.ServiceDate) {
		var ret string
		return ret
	}
	return *o.ServiceDate
}

// GetServiceDateOk returns a tuple with the ServiceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetServiceDateOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDate) {
		return nil, false
	}
	return o.ServiceDate, true
}

// HasServiceDate returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasServiceDate() bool {
	if o != nil && !IsNil(o.ServiceDate) {
		return true
	}

	return false
}

// SetServiceDate gets a reference to the given string and assigns it to the ServiceDate field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetServiceDate(v string) {
	o.ServiceDate = &v
}

// GetSalesTaxCodeId returns the SalesTaxCodeId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetSalesTaxCodeId() string {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		var ret string
		return ret
	}
	return *o.SalesTaxCodeId
}

// GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetSalesTaxCodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.SalesTaxCodeId) {
		return nil, false
	}
	return o.SalesTaxCodeId, true
}

// HasSalesTaxCodeId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasSalesTaxCodeId() bool {
	if o != nil && !IsNil(o.SalesTaxCodeId) {
		return true
	}

	return false
}

// SetSalesTaxCodeId gets a reference to the given string and assigns it to the SalesTaxCodeId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetSalesTaxCodeId(v string) {
	o.SalesTaxCodeId = &v
}

// GetOverrideItemAccountId returns the OverrideItemAccountId field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOverrideItemAccountId() string {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		var ret string
		return ret
	}
	return *o.OverrideItemAccountId
}

// GetOverrideItemAccountIdOk returns a tuple with the OverrideItemAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOverrideItemAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.OverrideItemAccountId) {
		return nil, false
	}
	return o.OverrideItemAccountId, true
}

// HasOverrideItemAccountId returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasOverrideItemAccountId() bool {
	if o != nil && !IsNil(o.OverrideItemAccountId) {
		return true
	}

	return false
}

// SetOverrideItemAccountId gets a reference to the given string and assigns it to the OverrideItemAccountId field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetOverrideItemAccountId(v string) {
	o.OverrideItemAccountId = &v
}

// GetOtherCustomField1 returns the OtherCustomField1 field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOtherCustomField1() string {
	if o == nil || IsNil(o.OtherCustomField1) {
		var ret string
		return ret
	}
	return *o.OtherCustomField1
}

// GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOtherCustomField1Ok() (*string, bool) {
	if o == nil || IsNil(o.OtherCustomField1) {
		return nil, false
	}
	return o.OtherCustomField1, true
}

// HasOtherCustomField1 returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasOtherCustomField1() bool {
	if o != nil && !IsNil(o.OtherCustomField1) {
		return true
	}

	return false
}

// SetOtherCustomField1 gets a reference to the given string and assigns it to the OtherCustomField1 field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetOtherCustomField1(v string) {
	o.OtherCustomField1 = &v
}

// GetOtherCustomField2 returns the OtherCustomField2 field value if set, zero value otherwise.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOtherCustomField2() string {
	if o == nil || IsNil(o.OtherCustomField2) {
		var ret string
		return ret
	}
	return *o.OtherCustomField2
}

// GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) GetOtherCustomField2Ok() (*string, bool) {
	if o == nil || IsNil(o.OtherCustomField2) {
		return nil, false
	}
	return o.OtherCustomField2, true
}

// HasOtherCustomField2 returns a boolean if a field has been set.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) HasOtherCustomField2() bool {
	if o != nil && !IsNil(o.OtherCustomField2) {
		return true
	}

	return false
}

// SetOtherCustomField2 gets a reference to the given string and assigns it to the OtherCustomField2 field.
func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) SetOtherCustomField2(v string) {
	o.OtherCustomField2 = &v
}

func (o QuickbooksDesktopInvoicesIdPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopInvoicesIdPostRequestLinesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	if !IsNil(o.OverrideUnitOfMeasureSetId) {
		toSerialize["overrideUnitOfMeasureSetId"] = o.OverrideUnitOfMeasureSetId
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopInvoicesIdPostRequestLinesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varQuickbooksDesktopInvoicesIdPostRequestLinesInner := _QuickbooksDesktopInvoicesIdPostRequestLinesInner{}

	err = json.Unmarshal(data, &varQuickbooksDesktopInvoicesIdPostRequestLinesInner)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopInvoicesIdPostRequestLinesInner(varQuickbooksDesktopInvoicesIdPostRequestLinesInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "itemId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "overrideUnitOfMeasureSetId")
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner struct {
	value *QuickbooksDesktopInvoicesIdPostRequestLinesInner
	isSet bool
}

func (v NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) Get() *QuickbooksDesktopInvoicesIdPostRequestLinesInner {
	return v.value
}

func (v *NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) Set(val *QuickbooksDesktopInvoicesIdPostRequestLinesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopInvoicesIdPostRequestLinesInner(val *QuickbooksDesktopInvoicesIdPostRequestLinesInner) *NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner {
	return &NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopInvoicesIdPostRequestLinesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


