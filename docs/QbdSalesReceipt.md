# QbdSalesReceipt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this sales receipt. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_sales_receipt\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this sales receipt was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this sales receipt was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this sales receipt object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdSalesReceiptCustomer**](QbdSalesReceiptCustomer.md) |  | 
**Class** | [**QbdSalesReceiptClass**](QbdSalesReceiptClass.md) |  | 
**DocumentTemplate** | [**QbdSalesReceiptDocumentTemplate**](QbdSalesReceiptDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this sales receipt, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this sales receipt, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BillingAddress** | [**QbdSalesReceiptBillingAddress**](QbdSalesReceiptBillingAddress.md) |  | 
**ShippingAddress** | [**QbdSalesReceiptShippingAddress**](QbdSalesReceiptShippingAddress.md) |  | 
**IsPending** | **bool** | Indicates whether this sales receipt has not been completed. | 
**CheckNumber** | **string** | The check number of a check received for this sales receipt. | 
**PaymentMethod** | [**QbdSalesReceiptPaymentMethod**](QbdSalesReceiptPaymentMethod.md) |  | 
**DueDate** | **string** | The date by which this sales receipt must be paid, in ISO 8601 format (YYYY-MM-DD).  **NOTE**: For sales receipts, this field is often &#x60;null&#x60; because sales receipts are generally used for point-of-sale payments, where full payment is received at the time of purchase. | 
**SalesRepresentative** | [**QbdSalesReceiptSalesRepresentative**](QbdSalesReceiptSalesRepresentative.md) |  | 
**ShippingDate** | **string** | The date when the products or services for this sales receipt were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | 
**ShippingMethod** | [**QbdSalesReceiptShippingMethod**](QbdSalesReceiptShippingMethod.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this sales receipt is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**Subtotal** | **string** | The subtotal of this sales receipt, which is the sum of all sales receipt lines before taxes and payments are applied, represented as a decimal string. | 
**SalesTaxItem** | [**QbdSalesReceiptSalesTaxItem**](QbdSalesReceiptSalesTaxItem.md) |  | 
**SalesTaxPercentage** | **string** | The sales tax percentage applied to this sales receipt, represented as a decimal string. | 
**SalesTaxTotal** | **string** | The total amount of sales tax charged for this sales receipt, represented as a decimal string. | 
**TotalAmount** | **string** | The total monetary amount of this sales receipt, equivalent to the sum of the amounts in &#x60;lines&#x60; and &#x60;lineGroups&#x60;, represented as a decimal string. | 
**Currency** | [**QbdSalesReceiptCurrency**](QbdSalesReceiptCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this sales receipt&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this sales receipt converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | A memo or note for this sales receipt that appears in reports, but not on the sales receipt. | 
**CustomerMessage** | [**QbdSalesReceiptCustomerMessage**](QbdSalesReceiptCustomerMessage.md) |  | 
**IsQueuedForPrint** | **bool** | Indicates whether this sales receipt is included in the queue of documents for QuickBooks to print. | 
**IsQueuedForEmail** | **bool** | Indicates whether this sales receipt is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdSalesReceiptSalesTaxCode**](QbdSalesReceiptSalesTaxCode.md) |  | 
**DepositToAccount** | [**QbdSalesReceiptDepositToAccount**](QbdSalesReceiptDepositToAccount.md) |  | 
**CreditCardTransaction** | [**QbdSalesReceiptCreditCardTransaction**](QbdSalesReceiptCreditCardTransaction.md) |  | 
**OtherCustomField** | **string** | A built-in custom field for additional information specific to this sales receipt. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all sales receipts for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**Lines** | [**[]QbdSalesReceiptLine**](QbdSalesReceiptLine.md) | The sales receipt&#39;s line items, each representing a single product or service sold. | 
**LineGroups** | [**[]QbdSalesReceiptLineGroup**](QbdSalesReceiptLineGroup.md) | The sales receipt&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the sales receipt object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdSalesReceipt

`func NewQbdSalesReceipt(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdSalesReceiptCustomer, class QbdSalesReceiptClass, documentTemplate QbdSalesReceiptDocumentTemplate, transactionDate string, refNumber string, billingAddress QbdSalesReceiptBillingAddress, shippingAddress QbdSalesReceiptShippingAddress, isPending bool, checkNumber string, paymentMethod QbdSalesReceiptPaymentMethod, dueDate string, salesRepresentative QbdSalesReceiptSalesRepresentative, shippingDate string, shippingMethod QbdSalesReceiptShippingMethod, shipmentOrigin string, subtotal string, salesTaxItem QbdSalesReceiptSalesTaxItem, salesTaxPercentage string, salesTaxTotal string, totalAmount string, currency QbdSalesReceiptCurrency, exchangeRate float32, totalAmountInHomeCurrency string, memo string, customerMessage QbdSalesReceiptCustomerMessage, isQueuedForPrint bool, isQueuedForEmail bool, salesTaxCode QbdSalesReceiptSalesTaxCode, depositToAccount QbdSalesReceiptDepositToAccount, creditCardTransaction QbdSalesReceiptCreditCardTransaction, otherCustomField string, externalId string, lines []QbdSalesReceiptLine, lineGroups []QbdSalesReceiptLineGroup, customFields []QbdCustomField, ) *QbdSalesReceipt`

NewQbdSalesReceipt instantiates a new QbdSalesReceipt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdSalesReceiptWithDefaults

`func NewQbdSalesReceiptWithDefaults() *QbdSalesReceipt`

NewQbdSalesReceiptWithDefaults instantiates a new QbdSalesReceipt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdSalesReceipt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdSalesReceipt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdSalesReceipt) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdSalesReceipt) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdSalesReceipt) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdSalesReceipt) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdSalesReceipt) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdSalesReceipt) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdSalesReceipt) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdSalesReceipt) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdSalesReceipt) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdSalesReceipt) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdSalesReceipt) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdSalesReceipt) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdSalesReceipt) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdSalesReceipt) GetCustomer() QbdSalesReceiptCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdSalesReceipt) GetCustomerOk() (*QbdSalesReceiptCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdSalesReceipt) SetCustomer(v QbdSalesReceiptCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdSalesReceipt) GetClass() QbdSalesReceiptClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdSalesReceipt) GetClassOk() (*QbdSalesReceiptClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdSalesReceipt) SetClass(v QbdSalesReceiptClass)`

SetClass sets Class field to given value.


### GetDocumentTemplate

`func (o *QbdSalesReceipt) GetDocumentTemplate() QbdSalesReceiptDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdSalesReceipt) GetDocumentTemplateOk() (*QbdSalesReceiptDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdSalesReceipt) SetDocumentTemplate(v QbdSalesReceiptDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdSalesReceipt) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdSalesReceipt) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdSalesReceipt) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdSalesReceipt) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdSalesReceipt) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdSalesReceipt) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBillingAddress

