# QuickbooksDesktopEstimatesIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the estimate object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**CustomerId** | Pointer to **string** | The customer or customer-job associated with this estimate. | [optional] 
**ClassId** | Pointer to **string** | The estimate&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this estimate&#39;s line items unless overridden at the line item level. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this estimate when printed or displayed. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this estimate, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this estimate, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**BillingAddress** | Pointer to [**QuickbooksDesktopEstimatesPostRequestBillingAddress**](QuickbooksDesktopEstimatesPostRequestBillingAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopEstimatesPostRequestShippingAddress**](QuickbooksDesktopEstimatesPostRequestShippingAddress.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this estimate is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**CreateChangeOrder** | Pointer to **bool** | When &#x60;true&#x60;, creates a \&quot;change order\&quot; that appears in this estimate&#39;s description field in QuickBooks&#39;s estimate form, specifying exactly what changed in this update request, the dollar amount of each change, and the net dollar change to this estimate. | [optional] 
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
**Lines** | Pointer to [**[]QuickbooksDesktopEstimatesIdPostRequestLinesInner**](QuickbooksDesktopEstimatesIdPostRequestLinesInner.md) | The estimate&#39;s line items, each representing a single product or service quoted.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the estimate with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner**](QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner.md) | The estimate&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line item groups for the estimate with this array. To keep any existing line item groups, you must include them in this array even if they have not changed. **Any line item groups not included will be removed.**  2. To add a new line item group, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line item groups, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopEstimatesIdPostRequest

`func NewQuickbooksDesktopEstimatesIdPostRequest(revisionNumber string, ) *QuickbooksDesktopEstimatesIdPostRequest`

NewQuickbooksDesktopEstimatesIdPostRequest instantiates a new QuickbooksDesktopEstimatesIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopEstimatesIdPostRequestWithDefaults

`func NewQuickbooksDesktopEstimatesIdPostRequestWithDefaults() *QuickbooksDesktopEstimatesIdPostRequest`

NewQuickbooksDesktopEstimatesIdPostRequestWithDefaults instantiates a new QuickbooksDesktopEstimatesIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomerId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetBillingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetBillingAddress() QuickbooksDesktopEstimatesPostRequestBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetBillingAddressOk() (*QuickbooksDesktopEstimatesPostRequestBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetBillingAddress(v QuickbooksDesktopEstimatesPostRequestBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShippingAddress() QuickbooksDesktopEstimatesPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopEstimatesPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetShippingAddress(v QuickbooksDesktopEstimatesPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCreateChangeOrder

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCreateChangeOrder() bool`

GetCreateChangeOrder returns the CreateChangeOrder field if non-nil, zero value otherwise.

### GetCreateChangeOrderOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCreateChangeOrderOk() (*bool, bool)`

GetCreateChangeOrderOk returns a tuple with the CreateChangeOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateChangeOrder

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCreateChangeOrder(v bool)`

SetCreateChangeOrder sets CreateChangeOrder field to given value.

### HasCreateChangeOrder

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCreateChangeOrder() bool`

HasCreateChangeOrder returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesRepresentativeId() string`

GetSalesRepresentativeId returns the SalesRepresentativeId field if non-nil, zero value otherwise.

### GetSalesRepresentativeIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesRepresentativeIdOk() (*string, bool)`

GetSalesRepresentativeIdOk returns a tuple with the SalesRepresentativeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesRepresentativeId(v string)`

SetSalesRepresentativeId sets SalesRepresentativeId field to given value.

### HasSalesRepresentativeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesRepresentativeId() bool`

HasSalesRepresentativeId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxItemId() string`

GetSalesTaxItemId returns the SalesTaxItemId field if non-nil, zero value otherwise.

### GetSalesTaxItemIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxItemIdOk() (*string, bool)`

GetSalesTaxItemIdOk returns a tuple with the SalesTaxItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesTaxItemId(v string)`

SetSalesTaxItemId sets SalesTaxItemId field to given value.

### HasSalesTaxItemId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesTaxItemId() bool`

HasSalesTaxItemId returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetCustomerMessageId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerMessageId() string`

GetCustomerMessageId returns the CustomerMessageId field if non-nil, zero value otherwise.

### GetCustomerMessageIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetCustomerMessageIdOk() (*string, bool)`

GetCustomerMessageIdOk returns a tuple with the CustomerMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessageId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetCustomerMessageId(v string)`

SetCustomerMessageId sets CustomerMessageId field to given value.

### HasCustomerMessageId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasCustomerMessageId() bool`

HasCustomerMessageId returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.

### HasOtherCustomField

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasOtherCustomField() bool`

HasOtherCustomField returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLines() []QuickbooksDesktopEstimatesIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopEstimatesIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetLines(v []QuickbooksDesktopEstimatesIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLineGroups() []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopEstimatesIdPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopEstimatesIdPostRequest) SetLineGroups(v []QuickbooksDesktopEstimatesIdPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopEstimatesIdPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


