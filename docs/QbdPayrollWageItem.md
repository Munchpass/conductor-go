# QbdPayrollWageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this payroll wage item. This ID is unique across all payroll wage items but not across different QuickBooks object types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_payroll_wage_item\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this payroll wage item was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this payroll wage item was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this payroll wage item object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | **string** | The case-insensitive unique name of this payroll wage item, unique across all payroll wage items.  **NOTE**: Payroll wage items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents. | 
**IsActive** | **bool** | Indicates whether this payroll wage item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**WageType** | **string** | Categorizes how this payroll wage item calculates pay - can be hourly (regular, overtime, sick, or vacation), salary (regular, sick, or vacation), bonus, or commission based. | 
**ExpenseAccount** | [**QbdPayrollWageItemExpenseAccount**](QbdPayrollWageItemExpenseAccount.md) |  | 

## Methods

### NewQbdPayrollWageItem

`func NewQbdPayrollWageItem(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, name string, isActive bool, wageType string, expenseAccount QbdPayrollWageItemExpenseAccount, ) *QbdPayrollWageItem`

NewQbdPayrollWageItem instantiates a new QbdPayrollWageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPayrollWageItemWithDefaults

`func NewQbdPayrollWageItemWithDefaults() *QbdPayrollWageItem`

NewQbdPayrollWageItemWithDefaults instantiates a new QbdPayrollWageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdPayrollWageItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdPayrollWageItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdPayrollWageItem) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdPayrollWageItem) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdPayrollWageItem) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdPayrollWageItem) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdPayrollWageItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdPayrollWageItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdPayrollWageItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdPayrollWageItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdPayrollWageItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdPayrollWageItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdPayrollWageItem) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdPayrollWageItem) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdPayrollWageItem) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QbdPayrollWageItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QbdPayrollWageItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QbdPayrollWageItem) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QbdPayrollWageItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdPayrollWageItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdPayrollWageItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetWageType

`func (o *QbdPayrollWageItem) GetWageType() string`

GetWageType returns the WageType field if non-nil, zero value otherwise.

### GetWageTypeOk

`func (o *QbdPayrollWageItem) GetWageTypeOk() (*string, bool)`

GetWageTypeOk returns a tuple with the WageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWageType

`func (o *QbdPayrollWageItem) SetWageType(v string)`

SetWageType sets WageType field to given value.


### GetExpenseAccount

`func (o *QbdPayrollWageItem) GetExpenseAccount() QbdPayrollWageItemExpenseAccount`

GetExpenseAccount returns the ExpenseAccount field if non-nil, zero value otherwise.

### GetExpenseAccountOk

`func (o *QbdPayrollWageItem) GetExpenseAccountOk() (*QbdPayrollWageItemExpenseAccount, bool)`

GetExpenseAccountOk returns a tuple with the ExpenseAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAccount

`func (o *QbdPayrollWageItem) SetExpenseAccount(v QbdPayrollWageItemExpenseAccount)`

SetExpenseAccount sets ExpenseAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


