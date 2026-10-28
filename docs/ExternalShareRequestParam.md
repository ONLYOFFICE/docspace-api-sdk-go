# ExternalShareRequestParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **NullableString** | The password chosen by the member who shared the entry, spelled exactly as they typed it. It is compared  against the stored value and never returned back; a mismatch is reported through the answer's status instead  of an error. | [optional] 

## Methods

### NewExternalShareRequestParam

`func NewExternalShareRequestParam() *ExternalShareRequestParam`

NewExternalShareRequestParam instantiates a new ExternalShareRequestParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalShareRequestParamWithDefaults

`func NewExternalShareRequestParamWithDefaults() *ExternalShareRequestParam`

NewExternalShareRequestParamWithDefaults instantiates a new ExternalShareRequestParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ExternalShareRequestParam) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExternalShareRequestParam) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExternalShareRequestParam) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ExternalShareRequestParam) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ExternalShareRequestParam) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ExternalShareRequestParam) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


