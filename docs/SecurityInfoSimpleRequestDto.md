# SecurityInfoSimpleRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Share** | Pointer to [**[]FileShareParams**](FileShareParams.md) | One record per account or group whose rights are being set, each naming the subject and the level it gets; a  level of `None` takes the access away. An empty collection makes the call change nothing. | [optional] 
**Notify** | Pointer to **bool** | Set to true to have every account named in `share` emailed about the access it just received; false changes  the rights without telling anyone. | [optional] 
**SharingMessage** | Pointer to **NullableString** | The text put into that email, ignored while `notify` is false. Markup is stripped before sending, so only the  plain text of the value survives. | [optional] 

## Methods

### NewSecurityInfoSimpleRequestDto

`func NewSecurityInfoSimpleRequestDto() *SecurityInfoSimpleRequestDto`

NewSecurityInfoSimpleRequestDto instantiates a new SecurityInfoSimpleRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityInfoSimpleRequestDtoWithDefaults

`func NewSecurityInfoSimpleRequestDtoWithDefaults() *SecurityInfoSimpleRequestDto`

NewSecurityInfoSimpleRequestDtoWithDefaults instantiates a new SecurityInfoSimpleRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShare

`func (o *SecurityInfoSimpleRequestDto) GetShare() []FileShareParams`

GetShare returns the Share field if non-nil, zero value otherwise.

### GetShareOk

`func (o *SecurityInfoSimpleRequestDto) GetShareOk() (*[]FileShareParams, bool)`

GetShareOk returns a tuple with the Share field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShare

`func (o *SecurityInfoSimpleRequestDto) SetShare(v []FileShareParams)`

SetShare sets Share field to given value.

### HasShare

`func (o *SecurityInfoSimpleRequestDto) HasShare() bool`

HasShare returns a boolean if a field has been set.

### SetShareNil

`func (o *SecurityInfoSimpleRequestDto) SetShareNil(b bool)`

 SetShareNil sets the value for Share to be an explicit nil

### UnsetShare
`func (o *SecurityInfoSimpleRequestDto) UnsetShare()`

UnsetShare ensures that no value is present for Share, not even an explicit nil
### GetNotify

`func (o *SecurityInfoSimpleRequestDto) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *SecurityInfoSimpleRequestDto) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *SecurityInfoSimpleRequestDto) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *SecurityInfoSimpleRequestDto) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetSharingMessage

`func (o *SecurityInfoSimpleRequestDto) GetSharingMessage() string`

GetSharingMessage returns the SharingMessage field if non-nil, zero value otherwise.

### GetSharingMessageOk

`func (o *SecurityInfoSimpleRequestDto) GetSharingMessageOk() (*string, bool)`

GetSharingMessageOk returns a tuple with the SharingMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingMessage

`func (o *SecurityInfoSimpleRequestDto) SetSharingMessage(v string)`

SetSharingMessage sets SharingMessage field to given value.

### HasSharingMessage

`func (o *SecurityInfoSimpleRequestDto) HasSharingMessage() bool`

HasSharingMessage returns a boolean if a field has been set.

### SetSharingMessageNil

`func (o *SecurityInfoSimpleRequestDto) SetSharingMessageNil(b bool)`

 SetSharingMessageNil sets the value for SharingMessage to be an explicit nil

### UnsetSharingMessage
`func (o *SecurityInfoSimpleRequestDto) UnsetSharingMessage()`

UnsetSharingMessage ensures that no value is present for SharingMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


