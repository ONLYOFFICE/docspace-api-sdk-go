# AiContextFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloud** | **string** | The account the rooms belong to; `docspace` for this portal. | 
**Rooms** | [**[]AiContextRoom**](AiContextRoom.md) | The rooms of that account the caller can connect. | 

## Methods

### NewAiContextFolder

`func NewAiContextFolder(cloud string, rooms []AiContextRoom, ) *AiContextFolder`

NewAiContextFolder instantiates a new AiContextFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiContextFolderWithDefaults

`func NewAiContextFolderWithDefaults() *AiContextFolder`

NewAiContextFolderWithDefaults instantiates a new AiContextFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloud

`func (o *AiContextFolder) GetCloud() string`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AiContextFolder) GetCloudOk() (*string, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AiContextFolder) SetCloud(v string)`

SetCloud sets Cloud field to given value.


### GetRooms

`func (o *AiContextFolder) GetRooms() []AiContextRoom`

GetRooms returns the Rooms field if non-nil, zero value otherwise.

### GetRoomsOk

`func (o *AiContextFolder) GetRoomsOk() (*[]AiContextRoom, bool)`

GetRoomsOk returns a tuple with the Rooms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRooms

`func (o *AiContextFolder) SetRooms(v []AiContextRoom)`

SetRooms sets Rooms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


