# CopyAsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestTitle** | **NullableString** | The title of the copy, extension included. That extension decides the format: the same one as the source  copies the content as it is, a different one has it converted first. | 
**DestFolderId** | [**CopyAsRequestDestFolderId**](CopyAsRequestDestFolderId.md) |  | 
**EnableExternalExt** | Pointer to **bool** | Whether the extension of the new title may be one the portal does not edit itself. | [optional] 
**Password** | Pointer to **NullableString** | The password that opens the source document, for a file that is protected by one. | [optional] 
**ToForm** | Pointer to **bool** | Whether the copy is to become a PDF form rather than a plain document, which the conversion supports for the  text formats it can read. | [optional] 

## Methods

### NewCopyAsRequest

`func NewCopyAsRequest(destTitle NullableString, destFolderId CopyAsRequestDestFolderId, ) *CopyAsRequest`

NewCopyAsRequest instantiates a new CopyAsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyAsRequestWithDefaults

`func NewCopyAsRequestWithDefaults() *CopyAsRequest`

NewCopyAsRequestWithDefaults instantiates a new CopyAsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestTitle

`func (o *CopyAsRequest) GetDestTitle() string`

GetDestTitle returns the DestTitle field if non-nil, zero value otherwise.

### GetDestTitleOk

`func (o *CopyAsRequest) GetDestTitleOk() (*string, bool)`

GetDestTitleOk returns a tuple with the DestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestTitle

`func (o *CopyAsRequest) SetDestTitle(v string)`

SetDestTitle sets DestTitle field to given value.


### SetDestTitleNil

`func (o *CopyAsRequest) SetDestTitleNil(b bool)`

 SetDestTitleNil sets the value for DestTitle to be an explicit nil

### UnsetDestTitle
`func (o *CopyAsRequest) UnsetDestTitle()`

UnsetDestTitle ensures that no value is present for DestTitle, not even an explicit nil
### GetDestFolderId

`func (o *CopyAsRequest) GetDestFolderId() CopyAsRequestDestFolderId`

GetDestFolderId returns the DestFolderId field if non-nil, zero value otherwise.

### GetDestFolderIdOk

`func (o *CopyAsRequest) GetDestFolderIdOk() (*CopyAsRequestDestFolderId, bool)`

GetDestFolderIdOk returns a tuple with the DestFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestFolderId

`func (o *CopyAsRequest) SetDestFolderId(v CopyAsRequestDestFolderId)`

SetDestFolderId sets DestFolderId field to given value.


### GetEnableExternalExt

`func (o *CopyAsRequest) GetEnableExternalExt() bool`

GetEnableExternalExt returns the EnableExternalExt field if non-nil, zero value otherwise.

### GetEnableExternalExtOk

`func (o *CopyAsRequest) GetEnableExternalExtOk() (*bool, bool)`

GetEnableExternalExtOk returns a tuple with the EnableExternalExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalExt

`func (o *CopyAsRequest) SetEnableExternalExt(v bool)`

SetEnableExternalExt sets EnableExternalExt field to given value.

### HasEnableExternalExt

`func (o *CopyAsRequest) HasEnableExternalExt() bool`

HasEnableExternalExt returns a boolean if a field has been set.

### GetPassword

`func (o *CopyAsRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CopyAsRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CopyAsRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CopyAsRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CopyAsRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CopyAsRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetToForm

`func (o *CopyAsRequest) GetToForm() bool`

GetToForm returns the ToForm field if non-nil, zero value otherwise.

### GetToFormOk

`func (o *CopyAsRequest) GetToFormOk() (*bool, bool)`

GetToFormOk returns a tuple with the ToForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToForm

`func (o *CopyAsRequest) SetToForm(v bool)`

SetToForm sets ToForm field to given value.

### HasToForm

`func (o *CopyAsRequest) HasToForm() bool`

HasToForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