`func (o *QbdSalesReceipt) GetBillingAddress() QbdSalesReceiptBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdSalesReceipt) GetBillingAddressOk() (*QbdSalesReceiptBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdSalesReceipt) SetBillingAddress(v QbdSalesReceiptBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *QbdSalesReceipt) GetShippingAddress() QbdSalesReceiptShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdSalesReceipt) GetShippingAddressOk() (*QbdSalesReceiptShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdSalesReceipt) SetShippingAddress(v QbdSalesReceiptShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetIsPending

`func (o *QbdSalesReceipt) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QbdSalesReceipt) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QbdSalesReceipt) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetCheckNumber

`func (o *QbdSalesReceipt) GetCheckNumber() string`

GetCheckNumber returns the CheckNumber field if non-nil, zero value otherwise.

### GetCheckNumberOk

`func (o *QbdSalesReceipt) GetCheckNumberOk() (*string, bool)`

GetCheckNumberOk returns a tuple with the CheckNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckNumber

`func (o *QbdSalesReceipt) SetCheckNumber(v string)`

SetCheckNumber sets CheckNumber field to given value.


### GetPaymentMethod

`func (o *QbdSalesReceipt) GetPaymentMethod() QbdSalesReceiptPaymentMethod`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *QbdSalesReceipt) GetPaymentMethodOk() (*QbdSalesReceiptPaymentMethod, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *QbdSalesReceipt) SetPaymentMethod(v QbdSalesReceiptPaymentMethod)`

SetPaymentMethod sets PaymentMethod field to given value.


### GetDueDate

`func (o *QbdSalesReceipt) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdSalesReceipt) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdSalesReceipt) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetSalesRepresentative

`func (o *QbdSalesReceipt) GetSalesRepresentative() QbdSalesReceiptSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdSalesReceipt) GetSalesRepresentativeOk() (*QbdSalesReceiptSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdSalesReceipt) SetSalesRepresentative(v QbdSalesReceiptSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetShippingDate

`func (o *QbdSalesReceipt) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QbdSalesReceipt) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QbdSalesReceipt) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.


### GetShippingMethod

`func (o *QbdSalesReceipt) GetShippingMethod() QbdSalesReceiptShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *QbdSalesReceipt) GetShippingMethodOk() (*QbdSalesReceiptShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *QbdSalesReceipt) SetShippingMethod(v QbdSalesReceiptShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.


### GetShipmentOrigin

`func (o *QbdSalesReceipt) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdSalesReceipt) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdSalesReceipt) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetSubtotal

