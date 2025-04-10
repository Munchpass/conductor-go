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

// checks if the QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction{}

// QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction The credit card transaction data for this receive-payment's payment when using QuickBooks Merchant Services (QBMS). If specifying this field, you must also specify the `paymentMethod` field.
type QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction struct {
	Request *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest `json:"request,omitempty"`
	Response *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse `json:"response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction

// NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction instantiates a new QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction() *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction {
	this := QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction{}
	return &this
}

// NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionWithDefaults() *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction {
	this := QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction{}
	return &this
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) GetRequest() QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest {
	if o == nil || IsNil(o.Request) {
		var ret QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) GetRequestOk() (*QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest and assigns it to the Request field.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) SetRequest(v QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionRequest) {
	o.Request = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) GetResponse() QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse {
	if o == nil || IsNil(o.Response) {
		var ret QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) GetResponseOk() (*QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse and assigns it to the Response field.
func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) SetResponse(v QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) {
	o.Response = &v
}

func (o QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) UnmarshalJSON(data []byte) (err error) {
	varQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction := _QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction{}

	err = json.Unmarshal(data, &varQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction)

	if err != nil {
		return err
	}

	*o = QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction(varQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "request")
		delete(additionalProperties, "response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction struct {
	value *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction
	isSet bool
}

func (v NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) Get() *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction {
	return v.value
}

func (v *NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) Set(val *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction(val *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) *NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction {
	return &NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction{value: val, isSet: true}
}

func (v NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


