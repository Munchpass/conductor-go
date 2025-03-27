# QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **float32** | The status code returned in the original QBMS transaction response for this credit card transaction. | [optional] 
**StatusMessage** | Pointer to **string** | The status message returned in the original QBMS transaction response for this credit card transaction. | [optional] 
**CreditCardTransactionId** | Pointer to **string** | The ID returned from the credit card processor for this credit card transaction. | [optional] 
**MerchantAccountNumber** | Pointer to **string** | The QBMS account number of the merchant who is running this transaction using the customer&#39;s credit card. | [optional] 
**AuthorizationCode** | Pointer to **string** | The authorization code returned from the credit card processor to indicate that this charge will be paid by the card issuer. | [optional] 
**AvsStreetStatus** | Pointer to **string** | Indicates whether the street address supplied in the transaction request matches the customer&#39;s address on file at the card issuer. | [optional] 
**AvsZipStatus** | Pointer to **string** | Indicates whether the customer postal ZIP code supplied in the transaction request matches the customer&#39;s postal code recognized at the card issuer. | [optional] 
**CardSecurityCodeMatch** | Pointer to **string** | Indicates whether the card security code supplied in the transaction request matches the card security code recognized for that credit card number at the card issuer. | [optional] 
**ReconBatchId** | Pointer to **string** | An internal ID returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | [optional] 
**PaymentGroupingCode** | Pointer to **float32** | An internal code returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | [optional] 
**PaymentStatus** | Pointer to **string** | Indicates whether this credit card transaction is known to have been successfully processed by the card issuer. | [optional] 
**TransactionAuthorizedAt** | Pointer to **string** | The date and time when the credit card processor authorized this credit card transaction. | [optional] 
**TransactionAuthorizationStamp** | Pointer to **float32** | An internal value for this credit card transaction, needed for the QuickBooks reconciliation feature. | [optional] 
**ClientTransactionId** | Pointer to **string** | A value returned from QBMS transactions for future use by the QuickBooks Reconciliation feature. | [optional] 

## Methods

### NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse

`func NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse() *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse`

NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponseWithDefaults

`func NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponseWithDefaults() *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse`

NewQuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponseWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusMessage

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetCreditCardTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetCreditCardTransactionId() string`

GetCreditCardTransactionId returns the CreditCardTransactionId field if non-nil, zero value otherwise.

### GetCreditCardTransactionIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetCreditCardTransactionIdOk() (*string, bool)`

GetCreditCardTransactionIdOk returns a tuple with the CreditCardTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetCreditCardTransactionId(v string)`

SetCreditCardTransactionId sets CreditCardTransactionId field to given value.

### HasCreditCardTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasCreditCardTransactionId() bool`

HasCreditCardTransactionId returns a boolean if a field has been set.

### GetMerchantAccountNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetMerchantAccountNumber() string`

GetMerchantAccountNumber returns the MerchantAccountNumber field if non-nil, zero value otherwise.

### GetMerchantAccountNumberOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetMerchantAccountNumberOk() (*string, bool)`

GetMerchantAccountNumberOk returns a tuple with the MerchantAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetMerchantAccountNumber(v string)`

SetMerchantAccountNumber sets MerchantAccountNumber field to given value.

### HasMerchantAccountNumber

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasMerchantAccountNumber() bool`

HasMerchantAccountNumber returns a boolean if a field has been set.

### GetAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAvsStreetStatus() string`

GetAvsStreetStatus returns the AvsStreetStatus field if non-nil, zero value otherwise.

### GetAvsStreetStatusOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAvsStreetStatusOk() (*string, bool)`

GetAvsStreetStatusOk returns a tuple with the AvsStreetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetAvsStreetStatus(v string)`

SetAvsStreetStatus sets AvsStreetStatus field to given value.

### HasAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasAvsStreetStatus() bool`

HasAvsStreetStatus returns a boolean if a field has been set.

### GetAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAvsZipStatus() string`

GetAvsZipStatus returns the AvsZipStatus field if non-nil, zero value otherwise.

### GetAvsZipStatusOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetAvsZipStatusOk() (*string, bool)`

GetAvsZipStatusOk returns a tuple with the AvsZipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetAvsZipStatus(v string)`

SetAvsZipStatus sets AvsZipStatus field to given value.

### HasAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasAvsZipStatus() bool`

HasAvsZipStatus returns a boolean if a field has been set.

### GetCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetCardSecurityCodeMatch() string`

GetCardSecurityCodeMatch returns the CardSecurityCodeMatch field if non-nil, zero value otherwise.

### GetCardSecurityCodeMatchOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetCardSecurityCodeMatchOk() (*string, bool)`

GetCardSecurityCodeMatchOk returns a tuple with the CardSecurityCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetCardSecurityCodeMatch(v string)`

SetCardSecurityCodeMatch sets CardSecurityCodeMatch field to given value.

### HasCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasCardSecurityCodeMatch() bool`

HasCardSecurityCodeMatch returns a boolean if a field has been set.

### GetReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetReconBatchId() string`

GetReconBatchId returns the ReconBatchId field if non-nil, zero value otherwise.

### GetReconBatchIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetReconBatchIdOk() (*string, bool)`

GetReconBatchIdOk returns a tuple with the ReconBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetReconBatchId(v string)`

SetReconBatchId sets ReconBatchId field to given value.

### HasReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasReconBatchId() bool`

HasReconBatchId returns a boolean if a field has been set.

### GetPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetPaymentGroupingCode() float32`

GetPaymentGroupingCode returns the PaymentGroupingCode field if non-nil, zero value otherwise.

### GetPaymentGroupingCodeOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetPaymentGroupingCodeOk() (*float32, bool)`

GetPaymentGroupingCodeOk returns a tuple with the PaymentGroupingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetPaymentGroupingCode(v float32)`

SetPaymentGroupingCode sets PaymentGroupingCode field to given value.

### HasPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasPaymentGroupingCode() bool`

HasPaymentGroupingCode returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetTransactionAuthorizedAt

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetTransactionAuthorizedAt() string`

GetTransactionAuthorizedAt returns the TransactionAuthorizedAt field if non-nil, zero value otherwise.

### GetTransactionAuthorizedAtOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetTransactionAuthorizedAtOk() (*string, bool)`

GetTransactionAuthorizedAtOk returns a tuple with the TransactionAuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizedAt

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetTransactionAuthorizedAt(v string)`

SetTransactionAuthorizedAt sets TransactionAuthorizedAt field to given value.

### HasTransactionAuthorizedAt

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasTransactionAuthorizedAt() bool`

HasTransactionAuthorizedAt returns a boolean if a field has been set.

### GetTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetTransactionAuthorizationStamp() float32`

GetTransactionAuthorizationStamp returns the TransactionAuthorizationStamp field if non-nil, zero value otherwise.

### GetTransactionAuthorizationStampOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetTransactionAuthorizationStampOk() (*float32, bool)`

GetTransactionAuthorizationStampOk returns a tuple with the TransactionAuthorizationStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetTransactionAuthorizationStamp(v float32)`

SetTransactionAuthorizationStamp sets TransactionAuthorizationStamp field to given value.

### HasTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasTransactionAuthorizationStamp() bool`

HasTransactionAuthorizationStamp returns a boolean if a field has been set.

### GetClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.

### HasClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsIdPostRequestCreditCardTransactionResponse) HasClientTransactionId() bool`

HasClientTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


