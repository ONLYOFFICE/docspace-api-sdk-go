# OwnerChangeInstructionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | Whether the letter was sent: `1` that it was, `0` that the request was turned down. A refusal comes back  with HTTP 200, so this field and not the status code is what says whether anything happened - the request  is turned down when the caller is not the portal owner and when the named member is unknown or inactive. | [optional] 
**Message** | Pointer to **NullableString** | The outcome spelled out in the portal language. On success it names the address the letter went to, and it  carries an HTML `mailto:` anchor rather than plain text, so it has to be rendered as markup or stripped;  on a refusal it is the localised reason. | [optional] 

## Methods

### NewOwnerChangeInstructionsDto

`func NewOwnerChangeInstructionsDto() *OwnerChangeInstructionsDto`

NewOwnerChangeInstructionsDto instantiates a new OwnerChangeInstructionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerChangeInstructionsDtoWithDefaults

`func NewOwnerChangeInstructionsDtoWithDefaults() *OwnerChangeInstructionsDto`

NewOwnerChangeInstructionsDtoWithDefaults instantiates a new OwnerChangeInstructionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OwnerChangeInstructionsDto) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OwnerChangeInstructionsDto) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OwnerChangeInstructionsDto) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OwnerChangeInstructionsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *OwnerChangeInstructionsDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OwnerChangeInstructionsDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OwnerChangeInstructionsDto) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OwnerChangeInstructionsDto) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *OwnerChangeInstructionsDto) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *OwnerChangeInstructionsDto) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


