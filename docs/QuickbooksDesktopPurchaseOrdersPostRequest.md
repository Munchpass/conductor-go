# QuickbooksDesktopPurchaseOrdersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VendorId** | Pointer to **string** | The vendor who sent this purchase order for goods or services purchased. | [optional] 
**ClassId** | Pointer to **string** | The purchase order&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this purchase order&#39;s line items unless overridden at the line item level. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this purchase order is stored. | [optional] 
**ShipToEntityId** | Pointer to **string** | The customer, vendor, employee, or other entity to whom this purchase order is to be shipped. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this purchase order when printed or displayed. | [optional] 
**TransactionDate** | **string** | The date of this purchase order, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this purchase order, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. When left blank in this create request, this field will be left blank in QuickBooks (i.e., it does *not* auto-increment). | [optional] 
**VendorAddress** | Pointer to [**QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress**](QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress**](QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress.md) |  | [optional] 
**TermsId** | Pointer to **string** | The purchase order&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this purchase order must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ExpectedDate** | Pointer to **string** | The date on which shipment of this purchase order is expected to be completed, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this purchase order, such as standard mail or overnight delivery. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this purchase order is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this purchase order that appears in reports, but not on the purchase order. | [optional] 
**VendorMessage** | Pointer to **string** | A message to be printed on this purchase order for the vendor to read. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this purchase order, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the purchase order&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this purchase order&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopPurchaseOrdersPostRequestLinesInner**](QuickbooksDesktopPurchaseOrdersPostRequestLinesInner.md) | The purchase order&#39;s line items, each representing a single product or service ordered.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a purchase order. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopPurchaseOrdersPostRequestLineGroupsInner**](QuickbooksDesktopPurchaseOrdersPostRequestLineGroupsInner.md) | The purchase order&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**: You must specify &#x60;lines&#x60;, &#x60;lineGroups&#x60;, or both when creating a purchase order. | [optional] 

## Methods

### NewQuickbooksDesktopPurchaseOrdersPostRequest

`func NewQuickbooksDesktopPurchaseOrdersPostRequest(transactionDate string, ) *QuickbooksDesktopPurchaseOrdersPostRequest`

NewQuickbooksDesktopPurchaseOrdersPostRequest instantiates a new QuickbooksDesktopPurchaseOrdersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPurchaseOrdersPostRequestWithDefaults

`func NewQuickbooksDesktopPurchaseOrdersPostRequestWithDefaults() *QuickbooksDesktopPurchaseOrdersPostRequest`

NewQuickbooksDesktopPurchaseOrdersPostRequestWithDefaults instantiates a new QuickbooksDesktopPurchaseOrdersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVendorId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShipToEntityId() string`

GetShipToEntityId returns the ShipToEntityId field if non-nil, zero value otherwise.

### GetShipToEntityIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShipToEntityIdOk() (*string, bool)`

GetShipToEntityIdOk returns a tuple with the ShipToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetShipToEntityId(v string)`

SetShipToEntityId sets ShipToEntityId field to given value.

### HasShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasShipToEntityId() bool`

HasShipToEntityId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorAddress() QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorAddressOk() (*QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetVendorAddress(v QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.

### HasVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasVendorAddress() bool`

HasVendorAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShippingAddress() QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShippingAddressOk() (*QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetShippingAddress(v QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExpectedDate() string`

GetExpectedDate returns the ExpectedDate field if non-nil, zero value otherwise.

### GetExpectedDateOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExpectedDateOk() (*string, bool)`

GetExpectedDateOk returns a tuple with the ExpectedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetExpectedDate(v string)`

SetExpectedDate sets ExpectedDate field to given value.

### HasExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasExpectedDate() bool`

HasExpectedDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorMessage() string`

GetVendorMessage returns the VendorMessage field if non-nil, zero value otherwise.

### GetVendorMessageOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetVendorMessageOk() (*string, bool)`

GetVendorMessageOk returns a tuple with the VendorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetVendorMessage(v string)`

SetVendorMessage sets VendorMessage field to given value.

### HasVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasVendorMessage() bool`

HasVendorMessage returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetLines() []QuickbooksDesktopPurchaseOrdersPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetLinesOk() (*[]QuickbooksDesktopPurchaseOrdersPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetLines(v []QuickbooksDesktopPurchaseOrdersPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetLineGroups() []QuickbooksDesktopPurchaseOrdersPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopPurchaseOrdersPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) SetLineGroups(v []QuickbooksDesktopPurchaseOrdersPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


