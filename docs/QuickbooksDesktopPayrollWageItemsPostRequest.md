# QuickbooksDesktopPayrollWageItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive unique name of this payroll wage item, unique across all payroll wage items.  **NOTE**: Payroll wage items do not have a &#x60;fullName&#x60; field because they are not hierarchical objects, which is why &#x60;name&#x60; is unique for them but not for objects that have parents.  Maximum length: 31 characters. | 
**IsActive** | Pointer to **bool** | Indicates whether this payroll wage item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**WageType** | **string** | Categorizes how this payroll wage item calculates pay - can be hourly (regular, overtime, sick, or vacation), salary (regular, sick, or vacation), bonus, or commission based. | 
**ExpenseAccountId** | **string** | The expense account used to track wage expenses paid through this payroll wage item. | 

## Methods

### NewQuickbooksDesktopPayrollWageItemsPostRequest

`func NewQuickbooksDesktopPayrollWageItemsPostRequest(name string, wageType string, expenseAccountId string, ) *QuickbooksDesktopPayrollWageItemsPostRequest`

NewQuickbooksDesktopPayrollWageItemsPostRequest instantiates a new QuickbooksDesktopPayrollWageItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPayrollWageItemsPostRequestWithDefaults

`func NewQuickbooksDesktopPayrollWageItemsPostRequestWithDefaults() *QuickbooksDesktopPayrollWageItemsPostRequest`

NewQuickbooksDesktopPayrollWageItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopPayrollWageItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetWageType

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetWageType() string`

GetWageType returns the WageType field if non-nil, zero value otherwise.

### GetWageTypeOk

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetWageTypeOk() (*string, bool)`

GetWageTypeOk returns a tuple with the WageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWageType

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) SetWageType(v string)`

SetWageType sets WageType field to given value.


### GetExpenseAccountId

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetExpenseAccountId() string`

GetExpenseAccountId returns the ExpenseAccountId field if non-nil, zero value otherwise.

### GetExpenseAccountIdOk

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) GetExpenseAccountIdOk() (*string, bool)`

GetExpenseAccountIdOk returns a tuple with the ExpenseAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAccountId

`func (o *QuickbooksDesktopPayrollWageItemsPostRequest) SetExpenseAccountId(v string)`

SetExpenseAccountId sets ExpenseAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


