# SalesRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | The name the sales team should address the reply to. It is sent as written and is not matched against any  portal account; an empty value fails the request with 400. | 
**Email** | **string** | The address the answer is sent to. It has to be a well-formed email address and need not be the caller portal  address; an empty or malformed value fails the request with 400. | 
**Message** | **string** | What is being asked of the sales team - a quote, an invoice, or a plan that cannot be bought online. An empty  value fails the request with 400. | 

## Methods

### NewSalesRequestDto

`func NewSalesRequestDto(userName string, email string, message string, ) *SalesRequestDto`

NewSalesRequestDto instantiates a new SalesRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesRequestDtoWithDefaults

`func NewSalesRequestDtoWithDefaults() *SalesRequestDto`

NewSalesRequestDtoWithDefaults instantiates a new SalesRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *SalesRequestDto) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SalesRequestDto) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SalesRequestDto) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetEmail

`func (o *SalesRequestDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SalesRequestDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SalesRequestDto) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetMessage

`func (o *SalesRequestDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SalesRequestDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SalesRequestDto) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


