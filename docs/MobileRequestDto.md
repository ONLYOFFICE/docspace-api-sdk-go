# MobileRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobilePhone** | Pointer to **NullableString** | The number the SMS codes are sent to, in international form with the leading `+` and no spaces. It is stored  as not yet activated and only becomes the confirmed number once a code sent to it is accepted; an already  activated number is not replaced this way and has to be erased first. | [optional] 

## Methods

### NewMobileRequestDto

`func NewMobileRequestDto() *MobileRequestDto`

NewMobileRequestDto instantiates a new MobileRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileRequestDtoWithDefaults

`func NewMobileRequestDtoWithDefaults() *MobileRequestDto`

NewMobileRequestDtoWithDefaults instantiates a new MobileRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobilePhone

`func (o *MobileRequestDto) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *MobileRequestDto) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *MobileRequestDto) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *MobileRequestDto) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### SetMobilePhoneNil

`func (o *MobileRequestDto) SetMobilePhoneNil(b bool)`

 SetMobilePhoneNil sets the value for MobilePhone to be an explicit nil

### UnsetMobilePhone
`func (o *MobileRequestDto) UnsetMobilePhone()`

UnsetMobilePhone ensures that no value is present for MobilePhone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


