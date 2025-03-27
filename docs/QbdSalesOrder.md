# QbdSalesOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales order. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_order\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this sales order was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this sales order was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this sales order object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdSalesOrderCustomer**](QbdSalesOrderCustomer.md) |  | 
**Class** | [**QbdSalesOrderClass**](QbdSalesOrderClass.md) |  | 
**DocumentTemplate** | [**QbdSalesOrderDocumentTemplate**](QbdSalesOrderDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this sales order, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this sales order, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BillingAddress** | [**QbdSalesOrderBillingAddress**](QbdSalesOrderBillingAddress.md) |  | 
**ShippingAddress** | [**QbdSalesOrderShippingAddress**](QbdSalesOrderShippingAddress.md) |  | 
**PurchaseOrderNumber** | **string** | The customer&#39;s Purchase Order (PO) number associated with this sales order. This field is often used to cross-reference the sales order with the customer&#39;s purchasing system. | 
**Terms** | [**QbdSalesOrderTerms**](QbdSalesOrderTerms.md) |  | 
**DueDate** | **string** | The date by which this sales order must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**SalesRepresentative** | [**QbdSalesOrderSalesRepresentative**](QbdSalesOrderSalesRepresentative.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this sales order is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**ShippingDate** | **string** | The date when the products or services for this sales order were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | 
**ShippingMethod** | [**QbdSalesOrderShippingMethod**](QbdSalesOrderShippingMethod.md) |  | 
**Subtotal** | **string** | The subtotal of this sales order, which is the sum of all sales order lines before taxes and payments are applied, represented as a decimal string. | 
**SalesTaxItem** | [**QbdSalesOrderSalesTaxItem**](QbdSalesOrderSalesTaxItem.md) |  | 
**SalesTaxPercentage** | **string** | The sales tax percentage applied to this sales order, represented as a decimal string. | 
**SalesTaxTotal** | **string** | The total amount of sales tax charged for this sales order, represented as a decimal string. | 
**TotalAmount** | **string** | The total monetary amount of this sales order, equivalent to the sum of the amounts in &#x60;lines&#x60; and &#x60;lineGroups&#x60;, represented as a decimal string. | 
**Currency** | [**QbdSalesOrderCurrency**](QbdSalesOrderCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this sales order&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this sales order converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**IsManuallyClosed** | **bool** | Indicates whether this sales order has been manually marked as closed, even if it has not been invoiced. | 
**IsFullyInvoiced** | **bool** | Indicates whether all items in this sales order have been invoiced. | 
**Memo** | **string** | A memo or note for this sales order. | 
**CustomerMessage** | [**QbdSalesOrderCustomerMessage**](QbdSalesOrderCustomerMessage.md) |  | 
**IsQueuedForPrint** | **bool** | Indicates whether this sales order is included in the queue of documents for QuickBooks to print. | 
**IsQueuedForEmail** | **bool** | Indicates whether this sales order is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdSalesOrderSalesTaxCode**](QbdSalesOrderSalesTaxCode.md) |  | 
**OtherCustomField** | **string** | A built-in custom field for additional information specific to this sales order. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales orders for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The sales order&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of sales orders to receive this field because it is not returned by default. | 
**Lines** | [**[]QbdSalesOrderLine**](QbdSalesOrderLine.md) | The sales order&#39;s line items, each representing a single product or service ordered. | 
**LineGroups** | [**[]QbdSalesOrderLineGroup**](QbdSalesOrderLineGroup.md) | The sales order&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales order object, added as user-defined data extensions, not included in the standard QuickBooks object. | 
**SalesChannelName** | **string** | The type of the sales channel for this sales order. | 
**SalesStoreName** | **string** | The name of the sales store for this sales order. | 
**SalesStoreType** | **string** | The type of the sales store for this sales order. | 

## Methods

### NewQbdSalesOrder

`func NewQbdSalesOrder(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdSalesOrderCustomer, class QbdSalesOrderClass, documentTemplate QbdSalesOrderDocumentTemplate, transactionDate string, refNumber string, billingAddress QbdSalesOrderBillingAddress, shippingAddress QbdSalesOrderShippingAddress, purchaseOrderNumber string, terms QbdSalesOrderTerms, dueDate string, salesRepresentative QbdSalesOrderSalesRepresentative, shipmentOrigin string, shippingDate string, shippingMethod QbdSalesOrderShippingMethod, subtotal string, salesTaxItem QbdSalesOrderSalesTaxItem, salesTaxPercentage string, salesTaxTotal string, totalAmount string, currency QbdSalesOrderCurrency, exchangeRate float32, totalAmountInHomeCurrency string, isManuallyClosed bool, isFullyInvoiced bool, memo string, customerMessage QbdSalesOrderCustomerMessage, isQueuedForPrint bool, isQueuedForEmail bool, salesTaxCode QbdSalesOrderSalesTaxCode, otherCustomField string, externalId string, linkedTransactions []QbdLinkedTransaction, lines []QbdSalesOrderLine, lineGroups []QbdSalesOrderLineGroup, customFields []QbdCustomField, salesChannelName string, salesStoreName string, salesStoreType string, ) *QbdSalesOrder`

NewQbdSalesOrder instantiates a new QbdSalesOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesOrderWithDefaults

`func NewQbdSalesOrderWithDefaults() *QbdSalesOrder`

NewQbdSalesOrderWithDefaults instantiates a new QbdSalesOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesOrder) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesOrder) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesOrder) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesOrder) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSalesOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSalesOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSalesOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSalesOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSalesOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSalesOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSalesOrder) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSalesOrder) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSalesOrder) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdSalesOrder) GetCustomer() QbdSalesOrderCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdSalesOrder) GetCustomerOk() (*QbdSalesOrderCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdSalesOrder) SetCustomer(v QbdSalesOrderCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdSalesOrder) GetClass() QbdSalesOrderClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdSalesOrder) GetClassOk() (*QbdSalesOrderClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdSalesOrder) SetClass(v QbdSalesOrderClass)`

SetClass sets Class field to given value.


### GetDocumentTemplate

`func (o *QbdSalesOrder) GetDocumentTemplate() QbdSalesOrderDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdSalesOrder) GetDocumentTemplateOk() (*QbdSalesOrderDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdSalesOrder) SetDocumentTemplate(v QbdSalesOrderDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdSalesOrder) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdSalesOrder) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdSalesOrder) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdSalesOrder) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdSalesOrder) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdSalesOrder) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBillingAddress

