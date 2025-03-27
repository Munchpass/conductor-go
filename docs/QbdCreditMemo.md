# QbdCreditMemo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this credit memo. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_credit_memo\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this credit memo was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this credit memo was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this credit memo object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdCreditMemoCustomer**](QbdCreditMemoCustomer.md) |  | 
**Class** | [**QbdCreditMemoClass**](QbdCreditMemoClass.md) |  | 
**ReceivablesAccount** | [**QbdCreditMemoReceivablesAccount**](QbdCreditMemoReceivablesAccount.md) |  | 
**DocumentTemplate** | [**QbdCreditMemoDocumentTemplate**](QbdCreditMemoDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this credit memo, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this credit memo, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BillingAddress** | [**QbdCreditMemoBillingAddress**](QbdCreditMemoBillingAddress.md) |  | 
**ShippingAddress** | [**QbdCreditMemoShippingAddress**](QbdCreditMemoShippingAddress.md) |  | 
**IsPending** | **bool** | Indicates whether this credit memo has not been completed. | 
**PurchaseOrderNumber** | **string** | The customer&#39;s Purchase Order (PO) number associated with this credit memo. This field is often used to cross-reference the credit memo with the customer&#39;s purchasing system. | 
**Terms** | [**QbdCreditMemoTerms**](QbdCreditMemoTerms.md) |  | 
**DueDate** | **string** | The date by which this credit memo must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**SalesRepresentative** | [**QbdCreditMemoSalesRepresentative**](QbdCreditMemoSalesRepresentative.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this credit memo is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**ShippingDate** | **string** | The date when the products or services for this credit memo were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | 
**ShippingMethod** | [**QbdCreditMemoShippingMethod**](QbdCreditMemoShippingMethod.md) |  | 
**Subtotal** | **string** | The subtotal of this credit memo, which is the sum of all credit memo lines before taxes and payments are applied, represented as a decimal string. | 
**SalesTaxItem** | [**QbdCreditMemoSalesTaxItem**](QbdCreditMemoSalesTaxItem.md) |  | 
**SalesTaxPercentage** | **string** | The sales tax percentage applied to this credit memo, represented as a decimal string. | 
**SalesTaxTotal** | **string** | The total amount of sales tax charged for this credit memo, represented as a decimal string. | 
**TotalAmount** | **string** | The total monetary amount of this credit memo, equivalent to the sum of the amounts in &#x60;lines&#x60; and &#x60;lineGroups&#x60;, represented as a decimal string. | 
**CreditRemaining** | **string** | The remaining balance of this credit memo that has not yet been applied to other transactions or refunded to the customer. Represented as a decimal string. | 
**Currency** | [**QbdCreditMemoCurrency**](QbdCreditMemoCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this credit memo&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**CreditRemainingInHomeCurrency** | **string** | The remaining balance of this credit memo converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | A memo or note for this credit memo that appears in the account register and customer register, but not on the credit memo itself. | 
**CustomerMessage** | [**QbdCreditMemoCustomerMessage**](QbdCreditMemoCustomerMessage.md) |  | 
**IsQueuedForPrint** | **bool** | Indicates whether this credit memo is included in the queue of documents for QuickBooks to print. | 
**IsQueuedForEmail** | **bool** | Indicates whether this credit memo is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdCreditMemoSalesTaxCode**](QbdCreditMemoSalesTaxCode.md) |  | 
**OtherCustomField** | **string** | A built-in custom field for additional information specific to this credit memo. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all credit memos for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The credit memo&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of credit memos to receive this field because it is not returned by default. | 
**Lines** | [**[]QbdCreditMemoLine**](QbdCreditMemoLine.md) | The credit memo&#39;s line items, each representing a single product or service sold. | 
**LineGroups** | [**[]QbdCreditMemoLineGroup**](QbdCreditMemoLineGroup.md) | The credit memo&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the credit memo object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdCreditMemo

`func NewQbdCreditMemo(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdCreditMemoCustomer, class QbdCreditMemoClass, receivablesAccount QbdCreditMemoReceivablesAccount, documentTemplate QbdCreditMemoDocumentTemplate, transactionDate string, refNumber string, billingAddress QbdCreditMemoBillingAddress, shippingAddress QbdCreditMemoShippingAddress, isPending bool, purchaseOrderNumber string, terms QbdCreditMemoTerms, dueDate string, salesRepresentative QbdCreditMemoSalesRepresentative, shipmentOrigin string, shippingDate string, shippingMethod QbdCreditMemoShippingMethod, subtotal string, salesTaxItem QbdCreditMemoSalesTaxItem, salesTaxPercentage string, salesTaxTotal string, totalAmount string, creditRemaining string, currency QbdCreditMemoCurrency, exchangeRate float32, creditRemainingInHomeCurrency string, memo string, customerMessage QbdCreditMemoCustomerMessage, isQueuedForPrint bool, isQueuedForEmail bool, salesTaxCode QbdCreditMemoSalesTaxCode, otherCustomField string, externalId string, linkedTransactions []QbdLinkedTransaction, lines []QbdCreditMemoLine, lineGroups []QbdCreditMemoLineGroup, customFields []QbdCustomField, ) *QbdCreditMemo`

NewQbdCreditMemo instantiates a new QbdCreditMemo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdCreditMemoWithDefaults

`func NewQbdCreditMemoWithDefaults() *QbdCreditMemo`

NewQbdCreditMemoWithDefaults instantiates a new QbdCreditMemo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdCreditMemo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdCreditMemo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdCreditMemo) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdCreditMemo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdCreditMemo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdCreditMemo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdCreditMemo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdCreditMemo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdCreditMemo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdCreditMemo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdCreditMemo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdCreditMemo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdCreditMemo) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdCreditMemo) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdCreditMemo) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdCreditMemo) GetCustomer() QbdCreditMemoCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdCreditMemo) GetCustomerOk() (*QbdCreditMemoCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdCreditMemo) SetCustomer(v QbdCreditMemoCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdCreditMemo) GetClass() QbdCreditMemoClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdCreditMemo) GetClassOk() (*QbdCreditMemoClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdCreditMemo) SetClass(v QbdCreditMemoClass)`

SetClass sets Class field to given value.


### GetReceivablesAccount

`func (o *QbdCreditMemo) GetReceivablesAccount() QbdCreditMemoReceivablesAccount`

GetReceivablesAccount returns the ReceivablesAccount field if non-nil, zero value otherwise.

### GetReceivablesAccountOk

`func (o *QbdCreditMemo) GetReceivablesAccountOk() (*QbdCreditMemoReceivablesAccount, bool)`

GetReceivablesAccountOk returns a tuple with the ReceivablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccount

`func (o *QbdCreditMemo) SetReceivablesAccount(v QbdCreditMemoReceivablesAccount)`

SetReceivablesAccount sets ReceivablesAccount field to given value.


### GetDocumentTemplate

`func (o *QbdCreditMemo) GetDocumentTemplate() QbdCreditMemoDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdCreditMemo) GetDocumentTemplateOk() (*QbdCreditMemoDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdCreditMemo) SetDocumentTemplate(v QbdCreditMemoDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdCreditMemo) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdCreditMemo) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdCreditMemo) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdCreditMemo) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdCreditMemo) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdCreditMemo) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBillingAddress

`func (o *QbdCreditMemo) GetBillingAddress() QbdCreditMemoBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdCreditMemo) GetBillingAddressOk() (*QbdCreditMemoBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdCreditMemo) SetBillingAddress(v QbdCreditMemoBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *QbdCreditMemo) GetShippingAddress() QbdCreditMemoShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdCreditMemo) GetShippingAddressOk() (*QbdCreditMemoShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdCreditMemo) SetShippingAddress(v QbdCreditMemoShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetIsPending

`func (o *QbdCreditMemo) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QbdCreditMemo) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QbdCreditMemo) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetPurchaseOrderNumber

`func (o *QbdCreditMemo) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QbdCreditMemo) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QbdCreditMemo) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.


### GetTerms

`func (o *QbdCreditMemo) GetTerms() QbdCreditMemoTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdCreditMemo) GetTermsOk() (*QbdCreditMemoTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdCreditMemo) SetTerms(v QbdCreditMemoTerms)`

SetTerms sets Terms field to given value.


### GetDueDate

`func (o *QbdCreditMemo) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdCreditMemo) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdCreditMemo) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetSalesRepresentative

`func (o *QbdCreditMemo) GetSalesRepresentative() QbdCreditMemoSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdCreditMemo) GetSalesRepresentativeOk() (*QbdCreditMemoSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdCreditMemo) SetSalesRepresentative(v QbdCreditMemoSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetShipmentOrigin

`func (o *QbdCreditMemo) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdCreditMemo) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdCreditMemo) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetShippingDate

`func (o *QbdCreditMemo) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QbdCreditMemo) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QbdCreditMemo) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.


### GetShippingMethod

`func (o *QbdCreditMemo) GetShippingMethod() QbdCreditMemoShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *QbdCreditMemo) GetShippingMethodOk() (*QbdCreditMemoShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *QbdCreditMemo) SetShippingMethod(v QbdCreditMemoShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.


### GetSubtotal

`func (o *QbdCreditMemo) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *QbdCreditMemo) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *QbdCreditMemo) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetSalesTaxItem

`func (o *QbdCreditMemo) GetSalesTaxItem() QbdCreditMemoSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdCreditMemo) GetSalesTaxItemOk() (*QbdCreditMemoSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdCreditMemo) SetSalesTaxItem(v QbdCreditMemoSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.


### GetSalesTaxPercentage

`func (o *QbdCreditMemo) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *QbdCreditMemo) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *QbdCreditMemo) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.


### GetSalesTaxTotal

`func (o *QbdCreditMemo) GetSalesTaxTotal() string`

GetSalesTaxTotal returns the SalesTaxTotal field if non-nil, zero value otherwise.

### GetSalesTaxTotalOk

`func (o *QbdCreditMemo) GetSalesTaxTotalOk() (*string, bool)`

GetSalesTaxTotalOk returns a tuple with the SalesTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxTotal

`func (o *QbdCreditMemo) SetSalesTaxTotal(v string)`

SetSalesTaxTotal sets SalesTaxTotal field to given value.


### GetTotalAmount

`func (o *QbdCreditMemo) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdCreditMemo) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdCreditMemo) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCreditRemaining

`func (o *QbdCreditMemo) GetCreditRemaining() string`

GetCreditRemaining returns the CreditRemaining field if non-nil, zero value otherwise.

### GetCreditRemainingOk

`func (o *QbdCreditMemo) GetCreditRemainingOk() (*string, bool)`

GetCreditRemainingOk returns a tuple with the CreditRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemaining

`func (o *QbdCreditMemo) SetCreditRemaining(v string)`

SetCreditRemaining sets CreditRemaining field to given value.


### GetCurrency

`func (o *QbdCreditMemo) GetCurrency() QbdCreditMemoCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdCreditMemo) GetCurrencyOk() (*QbdCreditMemoCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdCreditMemo) SetCurrency(v QbdCreditMemoCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdCreditMemo) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdCreditMemo) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdCreditMemo) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetCreditRemainingInHomeCurrency

`func (o *QbdCreditMemo) GetCreditRemainingInHomeCurrency() string`

GetCreditRemainingInHomeCurrency returns the CreditRemainingInHomeCurrency field if non-nil, zero value otherwise.

### GetCreditRemainingInHomeCurrencyOk

`func (o *QbdCreditMemo) GetCreditRemainingInHomeCurrencyOk() (*string, bool)`

GetCreditRemainingInHomeCurrencyOk returns a tuple with the CreditRemainingInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemainingInHomeCurrency

`func (o *QbdCreditMemo) SetCreditRemainingInHomeCurrency(v string)`

SetCreditRemainingInHomeCurrency sets CreditRemainingInHomeCurrency field to given value.


### GetMemo

`func (o *QbdCreditMemo) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdCreditMemo) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdCreditMemo) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetCustomerMessage

`func (o *QbdCreditMemo) GetCustomerMessage() QbdCreditMemoCustomerMessage`

GetCustomerMessage returns the CustomerMessage field if non-nil, zero value otherwise.

### GetCustomerMessageOk

`func (o *QbdCreditMemo) GetCustomerMessageOk() (*QbdCreditMemoCustomerMessage, bool)`

GetCustomerMessageOk returns a tuple with the CustomerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessage

`func (o *QbdCreditMemo) SetCustomerMessage(v QbdCreditMemoCustomerMessage)`

SetCustomerMessage sets CustomerMessage field to given value.


### GetIsQueuedForPrint

`func (o *QbdCreditMemo) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdCreditMemo) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdCreditMemo) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetIsQueuedForEmail

`func (o *QbdCreditMemo) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdCreditMemo) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdCreditMemo) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdCreditMemo) GetSalesTaxCode() QbdCreditMemoSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdCreditMemo) GetSalesTaxCodeOk() (*QbdCreditMemoSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdCreditMemo) SetSalesTaxCode(v QbdCreditMemoSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetOtherCustomField

`func (o *QbdCreditMemo) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QbdCreditMemo) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QbdCreditMemo) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.


### GetExternalId

`func (o *QbdCreditMemo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdCreditMemo) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdCreditMemo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdCreditMemo) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdCreditMemo) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdCreditMemo) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetLines

`func (o *QbdCreditMemo) GetLines() []QbdCreditMemoLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdCreditMemo) GetLinesOk() (*[]QbdCreditMemoLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdCreditMemo) SetLines(v []QbdCreditMemoLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdCreditMemo) GetLineGroups() []QbdCreditMemoLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdCreditMemo) GetLineGroupsOk() (*[]QbdCreditMemoLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdCreditMemo) SetLineGroups(v []QbdCreditMemoLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdCreditMemo) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdCreditMemo) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdCreditMemo) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


