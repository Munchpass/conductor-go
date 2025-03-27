# QuickbooksDesktopPurchaseOrdersIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the purchase order object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**VendorId** | Pointer to **string** | The vendor who sent this purchase order for goods or services purchased. | [optional] 
**ClassId** | Pointer to **string** | The purchase order&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. A class defined here is automatically used in this purchase order&#39;s line items unless overridden at the line item level. | [optional] 
**InventorySiteId** | Pointer to **string** | The site location where inventory for the item associated with this purchase order is stored. | [optional] 
**ShipToEntityId** | Pointer to **string** | The customer, vendor, employee, or other entity to whom this purchase order is to be shipped. | [optional] 
**DocumentTemplateId** | Pointer to **string** | The predefined template in QuickBooks that determines the layout and formatting for this purchase order when printed or displayed. | [optional] 
**TransactionDate** | Pointer to **string** | The date of this purchase order, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**RefNumber** | Pointer to **string** | The case-sensitive user-defined reference number for this purchase order, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | [optional] 
**VendorAddress** | Pointer to [**QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress**](QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress.md) |  | [optional] 
**ShippingAddress** | Pointer to [**QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress**](QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress.md) |  | [optional] 
**TermsId** | Pointer to **string** | The purchase order&#39;s payment terms, defining when payment is due and any applicable discounts. | [optional] 
**DueDate** | Pointer to **string** | The date by which this purchase order must be paid, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ExpectedDate** | Pointer to **string** | The date on which shipment of this purchase order is expected to be completed, in ISO 8601 format (YYYY-MM-DD). | [optional] 
**ShippingMethodId** | Pointer to **string** | The shipping method used for this purchase order, such as standard mail or overnight delivery. | [optional] 
**ShipmentOrigin** | Pointer to **string** | The origin location from where the product associated with this purchase order is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | [optional] 
**IsManuallyClosed** | Pointer to **bool** | Indicates whether this purchase order has been manually marked as closed, even if all items have not been received or the sale has not been cancelled. Once the purchase order is marked as closed, all of its line items become closed as well. You cannot change &#x60;isManuallyClosed&#x60; to &#x60;false&#x60; after the purchase order has been fully received. | [optional] 
**Memo** | Pointer to **string** | A memo or note for this purchase order that appears in reports, but not on the purchase order. | [optional] 
**VendorMessage** | Pointer to **string** | A message to be printed on this purchase order for the vendor to read. | [optional] 
**IsQueuedForPrint** | Pointer to **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to print. | [optional] 
**IsQueuedForEmail** | Pointer to **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to email to the customer. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The sales-tax code for this purchase order, determining whether it is taxable or non-taxable. If set, this overrides any sales-tax codes defined on the vendor. This can be overridden on the purchase order&#39;s individual lines.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**OtherCustomField1** | Pointer to **string** | A built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**OtherCustomField2** | Pointer to **string** | A second built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | [optional] 
**ExchangeRate** | Pointer to **float32** | The market exchange rate between this purchase order&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | [optional] 
**Lines** | Pointer to [**[]QuickbooksDesktopPurchaseOrdersIdPostRequestLinesInner**](QuickbooksDesktopPurchaseOrdersIdPostRequestLinesInner.md) | The purchase order&#39;s line items, each representing a single product or service ordered.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line items for the purchase order with this array. To keep any existing line items, you must include them in this array even if they have not changed. **Any line items not included will be removed.**  2. To add a new line item, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line items, omit this field entirely to keep them unchanged. | [optional] 
**LineGroups** | Pointer to [**[]QuickbooksDesktopPurchaseOrdersIdPostRequestLineGroupsInner**](QuickbooksDesktopPurchaseOrdersIdPostRequestLineGroupsInner.md) | The purchase order&#39;s line item groups, each representing a predefined set of related items.  **IMPORTANT**:  1. Including this array in your update request will **REPLACE** all existing line item groups for the purchase order with this array. To keep any existing line item groups, you must include them in this array even if they have not changed. **Any line item groups not included will be removed.**  2. To add a new line item group, include it here with the &#x60;id&#x60; field set to &#x60;-1&#x60;.  3. If you do not wish to modify any line item groups, omit this field entirely to keep them unchanged. | [optional] 

## Methods

### NewQuickbooksDesktopPurchaseOrdersIdPostRequest

`func NewQuickbooksDesktopPurchaseOrdersIdPostRequest(revisionNumber string, ) *QuickbooksDesktopPurchaseOrdersIdPostRequest`

NewQuickbooksDesktopPurchaseOrdersIdPostRequest instantiates a new QuickbooksDesktopPurchaseOrdersIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopPurchaseOrdersIdPostRequestWithDefaults

`func NewQuickbooksDesktopPurchaseOrdersIdPostRequestWithDefaults() *QuickbooksDesktopPurchaseOrdersIdPostRequest`

NewQuickbooksDesktopPurchaseOrdersIdPostRequestWithDefaults instantiates a new QuickbooksDesktopPurchaseOrdersIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendorId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetInventorySiteId() string`

GetInventorySiteId returns the InventorySiteId field if non-nil, zero value otherwise.

### GetInventorySiteIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetInventorySiteIdOk() (*string, bool)`

