# IsDefaultWhiteLabelLogosDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **NullableString** | The stable name of the slot, matching the `name` of the same slot in  `GET api/2.0/settings/whitelabel/logos` - `LightSmall`, `LoginPage`, `Favicon`, `DocsEditor` and the rest,  plus `Notification`, which that list leaves out. The wordmark check reports the fixed name `logotext`  instead of a slot. | 
**Default** | **bool** | Whether the slot has never been written for this portal, in which case the built-in image is what gets  rendered. It turns `false` once an image has been stored, for either the light or the dark theme, and back  to `true` after the matching restore operation. For `logotext` it stays `true` when the built-in wordmark  itself is saved, because saving that value counts as clearing the setting. | 

## Methods

### NewIsDefaultWhiteLabelLogosDto

`func NewIsDefaultWhiteLabelLogosDto(name NullableString, default_ bool, ) *IsDefaultWhiteLabelLogosDto`

NewIsDefaultWhiteLabelLogosDto instantiates a new IsDefaultWhiteLabelLogosDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIsDefaultWhiteLabelLogosDtoWithDefaults

`func NewIsDefaultWhiteLabelLogosDtoWithDefaults() *IsDefaultWhiteLabelLogosDto`

NewIsDefaultWhiteLabelLogosDtoWithDefaults instantiates a new IsDefaultWhiteLabelLogosDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IsDefaultWhiteLabelLogosDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IsDefaultWhiteLabelLogosDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IsDefaultWhiteLabelLogosDto) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *IsDefaultWhiteLabelLogosDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *IsDefaultWhiteLabelLogosDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDefault

`func (o *IsDefaultWhiteLabelLogosDto) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *IsDefaultWhiteLabelLogosDto) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *IsDefaultWhiteLabelLogosDto) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


