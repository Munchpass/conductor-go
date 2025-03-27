# QbdInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this invoice. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_invoice\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this invoice was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this invoice was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this invoice object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdInvoiceCustomer**](QbdInvoiceCustomer.md) |  | 
**Class** | [**QbdInvoiceClass**](QbdInvoiceClass.md) |  | 
**ReceivablesAccount** | [**QbdInvoiceReceivablesAccount**](QbdInvoiceReceivablesAccount.md) |  | 
**DocumentTemplate** | [**QbdInvoiceDocumentTemplate**](QbdInvoiceDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this invoice, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this invoice, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BillingAddress** | [**QbdInvoiceBillingAddress**](QbdInvoiceBillingAddress.md) |  | 
**ShippingAddress** | [**QbdInvoiceShippingAddress**](QbdInvoiceShippingAddress.md) |  | 
**IsPending** | **bool** | Indicates whether this invoice has not been completed or is in a draft version. | 
**IsFinanceCharge** | **bool** | Whether this invoice includes a finance charge. This field is immutable and can only be set during invoice creation. | 
**PurchaseOrderNumber** | **string** | The customer&#39;s Purchase Order (PO) number associated with this invoice. This field is often used to cross-reference the invoice with the customer&#39;s purchasing system. | 
**Terms** | [**QbdInvoiceTerms**](QbdInvoiceTerms.md) |  | 
**DueDate** | **string** | The date by which this invoice must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**SalesRepresentative** | [**QbdInvoiceSalesRepresentative**](QbdInvoiceSalesRepresentative.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this invoice is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**ShippingDate** | **string** | The date when the products or services for this invoice were shipped or are expected to be shipped, in ISO 8601 format (YYYY-MM-DD). | 
**ShippingMethod** | [**QbdInvoiceShippingMethod**](QbdInvoiceShippingMethod.md) |  | 
**Subtotal** | **string** | The subtotal of this invoice, which is the sum of all invoice lines before taxes and payments are applied, represented as a decimal string. | 
**SalesTaxItem** | [**QbdInvoiceSalesTaxItem**](QbdInvoiceSalesTaxItem.md) |  | 
**SalesTaxPercentage** | **string** | The sales tax percentage applied to this invoice, represented as a decimal string. | 
**SalesTaxTotal** | **string** | The total amount of sales tax charged for this invoice, represented as a decimal string. | 
**AppliedAmount** | **string** | The amount of credit applied to this invoice. This could include customer deposits, payments, or credits. Represented as a decimal string. | 
**BalanceRemaining** | **string** | The outstanding balance of this invoice after applying any credits or payments. Calculated as &#x60;subtotal&#x60; + &#x60;salesTaxTotal&#x60; - &#x60;appliedAmount&#x60;. Represented as a decimal string. | 
**Currency** | [**QbdInvoiceCurrency**](QbdInvoiceCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this invoice&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**BalanceRemainingInHomeCurrency** | **string** | The outstanding balance of this invoice converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | A memo or note for this invoice that appears in reports, but not on the invoice. Use &#x60;customerMessage&#x60; to add a note to this invoice. | 
**IsPaid** | **bool** | Indicates whether this invoice has been paid in full. When &#x60;true&#x60;, &#x60;openAmount&#x60; will be 0. | 
**CustomerMessage** | [**QbdInvoiceCustomerMessage**](QbdInvoiceCustomerMessage.md) |  | 
**IsQueuedForPrint** | **bool** | Indicates whether this invoice is included in the queue of documents for QuickBooks to print. | 
**IsQueuedForEmail** | **bool** | Indicates whether this invoice is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdInvoiceSalesTaxCode**](QbdInvoiceSalesTaxCode.md) |  | 
**SuggestedDiscountAmount** | **string** | The suggested discount amount for this invoice, represented as a decimal string. | 
**SuggestedDiscountDate** | **string** | The date when the &#x60;suggestedDiscountAmount&#x60; for this invoice would apply, in ISO 8601 format (YYYY-MM-DD). | 
**OtherCustomField** | **string** | A built-in custom field for additional information specific to this invoice. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all invoices for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The invoice&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of invoices to receive this field because it is not returned by default. | 
**Lines** | [**[]QbdInvoiceLine**](QbdInvoiceLine.md) | The invoice&#39;s line items, each representing a single product or service sold. | 
**LineGroups** | [**[]QbdInvoiceLineGroup**](QbdInvoiceLineGroup.md) | The invoice&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the invoice object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdInvoice

`func NewQbdInvoice(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdInvoiceCustomer, class QbdInvoiceClass, receivablesAccount QbdInvoiceReceivablesAccount, documentTemplate QbdInvoiceDocumentTemplate, transactionDate string, refNumber string, billingAddress QbdInvoiceBillingAddress, shippingAddress QbdInvoiceShippingAddress, isPending bool, isFinanceCharge bool, purchaseOrderNumber string, terms QbdInvoiceTerms, dueDate string, salesRepresentative QbdInvoiceSalesRepresentative, shipmentOrigin string, shippingDate string, shippingMethod QbdInvoiceShippingMethod, subtotal string, salesTaxItem QbdInvoiceSalesTaxItem, salesTaxPercentage string, salesTaxTotal string, appliedAmount string, balanceRemaining string, currency QbdInvoiceCurrency, exchangeRate float32, balanceRemainingInHomeCurrency string, memo string, isPaid bool, customerMessage QbdInvoiceCustomerMessage, isQueuedForPrint bool, isQueuedForEmail bool, salesTaxCode QbdInvoiceSalesTaxCode, suggestedDiscountAmount string, suggestedDiscountDate string, otherCustomField string, externalId string, linkedTransactions []QbdLinkedTransaction, lines []QbdInvoiceLine, lineGroups []QbdInvoiceLineGroup, customFields []QbdCustomField, ) *QbdInvoice`

NewQbdInvoice instantiates a new QbdInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdInvoiceWithDefaults

`func NewQbdInvoiceWithDefaults() *QbdInvoice`

NewQbdInvoiceWithDefaults instantiates a new QbdInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdInvoice) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdInvoice) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdInvoice) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdInvoice) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdInvoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdInvoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdInvoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdInvoice) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdInvoice) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdInvoice) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdInvoice) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdInvoice) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdInvoice) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdInvoice) GetCustomer() QbdInvoiceCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdInvoice) GetCustomerOk() (*QbdInvoiceCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdInvoice) SetCustomer(v QbdInvoiceCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdInvoice) GetClass() QbdInvoiceClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdInvoice) GetClassOk() (*QbdInvoiceClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdInvoice) SetClass(v QbdInvoiceClass)`

SetClass sets Class field to given value.


### GetReceivablesAccount

`func (o *QbdInvoice) GetReceivablesAccount() QbdInvoiceReceivablesAccount`

GetReceivablesAccount returns the ReceivablesAccount field if non-nil, zero value otherwise.

### GetReceivablesAccountOk

`func (o *QbdInvoice) GetReceivablesAccountOk() (*QbdInvoiceReceivablesAccount, bool)`

GetReceivablesAccountOk returns a tuple with the ReceivablesAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivablesAccount

`func (o *QbdInvoice) SetReceivablesAccount(v QbdInvoiceReceivablesAccount)`

SetReceivablesAccount sets ReceivablesAccount field to given value.


### GetDocumentTemplate

`func (o *QbdInvoice) GetDocumentTemplate() QbdInvoiceDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdInvoice) GetDocumentTemplateOk() (*QbdInvoiceDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdInvoice) SetDocumentTemplate(v QbdInvoiceDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdInvoice) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdInvoice) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdInvoice) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdInvoice) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdInvoice) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdInvoice) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBillingAddress

