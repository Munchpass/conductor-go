# QuickbooksDesktopEstimatesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer or customer-job associated with this estimate. | 
**ClassId** | Pointer to **string** | The estimate&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this estimate&#39;s line items unless overridden at the line item level. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this estimate when printed or displayed. | [optional] 
**TransactionDate** | **string** | The date of this estimate, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this estimate, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopEstimatesPostRequestBillingAddress**](QuickbooksDesktopEstimatesPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopEstimatesPostRequestShippingAddress**](QuickbooksDesktopEstimatesPostRequestShippingAddress.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this estimate is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**PurchaseOrderNumber** | Pointer to **string** | The customer&#39;s Purchase Order (PO) number associated with this estimate. This field is often used to cross-reference the estimate with the customer&#39;s purchasing system. | [optional] 
**TermsId** | Pointer to **string** | The estimate&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this estimate must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**SalesRepresentativeId** | Pointer to **string** | The estimate&#39;s sales representative. Sales representatives can be employees, vendors, or other names in QuickBooks. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this estimate is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**SalesTaxItemId** | Pointer to **string** | The sales-tax item used to calculate the actual tax amount for this estimate&#39;s transactions by applying a specific tax rate collected for a single tax agency. Unlike &#x60;salesTaxCode&#x60;, which only indicates general taxability, this field drives the actual tax calculation and reporting. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this estimate that appears in reports, but not on the estimate. Use &#x60;customerMessage&#x60; to add a note to this estimate. | [optional] 
**CustomerMessageId** | Pointer to **string** | The message to display to the customer on the estimate. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this estimate is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this estimate, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField** | Pointer to **string** | A built-in custom field for additional information specific to this estimate. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimates for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this estimate&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopEstimatesPostRequestLinesInner**](QuickbooksDesktopEstimatesPostRequestLinesInner.md) | The estimate&#39;s line items, each representing a single product or service quoted.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating an estimate. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopEstimatesPostRequestLineGroupsInner**](QuickbooksDesktopEstimatesPostRequestLineGroupsInner.md) | The estimate&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating an estimate. | [optional] 

## Methods

### NewQuickbooksDesktopEstimatesPostRequest

`func NewQuickbooksDesktopEstimatesPostRequest(customerId string, transactionDate string, ) *QuickbooksDesktopEstimatesPostRequest`

NewQuickbooksDesktopEstimatesPostRequest instantiates a new QuickbooksDesktopEstimatesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEstimatesPostRequestWithDefaults

`func NewQuickbooksDesktopEstimatesPostRequestWithDefaults() *QuickbooksDesktopEstimatesPostRequest`

NewQuickbooksDesktopEstimatesPostRequestWithDefaults instantiates a new QuickbooksDesktopEstimatesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetClassId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopEstimatesPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopEstimatesPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) GetBillingAddress() QuickbooksDesktopEstimatesPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetBillingAddressOk() (*QuickbooksDesktopEstimatesPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) SetBillingAddress(v QuickbooksDesktopEstimatesPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) GetShippingAddress() QuickbooksDesktopEstimatesPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetShippingAddressOk() (*QuickbooksDesktopEstimatesPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) SetShippingAddress(v QuickbooksDesktopEstimatesPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopEstimatesPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopEstimatesPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopEstimatesPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopEstimatesPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesPostRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopEstimatesPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopEstimatesPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopEstimatesPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopEstimatesPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopEstimatesPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopEstimatesPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopEstimatesPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopEstimatesPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopEstimatesPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopEstimatesPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopEstimatesPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopEstimatesPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopEstimatesPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopEstimatesPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopEstimatesPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopEstimatesPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopEstimatesPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopEstimatesPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopEstimatesPostRequest) GetLines() []QuickbooksDesktopEstimatesPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetLinesOk() (*[]QuickbooksDesktopEstimatesPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopEstimatesPostRequest) SetLines(v []QuickbooksDesktopEstimatesPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopEstimatesPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopEstimatesPostRequest) GetLineGroups() []QuickbooksDesktopEstimatesPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopEstimatesPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopEstimatesPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopEstimatesPostRequest) SetLineGroups(v []QuickbooksDesktopEstimatesPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopEstimatesPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