`func (o *QbdSalesOrder) GetBillingAddress() QbdSalesOrderBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdSalesOrder) GetBillingAddressOk() (*QbdSalesOrderBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdSalesOrder) SetBillingAddress(v QbdSalesOrderBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *QbdSalesOrder) GetShippingAddress() QbdSalesOrderShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdSalesOrder) GetShippingAddressOk() (*QbdSalesOrderShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdSalesOrder) SetShippingAddress(v QbdSalesOrderShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetPurchaseOrderNumber

`func (o *QbdSalesOrder) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QbdSalesOrder) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QbdSalesOrder) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.


### GetTerms

`func (o *QbdSalesOrder) GetTerms() QbdSalesOrderTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdSalesOrder) GetTermsOk() (*QbdSalesOrderTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdSalesOrder) SetTerms(v QbdSalesOrderTerms)`

SetTerms sets Terms field to given value.


### GetDueDate

`func (o *QbdSalesOrder) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdSalesOrder) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdSalesOrder) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetSalesRepresentative

`func (o *QbdSalesOrder) GetSalesRepresentative() QbdSalesOrderSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdSalesOrder) GetSalesRepresentativeOk() (*QbdSalesOrderSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdSalesOrder) SetSalesRepresentative(v QbdSalesOrderSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetShipmentOrigin

`func (o *QbdSalesOrder) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdSalesOrder) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdSalesOrder) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetShippingDate

`func (o *QbdSalesOrder) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QbdSalesOrder) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QbdSalesOrder) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.


### GetShippingMethod

`func (o *QbdSalesOrder) GetShippingMethod() QbdSalesOrderShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *QbdSalesOrder) GetShippingMethodOk() (*QbdSalesOrderShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *QbdSalesOrder) SetShippingMethod(v QbdSalesOrderShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.


### GetSubtotal

`func (o *QbdSalesOrder) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *QbdSalesOrder) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *QbdSalesOrder) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetSalesTaxItem

`func (o *QbdSalesOrder) GetSalesTaxItem() QbdSalesOrderSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdSalesOrder) GetSalesTaxItemOk() (*QbdSalesOrderSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdSalesOrder) SetSalesTaxItem(v QbdSalesOrderSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.


### GetSalesTaxPercentage

`func (o *QbdSalesOrder) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *QbdSalesOrder) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *QbdSalesOrder) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.


### GetSalesTaxTotal

`func (o *QbdSalesOrder) GetSalesTaxTotal() string`

GetSalesTaxTotal returns the SalesTaxTotal field if non-nil, zero value otherwise.

### GetSalesTaxTotalOk

`func (o *QbdSalesOrder) GetSalesTaxTotalOk() (*string, bool)`

GetSalesTaxTotalOk returns a tuple with the SalesTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxTotal

`func (o *QbdSalesOrder) SetSalesTaxTotal(v string)`

SetSalesTaxTotal sets SalesTaxTotal field to given value.


### GetTotalAmount

