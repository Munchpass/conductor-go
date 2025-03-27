# QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** | The status code returned in the original QBMS transaction response for this credit card transaction. | 
**StatusMessage** | **string** | The status message returned in the original QBMS transaction response for this credit card transaction. | 
**CreditCardTransactionId** | **string** | The ID returned from the credit card processor for this credit card transaction. | 
**MerchantAccountNumber** | **string** | The QBMS account number of the merchant who is running this transaction using the customer&#39;s credit card. | 
**AuthorizationCode** | Pointer to **string** | The authorization code returned from the credit card processor to indicate that this charge will be paid by the card issuer. | [optional] 
**AvsStreetStatus** | Pointer to **string** | Indicates whether the street address supplied in the transaction request matches the customer&#39;s address on file at the card issuer. | [optional] 
**AvsZipStatus** | Pointer to **string** | Indicates whether the customer postal ZIP code supplied in the transaction request matches the customer&#39;s postal code recognized at the card issuer. | [optional] 
**CardSecurityCodeMatch** | Pointer to **string** | Indicates whether the card security code supplied in the transaction request matches the card security code recognized for that credit card number at the card issuer. | [optional] 
**ReconBatchId** | Pointer to **string** | An internal ID returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | [optional] 
**PaymentGroupingCode** | Pointer to **float32** | An internal code returned by QuickBooks Merchant Services (QBMS) from the transaction request, needed for the QuickBooks reconciliation feature. | [optional] 
**PaymentStatus** | **string** | Indicates whether this credit card transaction is known to have been successfully processed by the card issuer. | 
**TransactionAuthorizedAt** | **string** | The date and time when the credit card processor authorized this credit card transaction. | 
**TransactionAuthorizationStamp** | Pointer to **float32** | An internal value for this credit card transaction, needed for the QuickBooks reconciliation feature. | [optional] 
**ClientTransactionId** | Pointer to **string** | A value returned from QBMS transactions for future use by the QuickBooks Reconciliation feature. | [optional] 

## Methods

### NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse

`func NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse(statusCode float32, statusMessage string, creditCardTransactionId string, merchantAccountNumber string, paymentStatus string, transactionAuthorizedAt string, ) *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse`

NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse instantiates a new QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponseWithDefaults

`func NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponseWithDefaults() *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse`

NewQuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponseWithDefaults instantiates a new QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetStatusMessage

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.


### GetCreditCardTransactionId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetCreditCardTransactionId() string`

GetCreditCardTransactionId returns the CreditCardTransactionId field if non-nil, zero value otherwise.

### GetCreditCardTransactionIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetCreditCardTransactionIdOk() (*string, bool)`

GetCreditCardTransactionIdOk returns a tuple with the CreditCardTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardTransactionId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetCreditCardTransactionId(v string)`

SetCreditCardTransactionId sets CreditCardTransactionId field to given value.


### GetMerchantAccountNumber

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetMerchantAccountNumber() string`

GetMerchantAccountNumber returns the MerchantAccountNumber field if non-nil, zero value otherwise.

### GetMerchantAccountNumberOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetMerchantAccountNumberOk() (*string, bool)`

GetMerchantAccountNumberOk returns a tuple with the MerchantAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantAccountNumber

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetMerchantAccountNumber(v string)`

SetMerchantAccountNumber sets MerchantAccountNumber field to given value.


### GetAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.

### HasAuthorizationCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasAuthorizationCode() bool`

HasAuthorizationCode returns a boolean if a field has been set.

### GetAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAvsStreetStatus() string`

GetAvsStreetStatus returns the AvsStreetStatus field if non-nil, zero value otherwise.

### GetAvsStreetStatusOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAvsStreetStatusOk() (*string, bool)`

GetAvsStreetStatusOk returns a tuple with the AvsStreetStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetAvsStreetStatus(v string)`

SetAvsStreetStatus sets AvsStreetStatus field to given value.

### HasAvsStreetStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasAvsStreetStatus() bool`

HasAvsStreetStatus returns a boolean if a field has been set.

### GetAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAvsZipStatus() string`

GetAvsZipStatus returns the AvsZipStatus field if non-nil, zero value otherwise.

### GetAvsZipStatusOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetAvsZipStatusOk() (*string, bool)`

GetAvsZipStatusOk returns a tuple with the AvsZipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetAvsZipStatus(v string)`

SetAvsZipStatus sets AvsZipStatus field to given value.

### HasAvsZipStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasAvsZipStatus() bool`

HasAvsZipStatus returns a boolean if a field has been set.

### GetCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetCardSecurityCodeMatch() string`

GetCardSecurityCodeMatch returns the CardSecurityCodeMatch field if non-nil, zero value otherwise.

### GetCardSecurityCodeMatchOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetCardSecurityCodeMatchOk() (*string, bool)`

GetCardSecurityCodeMatchOk returns a tuple with the CardSecurityCodeMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetCardSecurityCodeMatch(v string)`

SetCardSecurityCodeMatch sets CardSecurityCodeMatch field to given value.

### HasCardSecurityCodeMatch

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasCardSecurityCodeMatch() bool`

HasCardSecurityCodeMatch returns a boolean if a field has been set.

### GetReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetReconBatchId() string`

GetReconBatchId returns the ReconBatchId field if non-nil, zero value otherwise.

### GetReconBatchIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetReconBatchIdOk() (*string, bool)`

GetReconBatchIdOk returns a tuple with the ReconBatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetReconBatchId(v string)`

SetReconBatchId sets ReconBatchId field to given value.

### HasReconBatchId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasReconBatchId() bool`

HasReconBatchId returns a boolean if a field has been set.

### GetPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetPaymentGroupingCode() float32`

GetPaymentGroupingCode returns the PaymentGroupingCode field if non-nil, zero value otherwise.

### GetPaymentGroupingCodeOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetPaymentGroupingCodeOk() (*float32, bool)`

GetPaymentGroupingCodeOk returns a tuple with the PaymentGroupingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetPaymentGroupingCode(v float32)`

SetPaymentGroupingCode sets PaymentGroupingCode field to given value.

### HasPaymentGroupingCode

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasPaymentGroupingCode() bool`

HasPaymentGroupingCode returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetTransactionAuthorizedAt

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetTransactionAuthorizedAt() string`

GetTransactionAuthorizedAt returns the TransactionAuthorizedAt field if non-nil, zero value otherwise.

### GetTransactionAuthorizedAtOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetTransactionAuthorizedAtOk() (*string, bool)`

GetTransactionAuthorizedAtOk returns a tuple with the TransactionAuthorizedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizedAt

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetTransactionAuthorizedAt(v string)`

SetTransactionAuthorizedAt sets TransactionAuthorizedAt field to given value.


### GetTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetTransactionAuthorizationStamp() float32`

GetTransactionAuthorizationStamp returns the TransactionAuthorizationStamp field if non-nil, zero value otherwise.

### GetTransactionAuthorizationStampOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetTransactionAuthorizationStampOk() (*float32, bool)`

GetTransactionAuthorizationStampOk returns a tuple with the TransactionAuthorizationStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetTransactionAuthorizationStamp(v float32)`

SetTransactionAuthorizationStamp sets TransactionAuthorizationStamp field to given value.

### HasTransactionAuthorizationStamp

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasTransactionAuthorizationStamp() bool`

HasTransactionAuthorizationStamp returns a boolean if a field has been set.

### GetClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetClientTransactionId() string`

GetClientTransactionId returns the ClientTransactionId field if non-nil, zero value otherwise.

### GetClientTransactionIdOk

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) GetClientTransactionIdOk() (*string, bool)`

GetClientTransactionIdOk returns a tuple with the ClientTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) SetClientTransactionId(v string)`

SetClientTransactionId sets ClientTransactionId field to given value.

### HasClientTransactionId

`func (o *QuickbooksDesktopReceivePaymentsPostRequestCreditCardTransactionResponse) HasClientTransactionId() bool`

HasClientTransactionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


