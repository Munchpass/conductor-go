# QuickbooksDesktopInvoicesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer or customer-job associated with this invoice. | 
**ClassId** | Pointer to **string** | The invoice&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this invoice&#39;s line items unless overridden at the line item level. | [optional] 
**ReceivablesAccountId** | Pointer to **string** | The Accounts-Receivable (A/R) account to which this invoice is assigned, used to track the amount owed. If not specified, QuickBooks Desktop will use its default A/R account.  **IMPORTANT**: If this invoice is linked to other transactions, this A/R account must match the &#x60;receivablesAccount&#x60; used in all linked transactions. For example, when refunding a credit card payment, the A/R account must match the one used in the original credit transactions being refunded. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this invoice when printed or displayed. | [optional] 
**TransactionDate** | **string** | The date of this invoice, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this invoice, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopInvoicesPostRequestBillingAddress**](QuickbooksDesktopInvoicesPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopInvoicesPostRequestShippingAddress**](QuickbooksDesktopInvoicesPostRequestShippingAddress.md) |  | [optional] 
**IsPending** | Pointer to **bool** | Indicates whether this invoice has not been completed or is in a draft version. | [optional] 
**IsFinanceCharge** | Pointer to **bool** | Whether this invoice includes a finance charge. This field is immutable and can only be set during invoice creation. | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | The customer&#39;s Purchase Order (PO) number associated with this invoice. This field is often used to cross-reference the invoice with the customer&#39;s purchasing system. | [optional] 
**TermsId** | Pointer to **string** | The invoice&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this invoice must be paid, in ISO 8601 format (YYYY-MM-DD).  **NOTE**: If &#x60;dueDate&#x60; is excluded when creating this invoice, QuickBooks might determine the due date according to the terms set for this customer. | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The invoice&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this invoice is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**ShippingDate** | Pointer to **string** | The date when the products or services for this invoice were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this invoice, such as standard mail or overnight delivery. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this invoice&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting.  For invoices, while using this field to specify a single tax item/group that applies uniformly is recommended, complex tax scenarios may require alternative approaches. In such cases, you can set this field to a 0% tax item (conventionally named \&quot;Tax Calculated On Invoice\&quot;) and handle tax calculations through line items instead. When using line items for taxes, note that only individual tax items (not tax groups) can be used, subtotals can help apply a tax to multiple items but only the first tax line after a subtotal is calculated automatically (subsequent tax lines require manual amounts), and the rate column will always display the actual tax amount rather than the rate percentage. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this invoice that appears in reports, but not on the invoice. Use &#x60;customerMessage&#x60; to add a note to this invoice. | [optional] 
**CustomerMessageId** | Pointer to **string** | The message to display to the customer on the invoice. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this invoice is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this invoice is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this invoice, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this invoice. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all invoices for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this invoice&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**LinkToTransactionIds** | Pointer to **[]string** | IDs of existing transactions that you wish to link to this invoice, such as payments applied, credits used, or associated purchase orders. Note that this links entire transactions, not individual transaction lines. If you want to link individual lines in a transaction, instead use the field &#x60;linkToTransactionLine&#x60; on this invoice&#39;s lines, if available.  Transactions can only be linked when creating this invoice and cannot be unlinked later.  You can use both &#x60;linkToTransactionIds&#x60; (on this invoice) and &#x60;linkToTransactionLine&#x60; (on its transaction lines) as long as they do NOT link to the same transaction (otherwise, QuickBooks will return an error). QuickBooks will also return an error if you attempt to link a transaction that is empty or already closed.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the invoice and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of invoices, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**ApplyCredits** | Pointer to [**[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner**](QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner.md) | Credit memos to apply to this invoice, reducing its balance. This creates a link between this invoice and the specified credit memos.  **IMPORTANT**: By default, QuickBooks will not return any information about the linked transactions in this endpoint&#39;s response even when this request is successful. To see the transactions linked via this field, refetch the invoice and check the &#x60;linkedTransactions&#x60; response field. If fetching a list of invoices, you must also specify the parameter &#x60;includeLinkedTransactions&#x3D;true&#x60; to see the &#x60;linkedTransactions&#x60; response field. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopInvoicesPostRequestLinesInner**](QuickbooksDesktopInvoicesPostRequestLinesInner.md) | The invoice&#39;s line items, each representing a single product or service sold. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopInvoicesPostRequestLineGroupsInner**](QuickbooksDesktopInvoicesPostRequestLineGroupsInner.md) | The invoice&#39;s line item groups, each representing a predefined set of related items. | [optional] 

## Methods

### NewQuickbooksDesktopInvoicesPostRequest

`func NewQuickbooksDesktopInvoicesPostRequest(customerId string, transactionDate string, ) *QuickbooksDesktopInvoicesPostRequest`

NewQuickbooksDesktopInvoicesPostRequest instantiates a new QuickbooksDesktopInvoicesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopInvoicesPostRequestWithDefaults

