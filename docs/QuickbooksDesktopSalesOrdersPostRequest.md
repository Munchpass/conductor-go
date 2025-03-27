# QuickbooksDesktopSalesOrdersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer or customer-job associated with this sales order. | 
**ClassId** | Pointer to **string** | The sales order&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this sales order&#39;s line items unless overridden at the line item level. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this sales order when printed or displayed. | [optional] 
**TransactionDate** | **string** | The date of this sales order, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this sales order, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopSalesOrdersPostRequestBillingAddress**](QuickbooksDesktopSalesOrdersPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopSalesOrdersPostRequestShippingAddress**](QuickbooksDesktopSalesOrdersPostRequestShippingAddress.md) |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** | The customer&#39;s Purchase Order (PO) number associated with this sales order. This field is often used to cross-reference the sales order with the customer&#39;s purchasing system. | [optional] 
**TermsId** | Pointer to **string** | The sales order&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this sales order must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The sales order&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this sales order is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**ShippingDate** | Pointer to **string** | The date when the products or services for this sales order were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this sales order, such as standard mail or overnight delivery. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this sales order&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**IsManuallyClosed** | Pointer to **bool** | Indicates whether this sales order has been manually marked as closed, even if it has not been invoiced. | [optional] [default to false]
**Memo** | Pointer to **string** | A memo or note for this sales order. | [optional] 
**CustomerMessageId** | Pointer to **string** | The message to display to the customer on the sales order. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this sales order is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this sales order is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this sales order, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this sales order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales orders for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this sales order&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopSalesOrdersPostRequestLinesInner**](QuickbooksDesktopSalesOrdersPostRequestLinesInner.md) | The sales order&#39;s line items, each representing a single product or service ordered.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a sales order. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner**](QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner.md) | The sales order&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a sales order. | [optional] 
**SalesChannelName** | Pointer to **string** | The type of the sales channel for this sales order. | [optional] 
**SalesStoreName** | Pointer to **string** | The name of the sales store for this sales order. | [optional] 
**SalesStoreType** | Pointer to **string** | The type of the sales store for this sales order. | [optional] 

## Methods

### NewQuickbooksDesktopSalesOrdersPostRequest

`func NewQuickbooksDesktopSalesOrdersPostRequest(customerId string, transactionDate string, ) *QuickbooksDesktopSalesOrdersPostRequest`

NewQuickbooksDesktopSalesOrdersPostRequest instantiates a new QuickbooksDesktopSalesOrdersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopSalesOrdersPostRequestWithDefaults

`func NewQuickbooksDesktopSalesOrdersPostRequestWithDefaults() *QuickbooksDesktopSalesOrdersPostRequest`

NewQuickbooksDesktopSalesOrdersPostRequestWithDefaults instantiates a new QuickbooksDesktopSalesOrdersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetBillingAddress() QuickbooksDesktopSalesOrdersPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetBillingAddressOk() (*QuickbooksDesktopSalesOrdersPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetBillingAddress(v QuickbooksDesktopSalesOrdersPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingAddress() QuickbooksDesktopSalesOrdersPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingAddressOk() (*QuickbooksDesktopSalesOrdersPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetShippingAddress(v QuickbooksDesktopSalesOrdersPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetShippingDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.

### HasShippingDate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasShippingDate() bool`

HasShippingDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.

### HasIsManuallyClosed

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasIsManuallyClosed() bool`

HasIsManuallyClosed returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetLines() []QuickbooksDesktopSalesOrdersPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetLinesOk() (*[]QuickbooksDesktopSalesOrdersPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetLines(v []QuickbooksDesktopSalesOrdersPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetLineGroups() []QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetLineGroups(v []QuickbooksDesktopSalesOrdersPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.

### GetSalesChannelName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesChannelName() string`

GetSalesChannelName returns the SalesChannelName field if non-nil, zero value otherwise.

### GetSalesChannelNameOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesChannelNameOk() (*string, bool)`

GetSalesChannelNameOk returns a tuple with the SalesChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannelName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesChannelName(v string)`

SetSalesChannelName sets SalesChannelName field to given value.

### HasSalesChannelName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesChannelName() bool`

HasSalesChannelName returns a boolean if a field has been set.

### GetSalesStoreName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesStoreName() string`

GetSalesStoreName returns the SalesStoreName field if non-nil, zero value otherwise.

### GetSalesStoreNameOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesStoreNameOk() (*string, bool)`

GetSalesStoreNameOk returns a tuple with the SalesStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesStoreName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesStoreName(v string)`

SetSalesStoreName sets SalesStoreName field to given value.

### HasSalesStoreName

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesStoreName() bool`

HasSalesStoreName returns a boolean if a field has been set.

### GetSalesStoreType

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesStoreType() string`

GetSalesStoreType returns the SalesStoreType field if non-nil, zero value otherwise.

### GetSalesStoreTypeOk

`func (o *QuickbooksDesktopSalesOrdersPostRequest) GetSalesStoreTypeOk() (*string, bool)`

GetSalesStoreTypeOk returns a tuple with the SalesStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesStoreType

`func (o *QuickbooksDesktopSalesOrdersPostRequest) SetSalesStoreType(v string)`

SetSalesStoreType sets SalesStoreType field to given value.

### HasSalesStoreType

`func (o *QuickbooksDesktopSalesOrdersPostRequest) HasSalesStoreType() bool`

HasSalesStoreType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