`func (o *QbdInvoice) GetBillingAddress() QbdInvoiceBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdInvoice) GetBillingAddressOk() (*QbdInvoiceBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdInvoice) SetBillingAddress(v QbdInvoiceBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *QbdInvoice) GetShippingAddress() QbdInvoiceShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdInvoice) GetShippingAddressOk() (*QbdInvoiceShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdInvoice) SetShippingAddress(v QbdInvoiceShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetIsPending

`func (o *QbdInvoice) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *QbdInvoice) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *QbdInvoice) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetIsFinanceCharge

`func (o *QbdInvoice) GetIsFinanceCharge() bool`

GetIsFinanceCharge returns the IsFinanceCharge field if non-nil, zero value otherwise.

### GetIsFinanceChargeOk

`func (o *QbdInvoice) GetIsFinanceChargeOk() (*bool, bool)`

GetIsFinanceChargeOk returns a tuple with the IsFinanceCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinanceCharge

`func (o *QbdInvoice) SetIsFinanceCharge(v bool)`

SetIsFinanceCharge sets IsFinanceCharge field to given value.


### GetPurchaseOrderNumber

`func (o *QbdInvoice) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QbdInvoice) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QbdInvoice) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.


### GetTerms

`func (o *QbdInvoice) GetTerms() QbdInvoiceTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdInvoice) GetTermsOk() (*QbdInvoiceTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdInvoice) SetTerms(v QbdInvoiceTerms)`

SetTerms sets Terms field to given value.


### GetDueDate

`func (o *QbdInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetSalesRepresentative

`func (o *QbdInvoice) GetSalesRepresentative() QbdInvoiceSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdInvoice) GetSalesRepresentativeOk() (*QbdInvoiceSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdInvoice) SetSalesRepresentative(v QbdInvoiceSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetShipmentOrigin

`func (o *QbdInvoice) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdInvoice) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdInvoice) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetShippingDate

`func (o *QbdInvoice) GetShippingDate() string`

GetShippingDate returns the ShippingDate field if non-nil, zero value otherwise.

### GetShippingDateOk

`func (o *QbdInvoice) GetShippingDateOk() (*string, bool)`

GetShippingDateOk returns a tuple with the ShippingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingDate

`func (o *QbdInvoice) SetShippingDate(v string)`

SetShippingDate sets ShippingDate field to given value.


### GetShippingMethod

`func (o *QbdInvoice) GetShippingMethod() QbdInvoiceShippingMethod`

GetShippingMethod returns the ShippingMethod field if non-nil, zero value otherwise.

### GetShippingMethodOk

`func (o *QbdInvoice) GetShippingMethodOk() (*QbdInvoiceShippingMethod, bool)`

GetShippingMethodOk returns a tuple with the ShippingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingMethod

`func (o *QbdInvoice) SetShippingMethod(v QbdInvoiceShippingMethod)`

SetShippingMethod sets ShippingMethod field to given value.


### GetSubtotal

`func (o *QbdInvoice) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *QbdInvoice) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *QbdInvoice) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetSalesTaxItem

`func (o *QbdInvoice) GetSalesTaxItem() QbdInvoiceSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdInvoice) GetSalesTaxItemOk() (*QbdInvoiceSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdInvoice) SetSalesTaxItem(v QbdInvoiceSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.


### GetSalesTaxPercentage

`func (o *QbdInvoice) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *QbdInvoice) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *QbdInvoice) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.


### GetSalesTaxTotal

`func (o *QbdInvoice) GetSalesTaxTotal() string`

GetSalesTaxTotal returns the SalesTaxTotal field if non-nil, zero value otherwise.

### GetSalesTaxTotalOk

`func (o *QbdInvoice) GetSalesTaxTotalOk() (*string, bool)`

GetSalesTaxTotalOk returns a tuple with the SalesTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxTotal

`func (o *QbdInvoice) SetSalesTaxTotal(v string)`

SetSalesTaxTotal sets SalesTaxTotal field to given value.


### GetAppliedAmount

`func (o *QbdInvoice) GetAppliedAmount() string`

GetAppliedAmount returns the AppliedAmount field if non-nil, zero value otherwise.

### GetAppliedAmountOk

`func (o *QbdInvoice) GetAppliedAmountOk() (*string, bool)`

GetAppliedAmountOk returns a tuple with the AppliedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedAmount

`func (o *QbdInvoice) SetAppliedAmount(v string)`

SetAppliedAmount sets AppliedAmount field to given value.


### GetBalanceRemaining

`func (o *QbdInvoice) GetBalanceRemaining() string`

GetBalanceRemaining returns the BalanceRemaining field if non-nil, zero value otherwise.

### GetBalanceRemainingOk

`func (o *QbdInvoice) GetBalanceRemainingOk() (*string, bool)`

GetBalanceRemainingOk returns a tuple with the BalanceRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceRemaining

`func (o *QbdInvoice) SetBalanceRemaining(v string)`

SetBalanceRemaining sets BalanceRemaining field to given value.


### GetCurrency

`func (o *QbdInvoice) GetCurrency() QbdInvoiceCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdInvoice) GetCurrencyOk() (*QbdInvoiceCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdInvoice) SetCurrency(v QbdInvoiceCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdInvoice) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdInvoice) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdInvoice) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetBalanceRemainingInHomeCurrency

`func (o *QbdInvoice) GetBalanceRemainingInHomeCurrency() string`

GetBalanceRemainingInHomeCurrency returns the BalanceRemainingInHomeCurrency field if non-nil, zero value otherwise.

### GetBalanceRemainingInHomeCurrencyOk

