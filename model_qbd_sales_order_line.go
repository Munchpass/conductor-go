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

// checks if the QbdSalesOrderLine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdSalesOrderLine{}

// QbdSalesOrderLine struct for QbdSalesOrderLine
type QbdSalesOrderLine struct {
	// The unique identifier assigned by QuickBooks to this sales order line. This ID is unique across all transaction line types.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_sales_order_line\"`.
	ObjectType string `json:"objectType"`
	Item QbdSalesOrderLineItem `json:"item"`
	// A description of this sales order line.
	Description string `json:"description"`
	// The quantity of the item associated with this sales order line. This field cannot be cleared.  **NOTE**: Do not use this field if the associated item is a discount item.
	Quantity float32 `json:"quantity"`
	// The unit-of-measure used for the `quantity` in this sales order line. Must be a valid unit within the item's available units of measure.
	UnitOfMeasure string `json:"unitOfMeasure"`
	OverrideUnitOfMeasureSet QbdSalesOrderLineOverrideUnitOfMeasureSet `json:"overrideUnitOfMeasureSet"`
	// The price per unit for this sales order line. If both `rate` and `amount` are specified, `rate` will be ignored. If both `quantity` and `amount` are specified but not `rate`, QuickBooks will use them to calculate `rate`. Represented as a decimal string. This field cannot be cleared.
	Rate string `json:"rate"`
	// The price of this sales order line expressed as a percentage. Typically used for discount or markup items.
	RatePercent string `json:"ratePercent"`
	Class QbdSalesOrderLineClass `json:"class"`
	// The monetary amount of this sales order line, represented as a decimal string. If both `quantity` and `rate` are specified but not `amount`, QuickBooks will use them to calculate `amount`. If `amount`, `rate`, and `quantity` are all unspecified, then QuickBooks will calculate `amount` based on a `quantity` of `1` and the suggested `rate`. This field cannot be cleared.
	Amount string `json:"amount"`
	InventorySite QbdSalesOrderLineInventorySite `json:"inventorySite"`
	InventorySiteLocation QbdSalesOrderLineInventorySiteLocation `json:"inventorySiteLocation"`
	// The serial number of the item associated with this sales order line. This is used for tracking individual units of serialized inventory items.
	SerialNumber string `json:"serialNumber"`
	// The lot number of the item associated with this sales order line. Used for tracking groups of inventory items that are purchased or manufactured together.
	LotNumber string `json:"lotNumber"`
	// The expiration date for the serial number or lot number of the item associated with this sales order line, in ISO 8601 format (YYYY-MM-DD). This is particularly relevant for perishable or time-sensitive inventory items. Note that this field is only supported on QuickBooks Desktop 2023 or later.
	ExpirationDate string `json:"expirationDate"`
	SalesTaxCode QbdSalesOrderLineSalesTaxCode `json:"salesTaxCode"`
	// The number of units of this sales order line's `quantity` that have been invoiced.
	QuantityInvoiced float32 `json:"quantityInvoiced"`
	// Indicates whether this sales order line has been manually marked as closed, even if it has not been invoiced.
	IsManuallyClosed bool `json:"isManuallyClosed"`
	// A built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the `customFields` array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Developers often use this field for tracking information that doesn't fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI.
	OtherCustomField1 string `json:"otherCustomField1"`
	// A second built-in custom field for additional information specific to this sales order line. Unlike the user-defined fields in the `customFields` array, this is a standard QuickBooks field that exists for all sales order lines for convenience. Like `otherCustomField1`, developers often use this field for tracking information that doesn't fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI.
	OtherCustomField2 string `json:"otherCustomField2"`
	// The custom fields for the sales order line object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QbdCustomField `json:"customFields"`
	AdditionalProperties map[string]interface{}
}

type _QbdSalesOrderLine QbdSalesOrderLine

