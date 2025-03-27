# QbdEstimate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier assigned by QuickBooks to this estimate. This ID is unique across all transaction types. | 
**ObjectType** | **string** | The type of object. This value is always &#x60;\&quot;qbd_estimate\&quot;&#x60;. | 
**CreatedAt** | **string** | The date and time when this estimate was created, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**UpdatedAt** | **string** | The date and time when this estimate was last updated, in ISO 8601 format (YYYY-MM-DDThh:mm:ss±hh:mm). The time zone is the same as the user&#39;s time zone in QuickBooks. | 
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of this estimate object, which changes each time the object is modified. When updating this object, you must provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Customer** | [**QbdEstimateCustomer**](QbdEstimateCustomer.md) |  | 
**Class** | [**QbdEstimateClass**](QbdEstimateClass.md) |  | 
**DocumentTemplate** | [**QbdEstimateDocumentTemplate**](QbdEstimateDocumentTemplate.md) |  | 
**TransactionDate** | **string** | The date of this estimate, in ISO 8601 format (YYYY-MM-DD). | 
**RefNumber** | **string** | The case-sensitive user-defined reference number for this estimate, which can be used to identify the transaction in QuickBooks. This value is not required to be unique and can be arbitrarily changed by the QuickBooks user. | 
**BillingAddress** | [**QbdEstimateBillingAddress**](QbdEstimateBillingAddress.md) |  | 
**ShippingAddress** | [**QbdEstimateShippingAddress**](QbdEstimateShippingAddress.md) |  | 
**IsActive** | **bool** | Indicates whether this estimate is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | 
**PurchaseOrderNumber** | **string** | The customer&#39;s Purchase Order (PO) number associated with this estimate. This field is often used to cross-reference the estimate with the customer&#39;s purchasing system. | 
**Terms** | [**QbdEstimateTerms**](QbdEstimateTerms.md) |  | 
**DueDate** | **string** | The date by which this estimate must be paid, in ISO 8601 format (YYYY-MM-DD). | 
**SalesRepresentative** | [**QbdEstimateSalesRepresentative**](QbdEstimateSalesRepresentative.md) |  | 
**ShipmentOrigin** | **string** | The origin location from where the product associated with this estimate is shipped. This is the point at which ownership and liability for goods transfer from seller to buyer. Internally, QuickBooks uses the term \&quot;FOB\&quot; for this field, which stands for \&quot;freight on board\&quot;. This field is informational and has no accounting implications. | 
**Subtotal** | **string** | The subtotal of this estimate, which is the sum of all estimate lines before taxes and payments are applied, represented as a decimal string. | 
**SalesTaxItem** | [**QbdEstimateSalesTaxItem**](QbdEstimateSalesTaxItem.md) |  | 
**SalesTaxPercentage** | **string** | The sales tax percentage applied to this estimate, represented as a decimal string. | 
**SalesTaxTotal** | **string** | The total amount of sales tax charged for this estimate, represented as a decimal string. | 
**TotalAmount** | **string** | The total monetary amount of this estimate, equivalent to the sum of the amounts in &#x60;lines&#x60; and &#x60;lineGroups&#x60;, represented as a decimal string. | 
**Currency** | [**QbdEstimateCurrency**](QbdEstimateCurrency.md) |  | 
**ExchangeRate** | **float32** | The market exchange rate between this estimate&#39;s currency and the home currency in QuickBooks at the time of this transaction. Represented as a decimal value (e.g., 1.2345 for 1 EUR &#x3D; 1.2345 USD if USD is the home currency). | 
**TotalAmountInHomeCurrency** | **string** | The total monetary amount of this estimate converted to the home currency of the QuickBooks company file. Represented as a decimal string. | 
**Memo** | **string** | A memo or note for this estimate that appears in reports, but not on the estimate. Use &#x60;customerMessage&#x60; to add a note to this estimate. | 
**CustomerMessage** | [**QbdEstimateCustomerMessage**](QbdEstimateCustomerMessage.md) |  | 
**IsQueuedForEmail** | **bool** | Indicates whether this estimate is included in the queue of documents for QuickBooks to email to the customer. | 
**SalesTaxCode** | [**QbdEstimateSalesTaxCode**](QbdEstimateSalesTaxCode.md) |  | 
**OtherCustomField** | **string** | A built-in custom field for additional information specific to this estimate. Unlike the user-defined fields in the &#x60;customFields&#x60; array, this is a standard QuickBooks field that exists for all estimates for convenience. Developers often use this field for tracking information that doesn&#39;t fit into other standard QuickBooks fields. Unlike &#x60;otherCustomField1&#x60; and &#x60;otherCustomField2&#x60;, which are line item fields, this exists at the transaction level. Hidden by default in the QuickBooks UI. | 
**ExternalId** | **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation. | 
**LinkedTransactions** | [**[]QbdLinkedTransaction**](QbdLinkedTransaction.md) | The estimate&#39;s linked transactions, such as payments applied, credits used, or associated purchase orders.  **IMPORTANT**: You must specify the parameter &#x60;includeLinkedTransactions&#x60; when fetching a list of estimates to receive this field because it is not returned by default. | 
**Lines** | [**[]QbdEstimateLine**](QbdEstimateLine.md) | The estimate&#39;s line items, each representing a single product or service quoted. | 
**LineGroups** | [**[]QbdEstimateLineGroup**](QbdEstimateLineGroup.md) | The estimate&#39;s line item groups, each representing a predefined set of related items. | 
**CustomFields** | [**[]QbdCustomField**](QbdCustomField.md) | The custom fields for the estimate object, added as user-defined data extensions, not included in the standard QuickBooks object. | 

