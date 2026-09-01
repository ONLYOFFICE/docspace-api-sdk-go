# NotificationSettingsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**NotificationType**](NotificationType.md) | The notification type. | [optional] 
**IsEnabled** | Pointer to **bool** | Specifies if the notification type is enabled or not. | [optional] 

## Methods

### NewNotificationSettingsDto

`func NewNotificationSettingsDto() *NotificationSettingsDto`

NewNotificationSettingsDto instantiates a new NotificationSettingsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationSettingsDtoWithDefaults

`func NewNotificationSettingsDtoWithDefaults() *NotificationSettingsDto`

NewNotificationSettingsDtoWithDefaults instantiates a new NotificationSettingsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationSettingsDto) GetType() NotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationSettingsDto) GetTypeOk() (*NotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationSettingsDto) SetType(v NotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *NotificationSettingsDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsEnabled

`func (o *NotificationSettingsDto) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *NotificationSettingsDto) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *NotificationSettingsDto) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *NotificationSettingsDto) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


