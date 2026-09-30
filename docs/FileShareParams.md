# FileShareParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The address of somebody who has no portal account yet. An invitation is sent to it and an account is created  once it is accepted, so this is the field to use instead of an account identifier when the person is new to  the portal. | [optional] 
**ShareTo** | Pointer to **string** | The account or the group the entry is about, taken from the portal people and group listings. Leave it out and  give an email address instead to share with somebody who has no account yet. | [optional] 
**Access** | Pointer to [**FileShare**](FileShare.md) | What the subject may do with the shared item. The value 0 takes the access away again, and which of the other  levels are accepted depends on what is being shared. | [optional] 

## Methods

### NewFileShareParams

`func NewFileShareParams() *FileShareParams`

NewFileShareParams instantiates a new FileShareParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileShareParamsWithDefaults

`func NewFileShareParamsWithDefaults() *FileShareParams`

NewFileShareParamsWithDefaults instantiates a new FileShareParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *FileShareParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FileShareParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FileShareParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FileShareParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetShareTo

`func (o *FileShareParams) GetShareTo() string`

GetShareTo returns the ShareTo field if non-nil, zero value otherwise.

### GetShareToOk

`func (o *FileShareParams) GetShareToOk() (*string, bool)`

GetShareToOk returns a tuple with the ShareTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareTo

`func (o *FileShareParams) SetShareTo(v string)`

SetShareTo sets ShareTo field to given value.

### HasShareTo

`func (o *FileShareParams) HasShareTo() bool`

HasShareTo returns a boolean if a field has been set.

### GetAccess

`func (o *FileShareParams) GetAccess() FileShare`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FileShareParams) GetAccessOk() (*FileShare, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FileShareParams) SetAccess(v FileShare)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *FileShareParams) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


