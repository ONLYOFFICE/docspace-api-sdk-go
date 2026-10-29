# SendModuleEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Room** | Pointer to **NullableString** | A room of the module in this portal: everyone subscribed to `plugin-{module}-{room}` receives the event. Latin  letters, digits and the characters `_ . : -`. | [optional] 
**User** | Pointer to **NullableString** | The e-mail of the portal user who receives the event; with single sign-on it is the same e-mail the module's own  application knows the person by. Only that user's open sessions receive it. | [optional] 
**Event** | Pointer to **NullableString** | The event name. Browsers receive it as `plugin:{module}:{event}`. Latin letters, digits and the characters  `_ . : -`. | [optional] 
**Message** | Pointer to **NullableString** | Passed to the browser as it is, as `{ message }`. Meant as a signal to re-read, such as an unread count; use  JSON text for structured data and keep personal content out of it. | [optional] 

## Methods

### NewSendModuleEventRequest

`func NewSendModuleEventRequest() *SendModuleEventRequest`

NewSendModuleEventRequest instantiates a new SendModuleEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendModuleEventRequestWithDefaults

`func NewSendModuleEventRequestWithDefaults() *SendModuleEventRequest`

NewSendModuleEventRequestWithDefaults instantiates a new SendModuleEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoom

`func (o *SendModuleEventRequest) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *SendModuleEventRequest) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *SendModuleEventRequest) SetRoom(v string)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *SendModuleEventRequest) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### SetRoomNil

`func (o *SendModuleEventRequest) SetRoomNil(b bool)`

 SetRoomNil sets the value for Room to be an explicit nil

### UnsetRoom
`func (o *SendModuleEventRequest) UnsetRoom()`

UnsetRoom ensures that no value is present for Room, not even an explicit nil
### GetUser

`func (o *SendModuleEventRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SendModuleEventRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SendModuleEventRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SendModuleEventRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SendModuleEventRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SendModuleEventRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetEvent

`func (o *SendModuleEventRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *SendModuleEventRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *SendModuleEventRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *SendModuleEventRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *SendModuleEventRequest) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *SendModuleEventRequest) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetMessage

`func (o *SendModuleEventRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendModuleEventRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendModuleEventRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendModuleEventRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SendModuleEventRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SendModuleEventRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


