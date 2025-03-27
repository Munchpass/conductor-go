# QbdPurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this purchase order. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_purchase_order\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this purchase order was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this purchase order was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this purchase order object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Vendor** | [**QbdPurchaseOrderVendor**](QbdPurchaseOrderVendor.md) |  | 
**Class** | [**QbdPurchaseOrderClass**](QbdPurchaseOrderClass.md) |  | 
**InventorySite** | [**QbdPurchaseOrderInventorySite**](QbdPurchaseOrderInventorySite.md) |  | 
**ShipToEntity** | [**QbdPurchaseOrderShipToEntity**](QbdPurchaseOrderShipToEntity.md) |  | 
**DocumentTemplate** | [**QbdPurchaseOrderDocumentTemplate**](QbdPurchaseOrderDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this purchase order, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this purchase order, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**VendorAddress** | [**QbdPurchaseOrderVendorAddress**](QbdPurchaseOrderVendorAddress.md) |  | 
**ShippingAddress** | [**QbdPurchaseOrderShippingAddress**](QbdPurchaseOrderShippingAddress.md) |  | 
**Terms** | [**QbdPurchaseOrderTerms**](QbdPurchaseOrderTerms.md) |  | 
**DueDate** | **string** | The date by which this purchase order must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**ExpectedDate** | **string** | The date on which shipment of this purchase order is expected to be completed, in ISO 8601 format (YYYY-MM-DD). | 
**ShippingMethod** | [**QbdPurchaseOrderShippingMethod**](QbdPurchaseOrderShippingMethod.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this purchase order is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**TotalAmount** | **string** | The total monetary amount of this purchase order, equivalent to the sum of the amounts in &#x60;lines&#x60; and &#x60;lineGroups&#x60;, represented as a decimal string. | 
**Currency** | [**QbdPurchaseOrderCurrency**](QbdPurchaseOrderCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this purchase order&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this purchase order converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**IsManuallyClosed** | **bool** | Indicates whether this purchase order has been manually marked as closed, even if all items have not been received or the sale has not been cancelled. Once the purchase order is marked as closed, all of its line items become closed as well. You cannot change &#x60;isManuallyClosed&#x60; to &#x60;false&#x60; after the purchase order has been fully received. | 
**IsFullyReceived** | **bool** | Indicates whether all items in this purchase order have been received and none of them were closed manually. | 
**Memo** | **string** | A memo or note for this purchase order that appears in reports, but not on the purchase order. | 
**VendorMessage** | **string** | A message to be printed on this purchase order for the vendor to read. | 
**IsQueuedForPrint** | **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to print. | 
**IsQueuedForEmail** | **bool** | Indicates whether this purchase order is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdPurchaseOrderSalesTaxCode**](QbdPurchaseOrderSalesTaxCode.md) |  | 
**OtherCustomField1** | **string** | A built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**OtherCustomField2** | **string** | A second built-in custom field for additional information specific to this purchase order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all purchase orders for convenience. Like &#x60;otherCustomField1&#x60;, developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The purchase order&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of purchase orders to receive this field because it is not returned by default. | 
**Lines** | [**[]QbdPurchaseOrderLine**](QbdPurchaseOrderLine.md) | The purchase order&#39;s line items, each representing a single product or service ordered. | 
**LineGroups** | [**[]QbdPurchaseOrderLineGroup**](QbdPurchaseOrderLineGroup.md) | The purchase order&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the purchase order object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdPurchaseOrder

`func NewQbdPurchaseOrder(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, vendor QbdPurchaseOrderVendor, class QbdPurchaseOrderClass, inventorySite QbdPurchaseOrderInventorySite, shipToEntity QbdPurchaseOrderShipToEntity, documentTemplate QbdPurchaseOrderDocumentTemplate, transactionDate string, refNumber string, vendorAddress QbdPurchaseOrderVendorAddress, shippingAddress QbdPurchaseOrderShippingAddress, terms QbdPurchaseOrderTerms, dueDate string, expectedDate string, shippingMethod QbdPurchaseOrderShippingMethod, shipmentOrigin string, totalAmount string, currency QbdPurchaseOrderCurrency, exchangeRate float32, totalAmountInHomeCurrency string, isManuallyClosed bool, isFullyReceived bool, memo string, vendorMessage string, isQueuedForPrint bool, isQueuedForEmail bool, salesTaxCode QbdPurchaseOrderSalesTaxCode, otherCustomField1 string, otherCustomField2 string, externalId string, linkedTransactions []QbdLinkedTransaction, lines []QbdPurchaseOrderLine, lineGroups []QbdPurchaseOrderLineGroup, customFields []QbdCustomField, ) *QbdPurchaseOrder`

NewQbdPurchaseOrder instantiates a new QbdPurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdPurchaseOrderWithDefaults

`func NewQbdPurchaseOrderWithDefaults() *QbdPurchaseOrder`

NewQbdPurchaseOrderWithDefaults instantiates a new QbdPurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdPurchaseOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdPurchaseOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdPurchaseOrder) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdPurchaseOrder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdPurchaseOrder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdPurchaseOrder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdPurchaseOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdPurchaseOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdPurchaseOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdPurchaseOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdPurchaseOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdPurchaseOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdPurchaseOrder) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdPurchaseOrder) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdPurchaseOrder) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetVendor

`func (o *QbdPurchaseOrder) GetVendor() QbdPurchaseOrderVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *QbdPurchaseOrder) GetVendorOk() (*QbdPurchaseOrderVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *QbdPurchaseOrder) SetVendor(v QbdPurchaseOrderVendor)`

SetVendor sets Vendor field to given value.


### GetClass

`func (o *QbdPurchaseOrder) GetClass() QbdPurchaseOrderClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdPurchaseOrder) GetClassOk() (*QbdPurchaseOrderClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdPurchaseOrder) SetClass(v QbdPurchaseOrderClass)`

SetClass sets Class field to given value.


### GetInventorySite

`func (o *QbdPurchaseOrder) GetInventorySite() QbdPurchaseOrderInventorySite`

GetInventorySite returns the InventorySite field if non-nil, zero value otherwise.

### GetInventorySiteOk

`func (o *QbdPurchaseOrder) GetInventorySiteOk() (*QbdPurchaseOrderInventorySite, bool)`

GetInventorySiteOk returns a tuple with the InventorySite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySite

`func (o *QbdPurchaseOrder) SetInventorySite(v QbdPurchaseOrderInventorySite)`

SetInventorySite sets InventorySite field to given value.


### GetShipToEntity

`func (o *QbdPurchaseOrder) GetShipToEntity() QbdPurchaseOrderShipToEntity`

GetShipToEntity returns the ShipToEntity field if non-nil, zero value otherwise.

### GetShipToEntityOk

`func (o *QbdPurchaseOrder) GetShipToEntityOk() (*QbdPurchaseOrderShipToEntity, bool)`

GetShipToEntityOk returns a tuple with the ShipToEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToEntity

`func (o *QbdPurchaseOrder) SetShipToEntity(v QbdPurchaseOrderShipToEntity)`

SetShipToEntity sets ShipToEntity field to given value.


### GetDocumentTemplate

`func (o *QbdPurchaseOrder) GetDocumentTemplate() QbdPurchaseOrderDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdPurchaseOrder) GetDocumentTemplateOk() (*QbdPurchaseOrderDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdPurchaseOrder) SetDocumentTemplate(v QbdPurchaseOrderDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdPurchaseOrder) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdPurchaseOrder) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdPurchaseOrder) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdPurchaseOrder) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdPurchaseOrder) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdPurchaseOrder) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetVendorAddress

`func (o *QbdPurchaseOrder) GetVendorAddress() QbdPurchaseOrderVendorAddress`

GetVendorAddress returns the VendorAddress field if non-nil, zero value otherwise.

### GetVendorAddressOk

`func (o *QbdPurchaseOrder) GetVendorAddressOk() (*QbdPurchaseOrderVendorAddress, bool)`

GetVendorAddressOk returns a tuple with the VendorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorAddress

`func (o *QbdPurchaseOrder) SetVendorAddress(v QbdPurchaseOrderVendorAddress)`

SetVendorAddress sets VendorAddress field to given value.


### GetShippingAddress

`func (o *QbdPurchaseOrder) GetShippingAddress() QbdPurchaseOrderShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdPurchaseOrder) GetShippingAddressOk() (*QbdPurchaseOrderShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdPurchaseOrder) SetShippingAddress(v QbdPurchaseOrderShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetTerms

`func (o *QbdPurchaseOrder) GetTerms() QbdPurchaseOrderTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdPurchaseOrder) GetTermsOk() (*QbdPurchaseOrderTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdPurchaseOrder) SetTerms(v QbdPurchaseOrderTerms)`

SetTerms sets Terms field to given value.


### GetDueDate

`func (o *QbdPurchaseOrder) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdPurchaseOrder) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdPurchaseOrder) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetExpectedDate

`func (o *QbdPurchaseOrder) GetExpectedDate() string`

GetExpectedDate returns the ExpectedDate field if non-nil, zero value otherwise.

### GetExpectedDateOk

`func (o *QbdPurchaseOrder) GetExpectedDateOk() (*string, bool)`

GetExpectedDateOk returns a tuple with the ExpectedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDate

`func (o *QbdPurchaseOrder) SetExpectedDate(v string)`

SetExpectedDate sets ExpectedDate field to given value.


### GetShippingMethod

`func (o *QbdPurchaseOrder) GetShippingMethod() QbdPurchaseOrderShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *QbdPurchaseOrder) GetShippingMethodOk() (*QbdPurchaseOrderShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *QbdPurchaseOrder) SetShippingMethod(v QbdPurchaseOrderShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.


### GetShipmentOrigin

`func (o *QbdPurchaseOrder) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdPurchaseOrder) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdPurchaseOrder) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetTotalAmount

`func (o *QbdPurchaseOrder) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdPurchaseOrder) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdPurchaseOrder) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdPurchaseOrder) GetCurrency() QbdPurchaseOrderCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdPurchaseOrder) GetCurrencyOk() (*QbdPurchaseOrderCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdPurchaseOrder) SetCurrency(v QbdPurchaseOrderCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdPurchaseOrder) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdPurchaseOrder) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdPurchaseOrder) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdPurchaseOrder) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdPurchaseOrder) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdPurchaseOrder) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetIsManuallyClosed

