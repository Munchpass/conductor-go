# QuickbooksDesktopBillsIdPostRequestExpenseLinesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The QuickBooks-assigned unique identifier of an existing expense line you wish to retain or update.  **IMPORTANT**: Set this field to &#x60;-1&#x60; for new expense lines you wish to add. | 
**AccountId** | Pointer to **string** | The expense account being debited (increased) for this expense line. The corresponding account being credited is usually a liability account (e.g., Accounts-Payable) or an asset account (e.g., Cash), depending on the transaction type. | [optional] 
**Amount** | Pointer to **string** | The monetary amount of this expense line, represented as a decimal string. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this expense line. | [optional] 
**PayeeId** | Pointer to **string** | If &#x60;account&#x60; refers to an Accounts-Payable (A/P) account, &#x60;payee&#x60; refers to the expense&#39;s vendor (not the customer). If &#x60;account&#x60; refers to any other type of account, &#x60;payee&#x60; refers to the expense&#39;s customer (not the vendor). | [optional] 
**ClassId** | Pointer to **string** | The expense line&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. If a class is specified for the entire parent transaction, it is automatically applied to all expense lines unless overridden here, at the transaction line level. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this expense line, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the parent transaction or the associated item.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**BillingStatus** | Pointer to **string** | The billing status of this expense line. | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The expense line&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 

## Methods

### NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInner

`func NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInner(id string, ) *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner`

NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInner instantiates a new QuickbooksDesktopBillsIdPostRequestExpenseLinesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInnerWithDefaults

`func NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInnerWithDefaults() *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner`

NewQuickbooksDesktopBillsIdPostRequestExpenseLinesInnerWithDefaults instantiates a new QuickbooksDesktopBillsIdPostRequestExpenseLinesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetId(v string)`

SetId sets Id field to given value.


### GetAccountId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAmount

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetPayeeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetPayeeId() string`

GetPayeeId returns the PayeeId field if non-nil, zero value otherwise.

### GetPayeeIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetPayeeIdOk() (*string, bool)`

GetPayeeIdOk returns a tuple with the PayeeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayeeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetPayeeId(v string)`

SetPayeeId sets PayeeId field to given value.

### HasPayeeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasPayeeId() bool`

HasPayeeId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetBillingStatus

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetBillingStatus() string`

GetBillingStatus returns the BillingStatus field if non-nil, zero value otherwise.

### GetBillingStatusOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetBillingStatusOk() (*string, bool)`

GetBillingStatusOk returns a tuple with the BillingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatus

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetBillingStatus(v string)`

SetBillingStatus sets BillingStatus field to given value.

### HasBillingStatus

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasBillingStatus() bool`

HasBillingStatus returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopBillsIdPostRequestExpenseLinesInner) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