// NewQbdSalesOrderLine instantiates a new QbdSalesOrderLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdSalesOrderLine(id string, objectType string, item QbdSalesOrderLineItem, description string, quantity float32, unitOfMeasure string, overrideUnitOfMeasureSet QbdSalesOrderLineOverrideUnitOfMeasureSet, rate string, ratePercent string, class QbdSalesOrderLineClass, amount string, inventorySite QbdSalesOrderLineInventorySite, inventorySiteLocation QbdSalesOrderLineInventorySiteLocation, serialNumber string, lotNumber string, expirationDate string, salesTaxCode QbdSalesOrderLineSalesTaxCode, quantityInvoiced float32, isManuallyClosed bool, otherCustomField1 string, otherCustomField2 string, customFields []QbdCustomField) *QbdSalesOrderLine {
	this := QbdSalesOrderLine{}
	this.Id = id
	this.ObjectType = objectType
	this.Item = item
	this.Description = description
	this.Quantity = quantity
	this.UnitOfMeasure = unitOfMeasure
	this.OverrideUnitOfMeasureSet = overrideUnitOfMeasureSet
	this.Rate = rate
	this.RatePercent = ratePercent
	this.Class = class
	this.Amount = amount
	this.InventorySite = inventorySite
	this.InventorySiteLocation = inventorySiteLocation
	this.SerialNumber = serialNumber
	this.LotNumber = lotNumber
	this.ExpirationDate = expirationDate
	this.SalesTaxCode = salesTaxCode
	this.QuantityInvoiced = quantityInvoiced
	this.IsManuallyClosed = isManuallyClosed
	this.OtherCustomField1 = otherCustomField1
	this.OtherCustomField2 = otherCustomField2
	this.CustomFields = customFields
	return &this
}

// NewQbdSalesOrderLineWithDefaults instantiates a new QbdSalesOrderLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdSalesOrderLineWithDefaults() *QbdSalesOrderLine {
	this := QbdSalesOrderLine{}
	return &this
}

// GetId returns the Id field value
func (o *QbdSalesOrderLine) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdSalesOrderLine) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QbdSalesOrderLine) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QbdSalesOrderLine) SetObjectType(v string) {
	o.ObjectType = v
}

// GetItem returns the Item field value
func (o *QbdSalesOrderLine) GetItem() QbdSalesOrderLineItem {
	if o == nil {
		var ret QbdSalesOrderLineItem
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetItemOk() (*QbdSalesOrderLineItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *QbdSalesOrderLine) SetItem(v QbdSalesOrderLineItem) {
	o.Item = v
}

// GetDescription returns the Description field value
func (o *QbdSalesOrderLine) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *QbdSalesOrderLine) SetDescription(v string) {
	o.Description = v
}

// GetQuantity returns the Quantity field value
func (o *QbdSalesOrderLine) GetQuantity() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetQuantityOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *QbdSalesOrderLine) SetQuantity(v float32) {
	o.Quantity = v
}

// GetUnitOfMeasure returns the UnitOfMeasure field value
func (o *QbdSalesOrderLine) GetUnitOfMeasure() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnitOfMeasure
}

// GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetUnitOfMeasureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitOfMeasure, true
}

// SetUnitOfMeasure sets field value
func (o *QbdSalesOrderLine) SetUnitOfMeasure(v string) {
	o.UnitOfMeasure = v
}

// GetOverrideUnitOfMeasureSet returns the OverrideUnitOfMeasureSet field value
func (o *QbdSalesOrderLine) GetOverrideUnitOfMeasureSet() QbdSalesOrderLineOverrideUnitOfMeasureSet {
	if o == nil {
		var ret QbdSalesOrderLineOverrideUnitOfMeasureSet
		return ret
	}

	return o.OverrideUnitOfMeasureSet
}

// GetOverrideUnitOfMeasureSetOk returns a tuple with the OverrideUnitOfMeasureSet field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetOverrideUnitOfMeasureSetOk() (*QbdSalesOrderLineOverrideUnitOfMeasureSet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OverrideUnitOfMeasureSet, true
}

// SetOverrideUnitOfMeasureSet sets field value
func (o *QbdSalesOrderLine) SetOverrideUnitOfMeasureSet(v QbdSalesOrderLineOverrideUnitOfMeasureSet) {
	o.OverrideUnitOfMeasureSet = v
}

// GetRate returns the Rate field value
func (o *QbdSalesOrderLine) GetRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value
func (o *QbdSalesOrderLine) SetRate(v string) {
	o.Rate = v
}

// GetRatePercent returns the RatePercent field value
func (o *QbdSalesOrderLine) GetRatePercent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RatePercent
}

// GetRatePercentOk returns a tuple with the RatePercent field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetRatePercentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatePercent, true
}

// SetRatePercent sets field value
func (o *QbdSalesOrderLine) SetRatePercent(v string) {
	o.RatePercent = v
}

