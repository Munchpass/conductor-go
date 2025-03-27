# QuickbooksDesktopCreditMemosIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the credit memo object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**CustomerId** | Pointer to **string** | The customer or customer-job associated with this credit memo. | [optional] 
**ClassId** | Pointer to **string** | The credit memo&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this credit memo&#39;s line items unless overridden at the line item level. | [optional] 
**ReceivablesAccountId** | Pointer to **string** | The Accounts-Receivable (A/R) account to which this credit memo is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/R account.  **IMPORTANT**: If this credit memo is linked to other transactions, this A/R account must match the &#x60;receivablesAccount&#x60; used in all linked transactions. For example, when refunding a credit card payment, the A/R account must match the one used in the original credit transactions being refunded. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this credit memo when printed or displayed. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this credit memo, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this credit memo, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopCreditMemosPostRequestBillingAddress**](QuickbooksDesktopCreditMemosPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopCreditMemosPostRequestShippingAddress**](QuickbooksDesktopCreditMemosPostRequestShippingAddress.md) |  | [optional] 
**IsPending** | Pointer to **bool** | Indicates whether this credit memo has not been completed. | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | The customer&#39;s Purchase Order (PO) number associated with this credit memo. This field is often used to cross-reference the credit memo with the customer&#39;s purchasing system. | [optional] 
**TermsId** | Pointer to **string** | The credit memo&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this credit memo must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The credit memo&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this credit memo is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**ShippingDate** | Pointer to **string** | The date when the products or services for this credit memo were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this credit memo, such as standard mail or overnight delivery. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this credit memo&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this credit memo that appears in the account register and customer register, but not on the credit memo itself. | [optional] 
**CustomerMessageId** | Pointer to **string** | The message to display to the customer on the credit memo. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this credit memo is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this credit memo is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this credit memo, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this credit memo. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all credit memos for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this credit memo&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopCreditMemosIdPostRequestLinesInner**](QuickbooksDesktopCreditMemosIdPostRequestLinesInner.md) | The credit memo&#39;s line items, each representing a single product or service sold.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the credit memo with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopCreditMemosIdPostRequestLineGroupsInner**](QuickbooksDesktopCreditMemosIdPostRequestLineGroupsInner.md) | The credit memo&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line item groups for the credit memo with this array. To keep any existing line item groups, you must include them in this array even if they have not changed. **Any line item groups not included will be removed.**  2. To add a new line item group, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line item groups, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopCreditMemosIdPostRequest

`func NewQuickbooksDesktopCreditMemosIdPostRequest(revisionNumber string, ) *QuickbooksDesktopCreditMemosIdPostRequest`

NewQuickbooksDesktopCreditMemosIdPostRequest instantiates a new QuickbooksDesktopCreditMemosIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopCreditMemosIdPostRequestWithDefaults

`func NewQuickbooksDesktopCreditMemosIdPostRequestWithDefaults() *QuickbooksDesktopCreditMemosIdPostRequest`

NewQuickbooksDesktopCreditMemosIdPostRequestWithDefaults instantiates a new QuickbooksDesktopCreditMemosIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetReceivablesAccountId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetReceivablesAccountId() string`

GetReceivablesAccountId returns the ReceivablesAccountId field if non-nil, zero value otherwise.

### GetReceivablesAccountIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetReceivablesAccountIdOk() (*string, bool)`

GetReceivablesAccountIdOk returns a tuple with the ReceivablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccountId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetReceivablesAccountId(v string)`

SetReceivablesAccountId sets ReceivablesAccountId field to given value.

### HasReceivablesAccountId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasReceivablesAccountId() bool`

HasReceivablesAccountId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetBillingAddress() QuickbooksDesktopCreditMemosPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetBillingAddressOk() (*QuickbooksDesktopCreditMemosPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetBillingAddress(v QuickbooksDesktopCreditMemosPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingAddress() QuickbooksDesktopCreditMemosPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopCreditMemosPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetShippingAddress(v QuickbooksDesktopCreditMemosPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsPending

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.

### HasIsPending

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasIsPending() bool`

HasIsPending returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetShippingDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetLines() []QuickbooksDesktopCreditMemosIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopCreditMemosIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetLines(v []QuickbooksDesktopCreditMemosIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetLineGroups() []QuickbooksDesktopCreditMemosIdPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopCreditMemosIdPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) SetLineGroups(v []QuickbooksDesktopCreditMemosIdPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopCreditMemosIdPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


