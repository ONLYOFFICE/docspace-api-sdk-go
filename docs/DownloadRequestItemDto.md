# DownloadRequestItemDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**DownloadRequestItemDtoKey**](DownloadRequestItemDtoKey.md) |  | 
**Value** | **NullableString** | The target format or conversion type for the file download. | 
**Password** | Pointer to **NullableString** | The optional password for accessing protected files. | [optional] 

## Methods

### NewDownloadRequestItemDto

`func NewDownloadRequestItemDto(key DownloadRequestItemDtoKey, value NullableString, ) *DownloadRequestItemDto`

NewDownloadRequestItemDto instantiates a new DownloadRequestItemDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDownloadRequestItemDtoWithDefaults

`func NewDownloadRequestItemDtoWithDefaults() *DownloadRequestItemDto`

NewDownloadRequestItemDtoWithDefaults instantiates a new DownloadRequestItemDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DownloadRequestItemDto) GetKey() DownloadRequestItemDtoKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DownloadRequestItemDto) GetKeyOk() (*DownloadRequestItemDtoKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DownloadRequestItemDto) SetKey(v DownloadRequestItemDtoKey)`

SetKey sets Key field to given value.


### GetValue

`func (o *DownloadRequestItemDto) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DownloadRequestItemDto) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DownloadRequestItemDto) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *DownloadRequestItemDto) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *DownloadRequestItemDto) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetPassword

`func (o *DownloadRequestItemDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DownloadRequestItemDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DownloadRequestItemDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DownloadRequestItemDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *DownloadRequestItemDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *DownloadRequestItemDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