## Methods

### NewQbdEstimate

`func NewQbdEstimate(id string, objectType string, createdAt string, updatedAt string, revisionNumber string, customer QbdEstimateCustomer, class QbdEstimateClass, documentTemplate QbdEstimateDocumentTemplate, transactionDate string, refNumber string, billingAddress QbdEstimateBillingAddress, shippingAddress QbdEstimateShippingAddress, isActive bool, purchaseOrderNumber string, terms QbdEstimateTerms, dueDate string, salesRepresentative QbdEstimateSalesRepresentative, shipmentOrigin string, subtotal string, salesTaxItem QbdEstimateSalesTaxItem, salesTaxPercentage string, salesTaxTotal string, totalAmount string, currency QbdEstimateCurrency, exchangeRate float32, totalAmountInHomeCurrency string, memo string, customerMessage QbdEstimateCustomerMessage, isQueuedForEmail bool, salesTaxCode QbdEstimateSalesTaxCode, otherCustomField string, externalId string, linkedTransactions []QbdLinkedTransaction, lines []QbdEstimateLine, lineGroups []QbdEstimateLineGroup, customFields []QbdCustomField, ) *QbdEstimate`

NewQbdEstimate instantiates a new QbdEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQbdEstimateWithDefaults

`func NewQbdEstimateWithDefaults() *QbdEstimate`

NewQbdEstimateWithDefaults instantiates a new QbdEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QbdEstimate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QbdEstimate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QbdEstimate) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *QbdEstimate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *QbdEstimate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *QbdEstimate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCreatedAt

`func (o *QbdEstimate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QbdEstimate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QbdEstimate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *QbdEstimate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QbdEstimate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QbdEstimate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRevisionNumber

`func (o *QbdEstimate) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QbdEstimate) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QbdEstimate) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetCustomer

`func (o *QbdEstimate) GetCustomer() QbdEstimateCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *QbdEstimate) GetCustomerOk() (*QbdEstimateCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *QbdEstimate) SetCustomer(v QbdEstimateCustomer)`

SetCustomer sets Customer field to given value.


### GetClass

`func (o *QbdEstimate) GetClass() QbdEstimateClass`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *QbdEstimate) GetClassOk() (*QbdEstimateClass, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *QbdEstimate) SetClass(v QbdEstimateClass)`

SetClass sets Class field to given value.


### GetDocumentTemplate

`func (o *QbdEstimate) GetDocumentTemplate() QbdEstimateDocumentTemplate`

GetDocumentTemplate returns the DocumentTemplate field if non-nil, zero value otherwise.

### GetDocumentTemplateOk

`func (o *QbdEstimate) GetDocumentTemplateOk() (*QbdEstimateDocumentTemplate, bool)`

GetDocumentTemplateOk returns a tuple with the DocumentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTemplate

`func (o *QbdEstimate) SetDocumentTemplate(v QbdEstimateDocumentTemplate)`

SetDocumentTemplate sets DocumentTemplate field to given value.


### GetTransactionDate

`func (o *QbdEstimate) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QbdEstimate) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QbdEstimate) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetRefNumber

`func (o *QbdEstimate) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *QbdEstimate) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *QbdEstimate) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.


### GetBillingAddress

`func (o *QbdEstimate) GetBillingAddress() QbdEstimateBillingAddress`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *QbdEstimate) GetBillingAddressOk() (*QbdEstimateBillingAddress, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *QbdEstimate) SetBillingAddress(v QbdEstimateBillingAddress)`

SetBillingAddress sets BillingAddress field to given value.


### GetShippingAddress

