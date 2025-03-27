# QuickbooksDesktopSalesReceiptsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the sales receipt object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**CustomerId** | Pointer to **string** | The customer or customer-job to which the payment for this sales receipt is credited. | [optional] 
**ClassId** | Pointer to **string** | The sales receipt&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this sales receipt&#39;s line items unless overridden at the line item level. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this sales receipt when printed or displayed. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this sales receipt, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this sales receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestBillingAddress**](QuickbooksDesktopSalesReceiptsPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestShippingAddress**](QuickbooksDesktopSalesReceiptsPostRequestShippingAddress.md) |  | [optional] 
**IsPending** | Pointer to **bool** | Indicates whether this sales receipt has not been completed. | [optional] 
**CheckNumber** | Pointer to **string** | The check number of a check received for this sales receipt. | [optional] 
**PaymentMethodId** | Pointer to **string** | The sales receipt&#39;s payment method (e.g., cash, check, credit card). | [optional] 
**DueDate** | Pointer to **string** | The date by which this sales receipt must be paid, in ISO 8601 format (YYYY-MM-DD).  **NOTE**: For sales receipts, this field is often &#x60;null&#x60; because sales receipts are generally used for point-of-sale payments, where full payment is received at the time of purchase. | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The sales receipt&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**ShippingDate** | Pointer to **string** | The date when the products or services for this sales receipt were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this sales receipt, such as standard mail or overnight delivery. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this sales receipt is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this sales receipt&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting.  For sales receipts, while using this field to specify a single tax item/group that applies uniformly is recommended, complex tax scenarios may require alternative approaches. In such cases, you can set this field to a 0% tax item (conventionally named \&quot;Tax Calculated On Invoice\&quot;) and handle tax calculations through line items instead. When using line items for taxes, note that only individual tax items (not tax groups) can be used, subtotals can help apply a tax to multiple items but only the first tax line after a subtotal is calculated automatically (subsequent tax lines require manual amounts), and the rate column will always display the actual tax amount rather than the rate percentage. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this sales receipt that appears in reports, but not on the sales receipt. | [optional] 
**CustomerMessageId** | Pointer to **string** | The message to display to the customer on the sales receipt. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this sales receipt is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this sales receipt is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this sales receipt, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**DepositToAccountId** | Pointer to **string** | The account where the funds for this sales receipt will be or have been deposited. | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this sales receipt. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipts for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this sales receipt&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopSalesReceiptsIdPostRequestLinesInner**](QuickbooksDesktopSalesReceiptsIdPostRequestLinesInner.md) | The sales receipt&#39;s line items, each representing a single product or service sold.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the sales receipt with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopSalesReceiptsIdPostRequestLineGroupsInner**](QuickbooksDesktopSalesReceiptsIdPostRequestLineGroupsInner.md) | The sales receipt&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line item groups for the sales receipt with this array. To keep any existing line item groups, you must include them in this array even if they have not changed. **Any line item groups not included will be removed.**  2. To add a new line item group, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line item groups, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopSalesReceiptsIdPostRequest

`func NewQuickbooksDesktopSalesReceiptsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopSalesReceiptsIdPostRequest`

NewQuickbooksDesktopSalesReceiptsIdPostRequest instantiates a new QuickbooksDesktopSalesReceiptsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesReceiptsIdPostRequestWithDefaults

`func NewQuickbooksDesktopSalesReceiptsIdPostRequestWithDefaults() *QuickbooksDesktopSalesReceiptsIdPostRequest`

NewQuickbooksDesktopSalesReceiptsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesReceiptsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetBillingAddress() QuickbooksDesktopSalesReceiptsPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetBillingAddressOk() (*QuickbooksDesktopSalesReceiptsPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetBillingAddress(v QuickbooksDesktopSalesReceiptsPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingAddress() QuickbooksDesktopSalesReceiptsPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopSalesReceiptsPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetShippingAddress(v QuickbooksDesktopSalesReceiptsPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsPending

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.

### HasIsPending

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasIsPending() bool`

HasIsPending returns a boolean if a field has been set.

### GetCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCheckNumber() string`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCheckNumberOk() (*string, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetCheckNumber(v string)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShippingDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDepositToAccountId() string`

GetDepositToAccountId returns the DepositToAccountId field if non-nil, zero value otherwise.

### GetDepositToAccountIdOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetDepositToAccountIdOk() (*string, bool)`

GetDepositToAccountIdOk returns a tuple with the DepositToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetDepositToAccountId(v string)`

SetDepositToAccountId sets DepositToAccountId field to given value.

### HasDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasDepositToAccountId() bool`

HasDepositToAccountId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetLines() []QuickbooksDesktopSalesReceiptsIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopSalesReceiptsIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetLines(v []QuickbooksDesktopSalesReceiptsIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetLineGroups() []QuickbooksDesktopSalesReceiptsIdPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopSalesReceiptsIdPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) SetLineGroups(v []QuickbooksDesktopSalesReceiptsIdPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopSalesReceiptsIdPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


