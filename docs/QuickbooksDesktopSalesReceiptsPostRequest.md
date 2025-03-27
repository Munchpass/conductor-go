# QuickbooksDesktopSalesReceiptsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer or customer-job to which the payment for this sales receipt is credited. | 
**ClassId** | Pointer to **string** | The sales receipt&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this sales receipt&#39;s line items unless overridden at the line item level. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this sales receipt when printed or displayed. | [optional] 
**TransactionDate** | **string** | The date of this sales receipt, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this sales receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestBillingAddress**](QuickbooksDesktopSalesReceiptsPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestShippingAddress**](QuickbooksDesktopSalesReceiptsPostRequestShippingAddress.md) |  | [optional] 
**IsPending** | Pointer to **bool** | Indicates whether this sales receipt has not been completed. | [optional] 
**CheckNumber** | Pointer to **string** | The check number of a check received for this sales receipt. | [optional] 
**PaymentMethodId** | Pointer to **string** | The sales receipt&#39;s payment method (e.g., cash, check, credit card).  **NOTE**: If this sales receipt contains credit card transaction data supplied from QuickBooks Merchant Services (QBMS) transaction responses, you must specify a credit card payment method (e.g., \&quot;Visa\&quot;, \&quot;MasterCard\&quot;, etc.). | [optional] 
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
**CreditCardTransaction** | Pointer to [**QuickbooksDesktopSalesReceiptsPostRequestCreditCardTransaction**](QuickbooksDesktopSalesReceiptsPostRequestCreditCardTransaction.md) |  | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this sales receipt. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipts for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this sales receipt&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopSalesReceiptsPostRequestLinesInner**](QuickbooksDesktopSalesReceiptsPostRequestLinesInner.md) | The sales receipt&#39;s line items, each representing a single product or service sold.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a sales receipt. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopSalesReceiptsPostRequestLineGroupsInner**](QuickbooksDesktopSalesReceiptsPostRequestLineGroupsInner.md) | The sales receipt&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a sales receipt. | [optional] 

## Methods

### NewQuickbooksDesktopSalesReceiptsPostRequest

`func NewQuickbooksDesktopSalesReceiptsPostRequest(customerId string, transactionDate string, ) *QuickbooksDesktopSalesReceiptsPostRequest`

NewQuickbooksDesktopSalesReceiptsPostRequest instantiates a new QuickbooksDesktopSalesReceiptsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesReceiptsPostRequestWithDefaults

`func NewQuickbooksDesktopSalesReceiptsPostRequestWithDefaults() *QuickbooksDesktopSalesReceiptsPostRequest`

NewQuickbooksDesktopSalesReceiptsPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesReceiptsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetBillingAddress() QuickbooksDesktopSalesReceiptsPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetBillingAddressOk() (*QuickbooksDesktopSalesReceiptsPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetBillingAddress(v QuickbooksDesktopSalesReceiptsPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingAddress() QuickbooksDesktopSalesReceiptsPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingAddressOk() (*QuickbooksDesktopSalesReceiptsPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetShippingAddress(v QuickbooksDesktopSalesReceiptsPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsPending

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.

### HasIsPending

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasIsPending() bool`

HasIsPending returns a boolean if a field has been set.

### GetCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCheckNumber() string`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCheckNumberOk() (*string, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetCheckNumber(v string)`

SetCheckNumber sets CheckNumber field to given value.

### HasCheckNumber

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasCheckNumber() bool`

HasCheckNumber returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShippingDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDepositToAccountId() string`

GetDepositToAccountId returns the DepositToAccountId field if non-nil, zero value otherwise.

### GetDepositToAccountIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetDepositToAccountIdOk() (*string, bool)`

GetDepositToAccountIdOk returns a tuple with the DepositToAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetDepositToAccountId(v string)`

SetDepositToAccountId sets DepositToAccountId field to given value.

### HasDepositToAccountId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasDepositToAccountId() bool`

HasDepositToAccountId returns a boolean if a field has been set.

### GetCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCreditCardTransaction() QuickbooksDesktopSalesReceiptsPostRequestCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetCreditCardTransactionOk() (*QuickbooksDesktopSalesReceiptsPostRequestCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetCreditCardTransaction(v QuickbooksDesktopSalesReceiptsPostRequestCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.

### HasCreditCardTransaction

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasCreditCardTransaction() bool`

HasCreditCardTransaction returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetLines() []QuickbooksDesktopSalesReceiptsPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetLinesOk() (*[]QuickbooksDesktopSalesReceiptsPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetLines(v []QuickbooksDesktopSalesReceiptsPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetLineGroups() []QuickbooksDesktopSalesReceiptsPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopSalesReceiptsPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) SetLineGroups(v []QuickbooksDesktopSalesReceiptsPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopSalesReceiptsPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