`func (o *QbdEstimate) GetShippingAddress() QbdEstimateShippingAddress`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *QbdEstimate) GetShippingAddressOk() (*QbdEstimateShippingAddress, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *QbdEstimate) SetShippingAddress(v QbdEstimateShippingAddress)`

SetShippingAddress sets ShippingAddress field to given value.


### GetIsActive

`func (o *QbdEstimate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QbdEstimate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QbdEstimate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPurchaseOrderNumber

`func (o *QbdEstimate) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *QbdEstimate) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *QbdEstimate) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.


### GetTerms

`func (o *QbdEstimate) GetTerms() QbdEstimateTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *QbdEstimate) GetTermsOk() (*QbdEstimateTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *QbdEstimate) SetTerms(v QbdEstimateTerms)`

SetTerms sets Terms field to given value.


### GetDueDate

`func (o *QbdEstimate) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *QbdEstimate) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *QbdEstimate) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.


### GetSalesRepresentative

`func (o *QbdEstimate) GetSalesRepresentative() QbdEstimateSalesRepresentative`

GetSalesRepresentative returns the SalesRepresentative field if non-nil, zero value otherwise.

### GetSalesRepresentativeOk

`func (o *QbdEstimate) GetSalesRepresentativeOk() (*QbdEstimateSalesRepresentative, bool)`

GetSalesRepresentativeOk returns a tuple with the SalesRepresentative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepresentative

`func (o *QbdEstimate) SetSalesRepresentative(v QbdEstimateSalesRepresentative)`

SetSalesRepresentative sets SalesRepresentative field to given value.


### GetShipmentOrigin

`func (o *QbdEstimate) GetShipmentOrigin() string`

GetShipmentOrigin returns the ShipmentOrigin field if non-nil, zero value otherwise.

### GetShipmentOriginOk

`func (o *QbdEstimate) GetShipmentOriginOk() (*string, bool)`

GetShipmentOriginOk returns a tuple with the ShipmentOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOrigin

`func (o *QbdEstimate) SetShipmentOrigin(v string)`

SetShipmentOrigin sets ShipmentOrigin field to given value.


### GetSubtotal

`func (o *QbdEstimate) GetSubtotal() string`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *QbdEstimate) GetSubtotalOk() (*string, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *QbdEstimate) SetSubtotal(v string)`

SetSubtotal sets Subtotal field to given value.


### GetSalesTaxItem

`func (o *QbdEstimate) GetSalesTaxItem() QbdEstimateSalesTaxItem`

GetSalesTaxItem returns the SalesTaxItem field if non-nil, zero value otherwise.

### GetSalesTaxItemOk

`func (o *QbdEstimate) GetSalesTaxItemOk() (*QbdEstimateSalesTaxItem, bool)`

GetSalesTaxItemOk returns a tuple with the SalesTaxItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxItem

`func (o *QbdEstimate) SetSalesTaxItem(v QbdEstimateSalesTaxItem)`

SetSalesTaxItem sets SalesTaxItem field to given value.


### GetSalesTaxPercentage

`func (o *QbdEstimate) GetSalesTaxPercentage() string`

GetSalesTaxPercentage returns the SalesTaxPercentage field if non-nil, zero value otherwise.

### GetSalesTaxPercentageOk

`func (o *QbdEstimate) GetSalesTaxPercentageOk() (*string, bool)`

GetSalesTaxPercentageOk returns a tuple with the SalesTaxPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxPercentage

`func (o *QbdEstimate) SetSalesTaxPercentage(v string)`

SetSalesTaxPercentage sets SalesTaxPercentage field to given value.


### GetSalesTaxTotal

`func (o *QbdEstimate) GetSalesTaxTotal() string`

GetSalesTaxTotal returns the SalesTaxTotal field if non-nil, zero value otherwise.

### GetSalesTaxTotalOk

`func (o *QbdEstimate) GetSalesTaxTotalOk() (*string, bool)`

GetSalesTaxTotalOk returns a tuple with the SalesTaxTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxTotal

`func (o *QbdEstimate) SetSalesTaxTotal(v string)`

SetSalesTaxTotal sets SalesTaxTotal field to given value.


### GetTotalAmount

`func (o *QbdEstimate) GetTotalAmount() string`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *QbdEstimate) GetTotalAmountOk() (*string, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *QbdEstimate) SetTotalAmount(v string)`

SetTotalAmount sets TotalAmount field to given value.


### GetCurrency

`func (o *QbdEstimate) GetCurrency() QbdEstimateCurrency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QbdEstimate) GetCurrencyOk() (*QbdEstimateCurrency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QbdEstimate) SetCurrency(v QbdEstimateCurrency)`

SetCurrency sets Currency field to given value.


### GetExchangeRate