`func (o *QbdInvoice) GetBalanceRemainingInHomeCurrencyOk() (*string, bool)`

GetBalanceRemainingInHomeCurrencyOk returns a tuple with the BalanceRemainingInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceRemainingInHomeCurrency

`func (o *QbdInvoice) SetBalanceRemainingInHomeCurrency(v string)`

SetBalanceRemainingInHomeCurrency sets BalanceRemainingInHomeCurrency field to given value.


### GetMemo

`func (o *QbdInvoice) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdInvoice) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdInvoice) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetIsPaid

`func (o *QbdInvoice) GetIsPaid() bool`

GetIsPaid returns the IsPaid field if non-nil, zero value otherwise.

### GetIsPaidOk

`func (o *QbdInvoice) GetIsPaidOk() (*bool, bool)`

GetIsPaidOk returns a tuple with the IsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPaid

`func (o *QbdInvoice) SetIsPaid(v bool)`

SetIsPaid sets IsPaid field to given value.


### GetCustomerMessage

`func (o *QbdInvoice) GetCustomerMessage() QbdInvoiceCustomerMessage`

GetCustomerMessage returns the CustomerMessage field if non-nil, zero value otherwise.

### GetCustomerMessageOk

`func (o *QbdInvoice) GetCustomerMessageOk() (*QbdInvoiceCustomerMessage, bool)`

GetCustomerMessageOk returns a tuple with the CustomerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessage

`func (o *QbdInvoice) SetCustomerMessage(v QbdInvoiceCustomerMessage)`

SetCustomerMessage sets CustomerMessage field to given value.


### GetIsQueuedForPrint

`func (o *QbdInvoice) GetIsQueuedForPrint() bool`

GetIsQueuedForPrint returns the IsQueuedForPrint field if non-nil, zero value otherwise.

### GetIsQueuedForPrintOk

`func (o *QbdInvoice) GetIsQueuedForPrintOk() (*bool, bool)`

GetIsQueuedForPrintOk returns a tuple with the IsQueuedForPrint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForPrint

`func (o *QbdInvoice) SetIsQueuedForPrint(v bool)`

SetIsQueuedForPrint sets IsQueuedForPrint field to given value.


### GetIsQueuedForEmail

`func (o *QbdInvoice) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdInvoice) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdInvoice) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdInvoice) GetSalesTaxCode() QbdInvoiceSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdInvoice) GetSalesTaxCodeOk() (*QbdInvoiceSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdInvoice) SetSalesTaxCode(v QbdInvoiceSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetSuggestedDiscountAmount

`func (o *QbdInvoice) GetSuggestedDiscountAmount() string`

GetSuggestedDiscountAmount returns the SuggestedDiscountAmount field if non-nil, zero value otherwise.

### GetSuggestedDiscountAmountOk

`func (o *QbdInvoice) GetSuggestedDiscountAmountOk() (*string, bool)`

GetSuggestedDiscountAmountOk returns a tuple with the SuggestedDiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDiscountAmount

`func (o *QbdInvoice) SetSuggestedDiscountAmount(v string)`

SetSuggestedDiscountAmount sets SuggestedDiscountAmount field to given value.


### GetSuggestedDiscountDate

`func (o *QbdInvoice) GetSuggestedDiscountDate() string`

GetSuggestedDiscountDate returns the SuggestedDiscountDate field if non-nil, zero value otherwise.

### GetSuggestedDiscountDateOk

`func (o *QbdInvoice) GetSuggestedDiscountDateOk() (*string, bool)`

GetSuggestedDiscountDateOk returns a tuple with the SuggestedDiscountDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedDiscountDate

`func (o *QbdInvoice) SetSuggestedDiscountDate(v string)`

SetSuggestedDiscountDate sets SuggestedDiscountDate field to given value.


### GetOtherCustomField

`func (o *QbdInvoice) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QbdInvoice) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QbdInvoice) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.


### GetExternalId

`func (o *QbdInvoice) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdInvoice) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdInvoice) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdInvoice) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdInvoice) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdInvoice) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetLines

`func (o *QbdInvoice) GetLines() []QbdInvoiceLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdInvoice) GetLinesOk() (*[]QbdInvoiceLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdInvoice) SetLines(v []QbdInvoiceLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdInvoice) GetLineGroups() []QbdInvoiceLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdInvoice) GetLineGroupsOk() (*[]QbdInvoiceLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdInvoice) SetLineGroups(v []QbdInvoiceLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdInvoice) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdInvoice) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdInvoice) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