`func NewQuickbooksDesktopInvoicesPostRequestWithDefaults() *QuickbooksDesktopInvoicesPostRequest`

NewQuickbooksDesktopInvoicesPostRequestWithDefaults instantiates a new QuickbooksDesktopInvoicesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClassId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetReceivablesAccountId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetReceivablesAccountId() string`

GetReceivablesAccountId returns the ReceivablesAccountId field if non-nil, zero value otherwise.

### GetReceivablesAccountIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetReceivablesAccountIdOk() (*string, bool)`

GetReceivablesAccountIdOk returns a tuple with the ReceivablesAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccountId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetReceivablesAccountId(v string)`

SetReceivablesAccountId sets ReceivablesAccountId field to given value.

### HasReceivablesAccountId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasReceivablesAccountId() bool`

HasReceivablesAccountId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopInvoicesPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopInvoicesPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) GetBillingAddress() QuickbooksDesktopInvoicesPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetBillingAddressOk() (*QuickbooksDesktopInvoicesPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) SetBillingAddress(v QuickbooksDesktopInvoicesPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingAddress() QuickbooksDesktopInvoicesPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingAddressOk() (*QuickbooksDesktopInvoicesPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) SetShippingAddress(v QuickbooksDesktopInvoicesPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopInvoicesPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsPending

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QuickbooksDesktopInvoicesPostRequest) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.

### HasIsPending

`func (o *QuickbooksDesktopInvoicesPostRequest) HasIsPending() bool`

HasIsPending returns a boolean if a field has been set.

### GetIsFinanceCharge

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsFinanceCharge() bool`

GetIsFinanceCharge returns the IsFinanceCharge field if non-nil, zero value otherwise.

### GetIsFinanceChargeOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsFinanceChargeOk() (*bool, bool)`

GetIsFinanceChargeOk returns a tuple with the IsFinanceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinanceCharge

`func (o *QuickbooksDesktopInvoicesPostRequest) SetIsFinanceCharge(v bool)`

SetIsFinanceCharge sets IsFinanceCharge field to given value.

### HasIsFinanceCharge

`func (o *QuickbooksDesktopInvoicesPostRequest) HasIsFinanceCharge() bool`

HasIsFinanceCharge returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *QuickbooksDesktopInvoicesPostRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopInvoicesPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopInvoicesPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopInvoicesPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopInvoicesPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopInvoicesPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetShippingDate

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QuickbooksDesktopInvoicesPostRequest) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *QuickbooksDesktopInvoicesPostRequest) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopInvoicesPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopInvoicesPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopInvoicesPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopInvoicesPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopInvoicesPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopInvoicesPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopInvoicesPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopInvoicesPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopInvoicesPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopInvoicesPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopInvoicesPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopInvoicesPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopInvoicesPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopInvoicesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopInvoicesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopInvoicesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLinkToTransactionIds

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLinkToTransactionIds() []string`

GetLinkToTransactionIds returns the LinkToTransactionIds field if non-nil, zero value otherwise.

### GetLinkToTransactionIdsOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLinkToTransactionIdsOk() (*[]string, bool)`

GetLinkToTransactionIdsOk returns a tuple with the LinkToTransactionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToTransactionIds

`func (o *QuickbooksDesktopInvoicesPostRequest) SetLinkToTransactionIds(v []string)`

SetLinkToTransactionIds sets LinkToTransactionIds field to given value.

### HasLinkToTransactionIds

`func (o *QuickbooksDesktopInvoicesPostRequest) HasLinkToTransactionIds() bool`

HasLinkToTransactionIds returns a boolean if a field has been set.

### GetApplyCredits

`func (o *QuickbooksDesktopInvoicesPostRequest) GetApplyCredits() []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner`

GetApplyCredits returns the ApplyCredits field if non-nil, zero value otherwise.

### GetApplyCreditsOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetApplyCreditsOk() (*[]QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner, bool)`

GetApplyCreditsOk returns a tuple with the ApplyCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyCredits

`func (o *QuickbooksDesktopInvoicesPostRequest) SetApplyCredits(v []QuickbooksDesktopBillCheckPaymentsPostRequestApplyToTransactionsInnerApplyCreditsInner)`

SetApplyCredits sets ApplyCredits field to given value.

### HasApplyCredits

`func (o *QuickbooksDesktopInvoicesPostRequest) HasApplyCredits() bool`

HasApplyCredits returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLines() []QuickbooksDesktopInvoicesPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLinesOk() (*[]QuickbooksDesktopInvoicesPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopInvoicesPostRequest) SetLines(v []QuickbooksDesktopInvoicesPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopInvoicesPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLineGroups() []QuickbooksDesktopInvoicesPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopInvoicesPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopInvoicesPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopInvoicesPostRequest) SetLineGroups(v []QuickbooksDesktopInvoicesPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopInvoicesPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


