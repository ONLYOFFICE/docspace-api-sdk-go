# MentionMessageWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionLink** | Pointer to [**ActionLinkConfig**](ActionLinkConfig.md) | The place in the document the notification link should open at, as the editor reports it when the mention is  made. Left out, the link opens the file at its beginning. | [optional] 
**Emails** | Pointer to **[]string** | The addresses to notify. Only an address that belongs to a portal account receives a mail; an unknown address  is skipped, and the answer then carries the access list of the file so that the client can invite its owner. | [optional] 
**Message** | Pointer to **NullableString** | The note shown next to the link in the mail. Only its first 200 characters are sent, and a value longer than  the field allows is refused. | [optional] 

## Methods

### NewMentionMessageWrapper

`func NewMentionMessageWrapper() *MentionMessageWrapper`

NewMentionMessageWrapper instantiates a new MentionMessageWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionMessageWrapperWithDefaults

`func NewMentionMessageWrapperWithDefaults() *MentionMessageWrapper`

NewMentionMessageWrapperWithDefaults instantiates a new MentionMessageWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionLink

`func (o *MentionMessageWrapper) GetActionLink() ActionLinkConfig`

GetActionLink returns the ActionLink field if non-nil, zero value otherwise.

### GetActionLinkOk

`func (o *MentionMessageWrapper) GetActionLinkOk() (*ActionLinkConfig, bool)`

GetActionLinkOk returns a tuple with the ActionLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionLink

`func (o *MentionMessageWrapper) SetActionLink(v ActionLinkConfig)`

SetActionLink sets ActionLink field to given value.

### HasActionLink

`func (o *MentionMessageWrapper) HasActionLink() bool`

HasActionLink returns a boolean if a field has been set.

### GetEmails

`func (o *MentionMessageWrapper) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *MentionMessageWrapper) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *MentionMessageWrapper) SetEmails(v []string)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *MentionMessageWrapper) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### SetEmailsNil

`func (o *MentionMessageWrapper) SetEmailsNil(b bool)`

 SetEmailsNil sets the value for Emails to be an explicit nil

### UnsetEmails
`func (o *MentionMessageWrapper) UnsetEmails()`

UnsetEmails ensures that no value is present for Emails, not even an explicit nil
### GetMessage

`func (o *MentionMessageWrapper) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MentionMessageWrapper) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MentionMessageWrapper) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MentionMessageWrapper) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *MentionMessageWrapper) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *MentionMessageWrapper) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


