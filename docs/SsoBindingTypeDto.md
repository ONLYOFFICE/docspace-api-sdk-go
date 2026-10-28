# SsoBindingTypeDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saml20HttpPost** | Pointer to **NullableString** | The SAML 2.0 HTTP POST binding, which carries the request in a self-submitting form. It is what the  built-in configuration uses and the one to pick when requests are signed, since it has no length limit. | [optional] [readonly] 
**Saml20HttpRedirect** | Pointer to **NullableString** | The SAML 2.0 HTTP redirect binding, which carries the request in the query string and is therefore bound  by the length a URL may have. | [optional] [readonly] 

## Methods

### NewSsoBindingTypeDto

`func NewSsoBindingTypeDto() *SsoBindingTypeDto`

NewSsoBindingTypeDto instantiates a new SsoBindingTypeDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoBindingTypeDtoWithDefaults

`func NewSsoBindingTypeDtoWithDefaults() *SsoBindingTypeDto`

NewSsoBindingTypeDtoWithDefaults instantiates a new SsoBindingTypeDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaml20HttpPost

`func (o *SsoBindingTypeDto) GetSaml20HttpPost() string`

GetSaml20HttpPost returns the Saml20HttpPost field if non-nil, zero value otherwise.

### GetSaml20HttpPostOk

`func (o *SsoBindingTypeDto) GetSaml20HttpPostOk() (*string, bool)`

GetSaml20HttpPostOk returns a tuple with the Saml20HttpPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20HttpPost

`func (o *SsoBindingTypeDto) SetSaml20HttpPost(v string)`

SetSaml20HttpPost sets Saml20HttpPost field to given value.

### HasSaml20HttpPost

`func (o *SsoBindingTypeDto) HasSaml20HttpPost() bool`

HasSaml20HttpPost returns a boolean if a field has been set.

### SetSaml20HttpPostNil

`func (o *SsoBindingTypeDto) SetSaml20HttpPostNil(b bool)`

 SetSaml20HttpPostNil sets the value for Saml20HttpPost to be an explicit nil

### UnsetSaml20HttpPost
`func (o *SsoBindingTypeDto) UnsetSaml20HttpPost()`

UnsetSaml20HttpPost ensures that no value is present for Saml20HttpPost, not even an explicit nil
### GetSaml20HttpRedirect

`func (o *SsoBindingTypeDto) GetSaml20HttpRedirect() string`

GetSaml20HttpRedirect returns the Saml20HttpRedirect field if non-nil, zero value otherwise.

### GetSaml20HttpRedirectOk

`func (o *SsoBindingTypeDto) GetSaml20HttpRedirectOk() (*string, bool)`

GetSaml20HttpRedirectOk returns a tuple with the Saml20HttpRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaml20HttpRedirect

`func (o *SsoBindingTypeDto) SetSaml20HttpRedirect(v string)`

SetSaml20HttpRedirect sets Saml20HttpRedirect field to given value.

### HasSaml20HttpRedirect

`func (o *SsoBindingTypeDto) HasSaml20HttpRedirect() bool`

HasSaml20HttpRedirect returns a boolean if a field has been set.

### SetSaml20HttpRedirectNil

`func (o *SsoBindingTypeDto) SetSaml20HttpRedirectNil(b bool)`

 SetSaml20HttpRedirectNil sets the value for Saml20HttpRedirect to be an explicit nil

### UnsetSaml20HttpRedirect
`func (o *SsoBindingTypeDto) UnsetSaml20HttpRedirect()`

UnsetSaml20HttpRedirect ensures that no value is present for Saml20HttpRedirect, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


