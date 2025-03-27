# QuickbooksDesktopDiscountItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the discount item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this discount item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two discount items could both have the &#x60;name&#x60; \&quot;10% labor discount\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Discounts:10% labor discount\&quot; and \&quot;Promotions:10% labor discount\&quot;.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopDiscountItemsPostRequestBarcode**](QuickbooksDesktopDiscountItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this discount item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The discount item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent discount item one level above this one in the hierarchy. For example, if this discount item has a &#x60;fullName&#x60; of \&quot;Discounts:10% labor discount\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Discounts\&quot;. If this discount item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**Description** | Pointer to **string** | The discount item&#39;s description that will appear on sales forms that include this item. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this discount item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**DiscountRate** | Pointer to **string** | The monetary amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A flat rate discount applies to ALL lines recorded above it and distributes the discount amount equally across those lines, which affects tax calculations. For example, a $10 discount applied to a $100 taxable item and $100 non-taxable item would result in a $5 taxable discount and $5 non-taxable discount. | [optional] 
**DiscountRatePercent** | Pointer to **string** | The percentage amount to subtract from the total or subtotal when applying this discount item to a transaction.  **NOTE**: A percentage discount only applies to the line immediately above it, so tax implications only affect that specific line. | [optional] 
**AccountId** | Pointer to **string** | The posting account to which transactions involving this discount item are posted for tracking discounts. | [optional] 
**UpdateExistingTransactionsAccount** | Pointer to **bool** | When &#x60;true&#x60;, applies the new account (specified by the &#x60;accountId&#x60; field) to all existing transactions associated with this discount item. This updates historical data and should be used with caution. The update will fail if any affected transaction falls within a closed accounting period. If this parameter is not specified, QuickBooks will prompt the user before making any changes. | [optional] 

## Methods

### NewQuickbooksDesktopDiscountItemsIdPostRequest

`func NewQuickbooksDesktopDiscountItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopDiscountItemsIdPostRequest`

NewQuickbooksDesktopDiscountItemsIdPostRequest instantiates a new QuickbooksDesktopDiscountItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopDiscountItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopDiscountItemsIdPostRequestWithDefaults() *QuickbooksDesktopDiscountItemsIdPostRequest`

NewQuickbooksDesktopDiscountItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopDiscountItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetBarcode() QuickbooksDesktopDiscountItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopDiscountItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetBarcode(v QuickbooksDesktopDiscountItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetDescription

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetDiscountRate

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRate() string`

GetDiscountRate returns the DiscountRate field if non-nil, zero value otherwise.

### GetDiscountRateOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRateOk() (*string, bool)`

GetDiscountRateOk returns a tuple with the DiscountRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRate

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDiscountRate(v string)`

SetDiscountRate sets DiscountRate field to given value.

### HasDiscountRate

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDiscountRate() bool`

HasDiscountRate returns a boolean if a field has been set.

### GetDiscountRatePercent

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRatePercent() string`

GetDiscountRatePercent returns the DiscountRatePercent field if non-nil, zero value otherwise.

### GetDiscountRatePercentOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetDiscountRatePercentOk() (*string, bool)`

GetDiscountRatePercentOk returns a tuple with the DiscountRatePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRatePercent

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetDiscountRatePercent(v string)`

SetDiscountRatePercent sets DiscountRatePercent field to given value.

### HasDiscountRatePercent

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasDiscountRatePercent() bool`

HasDiscountRatePercent returns a boolean if a field has been set.

### GetAccountId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetUpdateExistingTransactionsAccount() bool`

GetUpdateExistingTransactionsAccount returns the UpdateExistingTransactionsAccount field if non-nil, zero value otherwise.

### GetUpdateExistingTransactionsAccountOk

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) GetUpdateExistingTransactionsAccountOk() (*bool, bool)`

GetUpdateExistingTransactionsAccountOk returns a tuple with the UpdateExistingTransactionsAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) SetUpdateExistingTransactionsAccount(v bool)`

SetUpdateExistingTransactionsAccount sets UpdateExistingTransactionsAccount field to given value.

### HasUpdateExistingTransactionsAccount

`func (o *QuickbooksDesktopDiscountItemsIdPostRequest) HasUpdateExistingTransactionsAccount() bool`

HasUpdateExistingTransactionsAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


