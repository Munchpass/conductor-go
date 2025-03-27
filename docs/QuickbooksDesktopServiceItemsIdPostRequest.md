# QuickbooksDesktopServiceItemsIdPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevisionNumber** | **string** | The current QuickBooks-assigned revision number of the service item object you are updating, which you can get by fetching the object first. Provide the most recent &#x60;revisionNumber&#x60; to ensure you&#39;re working with the latest data; otherwise, the update will return an error. | 
**Name** | Pointer to **string** | The case-insensitive name of this service item. Not guaranteed to be unique because it does not include the names of its hierarchical parent objects like &#x60;fullName&#x60; does. For example, two service items could both have the &#x60;name&#x60; \&quot;Web-Design\&quot;, but they could have unique &#x60;fullName&#x60; values, such as \&quot;Consulting:Web-Design\&quot; and \&quot;Contracting:Web-Design\&quot;.  Maximum length: 31 characters. | [optional] 
**Barcode** | Pointer to [**QuickbooksDesktopServiceItemsPostRequestBarcode**](QuickbooksDesktopServiceItemsPostRequestBarcode.md) |  | [optional] 
**IsActive** | Pointer to **bool** | Indicates whether this service item is active. Inactive objects are typically hidden from views and reports in QuickBooks. Defaults to &#x60;true&#x60;. | [optional] 
**ClassId** | Pointer to **string** | The service item&#39;s class. Classes can be used to categorize objects into meaningful segments, such as department, location, or type of work. In QuickBooks, class tracking is off by default. | [optional] 
**ParentId** | Pointer to **string** | The parent service item one level above this one in the hierarchy. For example, if this service item has a &#x60;fullName&#x60; of \&quot;Consulting:Web-Design\&quot;, its parent has a &#x60;fullName&#x60; of \&quot;Consulting\&quot;. If this service item is at the top level, this field will be &#x60;null&#x60;. | [optional] 
**UnitOfMeasureSetId** | Pointer to **string** | The unit-of-measure set associated with this service item, which consists of a base unit and related units. | [optional] 
**ForceUnitOfMeasureChange** | Pointer to **bool** | Indicates whether to allow changing the service item&#39;s unit-of-measure set (using the &#x60;unitOfMeasureSetId&#x60; field) when the base unit of the new unit-of-measure set does not match that of the currently assigned set. Without setting this field to &#x60;true&#x60; in this scenario, the request will fail with an error; hence, this field is equivalent to accepting the warning prompt in the QuickBooks UI.  NOTE: Changing the base unit requires you to update the item&#39;s quantities-on-hand and cost to reflect the new unit; otherwise, these values will be inaccurate. Alternatively, consider creating a new item with the desired unit-of-measure set and deactivating the old item. | [optional] 
**SalesTaxCodeId** | Pointer to **string** | The default sales-tax code for this service item, determining whether it is taxable or non-taxable. This can be overridden at the transaction-line level.  Default codes include \&quot;Non\&quot; (non-taxable) and \&quot;Tax\&quot; (taxable), but custom codes can also be created in QuickBooks. If QuickBooks is not set up to charge sales tax (via the \&quot;Do You Charge Sales Tax?\&quot; preference), it will assign the default non-taxable code to all sales. | [optional] 
**SalesOrPurchaseDetails** | Pointer to [**QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails**](QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails.md) |  | [optional] 
**SalesAndPurchaseDetails** | Pointer to [**QuickbooksDesktopServiceItemsIdPostRequestSalesAndPurchaseDetails**](QuickbooksDesktopServiceItemsIdPostRequestSalesAndPurchaseDetails.md) |  | [optional] 

## Methods

### NewQuickbooksDesktopServiceItemsIdPostRequest

`func NewQuickbooksDesktopServiceItemsIdPostRequest(revisionNumber string, ) *QuickbooksDesktopServiceItemsIdPostRequest`

NewQuickbooksDesktopServiceItemsIdPostRequest instantiates a new QuickbooksDesktopServiceItemsIdPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopServiceItemsIdPostRequestWithDefaults

`func NewQuickbooksDesktopServiceItemsIdPostRequestWithDefaults() *QuickbooksDesktopServiceItemsIdPostRequest`

NewQuickbooksDesktopServiceItemsIdPostRequestWithDefaults instantiates a new QuickbooksDesktopServiceItemsIdPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevisionNumber

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetRevisionNumber() string`

GetRevisionNumber returns the RevisionNumber field if non-nil, zero value otherwise.

### GetRevisionNumberOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetRevisionNumberOk() (*string, bool)`

