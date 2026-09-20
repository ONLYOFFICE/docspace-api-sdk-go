# NotificationSettingsRequestsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NotificationType**](NotificationType.md) | The kind of notification being switched. A value outside the defined set is echoed back while nothing is  stored, so confirm the result with `GET api/2.0/settings/notification/{type}` rather than trusting the  answer. | 
**IsEnabled** | Pointer to **bool** | Whether that kind reaches the calling account. It applies to the caller own account alone and to every room  at once; a single room is silenced with `POST api/2.0/settings/notification/rooms` instead. | [optional] 

## Methods

### NewNotificationSettingsRequestsDto

`func NewNotificationSettingsRequestsDto(type_ NotificationType, ) *NotificationSettingsRequestsDto`

NewNotificationSettingsRequestsDto instantiates a new NotificationSettingsRequestsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsRequestsDtoWithDefaults

`func NewNotificationSettingsRequestsDtoWithDefaults() *NotificationSettingsRequestsDto`

NewNotificationSettingsRequestsDtoWithDefaults instantiates a new NotificationSettingsRequestsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationSettingsRequestsDto) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationSettingsRequestsDto) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationSettingsRequestsDto) SetType(v NotificationType)`

SetType sets Type field to given value.


### GetIsEnabled

`func (o *NotificationSettingsRequestsDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationSettingsRequestsDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationSettingsRequestsDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationSettingsRequestsDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


