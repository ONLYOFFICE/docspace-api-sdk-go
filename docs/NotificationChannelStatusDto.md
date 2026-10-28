# NotificationChannelStatusDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]NotificationChannelDto**](NotificationChannelDto.md) | The channels the running installation is configured with. A channel appears only when the notification  service names a sender for it, so the list can be shorter than the channels this build implements, and an  empty list means the configuration names none of them. | [optional] 

## Methods

### NewNotificationChannelStatusDto

`func NewNotificationChannelStatusDto() *NotificationChannelStatusDto`

NewNotificationChannelStatusDto instantiates a new NotificationChannelStatusDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelStatusDtoWithDefaults

`func NewNotificationChannelStatusDtoWithDefaults() *NotificationChannelStatusDto`

NewNotificationChannelStatusDtoWithDefaults instantiates a new NotificationChannelStatusDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *NotificationChannelStatusDto) GetChannels() []NotificationChannelDto`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *NotificationChannelStatusDto) GetChannelsOk() (*[]NotificationChannelDto, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *NotificationChannelStatusDto) SetChannels(v []NotificationChannelDto)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *NotificationChannelStatusDto) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### SetChannelsNil

`func (o *NotificationChannelStatusDto) SetChannelsNil(b bool)`

 SetChannelsNil sets the value for Channels to be an explicit nil

### UnsetChannels
`func (o *NotificationChannelStatusDto) UnsetChannels()`

UnsetChannels ensures that no value is present for Channels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


