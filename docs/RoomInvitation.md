# RoomInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** | The email address. | [optional] 
**Id** | Pointer to **string** | The ID of the user to share a room with. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) |  | [optional] 

## Methods

### NewRoomInvitation

`func NewRoomInvitation() *RoomInvitation`

NewRoomInvitation instantiates a new RoomInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomInvitationWithDefaults

`func NewRoomInvitationWithDefaults() *RoomInvitation`

NewRoomInvitationWithDefaults instantiates a new RoomInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RoomInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RoomInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RoomInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RoomInvitation) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *RoomInvitation) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *RoomInvitation) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *RoomInvitation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoomInvitation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoomInvitation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RoomInvitation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccess

`func (o *RoomInvitation) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *RoomInvitation) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *RoomInvitation) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *RoomInvitation) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