GetRevisionNumberOk returns a tuple with the RevisionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionNumber

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetRevisionNumber(v string)`

SetRevisionNumber sets RevisionNumber field to given value.


### GetName

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBarcode

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetBarcode() QuickbooksDesktopServiceItemsPostRequestBarcode`

GetBarcode returns the Barcode field if non-nil, zero value otherwise.

### GetBarcodeOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetBarcodeOk() (*QuickbooksDesktopServiceItemsPostRequestBarcode, bool)`

GetBarcodeOk returns a tuple with the Barcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarcode

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetBarcode(v QuickbooksDesktopServiceItemsPostRequestBarcode)`

SetBarcode sets Barcode field to given value.

### HasBarcode

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasBarcode() bool`

HasBarcode returns a boolean if a field has been set.

### GetIsActive

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetClassId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetClassId(v string)`

SetClassId sets ClassId field to given value.

### HasClassId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasClassId() bool`

HasClassId returns a boolean if a field has been set.

### GetParentId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetUnitOfMeasureSetId() string`

GetUnitOfMeasureSetId returns the UnitOfMeasureSetId field if non-nil, zero value otherwise.

### GetUnitOfMeasureSetIdOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetUnitOfMeasureSetIdOk() (*string, bool)`

GetUnitOfMeasureSetIdOk returns a tuple with the UnitOfMeasureSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetUnitOfMeasureSetId(v string)`

SetUnitOfMeasureSetId sets UnitOfMeasureSetId field to given value.

### HasUnitOfMeasureSetId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasUnitOfMeasureSetId() bool`

HasUnitOfMeasureSetId returns a boolean if a field has been set.

### GetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetForceUnitOfMeasureChange() bool`

GetForceUnitOfMeasureChange returns the ForceUnitOfMeasureChange field if non-nil, zero value otherwise.

### GetForceUnitOfMeasureChangeOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetForceUnitOfMeasureChangeOk() (*bool, bool)`

GetForceUnitOfMeasureChangeOk returns a tuple with the ForceUnitOfMeasureChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUnitOfMeasureChange

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetForceUnitOfMeasureChange(v bool)`

SetForceUnitOfMeasureChange sets ForceUnitOfMeasureChange field to given value.

### HasForceUnitOfMeasureChange

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasForceUnitOfMeasureChange() bool`

HasForceUnitOfMeasureChange returns a boolean if a field has been set.

### GetSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesTaxCodeId() string`

GetSalesTaxCodeId returns the SalesTaxCodeId field if non-nil, zero value otherwise.

### GetSalesTaxCodeIdOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesTaxCodeIdOk() (*string, bool)`

GetSalesTaxCodeIdOk returns a tuple with the SalesTaxCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetSalesTaxCodeId(v string)`

SetSalesTaxCodeId sets SalesTaxCodeId field to given value.

### HasSalesTaxCodeId

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasSalesTaxCodeId() bool`

HasSalesTaxCodeId returns a boolean if a field has been set.

### GetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesOrPurchaseDetails() QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails`

GetSalesOrPurchaseDetails returns the SalesOrPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesOrPurchaseDetailsOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesOrPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails, bool)`

GetSalesOrPurchaseDetailsOk returns a tuple with the SalesOrPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetSalesOrPurchaseDetails(v QuickbooksDesktopServiceItemsIdPostRequestSalesOrPurchaseDetails)`

SetSalesOrPurchaseDetails sets SalesOrPurchaseDetails field to given value.

### HasSalesOrPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasSalesOrPurchaseDetails() bool`

HasSalesOrPurchaseDetails returns a boolean if a field has been set.

### GetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesAndPurchaseDetails() QuickbooksDesktopServiceItemsIdPostRequestSalesAndPurchaseDetails`

GetSalesAndPurchaseDetails returns the SalesAndPurchaseDetails field if non-nil, zero value otherwise.

### GetSalesAndPurchaseDetailsOk

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) GetSalesAndPurchaseDetailsOk() (*QuickbooksDesktopServiceItemsIdPostRequestSalesAndPurchaseDetails, bool)`

GetSalesAndPurchaseDetailsOk returns a tuple with the SalesAndPurchaseDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) SetSalesAndPurchaseDetails(v QuickbooksDesktopServiceItemsIdPostRequestSalesAndPurchaseDetails)`

SetSalesAndPurchaseDetails sets SalesAndPurchaseDetails field to given value.

### HasSalesAndPurchaseDetails

`func (o *QuickbooksDesktopServiceItemsIdPostRequest) HasSalesAndPurchaseDetails() bool`

HasSalesAndPurchaseDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


