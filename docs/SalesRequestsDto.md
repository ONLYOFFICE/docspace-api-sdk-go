# SalesRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | Pointer to **NullableString** | The name of the user submitting the sales request. | [optional] 
**Email** | **NullableString** | The contact email address for the sales inquiry. | 
**Message** | **NullableString** | The details of the sales inquiry or payment request. | 

## Methods

### NewSalesRequestsDto

`func NewSalesRequestsDto(email NullableString, message NullableString, ) *SalesRequestsDto`

NewSalesRequestsDto instantiates a new SalesRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesRequestsDtoWithDefaults

`func NewSalesRequestsDtoWithDefaults() *SalesRequestsDto`

NewSalesRequestsDtoWithDefaults instantiates a new SalesRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *SalesRequestsDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SalesRequestsDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SalesRequestsDto) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SalesRequestsDto) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *SalesRequestsDto) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *SalesRequestsDto) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetEmail

`func (o *SalesRequestsDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SalesRequestsDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SalesRequestsDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *SalesRequestsDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *SalesRequestsDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetMessage

`func (o *SalesRequestsDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SalesRequestsDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SalesRequestsDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *SalesRequestsDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SalesRequestsDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


