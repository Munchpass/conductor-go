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

// checks if the QbdSalesRepresentative type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QbdSalesRepresentative{}

// QbdSalesRepresentative struct for QbdSalesRepresentative
type QbdSalesRepresentative struct {
	// The unique identifier assigned by QuickBooks to this sales representative. This ID is unique across all sales representatives but not across different QuickBooks object types.
	Id string `json:"id"`
	// The type of object. This value is always `\"qbd_sales_representative\"`.
	ObjectType string `json:"objectType"`
	// The date and time when this sales representative was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	CreatedAt string `json:"createdAt"`
	// The date and time when this sales representative was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user's time zone in QuickBooks.
	UpdatedAt string `json:"updatedAt"`
	// The current QuickBooks-assigned revision number of this sales representative object, which changes each time the object is modified. When updating this object, you must provide the most recent `revisionNumber` to ensure you're working with the latest data; otherwise, the update will return an error.
	RevisionNumber string `json:"revisionNumber"`
	// The initials of this sales representative's name.
	Initial string `json:"initial"`
	// Indicates whether this sales representative is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to `true`.
	IsActive bool `json:"isActive"`
	Entity QbdSalesRepresentativeEntity `json:"entity"`
	AdditionalProperties map[string]interface{}
}

type _QbdSalesRepresentative QbdSalesRepresentative

// NewQbdSalesRepresentative instantiates a new QbdSalesRepresentative object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQbdSalesRepresentative(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, initial string, isActive bool, entity QbdSalesRepresentativeEntity) *QbdSalesRepresentative {
	this := QbdSalesRepresentative{}
	this.Id = id
	this.ObjectType = objectType
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.RevisionNumber = revisionNumber
	this.Initial = initial
	this.IsActive = isActive
	this.Entity = entity
	return &this
}

// NewQbdSalesRepresentativeWithDefaults instantiates a new QbdSalesRepresentative object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQbdSalesRepresentativeWithDefaults() *QbdSalesRepresentative {
	this := QbdSalesRepresentative{}
	return &this
}

// GetId returns the Id field value
func (o *QbdSalesRepresentative) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *QbdSalesRepresentative) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *QbdSalesRepresentative) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *QbdSalesRepresentative) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *QbdSalesRepresentative) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *QbdSalesRepresentative) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *QbdSalesRepresentative) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *QbdSalesRepresentative) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetRevisionNumber returns the RevisionNumber field value
func (o *QbdSalesRepresentative) GetRevisionNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RevisionNumber
}

// GetRevisionNumberOk returns a tuple with the RevisionNumber field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetRevisionNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RevisionNumber, true
}

// SetRevisionNumber sets field value
func (o *QbdSalesRepresentative) SetRevisionNumber(v string) {
	o.RevisionNumber = v
}

// GetInitial returns the Initial field value
func (o *QbdSalesRepresentative) GetInitial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Initial
}

// GetInitialOk returns a tuple with the Initial field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetInitialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Initial, true
}

// SetInitial sets field value
func (o *QbdSalesRepresentative) SetInitial(v string) {
	o.Initial = v
}

// GetIsActive returns the IsActive field value
func (o *QbdSalesRepresentative) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *QbdSalesRepresentative) SetIsActive(v bool) {
	o.IsActive = v
}

// GetEntity returns the Entity field value
func (o *QbdSalesRepresentative) GetEntity() QbdSalesRepresentativeEntity {
	if o == nil {
		var ret QbdSalesRepresentativeEntity
		return ret
	}

	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value
// and a boolean to check if the value has been set.
func (o *QbdSalesRepresentative) GetEntityOk() (*QbdSalesRepresentativeEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Entity, true
}

// SetEntity sets field value
func (o *QbdSalesRepresentative) SetEntity(v QbdSalesRepresentativeEntity) {
	o.Entity = v
}

func (o QbdSalesRepresentative) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QbdSalesRepresentative) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["objectType"] = o.ObjectType
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["revisionNumber"] = o.RevisionNumber
	toSerialize["initial"] = o.Initial
	toSerialize["isActive"] = o.IsActive
	toSerialize["entity"] = o.Entity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QbdSalesRepresentative) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"objectType",
		"createdAt",
		"updatedAt",
		"revisionNumber",
		"initial",
		"isActive",
		"entity",
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

	varQbdSalesRepresentative := _QbdSalesRepresentative{}

	err = json.Unmarshal(data, &varQbdSalesRepresentative)

	if err != nil {
		return err
	}

	*o = QbdSalesRepresentative(varQbdSalesRepresentative)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "objectType")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "revisionNumber")
		delete(additionalProperties, "initial")
		delete(additionalProperties, "isActive")
		delete(additionalProperties, "entity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQbdSalesRepresentative struct {
	value *QbdSalesRepresentative
	isSet bool
}

func (v NullableQbdSalesRepresentative) Get() *QbdSalesRepresentative {
	return v.value
}

func (v *NullableQbdSalesRepresentative) Set(val *QbdSalesRepresentative) {
	v.value = val
	v.isSet = true
}

func (v NullableQbdSalesRepresentative) IsSet() bool {
	return v.isSet
}

func (v *NullableQbdSalesRepresentative) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQbdSalesRepresentative(val *QbdSalesRepresentative) *NullableQbdSalesRepresentative {
	return &NullableQbdSalesRepresentative{value: val, isSet: true}
}

func (v NullableQbdSalesRepresentative) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQbdSalesRepresentative) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


