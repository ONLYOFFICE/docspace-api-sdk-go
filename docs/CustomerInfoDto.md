# CustomerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortalId** | Pointer to **NullableString** | The portal ID. | [optional] [readonly] 
**PaymentMethodStatus** | Pointer to [**PaymentMethodStatus**](PaymentMethodStatus.md) | The customer's payment method. | [optional] 
**Email** | Pointer to **NullableString** | The customer email address. | [optional] [readonly] 
**Payer** | Pointer to [**EmployeeDto**](EmployeeDto.md) | The paying user. | [optional] 

## Methods

### NewCustomerInfoDto

`func NewCustomerInfoDto() *CustomerInfoDto`

NewCustomerInfoDto instantiates a new CustomerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerInfoDtoWithDefaults

`func NewCustomerInfoDtoWithDefaults() *CustomerInfoDto`

NewCustomerInfoDtoWithDefaults instantiates a new CustomerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortalId

`func (o *CustomerInfoDto) GetPortalId() string`

GetPortalId returns the PortalId field if non-nil, zero value otherwise.

### GetPortalIdOk

`func (o *CustomerInfoDto) GetPortalIdOk() (*string, bool)`

GetPortalIdOk returns a tuple with the PortalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalId

`func (o *CustomerInfoDto) SetPortalId(v string)`

SetPortalId sets PortalId field to given value.

### HasPortalId

`func (o *CustomerInfoDto) HasPortalId() bool`

HasPortalId returns a boolean if a field has been set.

### SetPortalIdNil

`func (o *CustomerInfoDto) SetPortalIdNil(b bool)`

 SetPortalIdNil sets the value for PortalId to be an explicit nil

### UnsetPortalId
`func (o *CustomerInfoDto) UnsetPortalId()`

UnsetPortalId ensures that no value is present for PortalId, not even an explicit nil
### GetPaymentMethodStatus

`func (o *CustomerInfoDto) GetPaymentMethodStatus() PaymentMethodStatus`

GetPaymentMethodStatus returns the PaymentMethodStatus field if non-nil, zero value otherwise.

### GetPaymentMethodStatusOk

`func (o *CustomerInfoDto) GetPaymentMethodStatusOk() (*PaymentMethodStatus, bool)`

GetPaymentMethodStatusOk returns a tuple with the PaymentMethodStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodStatus

`func (o *CustomerInfoDto) SetPaymentMethodStatus(v PaymentMethodStatus)`

SetPaymentMethodStatus sets PaymentMethodStatus field to given value.

### HasPaymentMethodStatus

`func (o *CustomerInfoDto) HasPaymentMethodStatus() bool`

HasPaymentMethodStatus returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerInfoDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerInfoDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerInfoDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerInfoDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CustomerInfoDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CustomerInfoDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPayer

`func (o *CustomerInfoDto) GetPayer() EmployeeDto`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *CustomerInfoDto) GetPayerOk() (*EmployeeDto, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *CustomerInfoDto) SetPayer(v EmployeeDto)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *CustomerInfoDto) HasPayer() bool`

HasPayer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