GetInventorySiteIdOk returns a tuple with the InventorySiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetInventorySiteId(v string)`

SetInventorySiteId sets InventorySiteId field to given value.

### HasInventorySiteId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasInventorySiteId() bool`

HasInventorySiteId returns a boolean if a field has been set.

### GetShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShipToEntityId() string`

GetShipToEntityId returns the ShipToEntityId field if non-nil, zero value otherwise.

### GetShipToEntityIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShipToEntityIdOk() (*string, bool)`

GetShipToEntityIdOk returns a tuple with the ShipToEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetShipToEntityId(v string)`

SetShipToEntityId sets ShipToEntityId field to given value.

### HasShipToEntityId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasShipToEntityId() bool`

HasShipToEntityId returns a boolean if a field has been set.

### GetDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetDocumentTemplateId() string`

GetDocumentTemplateId returns the DocumentTemplateId field if non-nil, zero value otherwise.

### GetDocumentTemplateIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetDocumentTemplateIdOk() (*string, bool)`

GetDocumentTemplateIdOk returns a tuple with the DocumentTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetDocumentTemplateId(v string)`

SetDocumentTemplateId sets DocumentTemplateId field to given value.

### HasDocumentTemplateId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasDocumentTemplateId() bool`

HasDocumentTemplateId returns a boolean if a field has been set.

### GetTransactionDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorAddress() QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorAddressOk() (*QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetVendorAddress(v QuickbooksDesktopPurchaseOrdersPostRequestVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.

### HasVendorAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasVendorAddress() bool`

HasVendorAddress returns a boolean if a field has been set.

### GetShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShippingAddress() QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShippingAddressOk() (*QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetShippingAddress(v QuickbooksDesktopPurchaseOrdersPostRequestShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### GetTermsId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetTermsId() string`

GetTermsId returns the TermsId field if non-nil, zero value otherwise.

### GetTermsIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetTermsIdOk() (*string, bool)`

GetTermsIdOk returns a tuple with the TermsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetTermsId(v string)`

SetTermsId sets TermsId field to given value.

### HasTermsId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasTermsId() bool`

HasTermsId returns a boolean if a field has been set.

### GetDueDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetExpectedDate() string`

GetExpectedDate returns the ExpectedDate field if non-nil, zero value otherwise.

### GetExpectedDateOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetExpectedDateOk() (*string, bool)`

GetExpectedDateOk returns a tuple with the ExpectedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetExpectedDate(v string)`

SetExpectedDate sets ExpectedDate field to given value.

### HasExpectedDate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasExpectedDate() bool`

HasExpectedDate returns a boolean if a field has been set.

### GetShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShippingMethodId() string`

GetShippingMethodId returns the ShippingMethodId field if non-nil, zero value otherwise.

### GetShippingMethodIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShippingMethodIdOk() (*string, bool)`

GetShippingMethodIdOk returns a tuple with the ShippingMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetShippingMethodId(v string)`

SetShippingMethodId sets ShippingMethodId field to given value.

### HasShippingMethodId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasShippingMethodId() bool`

HasShippingMethodId returns a boolean if a field has been set.

### GetShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.

### HasShipmentOrigin

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasShipmentOrigin() bool`

HasShipmentOrigin returns a boolean if a field has been set.

### GetIsManuallyClosed

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.

### HasIsManuallyClosed

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasIsManuallyClosed() bool`

HasIsManuallyClosed returns a boolean if a field has been set.

### GetMemo

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorMessage() string`

GetVendorMessage returns the VendorMessage field if non-nil, zero value otherwise.

### GetVendorMessageOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetVendorMessageOk() (*string, bool)`

GetVendorMessageOk returns a tuple with the VendorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetVendorMessage(v string)`

SetVendorMessage sets VendorMessage field to given value.

### HasVendorMessage

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasVendorMessage() bool`

HasVendorMessage returns a boolean if a field has been set.

### GetIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.

### HasIsQueuedForPrint

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasIsQueuedForPrint() bool`

HasIsQueuedForPrint returns a boolean if a field has been set.

### GetIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.

### HasIsQueuedForEmail

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasIsQueuedForEmail() bool`

HasIsQueuedForEmail returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.

### HasOtherCustomField1

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasOtherCustomField1() bool`

HasOtherCustomField1 returns a boolean if a field has been set.

### GetOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.

### HasOtherCustomField2

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasOtherCustomField2() bool`

HasOtherCustomField2 returns a boolean if a field has been set.

### GetExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetLines

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetLines() []QuickbooksDesktopPurchaseOrdersIdPostRequestLinesInner`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetLinesOk() (*[]QuickbooksDesktopPurchaseOrdersIdPostRequestLinesInner, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetLines(v []QuickbooksDesktopPurchaseOrdersIdPostRequestLinesInner)`

SetLines sets Lines field to given value.

### HasLines

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetLineGroups() []QuickbooksDesktopPurchaseOrdersIdPostRequestLineGroupsInner`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) GetLineGroupsOk() (*[]QuickbooksDesktopPurchaseOrdersIdPostRequestLineGroupsInner, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) SetLineGroups(v []QuickbooksDesktopPurchaseOrdersIdPostRequestLineGroupsInner)`

SetLineGroups sets LineGroups field to given value.

### HasLineGroups

`func (o *QuickbooksDesktopPurchaseOrdersIdPostRequest) HasLineGroups() bool`

HasLineGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


