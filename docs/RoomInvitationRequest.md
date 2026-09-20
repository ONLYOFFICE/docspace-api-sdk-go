# RoomInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | Pointer to [**[]RoomInvitation**](RoomInvitation.md) | Who is added, changed or removed, one entry per subject. The same subject named twice keeps the level of the  last entry, and an empty list is accepted and changes nothing. | [optional] 
**Notify** | Pointer to **bool** | Whether the subjects that gained access are told about it by email. With it off the change is silent, which is  the usual choice when membership is synchronised from another system. | [optional] 
**Message** | Pointer to **NullableString** | The line added to the invitation email. It is used only while the notification is on, and it reaches nobody  whose access was removed. | [optional] 
**Culture** | Pointer to **NullableString** | The language of the invitation email, as a portal culture name such as en-US. Leaving it out sends each  message in the language of its recipient. | [optional] 
**Force** | Pointer to **bool** | Whether a member who still holds a role in an unfinished form is removed anyway. With it off such a removal is  refused and reported through the error of the answer, so the form can be reassigned first. | [optional] 

## Methods

### NewRoomInvitationRequest

`func NewRoomInvitationRequest() *RoomInvitationRequest`

NewRoomInvitationRequest instantiates a new RoomInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomInvitationRequestWithDefaults

`func NewRoomInvitationRequestWithDefaults() *RoomInvitationRequest`

NewRoomInvitationRequestWithDefaults instantiates a new RoomInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *RoomInvitationRequest) GetInvitations() []RoomInvitation`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *RoomInvitationRequest) GetInvitationsOk() (*[]RoomInvitation, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *RoomInvitationRequest) SetInvitations(v []RoomInvitation)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *RoomInvitationRequest) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### SetInvitationsNil

`func (o *RoomInvitationRequest) SetInvitationsNil(b bool)`

 SetInvitationsNil sets the value for Invitations to be an explicit nil

### UnsetInvitations
`func (o *RoomInvitationRequest) UnsetInvitations()`

UnsetInvitations ensures that no value is present for Invitations, not even an explicit nil
### GetNotify

`func (o *RoomInvitationRequest) GetNotify() bool`

GetNotify returns the Notify field if non-nil, zero value otherwise.

### GetNotifyOk

`func (o *RoomInvitationRequest) GetNotifyOk() (*bool, bool)`

GetNotifyOk returns a tuple with the Notify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotify

`func (o *RoomInvitationRequest) SetNotify(v bool)`

SetNotify sets Notify field to given value.

### HasNotify

`func (o *RoomInvitationRequest) HasNotify() bool`

HasNotify returns a boolean if a field has been set.

### GetMessage

`func (o *RoomInvitationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RoomInvitationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RoomInvitationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RoomInvitationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *RoomInvitationRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *RoomInvitationRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCulture

`func (o *RoomInvitationRequest) GetCulture() string`

GetCulture returns the Culture field if non-nil, zero value otherwise.

### GetCultureOk

`func (o *RoomInvitationRequest) GetCultureOk() (*string, bool)`

GetCultureOk returns a tuple with the Culture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCulture

`func (o *RoomInvitationRequest) SetCulture(v string)`

SetCulture sets Culture field to given value.

### HasCulture

`func (o *RoomInvitationRequest) HasCulture() bool`

HasCulture returns a boolean if a field has been set.

### SetCultureNil

`func (o *RoomInvitationRequest) SetCultureNil(b bool)`

 SetCultureNil sets the value for Culture to be an explicit nil

### UnsetCulture
`func (o *RoomInvitationRequest) UnsetCulture()`

UnsetCulture ensures that no value is present for Culture, not even an explicit nil
### GetForce

`func (o *RoomInvitationRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *RoomInvitationRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *RoomInvitationRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *RoomInvitationRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