`func (o *QbdSalesOrder) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdSalesOrder) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdSalesOrder) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdSalesOrder) GetCurrency() QbdSalesOrderCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdSalesOrder) GetCurrencyOk() (*QbdSalesOrderCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdSalesOrder) SetCurrency(v QbdSalesOrderCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdSalesOrder) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdSalesOrder) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdSalesOrder) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdSalesOrder) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdSalesOrder) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdSalesOrder) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetIsManuallyClosed

`func (o *QbdSalesOrder) GetIsManuallyClosed() bool`

GetIsManuallyClosed returns the IsManuallyClosed field if non-nil, zero value otherwise.

### GetIsManuallyClosedOk

`func (o *QbdSalesOrder) GetIsManuallyClosedOk() (*bool, bool)`

GetIsManuallyClosedOk returns a tuple with the IsManuallyClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManuallyClosed

`func (o *QbdSalesOrder) SetIsManuallyClosed(v bool)`

SetIsManuallyClosed sets IsManuallyClosed field to given value.


### GetIsFullyInvoiced

`func (o *QbdSalesOrder) GetIsFullyInvoiced() bool`

GetIsFullyInvoiced returns the IsFullyInvoiced field if non-nil, zero value otherwise.

### GetIsFullyInvoicedOk

`func (o *QbdSalesOrder) GetIsFullyInvoicedOk() (*bool, bool)`

GetIsFullyInvoicedOk returns a tuple with the IsFullyInvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFullyInvoiced

`func (o *QbdSalesOrder) SetIsFullyInvoiced(v bool)`

SetIsFullyInvoiced sets IsFullyInvoiced field to given value.


### GetMemo

`func (o *QbdSalesOrder) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdSalesOrder) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdSalesOrder) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetCustomerMessage

`func (o *QbdSalesOrder) GetCustomerMessage() QbdSalesOrderCustomerMessage`

GetCustomerMessage returns the CustomerMessage field if non-nil, zero value otherwise.

### GetCustomerMessageOk

`func (o *QbdSalesOrder) GetCustomerMessageOk() (*QbdSalesOrderCustomerMessage, bool)`

GetCustomerMessageOk returns a tuple with the CustomerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessage

`func (o *QbdSalesOrder) SetCustomerMessage(v QbdSalesOrderCustomerMessage)`

SetCustomerMessage sets CustomerMessage field to given value.


### GetIsQueuedForPrint

`func (o *QbdSalesOrder) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdSalesOrder) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdSalesOrder) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetIsQueuedForEmail

`func (o *QbdSalesOrder) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdSalesOrder) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdSalesOrder) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdSalesOrder) GetSalesTaxCode() QbdSalesOrderSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdSalesOrder) GetSalesTaxCodeOk() (*QbdSalesOrderSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdSalesOrder) SetSalesTaxCode(v QbdSalesOrderSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetOtherCustomField

`func (o *QbdSalesOrder) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QbdSalesOrder) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QbdSalesOrder) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.


### GetExternalId

`func (o *QbdSalesOrder) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdSalesOrder) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdSalesOrder) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdSalesOrder) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdSalesOrder) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdSalesOrder) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetLines

`func (o *QbdSalesOrder) GetLines() []QbdSalesOrderLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdSalesOrder) GetLinesOk() (*[]QbdSalesOrderLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdSalesOrder) SetLines(v []QbdSalesOrderLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdSalesOrder) GetLineGroups() []QbdSalesOrderLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdSalesOrder) GetLineGroupsOk() (*[]QbdSalesOrderLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdSalesOrder) SetLineGroups(v []QbdSalesOrderLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdSalesOrder) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesOrder) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesOrder) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.


### GetSalesChannelName

`func (o *QbdSalesOrder) GetSalesChannelName() string`

GetSalesChannelName returns the SalesChannelName field if non-nil, zero value otherwise.

### GetSalesChannelNameOk

`func (o *QbdSalesOrder) GetSalesChannelNameOk() (*string, bool)`

GetSalesChannelNameOk returns a tuple with the SalesChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannelName

`func (o *QbdSalesOrder) SetSalesChannelName(v string)`

SetSalesChannelName sets SalesChannelName field to given value.


### GetSalesStoreName

`func (o *QbdSalesOrder) GetSalesStoreName() string`

GetSalesStoreName returns the SalesStoreName field if non-nil, zero value otherwise.

### GetSalesStoreNameOk

`func (o *QbdSalesOrder) GetSalesStoreNameOk() (*string, bool)`

GetSalesStoreNameOk returns a tuple with the SalesStoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesStoreName

`func (o *QbdSalesOrder) SetSalesStoreName(v string)`

SetSalesStoreName sets SalesStoreName field to given value.


### GetSalesStoreType

`func (o *QbdSalesOrder) GetSalesStoreType() string`

GetSalesStoreType returns the SalesStoreType field if non-nil, zero value otherwise.

### GetSalesStoreTypeOk

`func (o *QbdSalesOrder) GetSalesStoreTypeOk() (*string, bool)`

GetSalesStoreTypeOk returns a tuple with the SalesStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesStoreType

`func (o *QbdSalesOrder) SetSalesStoreType(v string)`

SetSalesStoreType sets SalesStoreType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