`func (o *QbdSalesReceipt) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *QbdSalesReceipt) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *QbdSalesReceipt) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetSalesTaxItem

`func (o *QbdSalesReceipt) GetSalesTaxItem() QbdSalesReceiptSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdSalesReceipt) GetSalesTaxItemOk() (*QbdSalesReceiptSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdSalesReceipt) SetSalesTaxItem(v QbdSalesReceiptSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.


### GetSalesTaxPercentage

`func (o *QbdSalesReceipt) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *QbdSalesReceipt) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *QbdSalesReceipt) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.


### GetSalesTaxTotal

`func (o *QbdSalesReceipt) GetSalesTaxTotal() string`

GetSalesTaxTotal returns the SalesTaxTotal field if non-nil, zero value otherwise.

### GetSalesTaxTotalOk

`func (o *QbdSalesReceipt) GetSalesTaxTotalOk() (*string, bool)`

GetSalesTaxTotalOk returns a tuple with the SalesTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxTotal

`func (o *QbdSalesReceipt) SetSalesTaxTotal(v string)`

SetSalesTaxTotal sets SalesTaxTotal field to given value.


### GetTotalAmount

`func (o *QbdSalesReceipt) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdSalesReceipt) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdSalesReceipt) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdSalesReceipt) GetCurrency() QbdSalesReceiptCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdSalesReceipt) GetCurrencyOk() (*QbdSalesReceiptCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdSalesReceipt) SetCurrency(v QbdSalesReceiptCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdSalesReceipt) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdSalesReceipt) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdSalesReceipt) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdSalesReceipt) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdSalesReceipt) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdSalesReceipt) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetMemo

`func (o *QbdSalesReceipt) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdSalesReceipt) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdSalesReceipt) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetCustomerMessage

`func (o *QbdSalesReceipt) GetCustomerMessage() QbdSalesReceiptCustomerMessage`

GetCustomerMessage returns the CustomerMessage field if non-nil, zero value otherwise.

### GetCustomerMessageOk

`func (o *QbdSalesReceipt) GetCustomerMessageOk() (*QbdSalesReceiptCustomerMessage, bool)`

GetCustomerMessageOk returns a tuple with the CustomerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessage

`func (o *QbdSalesReceipt) SetCustomerMessage(v QbdSalesReceiptCustomerMessage)`

SetCustomerMessage sets CustomerMessage field to given value.


### GetIsQueuedForPrint

`func (o *QbdSalesReceipt) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdSalesReceipt) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdSalesReceipt) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetIsQueuedForEmail

`func (o *QbdSalesReceipt) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdSalesReceipt) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdSalesReceipt) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdSalesReceipt) GetSalesTaxCode() QbdSalesReceiptSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdSalesReceipt) GetSalesTaxCodeOk() (*QbdSalesReceiptSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdSalesReceipt) SetSalesTaxCode(v QbdSalesReceiptSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetDepositToAccount

`func (o *QbdSalesReceipt) GetDepositToAccount() QbdSalesReceiptDepositToAccount`

GetDepositToAccount returns the DepositToAccount field if non-nil, zero value otherwise.

### GetDepositToAccountOk

`func (o *QbdSalesReceipt) GetDepositToAccountOk() (*QbdSalesReceiptDepositToAccount, bool)`

GetDepositToAccountOk returns a tuple with the DepositToAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositToAccount

`func (o *QbdSalesReceipt) SetDepositToAccount(v QbdSalesReceiptDepositToAccount)`

SetDepositToAccount sets DepositToAccount field to given value.


### GetCreditCardTransaction

`func (o *QbdSalesReceipt) GetCreditCardTransaction() QbdSalesReceiptCreditCardTransaction`

GetCreditCardTransaction returns the CreditCardTransaction field if non-nil, zero value otherwise.

### GetCreditCardTransactionOk

`func (o *QbdSalesReceipt) GetCreditCardTransactionOk() (*QbdSalesReceiptCreditCardTransaction, bool)`

GetCreditCardTransactionOk returns a tuple with the CreditCardTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransaction

`func (o *QbdSalesReceipt) SetCreditCardTransaction(v QbdSalesReceiptCreditCardTransaction)`

SetCreditCardTransaction sets CreditCardTransaction field to given value.


### GetOtherCustomField

`func (o *QbdSalesReceipt) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QbdSalesReceipt) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QbdSalesReceipt) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.


### GetExternalId

`func (o *QbdSalesReceipt) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdSalesReceipt) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdSalesReceipt) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLines

`func (o *QbdSalesReceipt) GetLines() []QbdSalesReceiptLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdSalesReceipt) GetLinesOk() (*[]QbdSalesReceiptLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdSalesReceipt) SetLines(v []QbdSalesReceiptLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdSalesReceipt) GetLineGroups() []QbdSalesReceiptLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdSalesReceipt) GetLineGroupsOk() (*[]QbdSalesReceiptLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdSalesReceipt) SetLineGroups(v []QbdSalesReceiptLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdSalesReceipt) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdSalesReceipt) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdSalesReceipt) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


