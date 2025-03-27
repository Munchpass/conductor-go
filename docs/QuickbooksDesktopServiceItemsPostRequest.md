# QuickbooksDesktopServiceItemsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The case-insensitive name of this service item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two service items could both have the &#x60;name&#x60; \&quot;Web-Design\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Consulting:Web-Design\&quot; and \&quot;Contracting:Web-Design\&quot;.  Maximum length: 31 characters. | 
**Barcode** | Pointer to [**QuickbooksDesktopServiceItemsPostRequestBarcode**](QuickbooksDesktopServiceItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this service item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] [default to true]
**ClassId** | Pointer to **string** | The service item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent service item one level above this one in the hierarchy. For example, if this service item has a &#x60;fullName&#x60; of \&quot;Consulting:Web-Design\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Consulting\&quot;. If this service item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this service item, which consists of a base unit and related units. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this service item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesOrPurchaseDetails** | Pointer to [**QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails**](QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails.md) |  | [optional] 
**SalesAndPurchaseDetails** | Pointer to [**QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails**](QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails.md) |  | [optional] 
**ExternalId** | Pointer to **string** | A globally unique identifier (GUID) you, the developer, can provide for tracking this object in your external system. This field is immutable and can only be set during object creation.  **IMPORTANT**: This field must be formatted as a valid GUID; otherwise, QuickBooks will return an error. | [optional] 

## Methods

### NewQuickbooksDesktopServiceItemsPostRequest

`func NewQuickbooksDesktopServiceItemsPostRequest(name string, ) *QuickbooksDesktopServiceItemsPostRequest`

NewQuickbooksDesktopServiceItemsPostRequest instantiates a new QuickbooksDesktopServiceItemsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopServiceItemsPostRequestWithDefaults

`func NewQuickbooksDesktopServiceItemsPostRequestWithDefaults() *QuickbooksDesktopServiceItemsPostRequest`

NewQuickbooksDesktopServiceItemsPostRequestWithDefaults instantiates a new QuickbooksDesktopServiceItemsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBarcode

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetBarcode() QuickbooksDesktopServiceItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetBarcodeOk() (*QuickbooksDesktopServiceItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetBarcode(v QuickbooksDesktopServiceItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesOrPurchaseDetails() QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesOrPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesOrPurchaseDetails(v QuickbooksDesktopServiceItemsPostRequestSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.

### HasSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesOrPurchaseDetails() bool`

HasSalesOrPurchaseDetails returns a boolean if a field has been set.

### GetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesAndPurchaseDetails() QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetSalesAndPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetSalesAndPurchaseDetails(v QuickbooksDesktopServiceItemsPostRequestSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.

### HasSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasSalesAndPurchaseDetails() bool`

HasSalesAndPurchaseDetails returns a boolean if a field has been set.

### GetExternalId

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *QuickbooksDesktopServiceItemsPostRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *QuickbooksDesktopServiceItemsPostRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *QuickbooksDesktopServiceItemsPostRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