`func (o *QbdPurchaseOrder) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QbdPurchaseOrder) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QbdPurchaseOrder) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.


### GetIsFullyReceived

`func (o *QbdPurchaseOrder) GetIsFullyReceived() bool`

GetIsFullyReceived returns the IsFullyReceived field if non-nil, zero value otherwise.

### GetIsFullyReceivedOk

`func (o *QbdPurchaseOrder) GetIsFullyReceivedOk() (*bool, bool)`

GetIsFullyReceivedOk returns a tuple with the IsFullyReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullyReceived

`func (o *QbdPurchaseOrder) SetIsFullyReceived(v bool)`

SetIsFullyReceived sets IsFullyReceived field to given value.


### GetMemo

`func (o *QbdPurchaseOrder) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdPurchaseOrder) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdPurchaseOrder) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetVendorMessage

`func (o *QbdPurchaseOrder) GetVendorMessage() string`

GetVendorMessage returns the VendorMessage field if non-nil, zero value otherwise.

### GetVendorMessageOk

`func (o *QbdPurchaseOrder) GetVendorMessageOk() (*string, bool)`

GetVendorMessageOk returns a tuple with the VendorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorMessage

`func (o *QbdPurchaseOrder) SetVendorMessage(v string)`

SetVendorMessage sets VendorMessage field to given value.


### GetIsQueuedForPrint

`func (o *QbdPurchaseOrder) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdPurchaseOrder) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdPurchaseOrder) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetIsQueuedForEmail

`func (o *QbdPurchaseOrder) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdPurchaseOrder) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdPurchaseOrder) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdPurchaseOrder) GetSalesTaxCode() QbdPurchaseOrderSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdPurchaseOrder) GetSalesTaxCodeOk() (*QbdPurchaseOrderSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdPurchaseOrder) SetSalesTaxCode(v QbdPurchaseOrderSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetOtherCustomField1

`func (o *QbdPurchaseOrder) GetOtherCustomField1() string`

GetOtherCustomField1 returns the OtherCustomField1 field if non-nil, zero value otherwise.

### GetOtherCustomField1Ok

`func (o *QbdPurchaseOrder) GetOtherCustomField1Ok() (*string, bool)`

GetOtherCustomField1Ok returns a tuple with the OtherCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField1

`func (o *QbdPurchaseOrder) SetOtherCustomField1(v string)`

SetOtherCustomField1 sets OtherCustomField1 field to given value.


### GetOtherCustomField2

`func (o *QbdPurchaseOrder) GetOtherCustomField2() string`

GetOtherCustomField2 returns the OtherCustomField2 field if non-nil, zero value otherwise.

### GetOtherCustomField2Ok

`func (o *QbdPurchaseOrder) GetOtherCustomField2Ok() (*string, bool)`

GetOtherCustomField2Ok returns a tuple with the OtherCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField2

`func (o *QbdPurchaseOrder) SetOtherCustomField2(v string)`

SetOtherCustomField2 sets OtherCustomField2 field to given value.


### GetExternalId

`func (o *QbdPurchaseOrder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdPurchaseOrder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdPurchaseOrder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdPurchaseOrder) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdPurchaseOrder) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdPurchaseOrder) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetLines

`func (o *QbdPurchaseOrder) GetLines() []QbdPurchaseOrderLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdPurchaseOrder) GetLinesOk() (*[]QbdPurchaseOrderLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdPurchaseOrder) SetLines(v []QbdPurchaseOrderLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdPurchaseOrder) GetLineGroups() []QbdPurchaseOrderLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdPurchaseOrder) GetLineGroupsOk() (*[]QbdPurchaseOrderLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdPurchaseOrder) SetLineGroups(v []QbdPurchaseOrderLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdPurchaseOrder) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdPurchaseOrder) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdPurchaseOrder) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


