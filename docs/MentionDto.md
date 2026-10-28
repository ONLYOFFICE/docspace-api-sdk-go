# MentionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**PortalUserDto**](PortalUserDto.md) | The account itself, as the portal stores it. | [optional] 
**Email** | Pointer to **NullableString** | Where a mention notification for this user is delivered. | [optional] 
**Id** | Pointer to **NullableString** | The account id as text, the same value the account object carries; it is what identifies the user in a sharing  request built from this list. | [optional] 
**Image** | Pointer to **NullableString** | An absolute address of the medium-sized avatar. A generated default avatar is reported when the user never  uploaded one, so the field is never empty. | [optional] 
**HasAccess** | Pointer to **bool** | Not filled in by the operations that return this list: it always comes back false. Whether a user can already  open the document has to be read from the sharing settings of the file. | [optional] 
**Name** | Pointer to **NullableString** | The name to display, assembled the way the portal is configured to show names. | [optional] 

## Methods

### NewMentionDto

`func NewMentionDto() *MentionDto`

NewMentionDto instantiates a new MentionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionDtoWithDefaults

`func NewMentionDtoWithDefaults() *MentionDto`

NewMentionDtoWithDefaults instantiates a new MentionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *MentionDto) GetUser() PortalUserDto`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MentionDto) GetUserOk() (*PortalUserDto, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MentionDto) SetUser(v PortalUserDto)`

SetUser sets User field to given value.

### HasUser

`func (o *MentionDto) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetEmail

`func (o *MentionDto) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MentionDto) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MentionDto) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MentionDto) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MentionDto) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MentionDto) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *MentionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MentionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MentionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MentionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MentionDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MentionDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetImage

`func (o *MentionDto) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *MentionDto) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *MentionDto) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *MentionDto) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *MentionDto) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *MentionDto) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetHasAccess

`func (o *MentionDto) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *MentionDto) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *MentionDto) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *MentionDto) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetName

`func (o *MentionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MentionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MentionDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MentionDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MentionDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MentionDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


