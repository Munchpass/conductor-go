# QbdAccountantCopy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountantCopyExists** | **bool** | Indicates whether an accountant copy has been made for this company file. An accountant copy allows an accountant to work on the books while the business continues daily operations. | 
**DividingDate** | **NullableString** | The fiscal period dividing date for accountant work, in ISO 8601 format (YYYY-MM-DD). While an accountant copy exists, transactions within this period cannot be modified or created. New accounts can be added, but existing accounts cannot have new subaccounts, be edited, merged, or made inactive. List items cannot be deleted or merged. | 

## Methods

### NewQbdAccountantCopy

`func NewQbdAccountantCopy(accountantCopyExists bool, dividingDate NullableString, ) *QbdAccountantCopy`

NewQbdAccountantCopy instantiates a new QbdAccountantCopy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdAccountantCopyWithDefaults

`func NewQbdAccountantCopyWithDefaults() *QbdAccountantCopy`

NewQbdAccountantCopyWithDefaults instantiates a new QbdAccountantCopy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountantCopyExists

`func (o *QbdAccountantCopy) GetAccountantCopyExists() bool`

GetAccountantCopyExists returns the AccountantCopyExists field if non-nil, zero value otherwise.

### GetAccountantCopyExistsOk

`func (o *QbdAccountantCopy) GetAccountantCopyExistsOk() (*bool, bool)`

GetAccountantCopyExistsOk returns a tuple with the AccountantCopyExists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountantCopyExists

`func (o *QbdAccountantCopy) SetAccountantCopyExists(v bool)`

SetAccountantCopyExists sets AccountantCopyExists field to given value.


### GetDividingDate

`func (o *QbdAccountantCopy) GetDividingDate() string`

GetDividingDate returns the DividingDate field if non-nil, zero value otherwise.

### GetDividingDateOk

`func (o *QbdAccountantCopy) GetDividingDateOk() (*string, bool)`

GetDividingDateOk returns a tuple with the DividingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividingDate

`func (o *QbdAccountantCopy) SetDividingDate(v string)`

SetDividingDate sets DividingDate field to given value.


### SetDividingDateNil

`func (o *QbdAccountantCopy) SetDividingDateNil(b bool)`

 SetDividingDateNil sets the value for DividingDate to be an explicit nil

### UnsetDividingDate
`func (o *QbdAccountantCopy) UnsetDividingDate()`

UnsetDividingDate ensures that no value is present for DividingDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


