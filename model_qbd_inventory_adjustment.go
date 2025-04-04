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

// checks if the QbdInventoryAdjustment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdInventoryAdjustment{}

// QbdInventoryAdjustment struct for QbdInventoryAdjustment
type QbdInventoryAdjustment struct {
	// The unique identifier assigned by QuickBooks to this inventory adjustment. This ID is unique across all transaction types.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_inventory_adjustment\"`.
	ObjectType string `json:"objectType"`
	// The date and time when this inventory adjustment was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	CreatedAt string `json:"createdAt"`
	// The date and time when this inventory adjustment was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	UpdatedAt string `json:"updatedAt"`
	// The current QuickBooks-assigned revision number of this inventory adjustment object, which changes each time the object is modified. When updating this object, you must provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	Account QbdInventoryAdjustmentAccount `json:"account"`
	InventorySite QbdInventoryAdjustmentInventorySite `json:"inventorySite"`
	// The date of this inventory adjustment, in ISO 8601 format (YYYY-MM-DD).
	TransactionDate string `json:"transactionDate"`
	// The case-sensitive user-defined reference number for this inventory adjustment, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user.
	RefNumber string `json:"refNumber"`
	Customer QbdInventoryAdjustmentCustomer `json:"customer"`
	Class QbdInventoryAdjustmentClass `json:"class"`
	// A memo or note for this inventory adjustment.
	Memo string `json:"memo"`
	// A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.
	ExternalId string `json:"externalId"`
	// The inventory adjustment's item lines, each representing the adjustment of an inventory item's quantity, value, serial number, or lot number.
	Lines []QbdInventoryAdjustmentLine `json:"lines"`
	// The custom fields for the inventory adjustment object, added as user-defined data extensions, not included in the standard QuickBooks object.
	CustomFields []QbdCustomField `json:"customFields"`
	AdditionalProperties map[string]interface{}
}

type _QbdInventoryAdjustment QbdInventoryAdjustment

// NewQbdInventoryAdjustment instantiates a new QbdInventoryAdjustment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdInventoryAdjustment(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, account QbdInventoryAdjustmentAccount, inventorySite QbdInventoryAdjustmentInventorySite, transactionDate string, refNumber string, customer QbdInventoryAdjustmentCustomer, class QbdInventoryAdjustmentClass, memo string, externalId string, lines []QbdInventoryAdjustmentLine, customFields []QbdCustomField) *QbdInventoryAdjustment {
	this := QbdInventoryAdjustment{}
	this.Id = id
	this.ObjectType = objectType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.RevisionNumber = revisionNumber
	this.Account = account
	this.InventorySite = inventorySite
	this.TransactionDate = transactionDate
	this.RefNumber = refNumber
	this.Customer = customer
	this.Class = class
	this.Memo = memo
	this.ExternalId = externalId
	this.Lines = lines
	this.CustomFields = customFields
	return &this
}

// NewQbdInventoryAdjustmentWithDefaults instantiates a new QbdInventoryAdjustment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdInventoryAdjustmentWithDefaults() *QbdInventoryAdjustment {
	this := QbdInventoryAdjustment{}
	return &this
}

// GetId returns the Id field value
func (o *QbdInventoryAdjustment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdInventoryAdjustment) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QbdInventoryAdjustment) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QbdInventoryAdjustment) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *QbdInventoryAdjustment) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *QbdInventoryAdjustment) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *QbdInventoryAdjustment) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *QbdInventoryAdjustment) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QbdInventoryAdjustment) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QbdInventoryAdjustment) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetAccount returns the Account field value
func (o *QbdInventoryAdjustment) GetAccount() QbdInventoryAdjustmentAccount {
	if o == nil {
		var ret QbdInventoryAdjustmentAccount
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetAccountOk() (*QbdInventoryAdjustmentAccount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *QbdInventoryAdjustment) SetAccount(v QbdInventoryAdjustmentAccount) {
	o.Account = v
}

// GetInventorySite returns the InventorySite field value
func (o *QbdInventoryAdjustment) GetInventorySite() QbdInventoryAdjustmentInventorySite {
	if o == nil {
		var ret QbdInventoryAdjustmentInventorySite
		return ret
	}

	return o.InventorySite
}

// GetInventorySiteOk returns a tuple with the InventorySite field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetInventorySiteOk() (*QbdInventoryAdjustmentInventorySite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InventorySite, true
}

// SetInventorySite sets field value
func (o *QbdInventoryAdjustment) SetInventorySite(v QbdInventoryAdjustmentInventorySite) {
	o.InventorySite = v
}

// GetTransactionDate returns the TransactionDate field value
func (o *QbdInventoryAdjustment) GetTransactionDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetTransactionDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionDate, true
}

// SetTransactionDate sets field value
func (o *QbdInventoryAdjustment) SetTransactionDate(v string) {
	o.TransactionDate = v
}

// GetRefNumber returns the RefNumber field value
func (o *QbdInventoryAdjustment) GetRefNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefNumber
}

// GetRefNumberOk returns a tuple with the RefNumber field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetRefNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefNumber, true
}

