# QbdExpenseLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this expense line. This ID is unique across all transaction line types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_expense_line\&quot;&#x60;. | 
**Account** | [**QbdExpenseLineAccount**](QbdExpenseLineAccount.md) |  | 
**Amount** | **string** | The monetary amount of this expense line, represented as a decimal string. | 
**Memo** | **string** | A memo or note for this expense line. | 
**Payee** | [**QbdExpenseLinePayee**](QbdExpenseLinePayee.md) |  | 
**Class** | [**QbdExpenseLineClass**](QbdExpenseLineClass.md) |  | 
**SalesTaxCode** | [**QbdExpenseLineSalesTaxCode**](QbdExpenseLineSalesTaxCode.md) |  | 
**BillingStatus** | **string** | The billing status of this expense line. | 
**SalesRepresentative** | [**QbdExpenseLineSalesRepresentative**](QbdExpenseLineSalesRepresentative.md) |  | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the expense line object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdExpenseLine

`func NewQbdExpenseLine(id string, objectType string, account QbdExpenseLineAccount, amount string, memo string, payee QbdExpenseLinePayee, class QbdExpenseLineClass, salesTaxCode QbdExpenseLineSalesTaxCode, billingStatus string, salesRepresentative QbdExpenseLineSalesRepresentative, customFields []QbdCustomField, ) *QbdExpenseLine`

NewQbdExpenseLine instantiates a new QbdExpenseLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdExpenseLineWithDefaults

`func NewQbdExpenseLineWithDefaults() *QbdExpenseLine`

NewQbdExpenseLineWithDefaults instantiates a new QbdExpenseLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdExpenseLine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdExpenseLine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdExpenseLine) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdExpenseLine) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdExpenseLine) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdExpenseLine) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetAccount

`func (o *QbdExpenseLine) GetAccount() QbdExpenseLineAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *QbdExpenseLine) GetAccountOk() (*QbdExpenseLineAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *QbdExpenseLine) SetAccount(v QbdExpenseLineAccount)`

SetAccount sets Account field to given value.


### GetAmount

`func (o *QbdExpenseLine) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QbdExpenseLine) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QbdExpenseLine) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMemo

`func (o *QbdExpenseLine) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdExpenseLine) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdExpenseLine) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetPayee

`func (o *QbdExpenseLine) GetPayee() QbdExpenseLinePayee`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *QbdExpenseLine) GetPayeeOk() (*QbdExpenseLinePayee, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *QbdExpenseLine) SetPayee(v QbdExpenseLinePayee)`

SetPayee sets Payee field to given value.


### GetClass

`func (o *QbdExpenseLine) GetClass() QbdExpenseLineClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdExpenseLine) GetClassOk() (*QbdExpenseLineClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdExpenseLine) SetClass(v QbdExpenseLineClass)`

SetClass sets Class field to given value.


### GetSalesTaxCode

`func (o *QbdExpenseLine) GetSalesTaxCode() QbdExpenseLineSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdExpenseLine) GetSalesTaxCodeOk() (*QbdExpenseLineSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdExpenseLine) SetSalesTaxCode(v QbdExpenseLineSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetBillingStatus

`func (o *QbdExpenseLine) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QbdExpenseLine) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QbdExpenseLine) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.


### GetSalesRepresentative

`func (o *QbdExpenseLine) GetSalesRepresentative() QbdExpenseLineSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdExpenseLine) GetSalesRepresentativeOk() (*QbdExpenseLineSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdExpenseLine) SetSalesRepresentative(v QbdExpenseLineSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetCustomFields

`func (o *QbdExpenseLine) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdExpenseLine) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdExpenseLine) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