`func (o *QbdEstimate) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *QbdEstimate) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *QbdEstimate) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.


### GetTotalAmountInHomeCurrency

`func (o *QbdEstimate) GetTotalAmountInHomeCurrency() string`

GetTotalAmountInHomeCurrency returns the TotalAmountInHomeCurrency field if non-nil, zero value otherwise.

### GetTotalAmountInHomeCurrencyOk

`func (o *QbdEstimate) GetTotalAmountInHomeCurrencyOk() (*string, bool)`

GetTotalAmountInHomeCurrencyOk returns a tuple with the TotalAmountInHomeCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmountInHomeCurrency

`func (o *QbdEstimate) SetTotalAmountInHomeCurrency(v string)`

SetTotalAmountInHomeCurrency sets TotalAmountInHomeCurrency field to given value.


### GetMemo

`func (o *QbdEstimate) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *QbdEstimate) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *QbdEstimate) SetMemo(v string)`

SetMemo sets Memo field to given value.


### GetCustomerMessage

`func (o *QbdEstimate) GetCustomerMessage() QbdEstimateCustomerMessage`

GetCustomerMessage returns the CustomerMessage field if non-nil, zero value otherwise.

### GetCustomerMessageOk

`func (o *QbdEstimate) GetCustomerMessageOk() (*QbdEstimateCustomerMessage, bool)`

GetCustomerMessageOk returns a tuple with the CustomerMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMessage

`func (o *QbdEstimate) SetCustomerMessage(v QbdEstimateCustomerMessage)`

SetCustomerMessage sets CustomerMessage field to given value.


### GetIsQueuedForEmail

`func (o *QbdEstimate) GetIsQueuedForEmail() bool`

GetIsQueuedForEmail returns the IsQueuedForEmail field if non-nil, zero value otherwise.

### GetIsQueuedForEmailOk

`func (o *QbdEstimate) GetIsQueuedForEmailOk() (*bool, bool)`

GetIsQueuedForEmailOk returns a tuple with the IsQueuedForEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQueuedForEmail

`func (o *QbdEstimate) SetIsQueuedForEmail(v bool)`

SetIsQueuedForEmail sets IsQueuedForEmail field to given value.


### GetSalesTaxCode

`func (o *QbdEstimate) GetSalesTaxCode() QbdEstimateSalesTaxCode`

GetSalesTaxCode returns the SalesTaxCode field if non-nil, zero value otherwise.

### GetSalesTaxCodeOk

`func (o *QbdEstimate) GetSalesTaxCodeOk() (*QbdEstimateSalesTaxCode, bool)`

GetSalesTaxCodeOk returns a tuple with the SalesTaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCode

`func (o *QbdEstimate) SetSalesTaxCode(v QbdEstimateSalesTaxCode)`

SetSalesTaxCode sets SalesTaxCode field to given value.


### GetOtherCustomField

`func (o *QbdEstimate) GetOtherCustomField() string`

GetOtherCustomField returns the OtherCustomField field if non-nil, zero value otherwise.

### GetOtherCustomFieldOk

`func (o *QbdEstimate) GetOtherCustomFieldOk() (*string, bool)`

GetOtherCustomFieldOk returns a tuple with the OtherCustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherCustomField

`func (o *QbdEstimate) SetOtherCustomField(v string)`

SetOtherCustomField sets OtherCustomField field to given value.


### GetExternalId

`func (o *QbdEstimate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QbdEstimate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QbdEstimate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetLinkedTransactions

`func (o *QbdEstimate) GetLinkedTransactions() []QbdLinkedTransaction`

GetLinkedTransactions returns the LinkedTransactions field if non-nil, zero value otherwise.

### GetLinkedTransactionsOk

`func (o *QbdEstimate) GetLinkedTransactionsOk() (*[]QbdLinkedTransaction, bool)`

GetLinkedTransactionsOk returns a tuple with the LinkedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedTransactions

`func (o *QbdEstimate) SetLinkedTransactions(v []QbdLinkedTransaction)`

SetLinkedTransactions sets LinkedTransactions field to given value.


### GetLines

`func (o *QbdEstimate) GetLines() []QbdEstimateLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *QbdEstimate) GetLinesOk() (*[]QbdEstimateLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *QbdEstimate) SetLines(v []QbdEstimateLine)`

SetLines sets Lines field to given value.


### GetLineGroups

`func (o *QbdEstimate) GetLineGroups() []QbdEstimateLineGroup`

GetLineGroups returns the LineGroups field if non-nil, zero value otherwise.

### GetLineGroupsOk

`func (o *QbdEstimate) GetLineGroupsOk() (*[]QbdEstimateLineGroup, bool)`

GetLineGroupsOk returns a tuple with the LineGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineGroups

`func (o *QbdEstimate) SetLineGroups(v []QbdEstimateLineGroup)`

SetLineGroups sets LineGroups field to given value.


### GetCustomFields

`func (o *QbdEstimate) GetCustomFields() []QbdCustomField`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *QbdEstimate) GetCustomFieldsOk() (*[]QbdCustomField, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *QbdEstimate) SetCustomFields(v []QbdCustomField)`

SetCustomFields sets CustomFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