// SetRefNumber sets field value
func (o *QbdInventoryAdjustment) SetRefNumber(v string) {
	o.RefNumber = v
}

// GetCustomer returns the Customer field value
func (o *QbdInventoryAdjustment) GetCustomer() QbdInventoryAdjustmentCustomer {
	if o == nil {
		var ret QbdInventoryAdjustmentCustomer
		return ret
	}

	return o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetCustomerOk() (*QbdInventoryAdjustmentCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Customer, true
}

// SetCustomer sets field value
func (o *QbdInventoryAdjustment) SetCustomer(v QbdInventoryAdjustmentCustomer) {
	o.Customer = v
}

// GetClass returns the Class field value
func (o *QbdInventoryAdjustment) GetClass() QbdInventoryAdjustmentClass {
	if o == nil {
		var ret QbdInventoryAdjustmentClass
		return ret
	}

	return o.Class
}

// GetClassOk returns a tuple with the Class field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetClassOk() (*QbdInventoryAdjustmentClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Class, true
}

// SetClass sets field value
func (o *QbdInventoryAdjustment) SetClass(v QbdInventoryAdjustmentClass) {
	o.Class = v
}

// GetMemo returns the Memo field value
func (o *QbdInventoryAdjustment) GetMemo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Memo
}

// GetMemoOk returns a tuple with the Memo field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetMemoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Memo, true
}

// SetMemo sets field value
func (o *QbdInventoryAdjustment) SetMemo(v string) {
	o.Memo = v
}

// GetExternalId returns the ExternalId field value
func (o *QbdInventoryAdjustment) GetExternalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetExternalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalId, true
}

// SetExternalId sets field value
func (o *QbdInventoryAdjustment) SetExternalId(v string) {
	o.ExternalId = v
}

// GetLines returns the Lines field value
func (o *QbdInventoryAdjustment) GetLines() []QbdInventoryAdjustmentLine {
	if o == nil {
		var ret []QbdInventoryAdjustmentLine
		return ret
	}

	return o.Lines
}

// GetLinesOk returns a tuple with the Lines field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetLinesOk() ([]QbdInventoryAdjustmentLine, bool) {
	if o == nil {
		return nil, false
	}
	return o.Lines, true
}

// SetLines sets field value
func (o *QbdInventoryAdjustment) SetLines(v []QbdInventoryAdjustmentLine) {
	o.Lines = v
}

// GetCustomFields returns the CustomFields field value
func (o *QbdInventoryAdjustment) GetCustomFields() []QbdCustomField {
	if o == nil {
		var ret []QbdCustomField
		return ret
	}

	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value
// and a boolean to check if the value has been set.
func (o *QbdInventoryAdjustment) GetCustomFieldsOk() ([]QbdCustomField, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomFields, true
}

// SetCustomFields sets field value
func (o *QbdInventoryAdjustment) SetCustomFields(v []QbdCustomField) {
	o.CustomFields = v
}

func (o QbdInventoryAdjustment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdInventoryAdjustment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["revisionNumber"] = o.RevisionNumber
	toSerialize["account"] = o.Account
	toSerialize["inventorySite"] = o.InventorySite
	toSerialize["transactionDate"] = o.TransactionDate
	toSerialize["refNumber"] = o.RefNumber
	toSerialize["customer"] = o.Customer
	toSerialize["class"] = o.Class
	toSerialize["memo"] = o.Memo
	toSerialize["externalId"] = o.ExternalId
	toSerialize["lines"] = o.Lines
	toSerialize["customFields"] = o.CustomFields

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdInventoryAdjustment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"createdAt",
		"updatedAt",
		"revisionNumber",
		"account",
		"inventorySite",
		"transactionDate",
		"refNumber",
		"customer",
		"class",
		"memo",
		"externalId",
		"lines",
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

	varQbdInventoryAdjustment := _QbdInventoryAdjustment{}

	err = json.Unmarshal(data, &varQbdInventoryAdjustment)

	if err != nil {
		return err
	}

	*o = QbdInventoryAdjustment(varQbdInventoryAdjustment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "account")
		delete(additionalProperties, "inventorySite")
		delete(additionalProperties, "transactionDate")
		delete(additionalProperties, "refNumber")
		delete(additionalProperties, "customer")
		delete(additionalProperties, "class")
		delete(additionalProperties, "memo")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "lines")
		delete(additionalProperties, "customFields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdInventoryAdjustment struct {
	value *QbdInventoryAdjustment
	isSet bool
}

func (v NullableQbdInventoryAdjustment) Get() *QbdInventoryAdjustment {
	return v.value
}

func (v *NullableQbdInventoryAdjustment) Set(val *QbdInventoryAdjustment) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdInventoryAdjustment) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdInventoryAdjustment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdInventoryAdjustment(val *QbdInventoryAdjustment) *NullableQbdInventoryAdjustment {
	return &NullableQbdInventoryAdjustment{value: val, isSet: true}
}

func (v NullableQbdInventoryAdjustment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdInventoryAdjustment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


