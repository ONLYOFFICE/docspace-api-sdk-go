# PermissionsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **bool** | Defines if the document can be commented or not. | [optional] 
**Chat** | Pointer to **bool** | Defines if the chat functionality is enabled in the document or not. | [optional] 
**Download** | Pointer to **bool** | Defines if the document can be downloaded or only viewed or edited online. | [optional] 
**Edit** | Pointer to **bool** | Defines if the document can be edited or only viewed. | [optional] 
**FillForms** | Pointer to **bool** | Defines if the forms can be filled. | [optional] 
**ModifyFilter** | Pointer to **bool** | Defines if the filter can be applied globally (true) affecting all the other users,  or locally (false), i.e. for the current user only. | [optional] 
**Protect** | Pointer to **bool** | Defines if the Protection tab on the toolbar and the Protect button in the left menu are displayedor hidden. | [optional] 
**Print** | Pointer to **bool** | Defines if the document can be printed or not. | [optional] 
**Review** | Pointer to **bool** | Defines if the document can be reviewed or not. | [optional] 
**Copy** | Pointer to **bool** | Defines if the content can be copied to the clipboard or not. | [optional] 

## Methods

### NewPermissionsConfig

`func NewPermissionsConfig() *PermissionsConfig`

NewPermissionsConfig instantiates a new PermissionsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsConfigWithDefaults

`func NewPermissionsConfigWithDefaults() *PermissionsConfig`

NewPermissionsConfigWithDefaults instantiates a new PermissionsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *PermissionsConfig) GetComment() bool`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PermissionsConfig) GetCommentOk() (*bool, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PermissionsConfig) SetComment(v bool)`

SetComment sets Comment field to given value.

### HasComment

`func (o *PermissionsConfig) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetChat

`func (o *PermissionsConfig) GetChat() bool`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *PermissionsConfig) GetChatOk() (*bool, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *PermissionsConfig) SetChat(v bool)`

SetChat sets Chat field to given value.

### HasChat

`func (o *PermissionsConfig) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetDownload

`func (o *PermissionsConfig) GetDownload() bool`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *PermissionsConfig) GetDownloadOk() (*bool, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *PermissionsConfig) SetDownload(v bool)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *PermissionsConfig) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetEdit

`func (o *PermissionsConfig) GetEdit() bool`

GetEdit returns the Edit field if non-nil, zero value otherwise.

### GetEditOk

`func (o *PermissionsConfig) GetEditOk() (*bool, bool)`

GetEditOk returns a tuple with the Edit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdit

`func (o *PermissionsConfig) SetEdit(v bool)`

SetEdit sets Edit field to given value.

### HasEdit

`func (o *PermissionsConfig) HasEdit() bool`

HasEdit returns a boolean if a field has been set.

### GetFillForms

`func (o *PermissionsConfig) GetFillForms() bool`

GetFillForms returns the FillForms field if non-nil, zero value otherwise.

### GetFillFormsOk

`func (o *PermissionsConfig) GetFillFormsOk() (*bool, bool)`

GetFillFormsOk returns a tuple with the FillForms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillForms

`func (o *PermissionsConfig) SetFillForms(v bool)`

SetFillForms sets FillForms field to given value.

### HasFillForms

`func (o *PermissionsConfig) HasFillForms() bool`

HasFillForms returns a boolean if a field has been set.

### GetModifyFilter

`func (o *PermissionsConfig) GetModifyFilter() bool`

GetModifyFilter returns the ModifyFilter field if non-nil, zero value otherwise.

### GetModifyFilterOk

`func (o *PermissionsConfig) GetModifyFilterOk() (*bool, bool)`

GetModifyFilterOk returns a tuple with the ModifyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyFilter

`func (o *PermissionsConfig) SetModifyFilter(v bool)`

SetModifyFilter sets ModifyFilter field to given value.

### HasModifyFilter

`func (o *PermissionsConfig) HasModifyFilter() bool`

HasModifyFilter returns a boolean if a field has been set.

### GetProtect

`func (o *PermissionsConfig) GetProtect() bool`

GetProtect returns the Protect field if non-nil, zero value otherwise.

### GetProtectOk

`func (o *PermissionsConfig) GetProtectOk() (*bool, bool)`

GetProtectOk returns a tuple with the Protect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtect

`func (o *PermissionsConfig) SetProtect(v bool)`

SetProtect sets Protect field to given value.

### HasProtect

`func (o *PermissionsConfig) HasProtect() bool`

HasProtect returns a boolean if a field has been set.

### GetPrint

`func (o *PermissionsConfig) GetPrint() bool`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *PermissionsConfig) GetPrintOk() (*bool, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *PermissionsConfig) SetPrint(v bool)`

SetPrint sets Print field to given value.

### HasPrint

`func (o *PermissionsConfig) HasPrint() bool`

HasPrint returns a boolean if a field has been set.

### GetReview

`func (o *PermissionsConfig) GetReview() bool`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *PermissionsConfig) GetReviewOk() (*bool, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *PermissionsConfig) SetReview(v bool)`

SetReview sets Review field to given value.

### HasReview

`func (o *PermissionsConfig) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetCopy

`func (o *PermissionsConfig) GetCopy() bool`

GetCopy returns the Copy field if non-nil, zero value otherwise.

### GetCopyOk

`func (o *PermissionsConfig) GetCopyOk() (*bool, bool)`

GetCopyOk returns a tuple with the Copy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopy

`func (o *PermissionsConfig) SetCopy(v bool)`

SetCopy sets Copy field to given value.

### HasCopy

`func (o *PermissionsConfig) HasCopy() bool`

HasCopy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