// GetClass returns the Class field value
func (o *QbdSalesOrderLine) GetClass() QbdSalesOrderLineClass {
	if o == nil {
		var ret QbdSalesOrderLineClass
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetClassOk() (*QbdSalesOrderLineClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *QbdSalesOrderLine) SetClass(v QbdSalesOrderLineClass) {
	o.Class = v
}

// GetAmount returns the Amount field value
func (o *QbdSalesOrderLine) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *QbdSalesOrderLine) SetAmount(v string) {
	o.Amount = v
}

// GetInventorySite returns the InventorySite field value
func (o *QbdSalesOrderLine) GetInventorySite() QbdSalesOrderLineInventorySite {
	if o == nil {
		var ret QbdSalesOrderLineInventorySite
		return ret
	}

	return o.InventorySite
}

// GetInventorySiteOk returns a tuple with the InventorySite field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetInventorySiteOk() (*QbdSalesOrderLineInventorySite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventorySite, true
}

// SetInventorySite sets field value
func (o *QbdSalesOrderLine) SetInventorySite(v QbdSalesOrderLineInventorySite) {
	o.InventorySite = v
}

// GetInventorySiteLocation returns the InventorySiteLocation field value
func (o *QbdSalesOrderLine) GetInventorySiteLocation() QbdSalesOrderLineInventorySiteLocation {
	if o == nil {
		var ret QbdSalesOrderLineInventorySiteLocation
		return ret
	}

	return o.InventorySiteLocation
}

// GetInventorySiteLocationOk returns a tuple with the InventorySiteLocation field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetInventorySiteLocationOk() (*QbdSalesOrderLineInventorySiteLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventorySiteLocation, true
}

// SetInventorySiteLocation sets field value
func (o *QbdSalesOrderLine) SetInventorySiteLocation(v QbdSalesOrderLineInventorySiteLocation) {
	o.InventorySiteLocation = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *QbdSalesOrderLine) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *QbdSalesOrderLine) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetLotNumber returns the LotNumber field value
func (o *QbdSalesOrderLine) GetLotNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LotNumber
}

// GetLotNumberOk returns a tuple with the LotNumber field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetLotNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LotNumber, true
}

// SetLotNumber sets field value
func (o *QbdSalesOrderLine) SetLotNumber(v string) {
	o.LotNumber = v
}

// GetExpirationDate returns the ExpirationDate field value
func (o *QbdSalesOrderLine) GetExpirationDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetExpirationDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationDate, true
}

// SetExpirationDate sets field value
func (o *QbdSalesOrderLine) SetExpirationDate(v string) {
	o.ExpirationDate = v
}

// GetSalesTaxCode returns the SalesTaxCode field value
func (o *QbdSalesOrderLine) GetSalesTaxCode() QbdSalesOrderLineSalesTaxCode {
	if o == nil {
		var ret QbdSalesOrderLineSalesTaxCode
		return ret
	}

	return o.SalesTaxCode
}

// GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetSalesTaxCodeOk() (*QbdSalesOrderLineSalesTaxCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SalesTaxCode, true
}

// SetSalesTaxCode sets field value
func (o *QbdSalesOrderLine) SetSalesTaxCode(v QbdSalesOrderLineSalesTaxCode) {
	o.SalesTaxCode = v
}

// GetQuantityInvoiced returns the QuantityInvoiced field value
func (o *QbdSalesOrderLine) GetQuantityInvoiced() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.QuantityInvoiced
}

// GetQuantityInvoicedOk returns a tuple with the QuantityInvoiced field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetQuantityInvoicedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityInvoiced, true
}

// SetQuantityInvoiced sets field value
func (o *QbdSalesOrderLine) SetQuantityInvoiced(v float32) {
	o.QuantityInvoiced = v
}

// GetIsManuallyClosed returns the IsManuallyClosed field value
func (o *QbdSalesOrderLine) GetIsManuallyClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsManuallyClosed
}

// GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetIsManuallyClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsManuallyClosed, true
}

// SetIsManuallyClosed sets field value
func (o *QbdSalesOrderLine) SetIsManuallyClosed(v bool) {
	o.IsManuallyClosed = v
}

// GetOtherCustomField1 returns the OtherCustomField1 field value
func (o *QbdSalesOrderLine) GetOtherCustomField1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtherCustomField1
}

// GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetOtherCustomField1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherCustomField1, true
}

// SetOtherCustomField1 sets field value
func (o *QbdSalesOrderLine) SetOtherCustomField1(v string) {
	o.OtherCustomField1 = v
}

// GetOtherCustomField2 returns the OtherCustomField2 field value
func (o *QbdSalesOrderLine) GetOtherCustomField2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OtherCustomField2
}

// GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetOtherCustomField2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OtherCustomField2, true
}

// SetOtherCustomField2 sets field value
func (o *QbdSalesOrderLine) SetOtherCustomField2(v string) {
	o.OtherCustomField2 = v
}

// GetCustomFields returns the CustomFields field value
func (o *QbdSalesOrderLine) GetCustomFields() []QbdCustomField {
	if o == nil {
		var ret []QbdCustomField
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
func (o *QbdSalesOrderLine) GetCustomFieldsOk() ([]QbdCustomField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *QbdSalesOrderLine) SetCustomFields(v []QbdCustomField) {
	o.CustomFields = v
}

func (o QbdSalesOrderLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdSalesOrderLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["item"] = o.Item
	toSerialize["description"] = o.Description
	toSerialize["quantity"] = o.Quantity
	toSerialize["unitOfMeasure"] = o.UnitOfMeasure
	toSerialize["overrideUnitOfMeasureSet"] = o.OverrideUnitOfMeasureSet
	toSerialize["rate"] = o.Rate
	toSerialize["ratePercent"] = o.RatePercent
	toSerialize["class"] = o.Class
	toSerialize["amount"] = o.Amount
	toSerialize["inventorySite"] = o.InventorySite
	toSerialize["inventorySiteLocation"] = o.InventorySiteLocation
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["lotNumber"] = o.LotNumber
	toSerialize["expirationDate"] = o.ExpirationDate
	toSerialize["salesTaxCode"] = o.SalesTaxCode
	toSerialize["quantityInvoiced"] = o.QuantityInvoiced
	toSerialize["isManuallyClosed"] = o.IsManuallyClosed
	toSerialize["otherCustomField1"] = o.OtherCustomField1
	toSerialize["otherCustomField2"] = o.OtherCustomField2
	toSerialize["customFields"] = o.CustomFields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdSalesOrderLine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"item",
		"description",
		"quantity",
		"unitOfMeasure",
		"overrideUnitOfMeasureSet",
		"rate",
		"ratePercent",
		"class",
		"amount",
		"inventorySite",
		"inventorySiteLocation",
		"serialNumber",
		"lotNumber",
		"expirationDate",
		"salesTaxCode",
		"quantityInvoiced",
		"isManuallyClosed",
		"otherCustomField1",
		"otherCustomField2",
		"customFields",
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

	varQbdSalesOrderLine := _QbdSalesOrderLine{}

	err = json.Unmarshal(data, &varQbdSalesOrderLine)

	if err != nil {
		return err
	}

	*o = QbdSalesOrderLine(varQbdSalesOrderLine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "item")
		delete(additionalProperties, "description")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "unitOfMeasure")
		delete(additionalProperties, "overrideUnitOfMeasureSet")
		delete(additionalProperties, "rate")
		delete(additionalProperties, "ratePercent")
		delete(additionalProperties, "class")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "inventorySite")
		delete(additionalProperties, "inventorySiteLocation")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "lotNumber")
		delete(additionalProperties, "expirationDate")
		delete(additionalProperties, "salesTaxCode")
		delete(additionalProperties, "quantityInvoiced")
		delete(additionalProperties, "isManuallyClosed")
		delete(additionalProperties, "otherCustomField1")
		delete(additionalProperties, "otherCustomField2")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdSalesOrderLine struct {
	value *QbdSalesOrderLine
	isSet bool
}

func (v NullableQbdSalesOrderLine) Get() *QbdSalesOrderLine {
	return v.value
}

func (v *NullableQbdSalesOrderLine) Set(val *QbdSalesOrderLine) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdSalesOrderLine) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdSalesOrderLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdSalesOrderLine(val *QbdSalesOrderLine) *NullableQbdSalesOrderLine {
	return &NullableQbdSalesOrderLine{value: val, isSet: true}
}

func (v NullableQbdSalesOrderLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdSalesOrderLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


