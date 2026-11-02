# AiSkillsWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | Pointer to [**AiSkillsDto**](AiSkillsDto.md) | The AiSkillsDto object returned by the operation. | [optional] 
**Count** | Pointer to **int32** | How many items this response carries: the length of the page for a list, 1 for a single object | [optional] 
**Total** | Pointer to **int64** | How many items the whole list holds across all pages, for working out how many pages to request. Sent only by the operations that page their results, and left out of every other response | [optional] 
**Links** | Pointer to [**[]GetPortalPrices200ResponseLinksInner**](GetPortalPrices200ResponseLinksInner.md) | List of links related to the response | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response | [optional] 
**StatusCode** | Pointer to **int32** | HTTP status code of the response (duplicate of status) | [optional] 

## Methods

### NewAiSkillsWrapper

`func NewAiSkillsWrapper() *AiSkillsWrapper`

NewAiSkillsWrapper instantiates a new AiSkillsWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSkillsWrapperWithDefaults

`func NewAiSkillsWrapperWithDefaults() *AiSkillsWrapper`

NewAiSkillsWrapperWithDefaults instantiates a new AiSkillsWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *AiSkillsWrapper) GetResponse() AiSkillsDto`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *AiSkillsWrapper) GetResponseOk() (*AiSkillsDto, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *AiSkillsWrapper) SetResponse(v AiSkillsDto)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *AiSkillsWrapper) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetCount

`func (o *AiSkillsWrapper) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AiSkillsWrapper) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AiSkillsWrapper) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AiSkillsWrapper) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetTotal

`func (o *AiSkillsWrapper) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AiSkillsWrapper) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AiSkillsWrapper) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AiSkillsWrapper) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *AiSkillsWrapper) GetLinks() []GetPortalPrices200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AiSkillsWrapper) GetLinksOk() (*[]GetPortalPrices200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AiSkillsWrapper) SetLinks(v []GetPortalPrices200ResponseLinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AiSkillsWrapper) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetStatus

`func (o *AiSkillsWrapper) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AiSkillsWrapper) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AiSkillsWrapper) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AiSkillsWrapper) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *AiSkillsWrapper) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *AiSkillsWrapper) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *AiSkillsWrapper) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *AiSkillsWrapper) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


