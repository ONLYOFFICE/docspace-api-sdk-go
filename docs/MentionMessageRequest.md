# MentionMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionLink** | Pointer to [**ActionLinkRequest**](ActionLinkRequest.md) | The place in the document the notification link should open at, as the editor reports it when the mention is  made. Left out, the link opens the file at its beginning. | [optional] 
**Emails** | Pointer to **[]string** | The addresses to notify. Only an address that belongs to a portal account receives a mail; an unknown address  is skipped, and the answer then carries the access list of the file so that the client can invite its owner. | [optional] 
**Message** | Pointer to **NullableString** | The note shown next to the link in the mail. Only its first 200 characters are sent, and a value longer than  the field allows is refused. | [optional] 

## Methods

### NewMentionMessageRequest

`func NewMentionMessageRequest() *MentionMessageRequest`

NewMentionMessageRequest instantiates a new MentionMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionMessageRequestWithDefaults

`func NewMentionMessageRequestWithDefaults() *MentionMessageRequest`

NewMentionMessageRequestWithDefaults instantiates a new MentionMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionLink

`func (o *MentionMessageRequest) GetActionLink() ActionLinkRequest`

GetActionLink returns the ActionLink field if non-nil, zero value otherwise.

### GetActionLinkOk

`func (o *MentionMessageRequest) GetActionLinkOk() (*ActionLinkRequest, bool)`

GetActionLinkOk returns a tuple with the ActionLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionLink

`func (o *MentionMessageRequest) SetActionLink(v ActionLinkRequest)`

SetActionLink sets ActionLink field to given value.

### HasActionLink

`func (o *MentionMessageRequest) HasActionLink() bool`

HasActionLink returns a boolean if a field has been set.

### GetEmails

`func (o *MentionMessageRequest) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MentionMessageRequest) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MentionMessageRequest) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MentionMessageRequest) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *MentionMessageRequest) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *MentionMessageRequest) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetMessage

`func (o *MentionMessageRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MentionMessageRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MentionMessageRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MentionMessageRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MentionMessageRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MentionMessageRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


