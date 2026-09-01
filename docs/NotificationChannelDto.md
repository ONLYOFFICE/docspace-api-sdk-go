# NotificationChannelDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The notification channel name. | 
**IsEnabled** | **bool** | Specifies whether the notification channel is enabled. | 

## Methods

### NewNotificationChannelDto

`func NewNotificationChannelDto(name NullableString, isEnabled bool, ) *NotificationChannelDto`

NewNotificationChannelDto instantiates a new NotificationChannelDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelDtoWithDefaults

`func NewNotificationChannelDtoWithDefaults() *NotificationChannelDto`

NewNotificationChannelDtoWithDefaults instantiates a new NotificationChannelDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationChannelDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationChannelDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationChannelDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *NotificationChannelDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *NotificationChannelDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsEnabled

`func (o *NotificationChannelDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationChannelDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationChannelDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


