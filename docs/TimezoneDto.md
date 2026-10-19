# TimezoneDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | The IANA identifier of the time zone. This is the value the portal time zone is set to, so pass it on  unchanged to `PUT api/2.0/settings/timeandlanguage`. | 
**DisplayName** | **NullableString** | The label to show for the zone, carrying its UTC offset as it stood when the list was built. The offset is a  snapshot rather than a rule, so a zone observing daylight saving reads differently at other times of the  year; sort and match on `id` instead. | 

## Methods

### NewTimezoneDto

`func NewTimezoneDto(id NullableString, displayName NullableString, ) *TimezoneDto`

NewTimezoneDto instantiates a new TimezoneDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimezoneDtoWithDefaults

`func NewTimezoneDtoWithDefaults() *TimezoneDto`

NewTimezoneDtoWithDefaults instantiates a new TimezoneDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimezoneDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimezoneDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimezoneDto) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *TimezoneDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TimezoneDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDisplayName

`func (o *TimezoneDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TimezoneDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TimezoneDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### SetDisplayNameNil

`func (o *TimezoneDto) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *TimezoneDto) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


