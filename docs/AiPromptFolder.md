# AiPromptFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique folder identifier (UUID). | 
**Name** | **string** | Folder display name. | 
**CreatedAt** | **float32** | Timestamp (ms since epoch) when the folder was created. | 
**UpdatedAt** | **float32** | Timestamp (ms since epoch) of the last folder modification. | 

## Methods

### NewAiPromptFolder

`func NewAiPromptFolder(id string, name string, createdAt float32, updatedAt float32, ) *AiPromptFolder`

NewAiPromptFolder instantiates a new AiPromptFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPromptFolderWithDefaults

`func NewAiPromptFolderWithDefaults() *AiPromptFolder`

NewAiPromptFolderWithDefaults instantiates a new AiPromptFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AiPromptFolder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiPromptFolder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiPromptFolder) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AiPromptFolder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiPromptFolder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiPromptFolder) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *AiPromptFolder) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AiPromptFolder) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AiPromptFolder) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AiPromptFolder) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AiPromptFolder) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AiPromptFolder) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


